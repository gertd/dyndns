package cmd

import (
	"fmt"
	"strings"

	"github.com/gertd/dyndns/pkg/cc"
	"github.com/gertd/dyndns/pkg/ipify"
	"github.com/gertd/dyndns/pkg/namecheap"
)

type SetCmd struct {
	Passwd string `name:"passwd" required:"" env:"DYNDNS_PASSWD" help:"namecheap.com Dynamic DNS Password"`
	Host   string `name:"host" required:"" env:"DYNDNS_HOST" default:"@" help:"host name"`
	Domain string `name:"domain" required:"" env:"DYNDNS_DOMAIN" help:"domain name ex. yourdomain.tld"`
	IPAddr string `name:"ip-addr" optional:"" env:"DYNDNS_IP_ADDR" help:"IP address to be associated with the dynamic DNS entry, if not specified the current IP used will be used to set for the domain"`
	Info   bool   `name:"info" optional:"" help:"display settings, do not execute"`
}

func (cmd SetCmd) String() string {
	s := fmt.Sprintf("%s = %s\n", "passwd ", mask(cmd.Passwd))
	s += fmt.Sprintf("%s = %s\n", "host   ", cmd.Host)
	s += fmt.Sprintf("%s = %s\n", "domain ", cmd.Domain)
	s += fmt.Sprintf("%s = %s\n", "ip-addr", cmd.IPAddr)
	return s
}

func mask(s string) string {
	const maxExposure = 3
	l := len(s)
	if l <= (maxExposure*2)+1 {
		return strings.Repeat("*", l)
	}
	return s[0:maxExposure] + strings.Repeat("*", l-(maxExposure*2)) + s[l-maxExposure:l]
}

func (cmd *SetCmd) Run(c *cc.CommonCtx) error {
	if cmd.Info {
		fmt.Fprintf(c.OutWriter, "%s = %s\n", "passwd ", mask(cmd.Passwd))
		fmt.Fprintf(c.OutWriter, "%s = %s\n", "host   ", cmd.Host)
		fmt.Fprintf(c.OutWriter, "%s = %s\n", "domain ", cmd.Domain)
		fmt.Fprintf(c.OutWriter, "%s = %s\n", "ip-addr", cmd.IPAddr)
		if cmd.IPAddr == "" {
			if curIP, err := ipify.GetIPAddress(c); err == nil {
				fmt.Fprintf(c.OutWriter, "%s = %s\n", "current", curIP)
			}
		}
		return nil
	}

	ddc := namecheap.NewDynamicDNSConfigurator(cmd.Passwd, cmd.Host, cmd.Domain, cmd.IPAddr)
	if err := ddc.Set(c); err != nil {
		return err
	}

	return nil
}
