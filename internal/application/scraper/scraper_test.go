package scraper

import (
	"testing"

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

// func TestGetImageSrc(t *testing.T) {
// 	imageInputSrc :=
// }
