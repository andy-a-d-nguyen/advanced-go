package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"

	"golang.org/x/sync/errgroup"
)

type page struct {
	url      string
	Title    string
	Children []*page `json:",omitempty"`
}

func newPage(url string) *page {
	return &page{
		url:      url,
		Children: make([]*page, 0),
	}
}

var (
	titleRegexp = regexp.MustCompile(`<title.*>(.+?)</title>`)
	linkRegexp  = regexp.MustCompile(`<a.*href=\"(http.+?)\".*</a>`)
)

func main() {
	p := newPage("https://go.dev")
	g := new(errgroup.Group)
	g.Go(createGoFunc(p, g, 1))

	err := g.Wait()
	if err != nil {
		log.Fatal(err)
	}
	d, err := json.Marshal(p)
	f, err := os.OpenFile("./result.json", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0200)
	if err != nil {
		log.Printf("Couldn't create result file: %v", err)
		return
	}
	defer f.Close()
	f.Write(d)

}

const maxDepth = 3

func createGoFunc(p *page, g *errgroup.Group, currentDepth int) func() error {
	return func() error {
		return visitSite(p, g, currentDepth)
	}
}

func visitSite(p *page, g *errgroup.Group, currentDepth int) error {
	title, children, err := getTitleAndChildren(p.url)
	if err != nil {
		return err
	}

	p.Title = title

	if currentDepth < maxDepth {
		p.Children = children
		for _, c := range children {
			g.Go(createGoFunc(c, g, currentDepth+1))
		}
	} else {
		p.Children = nil
	}
	return nil
}

func getTitleAndChildren(url string) (title string, children []*page, err error) {
	log.Printf("Reading URL: '%v'\n", url)
	r, err := http.Get(url)
	if err != nil {
		return "", nil, err
	}

	body, err := io.ReadAll(r.Body)
	r.Body.Close()
	if err != nil {
		return "", nil, err
	}

	titleMatches := titleRegexp.FindSubmatch(body)
	if len(titleMatches) >= 1 {
		title = string(titleMatches[1])
	}
	children = make([]*page, 0)

loop:
	for _, l := range linkRegexp.FindAllSubmatch(body, -1) {
		if len(l) < 1 {
			continue
		}
		url := string(l[1])
		for _, c := range children {
			if c.url == url {
				continue loop
			}
		}
		children = append(children, newPage(url))
	}

	return title, children, nil
}
