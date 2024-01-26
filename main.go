package main

import "fmt"

type Manga struct {
	name   string
	author string
}

func main() {
	gabaivol1 := Manga{name: "foo", author: "bar"}
	fmt.Println(gabaivol1)
}
