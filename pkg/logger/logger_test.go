package logger

import (
	"testing"

	"github.com/pkg/errors"
)

func TestJSONLogger(t *testing.T) {
	logger, err := NewJSONLogger(
		WithField("defined_key", "defined_value"),
	)
	if err != nil {
		t.Fatal(err)
	}
	defer logger.Sync()

	err = errors.New("pkg error")
	logger.Error("err occurs", WrapMeta(nil, NewMeta("para1", "value1"), NewMeta("para2", "value2"))...)
	logger.Error("err occurs", WrapMeta(err, NewMeta("para1", "value1"), NewMeta("para2", "value2"))...)

}

func BenchmarkJsonLogger(b *testing.B) {
	b.ResetTimer()
	logger, err := NewJSONLogger(
		WithField("defined_key", "defined_value"),
	)
	if err != nil {
		b.Fatal(err)
	}

	defer logger.Sync()

}
