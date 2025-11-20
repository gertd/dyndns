package cmd

import (
	"fmt"

	"github.com/gertd/dyndns/pkg/cc"
	"github.com/gertd/dyndns/pkg/version"
	"github.com/gertd/dyndns/pkg/x"
)

type CLI struct {
	Get     GetCmd     `cmd:"" help:"get current local IP address using https://ipify.org"`
	Set     SetCmd     `cmd:"" help:"configure namecheap.com dynamic DNS IP address"`
	Version VersionCmd `cmd:"" help:"build version information"`
}

type VersionCmd struct{}

func (cmd *VersionCmd) Run(c *cc.CommonCtx) error {
	fmt.Fprintf(c.OutWriter, "%s - %s\n",
		x.AppName,
		version.GetInfo().String(),
	)

	return nil
}
