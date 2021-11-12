package scraper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExistsInSlice(t *testing.T) {
	testSlice := []string{"test", "test1"}
	testValue := "test3"
	expected := true
	actual := existsInSlice(testSlice, testValue)

	assert.Equal(t, expected, actual)
}
