package reddit


import (
        "net/http"
        "encoding/json"
        "fmt"
	"errors"
	"log"
)

type Item struct {
        Title string
        URL   string
        Comments int `json:"num_comments"`
	Score int
	Subreddit string
}


type Response struct {
        Data struct {
                Children []struct {
                        Data Item
                }
        }
}

func (i Item) String() string {
        com := ""
        switch i.Comments {
                case 0:
			com = " (No Comments Yet!) "
                case 1:
                        com = " (1 Comment) "
                default:
                        com = fmt.Sprintf("(%d Comments)", i.Comments)
        }
        return fmt.Sprintf("%s\n%s\t%d votes\tr/%s\n%s\n", i.Title, com, i.Score, i.Subreddit, i.URL)
}

func AcceptInput() {


}

func PrintResponse(items []Item, err error) {
	if err != nil {
		log.Fatal(err)
	}
	for _, item := range items {
		fmt.Println(item)
	}
}

func Get(reddit string) ([]Item, error) {
	client := &http.Client{}
	url := fmt.Sprintf("http://reddit.com/r/%s.json", reddit)
	request, err := http.NewRequest("GET", url, nil)
	request.Header.Set("User-Agent", "Subreddit Crawler/V1.0")
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		return nil, errors.New(response.Status)
	}
	r := new(Response)
	err = json.NewDecoder(response.Body).Decode(r)
	items := make([]Item, len(r.Data.Children))
	for i, child := range r.Data.Children {
		items[i] = child.Data
	}
	return items, nil
}
