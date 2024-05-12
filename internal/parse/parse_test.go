package parse

import (
	"net/http"
	"testing"
)

func TestParseDoc(t *testing.T) {
	// os.Setenv("HTTPS_PROXY", "socks5://127.0.0.1:1080")
	res, err := http.Get("https://core.telegram.org/bots/api")
	if err != nil {
		return
	}
	defer res.Body.Close()
	parse(res.Body)
}
