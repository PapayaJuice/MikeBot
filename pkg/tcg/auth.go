package tcg

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

const (
	tcgpURI       = "https://api.tcgplayer.com/token"
	bearerPayload = "grant_type=client_credentials&client_id=%s&client_secret=%s"
)

var (
	currToken = &BearerToken{}

	pubKey = flag.String("pub-key", "", "Public Key for TCG Player API")
	priKey = flag.String("pri-key", "", "Private Key for TCG Player API")
)

// BearerToken represents a valid Bearer Token response from TCG Player API
type BearerToken struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
	UserName    string `json:"userName"`
	Issued      string `json:".issued"`
	Expires     string `json:".expires"`
}

// TokenRefresh attempts to request a new token if we need one
func TokenRefresh() {
	for {
		t, err := requestToken()
		if err != nil || currToken == nil {
			log.Errorf("error requesting token: %v", err)
			time.Sleep(10 * time.Minute)
			continue
		}
		currToken = t

		// Set time to request a little early so we don't lose connection
		checkTime := time.Duration(currToken.ExpiresIn - 3600)
		time.Sleep(checkTime)
	}
}

func requestToken() (*BearerToken, error) {
	payload := fmt.Sprintf(bearerPayload, *pubKey, *priKey)
	req, err := http.NewRequest(http.MethodPost, tcgpURI, bytes.NewBuffer([]byte(payload)))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var bt BearerToken
	err = json.Unmarshal(b, &bt)
	return &bt, err
}
