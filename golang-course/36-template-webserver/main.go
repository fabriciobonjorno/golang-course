package main

import (
	"html/template"
	"net/http"
)

type Course struct {
	Name     string
	Price    int
	Platform string
}

type Courses []Course

func main() {
	courses := Courses{
		{
			Name:     "Go Programming",
			Price:    299,
			Platform: "Udemy",
		},
		{
			Name:     "Python Programming",
			Price:    199,
			Platform: "Coursera",
		},
		{
			Name:     "JavaScript Programming",
			Price:    149,
			Platform: "edX",
		},
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t := template.Must(template.ParseFiles("template.html"))
		err := t.Execute(w, courses)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	http.ListenAndServe(":8282", nil)

	// t := template.Must(template.ParseFiles("template.html"))

	// err := t.Execute(os.Stdout, courses)
	// if err != nil {
	// 	panic(err)
	// }

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
