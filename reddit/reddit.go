package reddit


import (
        "net/http"
        "encoding/json"
        "fmt"
	"errors"
)

type Item struct {
        Title string
        URL   string
        Comments int `json:"num_comments"`
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

                case 1:
                        com = " (1 Comment) "
                default:
                        com = fmt.Sprintf("(%d Comments)", i.Comments)
        }
        return fmt.Sprintf("%s\n%s\n%s", i.Title, com, i.URL)
}

func Get(reddit string) ([]Item, error) {
	url := fmt.Sprintf("http://reddit.com/r/%s.json", reddit)
        resp, err := http.Get(url)
        if err != nil {
                return nil, err
        }
        defer resp.Body.Close()
        if resp.StatusCode != http.StatusOK {
                return nil, errors.New(resp.Status)
        }
        r := new(Response)
        err = json.NewDecoder(resp.Body).Decode(r)
        if err != nil {
                return nil, err
        }
        items := make([]Item, len(r.Data.Children))
        for i, child := range r.Data.Children {
                items[i] = child.Data
        }
        return items, nil
}
