module github.com/OliverSchlueter/mauerstrassenloewen/monitoring

go 1.24.2

require (
	github.com/OliverSchlueter/mauerstrassenloewen/common v0.0.0-00010101000000-000000000000
	github.com/justinas/alice v1.2.0
	github.com/nats-io/nats.go v1.43.0
)

require (
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/klauspost/compress v1.18.0 // indirect
	github.com/munnerz/goautoneg v0.0.0-20191010083416-a7dc8b61c822 // indirect
	github.com/nats-io/nkeys v0.4.11 // indirect
	github.com/nats-io/nuid v1.0.1 // indirect
	github.com/prometheus/client_golang v1.22.0 // indirect
	github.com/prometheus/client_model v0.6.2 // indirect
	github.com/prometheus/common v0.64.0 // indirect
	github.com/prometheus/procfs v0.16.1 // indirect
	golang.org/x/crypto v0.38.0 // indirect
	golang.org/x/sys v0.33.0 // indirect
	google.golang.org/protobuf v1.36.6 // indirect
)

replace github.com/OliverSchlueter/mauerstrassenloewen/common => ../common
