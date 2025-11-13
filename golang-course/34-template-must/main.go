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
	t := template.Must(template.New("CourseTemplate").Parse(`
		<h1>{{.Name}}</h1>
		<p>Price: {{.Price}}</p>
		<p>Platform: {{.Platform}}</p>
	`))

	err := t.Execute(os.Stdout, course)
	if err != nil {
		panic(err)
	}

	// Alternatively, without using template.Must

	// tmp := template.New("CourseTemplate")
	// tmp, err := tmp.Parse(`
	// 	<h1>{{.Name}}</h1>
	// 	<p>Price: {{.Price}}</p>
	// 	<p>Platform: {{.Platform}}</p>
	// `)
	// if err != nil {
	// 	panic(err)
	// }

	// err = tmp.Execute(os.Stdout, course)
	// if err != nil {
	// 	panic(err)
	// }
}
