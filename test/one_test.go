package test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	one "pretest-requirement-backend-engineer/one"
)

func TestOne(t *testing.T) {
	input := "Malam"
	assert.True(t, true, one.IsPalindrome(input))
}
