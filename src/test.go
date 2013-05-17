/**
 * User: Baoxu
 * Date: 13-5-17
 * Time: 上午10:10
  */

package main

import (
	"html/template"
	"os"
)

type Person struct {
	UserName string
}

func test() {
	t := template.New("fieldname example")
	t, _ = t.Parse("hello {{.UserName}}!")
	p := Person{UserName: "Astaxie"}
	t.Execute(os.Stdout, p)
}

