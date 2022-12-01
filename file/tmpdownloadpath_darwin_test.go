//go:build darwin

package file

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_tmpDownloadPath(t *testing.T) {
	path := tmpDownloadPath("you-wont-find-me")

	assert.Equal(t, "you-wont-find-me.download/you-wont-find-me", path)
	err := os.Remove("you-wont-find-me.download")
	require.NoError(t, err)
	file, err := os.Create("find-me.download")
	defer func() {
		err = os.Remove(file.Name())
		assert.NoError(t, err)
	}()
	if err != nil {
		panic(err)
	}
	_, err = file.Write([]byte("hello"))
	require.NoError(t, err)
	err = file.Close()
	if err != nil {
		panic(err)
	}
	path = tmpDownloadPath("find-me")
	assert.Equal(t, fmt.Sprintf("find-me (1).download/find-me"), path, "it increments a number")
	err = os.Remove("find-me (1).download")
	assert.NoError(t, err)
}
