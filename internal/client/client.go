// Package client implement client for pow server.
package client

import (
	"bufio"
	"context"
	"fmt"
	"net"
	"strings"

	"go.uber.org/zap"
)

type Solver interface {
	SolvePoW(challenge string) string
}

type Client struct {
	serverAddr  string
	logger      *zap.Logger
	solver      Solver
	rateLimiter chan struct{}
}

func New(serverAddr string, logger *zap.Logger, solver Solver) *Client {
	return &Client{
		serverAddr:  serverAddr,
		logger:      logger,
		solver:      solver,
		rateLimiter: make(chan struct{}, 1),
	}
}

func (c *Client) Run(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			c.logger.Info("Client received shutdown signal")
			return
		case c.rateLimiter <- struct{}{}:
			err := c.sendRequest()
			if err != nil {
				c.logger.Error("Error sending request", zap.Error(err))
			}
			<-c.rateLimiter
		}
	}
}

func (c *Client) sendRequest() error {
	conn, err := net.Dial("tcp", c.serverAddr)
	if err != nil {
		return fmt.Errorf("failed to connect to server: %w", err)
	}
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			c.logger.Error("failed to close connection", zap.Error(err))
		}
	}(conn)

	reader := bufio.NewReader(conn)
	challengeMsg, err := reader.ReadString('\n')
	if err != nil {
		return fmt.Errorf("error reading from server: %w", err)
	}

	challenge := strings.TrimSpace(strings.Split(challengeMsg, ": ")[1])
	c.logger.Info("Received challenge", zap.String("challenge", challenge))

	nonce := c.solver.SolvePoW(challenge)
	c.logger.Info("Solved challenge", zap.String("nonce", nonce))

	_, err = conn.Write([]byte(nonce + "\n"))
	if err != nil {
		return fmt.Errorf("error writing to server: %w", err)
	}

	response, err := reader.ReadString('\n')
	if err != nil {
		return fmt.Errorf("error reading from server: %w", err)
	}

	c.logger.Info("Received response", zap.String("response", response))

	return nil
}
