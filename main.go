package main

import "fmt"

type Manga struct {
	name   string
	author string
}

func main() {
	gabaiVol1 := new(Manga)
	gabaiVol1.author = "Yoshichi Shimada"
	gabaiVol1.name = "Gabai Volume 1"
	fmt.Printf("%+v\n", gabaiVol1)
}
