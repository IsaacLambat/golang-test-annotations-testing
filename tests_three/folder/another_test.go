package folder

import (
	"github.com/stretchr/testify/require"
	"log"
	"testing"
)

func TestFailMultipleFoldersDeep(t *testing.T) {
	log.Println("Fail")
	require.True(t, false)
}
