package pingpong

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type PingResponse struct {
	Pings	int
}

type PingPongAPI struct {
	baseURL 	string
}

func New(url string) *PingPongAPI {
	return &PingPongAPI{baseURL: url}
}

func (p PingPongAPI) getUrl(path string) string {
	return fmt.Sprintf("%s%s", p.baseURL, path)
}

func (p PingPongAPI) Pings() (int, error) {
	resp, err := http.Get(p.getUrl("/pings"))
	if err != nil {
		return 0, err
	}
	if resp.StatusCode != 200 {
		return 0, errors.New("Ping-pong did not respond with 200")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}
	var result PingResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return 0, err
	}
	return result.Pings, nil
}
