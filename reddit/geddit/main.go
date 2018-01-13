package main

import (
	"log"
	"fmt"
	"github.com/xavier-berthiaume/reddit"
)

func main() {
	items, err := reddit.Get("bartenders")
	if err != nil {
		log.Fatal(err)
	}
	for _, item := range items {
		fmt.Println(item.Title)
	}	
}
