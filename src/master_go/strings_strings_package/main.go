package main

import (
	"fmt"
	"strings"
)

func main() {

	p := fmt.Println
	f := fmt.Printf
	result := strings.Contains("I love Go Programming", "love")
	p(result)

	result = strings.ContainsAny("success", "xys")
	p(result)

	result = strings.ContainsRune("golang", 'g')
	p(result)

	n := strings.Count("cheese", "e")
	p(n)

	n = strings.Count("Five", "")
	p(n)

	p(strings.ToLower("GO PyTHON jaVA"))
	p(strings.ToUpper("GO PyTHON jaVA"))

	p("go" == "go")
	p("GO" == "go")

	p(strings.ToLower("GO") == strings.ToLower("go"))

	p(strings.EqualFold("GO", "go"))

	myStr := strings.Repeat("ab", 10)
	p(myStr)

	myStr = strings.Replace("192.168.0.1", ".", ":", 2) //sostituisce i primi due . con :
	p(myStr)

	myStr = strings.Replace("192.168.0.1", ".", ":", -1) //sostituisce tutti i . con :
	p(myStr)

	myStr = strings.ReplaceAll("192.168.0.1", ".", ":")
	p(myStr)

	s := strings.Split("a,b,c", ",")
	f("%T\n", s)
	f("%#v\n", s)

	s = strings.Split("Go for Go!", "")
	f("%#v\n", s) //[]string{"G", "o", " ", "f", "o", "r", " ", "G", "o", "!"}

	s = []string{"I", "learn", "Golang"}
	myStr = strings.Join(s, "-")

	myStr = "Orange Green \n Blue Yellow"
	fields := strings.Fields(myStr)
	f("%#v\n", fields)

	s1 := strings.TrimSpace("\t   Goodby Windows, Welcome Linux!   \n")
	f("%q\n", s1)

	s2 := strings.Trim("...Hello Gopher!!!!?", ".!?")
	f(s2)
}
