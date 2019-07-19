package skutils

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCopy(t *testing.T) {
	copy("/etc/resolv.conf", "/tmp/cp")
	origData, err := ioutil.ReadFile("/tmp/cp")
	copyData, err := ioutil.ReadFile("/tmp/cp")

	assert.Equal(t, nil, err)
	assert.Equal(t, origData, copyData)
}
