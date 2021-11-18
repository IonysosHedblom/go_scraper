package scraper

import (
	"strings"

	"github.com/ionysoshedblom/go_scraper/internal/domain/entity"
	"golang.org/x/net/html"
)

type scraper struct{}

func New() *scraper {
	return &scraper{}
}

func (s *scraper) GetRecipeResults(n *html.Node) ([]entity.Recipe, error) {
	titles, descriptions, imageUrls, recipeIds, ingredients := findRecipeInformation(n)

	int64RecipeIds, err := mapIdsToInt64(recipeIds)
	if err != nil {
		return nil, err
	}

	recipes := mapSliceValuesToRecipe(titles, descriptions, imageUrls, int64RecipeIds, ingredients)
	return recipes, nil
}

func (s *scraper) GetRecipeDetails(n *html.Node) entity.RecipeDetails {
	u, i, c := findRecipeDetails(n)
	recipeDetails := mapDetailsToStruct(u, i, c)
	return recipeDetails
}

func mapDetailsToStruct(units, ingredients, checklist []string) entity.RecipeDetails {
	var recipeDetails entity.RecipeDetails
	var combinedUnitAndIngredient []string

	for i := 0; i < len(ingredients); i++ {
		combinedUnitAndIngredient = append(combinedUnitAndIngredient, ingredients[i]+" "+units[i])
	}

	recipeDetails.Ingredients = combinedUnitAndIngredient
	recipeDetails.Checklist = checklist

	return recipeDetails
}

func findRecipeInformation(n *html.Node) (t, d, i, ri []string, ing [][]string) {
	const imgRegex string = `\n\s+<img src=`
	var recipeIds []string
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
			recipeIds = appendNonDuplicates(recipeIds, n.Attr[2].Val)
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
	return titles, descriptions, imageUrls, recipeIds, ingredients
}

func findRecipeDetails(n *html.Node) (u, i, c []string) {
	var units []string
	var ingredients []string
	var checklist []string

	var visitNode func(n *html.Node)
	visitNode = func(n *html.Node) {
		isElementNode := n.Type == html.ElementNode
		isUnit := isElementNode && n.Parent.Data == "div" && n.Data == "span" && n.Attr[0].Val == "ingredients-list-group__card__qty"
		isIngredient := isElementNode && n.Data == "span" && n.Attr[0].Val == "ingredients-list-group__card__ingr"
		isChecklist := isElementNode && n.Data == "div" && n.Attr[0].Val == "cooking-steps-main__text"

		if isUnit {
			units = append(units, n.FirstChild.Data)
		} else if isIngredient {
			ingredients = append(ingredients, n.FirstChild.Data)
		} else if isChecklist {
			checklist = append(checklist, n.FirstChild.Data)
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			visitNode(c)
		}
	}

	forEachNode(n, visitNode, nil)
	return units, ingredients, checklist
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
