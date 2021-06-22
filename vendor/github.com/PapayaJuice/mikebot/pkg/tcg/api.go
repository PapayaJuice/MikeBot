package tcg

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

const (
	mtgSearchURI = "https://api.tcgplayer.com/v1.39.0/catalog/products?categoryId=1&productName=%s"
	mtgReqURI    = "https://api.tcgplayer.com/v1.39.0/catalog/products/%d?getExtendedFields=true"
)

var (
	c = http.Client{
		Timeout: 30 * time.Second,
	}
)

// Result ...
type Result struct {
	Success bool            `json:"success"`
	Errors  []string        `json:"errors"`
	Results []ProductResult `json:"results"`
}

// ProductResult ...
type ProductResult struct {
	ProductID int    `json:"productId"`
	Name      string `json:"name"`
	CleanName string `json:"cleanName"`

	ImageURL string `json:"imageUrl"`
	URL      string `json:"url"`

	ExtendedData []ResultExtended `json:"extendedData"`
}

// ResultExtended ...
type ResultExtended struct {
	Name        string `json:"name"`
	DisplayName string `json:"displayName"`
	Value       string `json:"value"`
}

func queryAndParse(uri string) (*ProductResult, error) {
	req, err := http.NewRequest(http.MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("bearer %s", currToken.AccessToken))

	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var sr Result
	err = json.Unmarshal(b, &sr)
	if err != nil {
		return nil, err
	}

	if !sr.Success || len(sr.Results) <= 0 {
		if len(sr.Errors) > 0 {
			return nil, errors.New(sr.Errors[0])
		}
		return nil, errors.New("unknown api error")
	}

	return &sr.Results[0], nil
}

func searchCard(card string) (int, error) {
	uCard := url.QueryEscape(card)
	uri := fmt.Sprintf(mtgSearchURI, uCard)

	res, err := queryAndParse(uri)
	if err != nil {
		return 0, err
	}
	return res.ProductID, nil
}

func requestCard(cardID int) (*ProductResult, error) {
	uri := fmt.Sprintf(mtgReqURI, cardID)

	res, err := queryAndParse(uri)
	return res, err
}
