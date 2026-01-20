package todos

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type TodosAPI struct {
	baseURL	string
}

func New(url string) *TodosAPI {
	return &TodosAPI{baseURL: url}
}

func (t TodosAPI) getUrl(path string) string {
	return fmt.Sprintf("%s%s", t.baseURL, path)
}

func (t TodosAPI) GetAll() ([]string, error) {
	resp, err := http.Get(t.getUrl("/todos"))
	if err != nil {
		return []string{}, err
	}
	if resp.StatusCode != 200 {
		return []string{}, errors.New("Server did not respond with 200")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []string{}, err
	}
	var result []string
	if err := json.Unmarshal(body, &result); err != nil {
		return []string{}, err
	}
	return result, nil
}
