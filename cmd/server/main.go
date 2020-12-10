package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/grpc"
	"net"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
	"todoapp/config"
	"todoapp/lib/errors"
	"todoapp/lib/mysql"
	"todoapp/server"
)

func main() {
	conf := config.Load()
	db := mysql.MustConnect(conf.MySQL)
	root := server.NewRoot(db)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	grpcServer := grpc.NewServer()

	mux := runtime.NewServeMux(
		runtime.WithProtoErrorHandler(errors.CustomHTTPError),
		runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{EmitDefaults: true}),
	)

	httpMux := http.NewServeMux()
	httpMux.Handle("/metrics", promhttp.Handler())
	httpMux.Handle("/", mux)

	httpServer := http.Server{
		Addr:    conf.Server.HTTP.String(),
		Handler: httpMux,
	}

	ctx := context.Background()
	root.Register(ctx, grpcServer, mux, conf.Server.GRPC.String(), []grpc.DialOption{grpc.WithInsecure()})

	//--------------------------------
	// Run HTTP & gRPC servers
	//--------------------------------
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()

		err := httpServer.ListenAndServe()
		if err != nil {
			panic(err)
		}
	}()

	go func() {
		defer wg.Done()

		listener, err := net.Listen("tcp", conf.Server.GRPC.String())
		if err != nil {
			panic(err)
		}

		err = grpcServer.Serve(listener)
		if err != nil {
			panic(err)
		}
	}()

	//--------------------------------
	// Graceful Shutdown
	//--------------------------------
	<-stop

	ctx = context.Background()
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	grpcServer.GracefulStop()
	err := httpServer.Shutdown(ctx)
	if err != nil {
		panic(err)
	}

	err = db.Close()
	if err != nil {
		panic(err)
	}
}
