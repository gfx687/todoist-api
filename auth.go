package todoistapi

import (
	"errors"
	"fmt"
	"io"
	"net/http"
)

var todoistToken string

func Auth(token string) error {
	if token == "" {
		return errors.New("Todoist authentication token cannot be empty")
	}

	err := verifyToken(token)
	if err != nil {
		return fmt.Errorf("Todoist token is invalid: %w", err)
	}

	todoistToken = token
	return nil
}

func verifyToken(token string) error {
	client := getHttpClient()

	req, err := http.NewRequest("GET", "https://api.todoist.com/rest/v2/projects", nil)
	if err != nil {
		return fmt.Errorf("error while creating a new HTTP request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+token)

	res, err := client.Do(req)
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
