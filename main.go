package main

import (
	"context"
	"fmt"
	"log"
	"net"
	config "octane/configure"
	"octane/controller"
	lpb "octane/grpc/lotto"
	"octane/handler"
	"octane/record"
	"octane/repository"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"go.uber.org/fx"
	"google.golang.org/grpc"
)

// NewServer ...
func NewServer(lc fx.Lifecycle, lotto handler.Lotto) *grpc.Server {
	s := grpc.NewServer()
	lpb.RegisterLottoServer(s, lotto)
	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			log.Println("Stopping GRPC server.")
			s.Stop()
			return nil
		},
	})
	return s
}

func main() {

	app := fx.New(
		fx.Provide(
			config.New,
			repository.NewLotto,
			handler.NewLotto,
			controller.NewLotto,
			NewServer,
		),
		fx.Invoke(func(s *grpc.Server, c *record.Config) error {
			port := fmt.Sprintf(":%s", c.GRPC.Port)
			lis, err := net.Listen("tcp", port)
			if err != nil {
				log.Fatalf("failed to listen: %v", err)
			} else {
				log.Printf("Listening on %s", port)
			}
			log.Println("Start GRPC server.")
			go s.Serve(lis)
			return nil
		}),
	)
	startCtx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	if err := app.Start(startCtx); err != nil {
		log.Fatal(err)
	}

	<-app.Done()
}
