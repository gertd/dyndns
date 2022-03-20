package namecheap

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"

	"github.com/gertd/dyndns/pkg/cc"
	"github.com/pkg/errors"
	"golang.org/x/net/html/charset"
)

type DynamicDNSConfigurator struct {
	Host   string
	Domain string
	Passwd string
	IPAddr string
}

type interfaceResponse struct {
	Command       string     `xml:"Command"`
	Language      string     `xml:"Language"`
	IP            string     `xml:"IP"`
	ErrCount      int        `xml:"ErrCount"`
	Errors        []xerror   `xml:"errors,omitempty"`
	ResponseCount int        `xml:"ResponseCount"`
	Responses     []response `xml:"responses>response,omitempty"`
	Done          bool       `xml:"Done"`
}

type xerror struct {
	ErrorString string `xml:"Err1"`
}

type response struct {
	ResponseNumber int    `xml:"ResponseNumber"`
	ResponseString string `xml:"ResponseString"`
}

const (
	dynamicdnsURL string = "https://dynamicdns.park-your-domain.com/update?host=%s&domain=%s&password=%s&ip=%s"
)

func NewDynamicDNSConfigurator(passwd, host, domain, ipAddr string) *DynamicDNSConfigurator {
	return &DynamicDNSConfigurator{
		Passwd: passwd,
		Host:   host,
		Domain: domain,
		IPAddr: ipAddr,
	}
}

func (ddc *DynamicDNSConfigurator) Set(c *cc.CommonCtx) error {
	url := fmt.Sprintf(dynamicdnsURL,
		ddc.Host,
		ddc.Domain,
		ddc.Passwd,
		ddc.IPAddr,
	)

	req, err := http.NewRequestWithContext(c.Context, http.MethodGet, url, http.NoBody)
	if err != nil {
		return err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	buf, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	r, err := charset.NewReader(bytes.NewReader(buf), "utf-16")
	if err != nil {
		return err
	}

	var ir interfaceResponse
	dec := xml.NewDecoder(r)
	dec.CharsetReader = identReader
	if err := dec.Decode(&ir); err != nil {
		return err
	}

	if ir.ErrCount == 0 && ir.ResponseCount == 0 && len(ir.IP) > 0 {
		return nil
	}

	return errors.Errorf("failed to set Dynamic DNS entry %s - %s", ddc.IPAddr, ir.Errors[0].ErrorString)
}

func identReader(encoding string, input io.Reader) (io.Reader, error) {
	return input, nil
}
