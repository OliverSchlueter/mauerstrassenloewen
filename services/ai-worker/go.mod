module github.com/OliverSchlueter/mauerstrassenloewen/ai-worker

go 1.24.2

require (
	common v0.0.0
	github.com/OliverSchlueter/sloki v0.0.1
	github.com/google/uuid v1.6.0
	github.com/nats-io/nats.go v1.41.0
	github.com/ollama/ollama v0.6.4
	github.com/sashabaranov/go-openai v1.38.1
)

require (
	github.com/klauspost/compress v1.18.0 // indirect
	github.com/nats-io/nkeys v0.4.9 // indirect
	github.com/nats-io/nuid v1.0.1 // indirect
	golang.org/x/crypto v0.33.0 // indirect
	golang.org/x/sys v0.30.0 // indirect
)

replace common => ../common
