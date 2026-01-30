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

type Todo struct {
	ID		string	`json:"id"`
	IsDone	bool	`json:"is_done"`
	Content	string	`json:"content"`
}

func New(url string) *TodosAPI {
	return &TodosAPI{baseURL: url}
}

func (t TodosAPI) getUrl(path string) string {
	return fmt.Sprintf("%s%s", t.baseURL, path)
}

func (t TodosAPI) GetAll() ([]Todo, error) {
	resp, err := http.Get(t.getUrl("/todos"))
	if err != nil {
		return []Todo{}, err
	}
	if resp.StatusCode != 200 {
		return []Todo{}, errors.New("Server did not respond with 200")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []Todo{}, err
	}
	var result []Todo
	if err := json.Unmarshal(body, &result); err != nil {
		return []Todo{}, err
	}
	return result, nil
}
