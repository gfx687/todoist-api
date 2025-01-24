package todoistapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (todoistClient *TodoistClient) GetProjectList() ([]Project, error) {
	req, err := http.NewRequest("GET", "https://api.todoist.com/rest/v2/projects", nil)
	if err != nil {
		return nil, fmt.Errorf("error while creating a new HTTP request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+todoistClient.token)

	res, err := todoistClient.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error while making an HTTP request: %w", err)
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("error while reading HTTP response body: %w", err)
	}

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("response status code is not 200. Status: %d, Body: %v", res.StatusCode, string(body))
	}

	projects := []Project{}
	err = json.Unmarshal(body, &projects)
	if err != nil {
		return nil, fmt.Errorf("error while parsing HTTP response body: %w", err)
	}

	return projects, nil
}
