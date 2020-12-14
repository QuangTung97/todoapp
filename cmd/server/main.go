package main

import (
	"context"
	"fmt"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
	"time"
	"todoapp/config"
	"todoapp/lib/dblib"
	"todoapp/lib/errors"
	"todoapp/lib/mysql"
	"todoapp/server"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dblib.FinishRegisterQueries()

	rootCmd := &cobra.Command{
		Use: "server",
	}

	rootCmd.AddCommand(
		startServerCommand(),
		checkSQLCommand(),
	)

	err := rootCmd.Execute()
	if err != nil {
		fmt.Println(err)
	}
}

func startServerCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "start",
		Short: "start the server",
		Run: func(cmd *cobra.Command, args []string) {
			startServer()
		},
	}
}

func checkSQLCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "check-sql [filters]",
		Short: "check the syntax of all SQL queries",
		Run: func(cmd *cobra.Command, args []string) {
			conf := config.Load()
			db := mysql.MustConnect(conf.MySQL)

			if len(args) > 0 {
				filter := strings.Join(args, " ")
				dblib.CheckQueries(db, dblib.CheckOptions{
					Filter:       filter,
					EnablePrint:  true,
					DisableColor: false,
				})
			} else {
				dblib.CheckQueries(db, dblib.CheckOptions{
					DisableColor: false,
				})
			}
		},
	}
}

func startServer() {
	conf := config.Load()
	root := server.NewRoot(conf)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	grpcServer := grpc.NewServer(
		root.UnaryInterceptor(),
		root.StreamInterceptor(),
	)

	mux := runtime.NewServeMux(
		runtime.WithErrorHandler(errors.CustomHTTPError),
		runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{}),
	)

	ctx := context.Background()
	root.Register(ctx, grpcServer, mux, conf.Server.GRPC.String(), []grpc.DialOption{grpc.WithInsecure()})

	grpc_prometheus.EnableHandlingTimeHistogram()
	grpc_prometheus.Register(grpcServer)

	httpMux := http.NewServeMux()
	httpMux.Handle("/metrics", promhttp.Handler())
	httpMux.Handle("/", mux)

	httpServer := http.Server{
		Addr:    conf.Server.HTTP.String(),
		Handler: httpMux,
	}

	//--------------------------------
	// Run HTTP & gRPC servers
	//--------------------------------
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()

		err := httpServer.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
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

	root.Shutdown()
}
