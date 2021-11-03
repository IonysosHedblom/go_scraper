package scraper

import (
	"bytes"
	"fmt"
	"regexp"
	"strings"

	"github.com/ionysoshedblom/go_scraper/internal/domain/entity"
	"golang.org/x/net/html"
)

type Scraper struct{}

func New() *Scraper {
	return &Scraper{}
}

// Img format in node - massage this later
// "\n                    <img src=\"//assets.icanet.se/t_ICAseAbsoluteUrl/imagevaultfiles/id_63195/cf_5291/paj_med_adelost_och_purjolok.jpg\" alt=\"Paj med ädelost och purjolök\" class=\"lazyNoscriptActive\" />\n                "
// 20 space
var ImgRegex string = `\n\s+<img`

func (s Scraper) HandleSource(n *html.Node) ([]entity.Recipe, error) {
	var titles []*bytes.Buffer
	var desc []*bytes.Buffer
	var imageUrls []*bytes.Buffer
	var visitNode func(*html.Node)

	visitNode = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Parent.Data == "h2" {
			titleBuf := &bytes.Buffer{}
			writeNodeContentToBuffer(n, titleBuf)
			titles = append(titles, titleBuf)
		}

		if n.Type == html.ElementNode && n.Parent.Data == "a" && n.Data == "p" {
			dBuf := &bytes.Buffer{}
			writeNodeContentToBuffer(n, dBuf)
			desc = append(desc, dBuf)
		}

		rx, _ := regexp.Compile(ImgRegex)
		match := rx.MatchString(n.Data)

		isImgTag := strings.Contains(n.Data, "\n                    <img")

		if match {
			fmt.Println(n.Data)
		}
		if n.Type == html.ElementNode && n.Parent.Data == "noscript" && isImgTag {
			iBuf := &bytes.Buffer{}
			writeNodeContentToBuffer(n, iBuf)
			imageUrls = append(imageUrls, iBuf)
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			visitNode(c)
		}
	}

	forEachNode(n, visitNode, nil)
	recipes := mapBufValuesToStruct(titles, desc)
	return recipes, nil
}

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		post(n)
	}
}

func writeNodeContentToBuffer(n *html.Node, buf *bytes.Buffer) {
	if n.Type == html.TextNode {
		buf.WriteString(n.Data)
	}

	if n.Type == html.ElementNode {
		buf.WriteString(n.Data)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		writeNodeContentToBuffer(c, buf)
	}
}

func mapBufValuesToStruct(titles []*bytes.Buffer, descriptions []*bytes.Buffer) []entity.Recipe {
	var out []entity.Recipe
	for i := 0; i < len(titles); i++ {
		recipe := &entity.Recipe{Title: titles[i].String(), Description: descriptions[i].String()}
		out = append(out, *recipe)
	}
	return out
}
