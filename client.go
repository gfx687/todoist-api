package todoistapi

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

type TodoistClient struct {
	token      string
	httpClient *http.Client
}

// Returns a pointer to a new [TodoistClient] or an error if provided token is not valid
func NewClient(token string) (*TodoistClient, error) {
	if token == "" {
		return nil, errors.New("Todoist authentication token cannot be empty")
	}

	httpClient := &http.Client{Timeout: 10 * time.Second}

	err := verifyToken(httpClient, token)
	if err != nil {
		return nil, fmt.Errorf("Todoist token is invalid: %w", err)
	}

	return &TodoistClient{token: token, httpClient: httpClient}, nil
}

func verifyToken(httpClient *http.Client, token string) error {
	req, err := http.NewRequest("GET", "https://api.todoist.com/rest/v2/projects", nil)
	if err != nil {
		return fmt.Errorf("error while creating a new HTTP request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+token)

	res, err := httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("error while making an HTTP request: %w", err)
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("error while reading HTTP response body: %w", err)
	}

	if res.StatusCode == 401 || res.StatusCode == 403 {
		return fmt.Errorf("Todoist token is invalid. Status: %d, Body: %v", res.StatusCode, string(body))
	}

	if res.StatusCode != 200 {
		return fmt.Errorf("response status code is not 200. Status: %d, Body: %v", res.StatusCode, string(body))
	}

	return nil
}
