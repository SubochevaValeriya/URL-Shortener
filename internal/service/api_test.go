package service

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidateURL(t *testing.T) {
	t.Parallel()
	type testCase struct {
		url   string
		error bool
	}

	testCases := []testCase{
		{"https://www.google.com/", false},
		{"hsdh", true},
		{"]]]18", true},
	}

	for _, testCase := range testCases {
		_, err := validateURL(testCase.url)
		if testCase.error == false {
			assert.NoError(t, err)
		} else {
			assert.Error(t, err)
		}
	}
}
