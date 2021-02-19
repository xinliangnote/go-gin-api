package httpclient

import (
	"bytes"
	"context"
	"crypto/tls"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/xinliangnote/go-gin-api/pkg/trace"

	"github.com/pkg/errors"
	"go.uber.org/zap"
)

const (
	// _StatusReadRespErr read resp body err, should re-call doHTTP again.
	_StatusReadRespErr = -204
	// _StatusDoReqErr do req err, should re-call doHTTP again.
	_StatusDoReqErr = -500
)

var defaultClient = &http.Client{
	Transport: &http.Transport{
		DisableKeepAlives:  true,
		DisableCompression: true,
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
		MaxIdleConns:        100,
		MaxConnsPerHost:     100,
		MaxIdleConnsPerHost: 100,
	},
}

func doHTTP(ctx context.Context, method, url string, payload []byte, opt *option) ([]byte, int, error) {
	ts := time.Now()

	if mock := opt.mock; mock != nil {
		if opt.dialog != nil {
			opt.dialog.AppendResponse(&trace.Response{
				HttpCode:    http.StatusOK,
				HttpCodeMsg: http.StatusText(http.StatusOK),
				Body:        string(mock()),
				CostSeconds: time.Since(ts).Seconds(),
			})
		}
		return mock(), http.StatusOK, nil
	}

	req, err := http.NewRequestWithContext(ctx, method, url, bytes.NewReader(payload))
	if err != nil {
		return nil, -1, errors.Wrapf(err, "new request [%s %s] err", method, url)
	}

	for key, value := range opt.header {
		req.Header.Set(key, value[0])
	}

	resp, err := defaultClient.Do(req)
	if err != nil {
		err = errors.Wrapf(err, "do request [%s %s] err", method, url)
		if opt.dialog != nil {
			opt.dialog.AppendResponse(&trace.Response{
				Body:        err.Error(),
				CostSeconds: time.Since(ts).Seconds(),
			})
		}

		if opt.logger != nil {
			opt.logger.Warn("doHTTP got err", zap.Error(err))
		}
		return nil, _StatusDoReqErr, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		err = errors.Wrapf(err, "read resp body from [%s %s] err", method, url)
		if opt.dialog != nil {
			opt.dialog.AppendResponse(&trace.Response{
				Body:        err.Error(),
				CostSeconds: time.Since(ts).Seconds(),
			})
		}

		if opt.logger != nil {
			opt.logger.Warn("doHTTP got err", zap.Error(err))
		}
		return nil, _StatusReadRespErr, err
	}

	defer func() {
		if opt.dialog != nil {
			opt.dialog.AppendResponse(&trace.Response{
				Header:      resp.Header,
				HttpCode:    resp.StatusCode,
				HttpCodeMsg: resp.Status,
				Body:        string(body), // unsafe
				CostSeconds: time.Since(ts).Seconds(),
			})
		}
	}()

	if resp.StatusCode != http.StatusOK {
		return nil, resp.StatusCode, newReplyErr(
			resp.StatusCode,
			body,
			errors.Errorf("do [%s %s] return code: %d message: %s", method, url, resp.StatusCode, string(body)),
		)
	}

	return body, http.StatusOK, nil
}

// addFormValuesIntoURL append url.Values into url string
func addFormValuesIntoURL(rawURL string, form url.Values) (string, error) {
	if rawURL == "" {
		return "", errors.New("rawURL required")
	}
	if len(form) == 0 {
		return "", errors.New("form required")
	}

	target, err := url.Parse(rawURL)
	if err != nil {
		return "", errors.Wrapf(err, "parse rawURL `%s` err", rawURL)
	}

	urlValues := target.Query()
	for key, values := range form {
		for _, value := range values {
			urlValues.Add(key, value)
		}
	}

	target.RawQuery = urlValues.Encode()
	return target.String(), nil
}
