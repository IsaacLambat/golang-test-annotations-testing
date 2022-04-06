package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"log"
	"testing"
	"time"
)

func TestSuccess(t *testing.T) {
	t.Error("Test")
}

func TestReproLog(t *testing.T) {
	log.Println("This is a test")
	require.True(t, false)
}

func TestReproPrint(t *testing.T) {
	fmt.Println("This is a test")
	require.True(t, false)
}

func TestRepro(t *testing.T) {
	ISO8601Format := "2006-01-02T15:04:05.000Z07:00"
	timestampPrefix := time.Now().Format(ISO8601Format)
	fmt.Println(timestampPrefix, "This is a test")
	require.True(t, false)

}

func TestMultipleAsserts(t *testing.T) {
	fmt.Println("test")
	assert.True(t, false)
	log.Println("test")
	fmt.Println("test")
	assert.True(t, false)
}

func TestPRFail(t *testing.T) {
	log.Println("This test will fail intentionally")
	assert.True(t, false)
}
