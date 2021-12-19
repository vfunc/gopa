package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "Welcome!")
	})

	srv := http.Server{Addr: ":8002", Handler: mux}

	g, ctx := errgroup.WithContext(context.Background())

	g.Go(func() error {
		return srv.ListenAndServe()
	})

	g.Go(func() error {
		go func() {
			log.Println(http.ListenAndServe(":6060", nil))
		}()
		return nil
	})

	g.Go(func() error {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-c:
			return srv.Shutdown(ctx)
		}
	})

	if err := g.Wait(); err != nil {
		fmt.Println(err)
	}
}
