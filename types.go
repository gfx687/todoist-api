package todoistapi

// TODO: auth

func GetProjects() ([]ProjectSimple, error) {
	return []ProjectSimple{{Id: "123123", Name: "Inbox"}}, nil
}

type ProjectSimple struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}
