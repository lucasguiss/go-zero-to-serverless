package main

import "github.com/lucasguiss/go-zero-to-serverless/server"

func main() {
	server := server.NewServer()
	server.Run()
}
