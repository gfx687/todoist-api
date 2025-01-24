package todoistapi

type Project struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Task struct {
	Id      string   `json:"id"`
	Content string   `json:"content"`
	Labels  []string `json:"labels"`
}

type TaskCreate struct {
	Content     string   `json:"content"`
	Description string   `json:"description"`
	DueString   string   `json:"due_string"`
	Labels      []string `json:"labels"`
}
