package errors

import (
	"errors"
	"testing"

	"go.uber.org/zap"
)

func TestErr(t *testing.T) {
	logger, _ := zap.NewProduction()

	logger.Info("errorf", zap.Error(Errorf("%s %d", "127.0.0.1", 80)))

	err := New("a dummy err")
	logger.Info("new", zap.Error(err))

	err = Wrap(err, "ping timeout err")
	logger.Info("wrap", zap.Error(err))

	err = Wrapf(err, "ip: %s port: %d", "localhost", 80)
	logger.Info("wrapf", zap.Error(err))

	err = WithStack(err)
	logger.Info("withstack", zap.Error(err))

	logger.Info("wrap std", zap.Error(Wrap(errors.New("std err"), "some err occurs")))

	logger.Info("wrapf std", zap.Error(Wrapf(errors.New("std err"), "ip: %s port: %d", "localhost", 80)))

	logger.Info("withstack std", zap.Error(WithStack(errors.New("std err"))))

	t.Logf("%+v", New("a dummy error"))
}
