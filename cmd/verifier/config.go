package main

// Config represents the configuration for the POW server.
// nolint:lll
type Config struct {
	Port       int    `long:"port" env:"POW_SERVER_HTTP_PORT" default:"9080" description:"HTTP server bind port"`
	Difficulty int    `long:"difficulty" env:"POW_SERVER_DIFFICULTY" default:"5" description:"Proof of work difficulty"`
	QuotesPath string `long:"quotes_path" env:"POW_SERVER_QUOTES_PATH" default:"./quotes.txt" description:"Path to quotes file"`
}
