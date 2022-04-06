package tests_two

import (
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestFailAsWell(t *testing.T) {
	log.Println("This test will fail intentionally")
	assert.True(t, false)
}
