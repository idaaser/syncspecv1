package syncspec

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPointer(t *testing.T) {
	s := "johndoe"
	sp := Pointer(s)
	assert.Equal(t, s, *sp)

	i := 10
	ip := Pointer(i)
	assert.Equal(t, i, *ip)
}
