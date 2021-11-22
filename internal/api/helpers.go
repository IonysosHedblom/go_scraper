package api

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/ionysoshedblom/go_scraper/internal/domain/entity"
	"golang.org/x/net/html"
)

func (s *api) CallSource(url string) (*html.Node, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("getting %s: %s", url, res.Status)
	}

	doc, err := html.Parse(res.Body)

	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
	}

	return doc, nil
}

func (s *api) TrimHtmlAndCallSrc(url string) (*html.Node, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("getting %s: %s", url, res.Status)
	}

	defer res.Body.Close()

	src, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return nil, fmt.Errorf("reading %s: %s", url, res.Status)
	}

	strToReplace := `<li style="display:;"><a href="/handla/" class="navigation__item">Handla online</a> <!----></li>`

	cleanSrcDoc := strings.ReplaceAll(string(src), strToReplace, "")
	splitSrcDoc := strings.Split(cleanSrcDoc, `<div class="comment-section__wrapper extra-padding"><div class="comment-section__write-comment"><div aria-label="Skriv din kommentar" class="input-textarea input-textarea--56 input-textarea--simple">`)
	completeSrcDoc := splitSrcDoc[0] + "</body>" + "\n" + "</html>"
	doc, err := html.Parse(strings.NewReader(completeSrcDoc))

	if err != nil {
		return nil, err
	}

	return doc, nil
}

func buildUrlWithIngredientsQuery(ingredients []string) string {
	var queryString string = "https://www.ica.se/Templates/ajaxresponse.aspx?ajaxFunction=RecipeListMdsa&num=20&sortbymetadata=Relevance&id=12&mdsarowentityid=ca2947b2-0c0b-4936-b300-a42700eb2734"

	for _, ingredient := range ingredients {
		queryString += fmt.Sprintf("&filter=Ingrediens%%3A%v", strings.Title(ingredient))
	}

	return queryString
}

func buildQueryUrl(query string) string {
	query = strings.ReplaceAll(query, " ", "+")
	url := fmt.Sprintf("https://www.ica.se/Templates/ajaxresponse.aspx?ajaxFunction=RecipeListMdsa&mdsarowentityid=&num=16&query=%s&sortbymetadata=Relevance&id=12", query)
	return url
}

func buildRecipePageUrl(title string, id string) string {
	title = strings.ReplaceAll(title, " ", "-")
	title = strings.ReplaceAll(title, "ä", "a")
	title = strings.ReplaceAll(title, "ö", "o")
	title = strings.ReplaceAll(title, "å", "a")

	url := fmt.Sprintf("https://www.ica.se/recept/%s-%s/", title, id)
	return url
}

func getInstructions(in []entity.Instruction) []string {
	var out []string

	for _, i := range in {
		out = append(out, i.Text)
	}
	return out
}
