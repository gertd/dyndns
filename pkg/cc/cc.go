package cc

import (
	"context"
	"io"
	"log"
	"os"
)

type CommonCtx struct {
	Context   context.Context
	OutWriter io.Writer
	ErrWriter io.Writer
}

func New(ctx context.Context) *CommonCtx {
	log.SetOutput(io.Discard)
	log.SetPrefix("")
	log.SetFlags(log.LstdFlags)

	return &CommonCtx{
		Context:   ctx,
		OutWriter: os.Stdout,
		ErrWriter: os.Stderr,
	}
}
