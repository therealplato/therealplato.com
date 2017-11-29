package main

import (
	"fmt"
	"image/png"
	"log"
	"os"
)

func main() {
	f, err := os.Open("./source.png")
	if err != nil {
		log.Fatal(err)
	}
	i, err := png.Decode(f)
	if err != nil {
		log.Fatal(err)
	}
	results := slice(i)
	os.Mkdir("./results", os.ModeDir)
	for j, r := range results {
		f, err := os.Create(fmt.Sprintf("./results/comment-%03d.png", j+1))
		if err != nil {
			log.Fatal(err)
		}
		png.Encode(f, r)
	}
	fmt.Println(len(results))
}
