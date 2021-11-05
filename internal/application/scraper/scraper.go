package scraper

import (
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

var ImgRegex string = `\n\s+<img src=`

func (s Scraper) HandleSource(n *html.Node) ([]entity.Recipe, error) {
	var titles []string
	var desc []string
	var imageUrls []string
	var ingredients [][]string

	var visitNode func(*html.Node)

	visitNode = func(n *html.Node) {
		isTitle := n.Type == html.ElementNode && n.Parent.Data == "h2"
		isDescription := n.Type == html.ElementNode && n.Parent.Data == "a" && n.Data == "p"
		isIngredientsList := n.Type == html.ElementNode && n.Parent.Data == "li" && n.Data == "span" && n.Attr[1].Val == "ingredients"
		isImage := isRegexMatch(ImgRegex, n.Data)

		if isImage {
			n.Data = strings.TrimSpace(n.Data)
			n.Data = getImageSrc(n.Data)
			imageUrls = append(imageUrls, n.Data)
		} else if isTitle {
			stringExists := existsInSlice(titles, n.FirstChild.Data)

			if !stringExists {
				titles = append(titles, n.FirstChild.Data)
			}

		} else if isDescription {
			stringExists := existsInSlice(desc, n.FirstChild.Data)

			if !stringExists {
				desc = append(desc, n.FirstChild.Data)
			}

		} else if isIngredientsList {
			ingredientsSlice := strings.Split(n.Attr[0].Val, "\n")
			ingredients = append(ingredients, ingredientsSlice)
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			visitNode(c)
		}
	}

	forEachNode(n, visitNode, nil)

	recipes := mapSliceValuesToRecipe(titles, desc, imageUrls, ingredients)
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

func mapSliceValuesToRecipe(titles []string, descriptions []string, imageUrls []string, ingredients [][]string) []entity.Recipe {
	var out []entity.Recipe

	for i := 0; i < len(titles); i++ {
		recipe := &entity.Recipe{Title: titles[i], Description: descriptions[i], ImageUrl: imageUrls[i], Ingredients: ingredients[i]}
		out = append(out, *recipe)
	}

	return out
}

func existsInSlice(slice []string, value string) bool {
	for _, b := range slice {
		if b == value {
			return true
		}
	}
	return false
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

func isRegexMatch(regex string, target string) bool {
	rx, err := regexp.Compile(regex)
	if err != nil {
		fmt.Println(err)
	}

	match := rx.MatchString(target)

	return match
}
