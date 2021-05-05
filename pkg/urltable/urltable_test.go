package urltable

import (
	"strings"
	"testing"
)

func TestFormat(t *testing.T) {
	pattern, err := Format(" view  / a / b / c   ")
	if err != nil {
		t.Fatal(err)
	}

	if pattern != "VIEW/a/b/c" {
		t.Fatal("format failed")
	}
}

func TestParse(t *testing.T) {
	for i, pattern := range []string{
		"get/ a / b / c   ",
		"get/ a / b / * / **  ",
		"get/ a / b / * / c / **  ",
		"get/ a / b / * / * / c/ **  ",
		"get/ a / b / * / * / c/   ",
	} {
		paths, err := parse(pattern)
		if err != nil {
			t.Fatal(pattern, "should be legal; err: ", err.Error())
		}
		t.Log(i, strings.Join(paths, delimiter))
	}

	for _, pattern := range []string{
		"   ",
		" /  ",
		" x /  ",
		"get/  ",
		"get/ * ",
		"get/ ** ",
		"get/ ** / * ",
		"get/ ** / ** ",
		"get/ a / **  / * ",
		"get/ a /   / * ",
		"get/ a /   / ** ",
		"get/ a /   /  ",
		"get/ a /  * /   / ",
		"get/ a /  * / **  / ",
	} {
		if _, err := parse(pattern); err == nil {
			t.Fatal(pattern, "should be illegal")
		}
	}
}

func TestAppend(t *testing.T) {
	table := NewTable()

	if err := table.Append("get/a/b"); err != nil {
		t.Fatal("shouldn't be err")
	}

	if err := table.Append("get/a/b/*"); err != nil {
		t.Fatal("shouldn't be err")
	}
	if err := table.Append("get/a/b/*/**"); err != nil {
		t.Fatal("shouldn't be err")
	}

	if err := table.Append("get/a/b/c/*"); err != nil {
		t.Fatal("shouldn't be err")
	}
	if err := table.Append("get/a/b/c/*/**"); err != nil {
		t.Fatal("shouldn't be err")
	}

	t.Log(table.Size())
	if err := table.Append("get/a/b/c/*"); err != nil {
		t.Fatal("shouldn't be err")
	}
	t.Log(table.Size())

	if err := table.Append("get/a/b/**"); err == nil {
		t.Fatal("should be err")
	}
	if err := table.Append("get/a/b/*/*"); err == nil {
		t.Fatal("should be err")
	}
	if err := table.Append("get/a/b/*/c"); err == nil {
		t.Fatal("should be err")
	}
}

func TestMapping(t *testing.T) {
	table := NewTable()

	if err := table.Append("get/a/b"); err != nil {
		t.Fatal("shouldn't be err")
	}

	if err := table.Append("get/a/b/*"); err != nil {
		t.Fatal("shouldn't be err")
	}
	if err := table.Append("get/a/b/*/**"); err != nil {
		t.Fatal("shouldn't be err")
	}

	if err := table.Append("get/a/b/c/*"); err != nil {
		t.Fatal("shouldn't be err")
	}
	if err := table.Append("get/a/b/c/*/**"); err != nil {
		t.Fatal("shouldn't be err")
	}

	if pattern, _ := table.Mapping("get/a/b"); pattern == "" {
		t.Fatal("should contains")
	} else {
		t.Log("get/a/b", ">>", pattern)
	}
	if pattern, _ := table.Mapping("get/a/b/"); pattern == "" {
		t.Fatal("should contains")
	} else {
		t.Log("get/a/b/", ">>", pattern)
	}
	if pattern, _ := table.Mapping("get/a/b/x"); pattern == "" {
		t.Fatal("should contains")
	} else {
		t.Log("get/a/b/x", ">>", pattern)
	}
	if pattern, _ := table.Mapping("get/a/b/x/y/z"); pattern == "" {
		t.Fatal("should contains")
	} else {
		t.Log("get/a/b/x/y/z", ">>", pattern)
	}

	if pattern, _ := table.Mapping("get/a/b/c/"); pattern == "" {
		t.Fatal("should contains")
	} else {
		t.Log("get/a/b/c/", ">>", pattern)
	}
	if pattern, _ := table.Mapping("get/a/b/c/d/e/f"); pattern == "" {
		t.Fatal("should contains")
	} else {
		t.Log("get/a/b/c/d/e/f", ">>", pattern)
	}

	if pattern, _ := table.Mapping("get/a/"); pattern != "" {
		t.Fatal("shouldn't contains")
	} else {
		t.Log("get/a/", ">>", pattern)
	}
	if pattern, _ := table.Mapping("get/a/c"); pattern != "" {
		t.Fatal("shouldn't contains")
	} else {
		t.Log("get/a/c", ">>", pattern)
	}
}
