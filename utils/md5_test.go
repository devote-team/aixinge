package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetFileMd5(t *testing.T) {
	assert.Equal(t, "ed7fb0c7ce9e95343016a3e7f6be70dd", GetFileMd5("./md5.go"))
}
