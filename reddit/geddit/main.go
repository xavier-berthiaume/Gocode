package main

import (
	"github.com/xavier-berthiaume/reddit"
	"fmt"
)

func main() {
	subreddit := reddit.AcceptInput()
	reddit.PrintResponse(reddit.Get(subreddit))
}
