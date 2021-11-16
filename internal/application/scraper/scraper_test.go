package scraper

import (
	"testing"

	"github.com/ionysoshedblom/go_scraper/internal/domain/entity"
	"github.com/stretchr/testify/assert"
)

func TestExistsInSlice(t *testing.T) {
	testSlice := []string{"test1", "test2", "test2"}
	testValue := "test3"
	expected := false
	actual := existsInSlice(testSlice, testValue)

	assert.Equal(t, expected, actual)
}

func TestAppendNonDuplicates(t *testing.T) {
	testSlice := []string{"test", "test1", "test2"}
	testValue := "test2"

	expected := []string{"test", "test1", "test2"}
	actual := appendNonDuplicates(testSlice, testValue)

	assert.Equal(t, expected, actual)
}

func TestGetImageSrc(t *testing.T) {
	testImageTag := "\n<img src=\"//assets.icanet.se/t_ICAseAbsoluteUrl/imagevaultfiles/id_135889/cf_5291/halstrad_tonfisk_med_avokadohummus.jpg\" alt=\"Halstrad tonfisk med avokadohummus\" class=\"lazyNoscriptActive\" />\n"

	expected := "https://assets.icanet.se/t_ICAseAbsoluteUrl/imagevaultfiles/id_135889/cf_5291/halstrad_tonfisk_med_avokadohummus.jpg"
	actual := getImageSrc(testImageTag)

	assert.Equal(t, expected, actual)
}

func TestMapSliceValuesToRecipe(t *testing.T) {
	titles := []string{"title1", "title2"}
	descriptions := []string{"desc1", "desc2"}
	imageUrls := []string{"https://imageUrl1.jpg", "https://imageUrl2.jpg"}
	ingredients := [][]string{{"salt", "pepper"}, {"ketchup", "mustard"}}
	recipeIds := []int64{1, 2, 3}

	expected := []entity.Recipe{
		{Id: recipeIds[0], Title: titles[0], Description: descriptions[0], ImageUrl: imageUrls[0], Ingredients: ingredients[0]},
		{Id: recipeIds[1], Title: titles[1], Description: descriptions[1], ImageUrl: imageUrls[1], Ingredients: ingredients[1]},
	}
	actual := mapSliceValuesToRecipe(titles, descriptions, imageUrls, recipeIds, ingredients)

	assert.Equal(t, expected, actual)
}

func TestIsRegexMatch(t *testing.T) {
	tests := []struct {
		regex  string
		target string
		match  bool
	}{
		{`\n\s+<img src=`, "\n       <img src=\"//assets.icanet.se/halstrad_tonfisk_med_avokadohummus.jpg\" />\n", true},
		{`\n\s+<img src=`, "assets.icanet.se/halstrad_tonfisk_med_avokadohummus.jpg\" />\n", false},
	}

	for _, test := range tests {
		actual := isRegexMatch(test.regex, test.target)
		assert.Equal(t, test.match, actual)
	}
}
