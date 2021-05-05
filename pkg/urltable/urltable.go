package urltable

import (
	"net/http"
	"strings"

	"github.com/xinliangnote/go-gin-api/pkg/errors"
)

const (
	empty      = ""
	fuzzy      = "*"
	omitted    = "**"
	delimiter  = "/"
	methodView = "VIEW"
)

// parse and validate pattern
func parse(pattern string) ([]string, error) {
	const format = "[get, post, put, patch, delete, view]/{a-Z}+/{*}+/{**}"

	if pattern = strings.TrimLeft(strings.TrimSpace(pattern), delimiter); pattern == "" {
		return nil, errors.Errorf("pattern illegal, should in format of %s", format)
	}

	paths := strings.Split(pattern, delimiter)
	if len(paths) < 2 {
		return nil, errors.Errorf("pattern illegal, should in format of %s", format)
	}

	for i := range paths {
		paths[i] = strings.TrimSpace(paths[i])
	}

	// likes get/ get/* get/**
	if len(paths) == 2 && (paths[1] == empty || paths[1] == fuzzy || paths[1] == omitted) {
		return nil, errors.New("illegal wildcard")
	}

	switch paths[0] = strings.ToUpper(paths[0]); paths[0] {
	case http.MethodGet,
		http.MethodPost,
		http.MethodPut,
		http.MethodPatch,
		http.MethodDelete,
		methodView:
	default:
		return nil, errors.Errorf("only supports [%s %s %s %s %s %s]",
			http.MethodGet, http.MethodPost, http.MethodPut, http.MethodPatch, http.MethodDelete, methodView)
	}

	for k := 1; k < len(paths); k++ {
		if paths[k] == empty && k+1 != len(paths) {
			return nil, errors.New("pattern contains illegal empty path")
		}

		if paths[k] == omitted && k+1 != len(paths) {
			return nil, errors.New("pattern contains illegal omitted path")
		}
	}

	return paths, nil
}

// Format pattern
func Format(pattern string) (string, error) {
	paths, err := parse(pattern)
	if err != nil {
		return "", err
	}

	return strings.Join(paths, delimiter), nil
}

type section struct {
	leaf    bool
	mapping map[string]*section
}

func newSection() *section {
	return &section{mapping: make(map[string]*section)}
}

// Table a (thread unsafe) table to store urls
type Table struct {
	size int
	root *section
}

// NewTable create a table
func NewTable() *Table {
	return &Table{root: newSection()}
}

// Size contains how many urls
func (t *Table) Size() int {
	return t.size
}

// Append pattern
func (t *Table) Append(pattern string) error {
	paths, err := parse(pattern)
	if err != nil {
		return err
	}

	insert := false
	root := t.root
	for i, path := range paths {
		if (path == fuzzy && root.mapping[omitted] != nil) ||
			(path == omitted && root.mapping[fuzzy] != nil) ||
			(path != omitted && root.mapping[omitted] != nil) {
			return errors.Errorf("conflict at %s", strings.Join(paths[:i], delimiter))
		}

		next := root.mapping[path]
		if next == nil {
			next = newSection()
			root.mapping[path] = next
			insert = true
		}
		root = next
	}

	if insert {
		t.size++
	}

	root.leaf = true
	return nil
}

// Mapping url to pattern
func (t *Table) Mapping(url string) (string, error) {
	paths, err := parse(url)
	if err != nil {
		return "", err
	}

	pattern := make([]string, 0, len(paths))

	root := t.root
	for _, path := range paths {
		next := root.mapping[path]
		if next == nil {
			nextFuzzy := root.mapping[fuzzy]
			nextOmitted := root.mapping[omitted]
			if nextFuzzy == nil && nextOmitted == nil {
				return "", nil
			}

			if nextOmitted != nil {
				pattern = append(pattern, omitted)
				return strings.Join(pattern, delimiter), nil
			}

			next = nextFuzzy
			path = fuzzy
		}

		root = next
		pattern = append(pattern, path)
	}

	if root.leaf {
		return strings.Join(pattern, delimiter), nil
	}
	return "", nil
}
