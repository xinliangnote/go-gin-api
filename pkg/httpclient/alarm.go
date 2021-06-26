package httpclient

import (
	"bufio"
	"bytes"

	"go.uber.org/zap"
)

// AlarmVerify Verify parse the body and verify that it is correct
type AlarmVerify func(body []byte) (shouldAlarm bool)

type AlarmObject interface {
	Send(subject, body string) error
}

func onFailedAlarm(title string, raw []byte, logger *zap.Logger, alarmObject AlarmObject) {
	buf := bytes.NewBuffer(nil)

	scanner := bufio.NewScanner(bytes.NewReader(raw))
	for scanner.Scan() {
		buf.WriteString(scanner.Text())
		buf.WriteString(" <br/>")
	}

	if err := alarmObject.Send(title, buf.String()); err != nil && logger != nil {
		logger.Error("calls failed alarm err", zap.Error(err))
	}
}
