package main

import (
	"context"
	"time"

	"github.com/alecthomas/kong"
	"github.com/gertd/dyndns/pkg/cc"
	"github.com/gertd/dyndns/pkg/cmd"
	"github.com/gertd/dyndns/pkg/x"
	_ "github.com/joho/godotenv/autoload"
)

const commandTimeOut = time.Second * 5

func main() {
	ct1, cancel := context.WithTimeout(context.Background(), commandTimeOut)
	defer cancel()

	c := cc.New(ct1)

	cli := cmd.CLI{}
	ctx := kong.Parse(&cli,
		kong.Name(x.AppName),
		kong.Description(x.AppDescription),
		kong.UsageOnError(),
		kong.ConfigureHelp(kong.HelpOptions{
			NoAppSummary:        false,
			Summary:             true,
			Compact:             true,
			Tree:                false,
			FlagsLast:           true,
			Indenter:            kong.SpaceIndenter,
			NoExpandSubcommands: false,
		}),
		kong.Bind(&cli),
		kong.Bind(c),
	)

	err := ctx.Run(c)
	ctx.FatalIfErrorf(err)
}
