package mysim

import (
	Context "golang.org/x/net/context"
)

// Server for MysimServer
type Server struct{}

// PokerCalculator Grpc呼叫的入口 MysimServer.PokerCalculator()
func (s *Server) PokerCalculator(ctx Context.Context, in *IntoData) (*OutData, error) {
	return Service(ctx, in)
}
