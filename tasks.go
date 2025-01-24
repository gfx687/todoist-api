package todoistapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetTasksByProject(projectId string) ([]Task, error) {
	client := getHttpClient()

	req, err := http.NewRequest("GET", "https://api.todoist.com/rest/v2/tasks?project_id="+projectId, nil)
	if err != nil {
		return nil, fmt.Errorf("error while creating a new HTTP request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+todoistToken)

	res, err := client.Do(req)
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

	tasks := []Task{}
	err = json.Unmarshal(body, &tasks)
	if err != nil {
		return nil, fmt.Errorf("error while parsing HTTP response body: %w", err)
	}

	return tasks, nil
}

func GetTasksByLabel(label string) ([]Task, error) {
	req, err := http.NewRequest("GET", "https://api.todoist.com/rest/v2/tasks?label="+label, nil)
	if err != nil {
		return nil, fmt.Errorf("error while creating a new HTTP request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+todoistToken)

	client := getHttpClient()

	res, err := client.Do(req)
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

	tasks := []Task{}
	err = json.Unmarshal(body, &tasks)
	if err != nil {
		return nil, fmt.Errorf("error while unmarshalling HTTP response body: %w", err)
	}

	return tasks, nil
}

func CreateTask(content string, description string, dueString string) error {
	task := TaskCreate{
		Content:     content,
		Description: description,
		DueString:   dueString,
		Labels:      []string{"automation"},
	}

	reqBody, err := json.Marshal(&task)
	if err != nil {
		return fmt.Errorf("error while json marshalling: %w", err)
	}

	reqBodyBytes := bytes.NewBuffer(reqBody)

	req, err := http.NewRequest("POST", "https://api.todoist.com/rest/v2/tasks", reqBodyBytes)
	if err != nil {
		return fmt.Errorf("error while creating a new HTTP request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+todoistToken)
	req.Header.Set("Content-Type", "application/json")

	client := getHttpClient()

	res, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error while making an HTTP request: %w", err)
	}
	defer res.Body.Close()

	resBodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("error while reading HTTP response body: %w", err)
	}

	if res.StatusCode != 200 {
		return fmt.Errorf("response status code is not 200. Status: %d, Body: %v", res.StatusCode, string(resBodyBytes))
	}

	return nil
}
