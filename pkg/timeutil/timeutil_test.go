package timeutil

import "testing"

func TestRFC3339ToCSTLayout(t *testing.T) {
	t.Log(RFC3339ToCSTLayout("2020-11-08T08:18:46+08:00"))
}

func TestCSTLayoutString(t *testing.T) {
	t.Log(CSTLayoutString())
}

func TestCSTLayoutStringToUnix(t *testing.T) {
	t.Log(CSTLayoutStringToUnix("2020-01-24 21:11:11"))
}

func TestGMTLayoutString(t *testing.T) {
	t.Log(GMTLayoutString())
}
