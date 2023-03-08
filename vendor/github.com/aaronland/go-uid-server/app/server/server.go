package server

import (
	"context"
	"log"
	"fmt"
	"flag"
	"net/http"

	aa_log "github.com/aaronland/go-log/v2"	
	"github.com/aaronland/go-http-server"
	"github.com/aaronland/go-uid"
	"github.com/aaronland/go-uid-server/http/api"			
	"github.com/sfomuseum/go-http-auth"
	"github.com/sfomuseum/go-flags/flagset"
)

func Run(ctx context.Context, logger *log.Logger) error {
	fs := DefaultFlagSet()
	return RunWithFlagSet(ctx, fs, logger)
}

func RunWithFlagSet(ctx context.Context, fs *flag.FlagSet, logger *log.Logger) error {

	flagset.Parse(fs)

	pr, err := uid.NewProvider(ctx, provider_uri)

	if err != nil {
		return fmt.Errorf("Failed to create proxy provider, %w", err)
	}

	pr.SetLogger(ctx, logger)

	authenticator, err := auth.NewAuthenticator(ctx, authenticator_uri)

	if err != nil {
		return fmt.Errorf("Failed to create authenticator, %w", err)
	}
	
	mux := http.NewServeMux()

	uid_opts := &api.UIDHandlerOptions{
		Provider: pr,
		Authenticator: authenticator,
	}
	
	uid_handler, err := api.UIDHandler(uid_opts)

	if err != nil {
		return fmt.Errorf("Failed to create UID handler, %w", err)
	}

	uid_handler = authenticator.WrapHandler(uid_handler)
	
	mux.Handle("/", uid_handler)

	s, err := server.NewServer(ctx, server_uri)

	if err != nil {
		return fmt.Errorf("Failed to create server, %w", err)
	}

	aa_log.Info(logger, "Listening for requests at %s\n", s.Address())

	err = s.ListenAndServe(ctx, mux)

	if err != nil {
		return fmt.Errorf("Failed to serve requests, %w", err)
	}

	return nil
}
