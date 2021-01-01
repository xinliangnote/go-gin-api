package env

import "testing"

func TestActive(t *testing.T) {
	t.Log(Active().Value())
}
