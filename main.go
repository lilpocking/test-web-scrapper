package main

import (
	"home/pkg"
	"log"
)

func main() {
	if err := pkg.StartReadAndSearchInn(); err != nil {
		log.Panic(err)
	}
	//fmt.Println(pkg.FindInn("TNMK.ru"))
	//fmt.Println(pkg.FindInn("tnmk-96.ru"))

}
