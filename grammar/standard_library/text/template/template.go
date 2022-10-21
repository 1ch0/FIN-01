package main

import (
	"os"
	"text/template"
	"time"
)

type User struct {
	Username, Password string
	RegTime            time.Time
}

func ShowTime(t time.Time, format string) string {
	return t.Format(format)
}

func Time(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

func main() {
	u := User{"user01", "xxxx", time.Now()}
	t, err := template.New("text").Funcs(template.FuncMap{"time": Time}).
		Parse(`{{.Password}} {{.Password}} {{.Password}}  {{.Username}}  {{.RegTime.Format "2006-01-02 15:04:05"}}

{{.Username}}|{{.Password}}|{{time .RegTime}}
`)
	if err != nil {
		panic(err)
	}
	t.Execute(os.Stdout, u)
}
