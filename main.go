package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"golang.org/x/net/html"
)

var dir string

type img struct {
	p string
	t time.Time
}

func findLink(n *html.Node) string {
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		fmt.Println(c.Data)
		if c.Type == html.ElementNode && (c.Data == "video" || c.Data == "a") {
			for _, a := range c.Attr {
				if (c.Data == "video" && a.Key == "src") || (c.Data == "a" && a.Key == "href") {
					return a.Val
				}
			}
		}
	}
	return ""
}

func findTime(n *html.Node) time.Time {
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if c.Type == html.ElementNode && c.Data == "a" {
			for tc := c.FirstChild; tc != nil; tc = tc.NextSibling {
				if tc.Type == html.TextNode {
					t, err := time.Parse("Jan 2, 2006, 3:04 PM", tc.Data)
					if err != nil {
						log.Fatal(err)
					}
					return t
				}
			}
		}
	}
	return time.Time{}
}

func findImg(n *html.Node) img {
	var l string
	var t time.Time

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		for _, a := range c.Attr {
			if a.Key != "class" {
				continue
			}

			switch a.Val {
			case "_3-96 _2let":
				l = filepath.Join(dir, findLink(c))
			case "_3-94 _2lem":
				t = findTime(c)
			}

			if !t.IsZero() && l != "" {
				return img{p: l, t: t}
			}
		}
	}

	panic("unreachable")
}

func findImgs(n *html.Node) (is []img) {
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if c.Type != html.ElementNode || c.Data != "div" {
			continue
		}

		for _, a := range c.Attr {
			if a.Key == "class" && a.Val == "pam _3-95 _2pi0 _2lej uiBoxWhite noborder" {
				i := findImg(c)
				is = append(is, i)
			}
		}
	}
	return
}

func findMain(n *html.Node) *html.Node {
	if n.Type == html.ElementNode {
		f := false
		for _, a := range n.Attr {
			if a.Key == "class" && a.Val == "_4t5n" {
				f = true
			}
		}
		if f {
			return n
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		main := findMain(c)
		if main != nil {
			return main
		}
	}
	return nil
}

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "usage: %s dir album.html\n", os.Args[0])
		os.Exit(1)
	}
	dir = os.Args[1]

	f, err := os.Open(os.Args[2])
	if err != nil {
		log.Fatal(err)
	}

	doc, err := html.Parse(f)
	if err != nil {
		log.Fatal(err)
	}

	m := findMain(doc)
	is := findImgs(m)

	for _, i := range is {
		fmt.Println(i.p, i.t.Format(time.DateTime))
	}
}
