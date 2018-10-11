package grpcinit

import (
	Flopplayer2 "git.com.ggttoo44/src/service/flopplayer2"
	Mysim "git.com.ggttoo44/src/service/mysim"
	grpc "google.golang.org/grpc"
)

// RegisterAllServer 在這裡註冊所有的GRPC Server
func RegisterAllServer(s *grpc.Server) {
	Mysim.RegisterMysimServer(s, &Mysim.Server{})
	Flopplayer2.RegisterFlopplayer2Server(s, &Flopplayer2.Server{})
}
