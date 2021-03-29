package main

type any interface{}

func main() {
	var a any
	b := a.(int)
}