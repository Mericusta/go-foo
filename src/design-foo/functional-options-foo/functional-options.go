package functionalOptionFoo

import "time"

type Option func(*Server)

type Server struct {
	host    string
	port    int
	timeout time.Duration
	maxConn int
}

func New(opts ...Option) *Server {
	s := &Server{}
	for _, opt := range opts {
		opt(s)
	}
	return s
}

func WithHost(host string) Option {
	return func(s *Server) { s.host = host }
}

func WithPort(port int) Option {
	return func(s *Server) { s.port = port }
}

func WithTimeout(timeout time.Duration) Option {
	return func(s *Server) { s.timeout = timeout }
}

func WithMaxConn(maxConn int) Option {
	return func(s *Server) { s.maxConn = maxConn }
}
