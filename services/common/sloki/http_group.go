package sloki

import (
	"log/slog"
	"net/http"
)

func WrapRequest(r *http.Request) slog.Attr {
	method := slog.Any("method", r.Method)
	url := slog.Any("url", r.URL.String())
	userAgent := slog.Any("userAgent", r.UserAgent())
	headers := slog.Any("headers", r.Header)
	referrer := slog.Any("referrer", r.Referer())
	body := slog.Any("body", r.Body)

	return slog.Group("request", method, url, userAgent, headers, referrer, body)
}

func WrapError(err error) slog.Attr {
	return slog.Any("error", err)
}
