package sloki

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"strconv"
)

type Service struct {
	url          string
	service      string
	consoleLevel slog.Level
	lokiLevel    slog.Level
	httpClient   *http.Client
}

type Configuration struct {
	URL          string
	Service      string
	ConsoleLevel slog.Level
	LokiLevel    slog.Level
}

func NewService(cfg Configuration) *Service {
	return &Service{
		url:          cfg.URL,
		service:      cfg.Service,
		consoleLevel: cfg.ConsoleLevel,
		lokiLevel:    cfg.LokiLevel,
		httpClient:   &http.Client{},
	}
}

func (s *Service) Enabled(_ context.Context, _ slog.Level) bool {
	return true
}

func (s *Service) printToConsole(level slog.Level) bool {
	return level >= s.consoleLevel
}

func (s *Service) sendToLoki(level slog.Level) bool {
	return level >= s.lokiLevel
}

func (s *Service) Handle(ctx context.Context, r slog.Record) error {
	if !s.printToConsole(r.Level) {
		return nil
	}

	attrs := map[string]string{}
	r.Attrs(func(a slog.Attr) bool {
		attrs[a.Key] = fmt.Sprint(a.Value)
		return true
	})

	var attrJson []byte
	if len(attrs) > 0 {
		attrJson, _ = json.Marshal(attrs)
		attrJson = append([]byte(" "), attrJson...)
	}

	fmt.Printf("%s [%s] %s%s\n",
		r.Time.Format("2006-01-02 15:04:05"),
		r.Level.String(),
		r.Message,
		string(attrJson),
	)

	if !s.sendToLoki(r.Level) {
		return nil
	}

	unixTimestamp := strconv.FormatInt(r.Time.UnixNano(), 10)
	if err := s.pushLogToLoki(unixTimestamp, r.Level.String(), r.Message, attrs); err != nil {
		fmt.Printf("Failed to send log to Loki: %v\n", err)
		return err
	}

	return nil
}

func (s *Service) WithAttrs(_ []slog.Attr) slog.Handler {
	return s
}

func (s *Service) WithGroup(_ string) slog.Handler {
	return s
}

func (s *Service) pushLogToLoki(timestamp, level, message string, attrs map[string]string) error {
	labels := map[string]string{
		"service": s.service,
		"level":   level,
	}
	for k, v := range attrs {
		labels[k] = v
	}

	req := PushLogsRequest{
		Streams: []Stream{
			{
				Labels: labels,
				Values: [][]string{
					{timestamp, message},
				},
			},
		},
	}

	reqJson, err := json.Marshal(req)
	if err != nil {
		return fmt.Errorf("failed to marshal request: %w", err)
	}

	resp, err := s.httpClient.Post(s.url, "application/json", bytes.NewReader(reqJson))
	if err != nil {
		return fmt.Errorf("failed to send request to Loki: %w", err)
	}

	if resp.StatusCode != http.StatusNoContent {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("loki responded with status %d: %s", resp.StatusCode, string(body))
	}

	return nil
}
