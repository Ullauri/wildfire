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

type JokesClient interface {
	GetJoke(ctx context.Context, firstName *string, lastName *string) (*schema.GetJokeResponse, error)
}

type jokesClient struct {
	httpClient *http.Client
}

func NewJokesClient(httpClient *http.Client) JokesClient {
	return &jokesClient{httpClient: httpClient}
}

func (c jokesClient) BaseUrl() string {
	return common.Config().JokesClient.BaseUrl
}

func (c *jokesClient) GetJoke(ctx context.Context, firstName *string, lastName *string) (*schema.GetJokeResponse, error) {
	// TODO: limitTo=nerdy hardcoded for now
	url := fmt.Sprintf("%sjoke?limitTo=nerdy&firstName=%s&lastName=%s", c.BaseUrl(), *firstName, *lastName)
	req, _ := http.NewRequest("GET", url, nil)

	log.Info().Str("func", "GetJoke").Str("request_url", url).Msg("sending request")

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

	log.Info().Str("func", "GetJoke").Int("status_code", resp.StatusCode).Str("response_body", string(body)).Msg("response received")

	if resp.StatusCode != 200 {
		switch resp.StatusCode {
		case 404:
			return nil, fmt.Errorf("joke not found")
		default:
			return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
		}
	}

	var getJokeResp schema.GetJokeResponse
	err = Jsoniter.Unmarshal(body, &getJokeResp)
	if err != nil {
		return nil, err
	}

	log.Info().Str("func", "GetJoke").Msg("result returned")

	return &getJokeResp, nil
}
