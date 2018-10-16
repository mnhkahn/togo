// Package util
package jsontothrift

import (
	"testing"

	"github.com/mnhkahn/gogogo/logger"
	"github.com/stretchr/testify/assert"
)

func TestJsonToThrift(t *testing.T) {
	res, err := JsonToThrift(`{"id":1,"type":"2"}`)
	assert.Nil(t, err)
	assert.NotEmpty(t, res)
	logger.Info(string(res))
}
