package main

import "flag"

func main() {
	name := flag.String("name", "anonymous", "your name")
	flag.Parse()

	println("Hello", *name)
}
