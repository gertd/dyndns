package ipify

import (
	"io"
	"net/http"

	"github.com/gertd/dyndns/pkg/cc"
)

const (
	ipifyURL string = "https://api.ipify.org"
)

func GetIPAddress(c *cc.CommonCtx) (string, error) {
	req, err := http.NewRequestWithContext(c.Context, http.MethodGet, ipifyURL, http.NoBody)
	if err != nil {
		return "", err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}

	if res.StatusCode != http.StatusOK {
		return "", err
	}
	defer res.Body.Close()

	buf, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	return string(buf), nil
}
