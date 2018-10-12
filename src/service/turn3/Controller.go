package turn3

import (
	Context "golang.org/x/net/context"
)

// Server for MysimServer
type Server struct{}

// CalculatorFlop Grpc呼叫的入口 MysimServer.PokerCalculator()
func (s *Server) CalculatorFlop(ctx Context.Context, in *IntoData) (*OutData, error) {
	return Service(ctx, in)
}
