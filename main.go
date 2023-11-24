package main

import (
	"fmt"

	"github.com/MJurayev/gitlab-client.git/projects"
)

func main() {
	response := projects.GetAll()
	fmt.Println(response)
}