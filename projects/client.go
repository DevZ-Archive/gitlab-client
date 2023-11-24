package projects

import (
	"io"
	"net/http"
)

type project struct {
}
// glpat-Eq6raxTjumH6QhYwYzaq
func GetAll() string {
	resp, err := http.Get("https://gitlab.com/api/v4/projects")

	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body,err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	return string(body)
}
