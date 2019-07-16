package main

import (
	"html/template"
	"os"
)

type Person struct {
	Name string
}

func main() {
	t := template.New("string template")
	// {{.}} 意思是this, {{.Name}} this.Name， Name必须是可导出的
	t, _ = t.Parse("hello {{.Name}}")
	p := Person{"ben"}
	t.Execute(os.Stdout, p)
}
