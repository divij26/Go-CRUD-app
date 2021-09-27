package main

import (
	"fmt"
)

type album struct {
	ID    string `json: "id"`
	Title string `json: "title"`
	Price string `json: "price"`
}

var albums = []album{
	{ID: "1", Title: "All of me", Artist: "John Legend", Price: 1000}
	{ID: "2", Title: "Lover", Artist: "Taylor Swift", Price: 900}
}

func main() {
	fmt.Printf("Test \n")
}
