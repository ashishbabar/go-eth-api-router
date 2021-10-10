module github.com/ashishbabar/go-eth-api-router

go 1.15

require (
	github.com/gorilla/mux v1.8.0
	github.com/spf13/viper v1.9.0
	go.uber.org/atomic v1.9.0 // indirect
	go.uber.org/multierr v1.7.0 // indirect
	go.uber.org/zap v1.19.1
)

replace github.com/ashishbabar/go-eth-api-router/handlers => ./handlers
