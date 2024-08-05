// Package server provides a simple proof-of-work server.
package server

import (
	"context"
	"fmt"
	"io"
	"net"
	"strconv"
	"strings"
	"sync"
	"time"

	"go.uber.org/zap"
)

type QuotesService interface {
	GetRandomQuote() (string, error)
}

type HashcashService interface {
	GenerateChallenge() string
	VerifyPoW(challenge, nonce string) bool
}

type POWServer struct {
	port            int
	logger          *zap.Logger
	quotesService   QuotesService
	hashcashService HashcashService
	listener        net.Listener
	sessions        sync.Map
	wg              sync.WaitGroup
	waitSolving     time.Duration
}

// NewPOWServer creates a new POWServer instance.
//
//nolint:exhaustruct,lll
func NewPOWServer(port int, waitSolving time.Duration, logger *zap.Logger, quotesService QuotesService, hashcashService HashcashService) *POWServer {
	return &POWServer{
		port:            port,
		logger:          logger,
		quotesService:   quotesService,
		hashcashService: hashcashService,
		waitSolving:     waitSolving,
	}
}

// handleConnection handles an individual client connection.
//
//nolint:cyclop,errorlint,funlen
func (s *POWServer) handleConnection(ctx context.Context, conn net.Conn) {
	defer s.wg.Done()
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			s.logger.Error("Error closing connection", zap.Error(err))
		}
	}(conn)

	clientAddr := conn.RemoteAddr().String()
	challenge := s.hashcashService.GenerateChallenge()
	s.sessions.Store(clientAddr, challenge)

	_, err := conn.Write([]byte("Challenge: " + challenge + "\n"))
	if err != nil {
		s.logger.Error("Error writing to connection", zap.Error(err))
		s.sessions.Delete(clientAddr)
		return
	}

	buf := make([]byte, 1024)
	connCtx, cancel := context.WithTimeout(ctx, s.waitSolving)
	defer cancel()

	for {
		select {
		case <-ctx.Done():
			s.logger.Info("Server is shutting down, closing connection", zap.String("clientAddr", clientAddr))
			return
		case <-connCtx.Done():
			s.logger.Info("Connection timed out, closing connection", zap.String("clientAddr", clientAddr))
			return
		default:
			n, err := conn.Read(buf)
			if err != nil {
				if err != io.EOF {
					s.logger.Error("Error reading from connection", zap.Error(err))
				}
				s.sessions.Delete(clientAddr)
				return
			}
			nonce := strings.TrimSpace(string(buf[:n]))

			storedChallenge, ok := s.sessions.Load(clientAddr)
			if !ok {
				s.logger.Error("Session not found", zap.String("clientAddr", clientAddr))
				return
			}

			if s.hashcashService.VerifyPoW(storedChallenge.(string), nonce) {
				quote, err := s.quotesService.GetRandomQuote()
				if err != nil {
					s.logger.Error("Error getting random quote", zap.Error(err))
					return
				}
				_, err = conn.Write([]byte("Quote: " + quote + "\n"))
				if err != nil {
					s.logger.Error("Error writing to connection", zap.Error(err))
				}
				s.logger.Info("PoW verified and quote sent", zap.String("quote", quote))
			} else {
				_, err := conn.Write([]byte("Invalid PoW\n"))
				if err != nil {
					s.logger.Error("Error writing to connection", zap.Error(err))
				}
				s.logger.Info("Invalid PoW", zap.String("challenge", storedChallenge.(string)), zap.String("nonce", nonce))
			}

			s.sessions.Delete(clientAddr)
			return
		}
	}
}

// Start starts the POWServer and listens for incoming connections.
func (s *POWServer) Start(ctx context.Context) error {
	addr := ":" + strconv.Itoa(s.port)
	var err error
	s.listener, err = net.Listen("tcp", addr)
	if err != nil {
		return fmt.Errorf("error starting server: %w", err)
	}

	s.logger.Info("Server is listening", zap.String("address", addr))

	go func() {
		<-ctx.Done()
		s.logger.Info("Received shutdown signal, shutting down server...")
		err := s.listener.Close()
		if err != nil {
			s.logger.Error("Error stopping server", zap.Error(err))
			return
		}
		s.logger.Info("Server stopped")
	}()

	for {
		conn, err := s.listener.Accept()
		if err != nil {
			select {
			case <-ctx.Done():
				s.logger.Info("Server is shutting down, stopped accepting new connections")
				return nil
			default:
				s.logger.Error("Error accepting connection", zap.Error(err))
			}
			continue
		}
		s.wg.Add(1)
		go s.handleConnection(ctx, conn)
	}
}

// Stop stops the POWServer gracefully.
func (s *POWServer) Stop() {
	if s.listener != nil {
		err := s.listener.Close()
		if err != nil {
			s.logger.Error("Error stopping server", zap.Error(err))
		} else {
			s.logger.Info("Server stopped")
		}
	}
	s.wg.Wait()
}
