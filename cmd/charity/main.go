package main

import (
	"net"

	"github.com/GP-Hacks/charity/internal/config"
	"github.com/GP-Hacks/charity/internal/service_provider"
	"github.com/GP-Hacks/charity/internal/utils/logger"
	proto "github.com/GP-Hacks/proto/pkg/api/charity"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

func main() {
	config.LoadConfig("./config")
	logger.SetupLogger()
	serviceProvider := service_provider.NewServiceProvider()

	log.Info().Msg("Init app")

	charityController := serviceProvider.CharityController()

	grpcServer := grpc.NewServer(grpc.Creds(insecure.NewCredentials()))
	reflection.Register(grpcServer)

	proto.RegisterCharityServiceServer(grpcServer, charityController)

	list, err := net.Listen("tcp", ":"+config.Cfg.Grpc.Port)
	if err != nil {
		log.Fatal().Msg("Failed start listen port")
	}

	err = grpcServer.Serve(list)
	if err != nil {
		log.Fatal().Msg("Failed serve grpc")
	}
}
