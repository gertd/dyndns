package cmd

import (
	"fmt"

	"github.com/gertd/dyndns/pkg/cc"
	"github.com/gertd/dyndns/pkg/ipify"
)

type GetCmd struct {
	Short bool `name:"short" short:"s" help:"short output"`
}

func (cmd *GetCmd) Run(c *cc.CommonCtx) error {
	ipAddr, err := ipify.GetIPAddress(c)
	if err != nil {
		return err
	}

	if cmd.Short {
		fmt.Fprintln(c.OutWriter, ipAddr)
		return nil
	}

	fmt.Fprintf(c.OutWriter, "current IP address = [%s]\n", ipAddr)
	return nil
}
