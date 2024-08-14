package main

import (
	"api-service/config"
	"api-service/server/api"
	"api-service/server/db/sql"
	"net"
	"os"
	"os/signal"

	pb "api-service/server/proto"

	"go.uber.org/zap"
	"google.golang.org/grpc"
)

var dsn = "host=localhost user=postgres password=postgres dbname=commiss_art_app port=5432 sslmode=disable Timezone=Asia/Jakarta"

func main() {

	log := zap.Must(zap.NewProduction())

	log.Info("Starting gRPC Server...")

	listener, err := net.Listen("tcp", "localhost:9000")
	if err != nil {
		log.Debug("Couldn't listen tcp in port 9000")
		return
	}

	gormConf, err := config.NewGorm(dsn)
	if err != nil {
		os.Exit(1)
	}

	repos := sql.NewRepository(gormConf)

	srv := api.New(repos)

	grpcServer := grpc.NewServer()
	pb.RegisterApiServiceServer(
		grpcServer, srv,
	)

	log.Info("gRPC Server is running on port 9000")

	go func() {
		err = grpcServer.Serve(listener)
		if err != nil {
			log.Debug("Couldn't serve grpc listener")
			return
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	<-c

	log.Info("Stopping gRPC server...")
	grpcServer.Stop()

	log.Sync()

}
