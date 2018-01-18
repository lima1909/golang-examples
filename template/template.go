package main

import (
	"fmt"
	"html/template"
	"os"
	"strings"
	"time"
)

type person struct {
	Name      string
	BirthDate time.Time
}

func (p person) formatBirthDate() string {
	return fmt.Sprintf("%d-%d-%d",
		p.BirthDate.Day(),
		p.BirthDate.Month(),
		p.BirthDate.Year())
}
func formatBirthDate(p person) string {
	return p.BirthDate.Format(time.UnixDate)
}

func simpleWithFunction(p person) {
	funcMap := template.FuncMap{
		"format":  p.formatBirthDate,
		"format2": formatBirthDate,
		"upper":   strings.ToUpper,
	}
	tmpl, err := template.New("simple").Funcs(funcMap).Parse(`--> ` +
		`"{{ .Name }}" or {{ upper .Name }} ` +
		`is born at "{{ format }}" OR "{{ format2 . }}"`)
	if err != nil {
		panic(err)
	}

	if err := tmpl.Execute(os.Stdout, p); err != nil {
		panic(err)
	}
}

func main() {
	p := person{Name: "Me", BirthDate: time.Now()}
	simpleWithFunction(p)
}
