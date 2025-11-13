package main

import (
	"html/template"
	"os"
)

type Course struct {
	Name     string
	Price    int
	Platform string
}

func main() {
	course := Course{
		Name:     "Go Programming",
		Price:    100,
		Platform: "Online",
	}

	tmp := template.New("CourseTemplate")
	tmp, err := tmp.Parse(`
		<h1>{{.Name}}</h1>
		<p>Price: {{.Price}}</p>
		<p>Platform: {{.Platform}}</p>
	`)
	if err != nil {
		panic(err)
	}

	err = tmp.Execute(os.Stdout, course)
	if err != nil {
		panic(err)
	}
}
