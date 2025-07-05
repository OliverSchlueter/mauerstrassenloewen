package telemetry

import (
	"context"
	"github.com/OliverSchlueter/mauerstrassenloewen/common/sloki"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/ollama/ollama/api"
	"log/slog"
	"time"
)

type Service struct {
	postgre *pgx.Conn
}

func NewService(conn *pgx.Conn) (*Service, error) {
	chatResponsesTable := `
	CREATE TABLE IF NOT EXISTS chat_responses (
	    timestamp TIMESTAMP,
	    model VARCHAR(255),
	    total_duration BIGINT,
	    load_duration BIGINT,
	    prompt_eval_count INT,
	    prompt_eval_duration BIGINT,
	    eval_count INT,
	    eval_duration BIGINT
    )
`
	if _, err := conn.Exec(context.Background(), chatResponsesTable); err != nil {
		return nil, err
	}

	chatMessagesTable := `
	CREATE TABLE IF NOT EXISTS chat_messages (
	    timestamp TIMESTAMP,
	    chat_id VARCHAR(255),
	    message_id VARCHAR(255)
    )
`
	if _, err := conn.Exec(context.Background(), chatMessagesTable); err != nil {
		return nil, err
	}

	return &Service{
		postgre: conn,
	}, nil
}

func (s *Service) TrackNewChatMessage(chatId string) {
	insert := `
	INSERT INTO chat_messages VALUES(
		  @timestamp, 
		  @chat_id, 
		  @message_id
	)
	`
	args := pgx.NamedArgs{
		"timestamp":  time.Now().Format(time.RFC3339),
		"chat_id":    chatId,
		"message_id": uuid.New().String(),
	}
	if _, err := s.postgre.Exec(context.Background(), insert, args); err != nil {
		slog.Error("failed to insert telemetry data", sloki.WrapError(err))
		return
	}

	slog.Debug("successfully inserted telemetry data")
	return
}

func (s *Service) TrackOllamaResponse(resp *api.ChatResponse) {
	insert := `
	INSERT INTO chat_responses VALUES(
		  @timestamp, 
		  @model, 
		  @total_duration, 
		  @load_duration, 
		  @prompt_eval_count, 
		  @prompt_eval_duration, 
		  @eval_count, 
		  @eval_duration
	)
	`
	args := pgx.NamedArgs{
		"timestamp":            time.Now().Format(time.RFC3339),
		"model":                resp.Model,
		"total_duration":       resp.TotalDuration.Milliseconds(),
		"load_duration":        resp.LoadDuration.Milliseconds(),
		"prompt_eval_count":    resp.PromptEvalCount,
		"prompt_eval_duration": resp.PromptEvalDuration.Milliseconds(),
		"eval_count":           resp.EvalCount,
		"eval_duration":        resp.EvalDuration.Milliseconds(),
	}
	if _, err := s.postgre.Exec(context.Background(), insert, args); err != nil {
		slog.Error("failed to insert telemetry data", sloki.WrapError(err))
		return
	}

	slog.Debug("successfully inserted telemetry data")
	return
}
