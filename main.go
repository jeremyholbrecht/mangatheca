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
	fmt.Println(Manga{"foo", "bar"})
	fmt.Println(gabaiVol1.name)
	fmt.Println(gabaiVol1.author)

	crowsVolume1 := Manga{"Crows", "Hiroshi Takahashi"}
	crowsVolume2 := Manga{"Crows", "Hiroshi Takahashi"}
	fmt.Println(crowsVolume1)
	fmt.Println(crowsVolume1.author)
	fmt.Println(crowsVolume1.name)
	fmt.Println(crowsVolume2.name)

}
