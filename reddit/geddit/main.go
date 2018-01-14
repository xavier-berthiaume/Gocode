package main

import (
	"github.com/xavier-berthiaume/reddit"
)

func main() {
	reddit.PrintResponse(reddit.Get("bartenders"))
}
