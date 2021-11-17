package scraper

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/ionysoshedblom/go_scraper/internal/domain/entity"
	"github.com/ionysoshedblom/go_scraper/internal/shared"
)

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

func mapSliceValuesToRecipe(
	titles,
	descriptions,
	imageUrls []string,
	recipeIds []int64,
	ingredients [][]string) []entity.Recipe {

	var recipes []entity.Recipe

	for i := 0; i < len(titles); i++ {
		recipe := &entity.Recipe{
			Id:          recipeIds[i],
			Title:       titles[i],
			Description: descriptions[i],
			ImageUrl:    imageUrls[i],
			Ingredients: ingredients[i],
		}
		recipes = append(recipes, *recipe)
	}

	return recipes
}

func mapIdsToInt64(idSlice []string) ([]int64, error) {
	var int64Slice []int64

	for _, str := range idSlice {
		num, err := shared.ConvertStringToInt64(str)
		if err != nil {
			fmt.Println("String is not convertible to int64")
			return nil, err
		}

		int64Slice = append(int64Slice, *num)
	}

	return int64Slice, nil
}
