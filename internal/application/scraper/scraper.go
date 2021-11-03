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

// <img src="//assets.icanet.se/t_ICAseAbsoluteUrl/imagevaultfiles/id_228677/cf_5291/paj_med_vegobacon_och_tomatsallad.jpg" alt="Paj med vegobacon och tomatsallad" class="lazyNoscriptActive" />

var ImgRegex string = `\n\s+<img src=`

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

		if match {
			n.Data = strings.TrimSpace(n.Data)
			n.Data = getImageSrc(n.Data)
			iBuf := &bytes.Buffer{}
			iBuf.WriteString(n.Data)
			imageUrls = append(imageUrls, iBuf)
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			visitNode(c)
		}
	}

	forEachNode(n, visitNode, nil)

	fmt.Printf("IMAGE URLS %v", len(imageUrls))
	fmt.Printf("TITLES %v", len(titles))
	fmt.Printf("DESCRIPTIONS %v", len(desc))
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

func getImageSrc(tag string) string {
	var out string = "https:"

	for idx, char := range tag {
		if string(char) == `"` && idx > 10 {
			return out
		}

		if idx > 9 {
			out += string(char)
		}
	}

	return out
}
