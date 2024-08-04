package main

type Config struct {
	ServerAddr string `long:"address" env:"PROVER_SERVER_ADDR" description:"Server address" required:"true"`
	Difficulty int    `long:"difficulty" env:"PROVER_DIFFICULTY" description:"Proof of Work difficulty level" default:"4"`
}
