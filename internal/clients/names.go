package clients

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"wildfire/internal/clients/schema"
	"wildfire/internal/common"

	"github.com/rs/zerolog/log"
)

type NamesClient interface {
	GetRandomName(ctx context.Context) (*schema.GetRandomNameResponse, error)
}

type namesClient struct {
	httpClient *http.Client
}

func NewNamesClient(httpClient *http.Client) NamesClient {
	return &namesClient{httpClient: httpClient}
}

func (c namesClient) BaseUrl() string {
	return common.Config().NamesClient.BaseUrl
}

func (c *namesClient) GetRandomName(ctx context.Context) (*schema.GetRandomNameResponse, error) {

	url := c.BaseUrl() + "v0/"
	req, _ := http.NewRequest("GET", url, nil)

	log.Info().Str("func", "GetRandomName").Str("request_url", url).Msg("sending request")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var body []byte
	if resp.Body != nil {
		body, err = io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
	}

	log.Info().Str("func", "GetRandomName").Int("status_code", resp.StatusCode).Str("response_body", string(body)).Msg("response received")

	if resp.StatusCode != 200 {
		switch resp.StatusCode {
		default:
			return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
		}
	}

	var getNameResp schema.GetRandomNameResponse
	err = Jsoniter.Unmarshal(body, &getNameResp)
	if err != nil {
		return nil, err
	}

	log.Info().Str("func", "GetRandomName").Msg("result returned")

	return &getNameResp, nil
}
