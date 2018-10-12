package grpcinit

import (
	Flopplayer2 "git.com.ggttoo44/src/service/flopplayer2"
	Flopplayer3 "git.com.ggttoo44/src/service/flopplayer3"
	Mysim "git.com.ggttoo44/src/service/mysim"
	Turn2 "git.com.ggttoo44/src/service/turn2"
	Turn3 "git.com.ggttoo44/src/service/turn3"
	grpc "google.golang.org/grpc"
)

// RegisterAllServer 在這裡註冊所有的GRPC Server
func RegisterAllServer(s *grpc.Server) {
	Mysim.RegisterMysimServer(s, &Mysim.Server{})
	Turn2.RegisterTurn2Server(s, &Turn2.Server{})
	Turn3.RegisterTurn3Server(s, &Turn3.Server{})
	Flopplayer2.RegisterFlopplayer2Server(s, &Flopplayer2.Server{})
	Flopplayer3.RegisterFlopplayer3Server(s, &Flopplayer3.Server{})
}
