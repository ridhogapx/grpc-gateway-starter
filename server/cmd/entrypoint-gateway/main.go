package main

import (
	"context"
	"net/http"

	pb "api-service/server/proto"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.uber.org/zap"
	"google.golang.org/grpc"

	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	log := zap.Must(zap.NewProduction())

	log.Info("Starting gateway server", zap.String("addr", ":3000"))

	mux := runtime.NewServeMux()

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	ctx := context.Background()

	err := pb.RegisterApiServiceHandlerFromEndpoint(ctx, mux, "localhost:9000", opts)
	if err != nil {
		log.Debug("Couldn't register service from endpoint", zap.Error(err))
		return
	}

	http.ListenAndServe(":3000", mux)

}
