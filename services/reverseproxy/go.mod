module github.com/OliverSchlueter/mauerstrassenloewen/reverseproxy

go 1.24.2

replace github.com/OliverSchlueter/mauerstrassenloewen/common => ../common

require (
	github.com/OliverSchlueter/mauerstrassenloewen/common v0.0.0-00010101000000-000000000000
	github.com/justinas/alice v1.2.0
	github.com/nats-io/nats.go v1.41.1
)

require (
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/klauspost/compress v1.18.0 // indirect
	github.com/munnerz/goautoneg v0.0.0-20191010083416-a7dc8b61c822 // indirect
	github.com/nats-io/nkeys v0.4.9 // indirect
	github.com/nats-io/nuid v1.0.1 // indirect
	github.com/prometheus/client_golang v1.22.0 // indirect
	github.com/prometheus/client_model v0.6.2 // indirect
	github.com/prometheus/common v0.63.0 // indirect
	github.com/prometheus/procfs v0.16.0 // indirect
	golang.org/x/crypto v0.31.0 // indirect
	golang.org/x/sys v0.32.0 // indirect
	google.golang.org/protobuf v1.36.6 // indirect
)
