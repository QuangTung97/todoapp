package mysql

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConfig_DSN(t *testing.T) {
	s := DefaultConfig.DSN()
	expected := "username:password@tcp(localhost:3306)/sample?parseTime=true&loc=Asia%2FHo_Chi_Minh"
	assert.Equal(t, expected, s)
}
