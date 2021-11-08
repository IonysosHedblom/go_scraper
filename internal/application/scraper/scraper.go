package scraper

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/ionysoshedblom/go_scraper/internal/domain/entity"
	"golang.org/x/net/html"
)

type scraper struct{}

func New() *scraper {
	return &scraper{}
}

func (s scraper) GetRecipeResults(n *html.Node) []entity.Recipe {
	titles, descriptions, imageUrls, ingredients := findRecipeInformation(n)
	recipes := mapSliceValuesToRecipe(titles, descriptions, imageUrls, ingredients)
	return recipes
}

func findRecipeInformation(n *html.Node) (t, d, i []string, ing [][]string) {
	const imgRegex string = `\n\s+<img src=`
	var titles []string
	var descriptions []string
	var imageUrls []string
	var ingredients [][]string
	var visitNode func(n *html.Node)

	visitNode = func(n *html.Node) {
		isElementNode := n.Type == html.ElementNode
		isTitle := isElementNode && n.Parent.Data == "h2"
		isDescription := isElementNode && n.Parent.Data == "a" && n.Data == "p"
		isIngredientsList := isElementNode && n.Parent.Data == "li" && n.Data == "span" && n.Attr[1].Val == "ingredients"
		isImage := isRegexMatch(imgRegex, n.Data)

		if isImage {
			n.Data = getImageSrc(n.Data)
			imageUrls = append(imageUrls, n.Data)
		} else if isTitle {
			titles = appendNonDuplicates(titles, n.FirstChild.Data)
		} else if isDescription {
			descriptions = appendNonDuplicates(descriptions, n.FirstChild.Data)
		} else if isIngredientsList {
			ingredientsSlice := strings.Split(n.Attr[0].Val, "\n")
			ingredients = append(ingredients, ingredientsSlice)
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			visitNode(c)
		}
	}

	forEachNode(n, visitNode, nil)
	return titles, descriptions, imageUrls, ingredients
}

func isRegexMatch(regex string, target string) bool {
	rx, err := regexp.Compile(regex)
	if err != nil {
		fmt.Print("Could not compile regex", err)
	}

	match := rx.MatchString(target)

	return match
}

func getImageSrc(tag string) string {
	tag = strings.TrimSpace(tag)
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

func appendNonDuplicates(targetSlice []string, value string) []string {
	stringExists := existsInSlice(targetSlice, value)

	if !stringExists {
		targetSlice = append(targetSlice, value)
	}

	return targetSlice
}

func existsInSlice(slice []string, value string) bool {
	for _, b := range slice {
		if b == value {
			return true
		}
	}
	return false
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

func mapSliceValuesToRecipe(
	titles,
	descriptions,
	imageUrls []string,
	ingredients [][]string) []entity.Recipe {

	var recipes []entity.Recipe

	for i := 0; i < len(titles); i++ {
		recipe := &entity.Recipe{
			Title:       titles[i],
			Description: descriptions[i],
			ImageUrl:    imageUrls[i],
			Ingredients: ingredients[i],
		}
		recipes = append(recipes, *recipe)
	}

	return recipes
}
