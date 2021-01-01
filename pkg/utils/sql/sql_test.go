package sql

import "testing"

func TestEscapeString(t *testing.T) {
	str := "' OR ''=' union select 1,database(),2#"
	t.Log(EscapeString(str))
}
