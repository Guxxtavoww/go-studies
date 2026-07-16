package main

import (
	"os"
	"text/template"
)

type Course struct {
	Name     string
	Workload int
}


// Must unwraps a call to a function returning (*Template, error) and panics if the error is non-nil. It is intended for use in variable initializations.
func mustTemplate() *template.Template {
	template := template.Must(template.New("CourseTemplate").Parse("Course Name: {{.Name}}, Workload: {{.Workload}} hours"))

	return template
}

func main() {
	course := &Course{
		Name:     "Go",
		Workload: 40,
	}

	template := template.New("CourseTemplate")
	// Ponto é usado para acessar os campos da struct
	template, _ = template.Parse("Course Name: {{.Name}}, Workload: {{.Workload}} hours")

	err := template.Execute(os.Stdout, course)

	if err != nil {
		panic(err)
	}
}
