package accept

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanSatisfy(t *testing.T) {
	assert.False(t, CanSatisfy(""))
	assert.False(t, CanSatisfy("*/*"))
	assert.False(t, CanSatisfy("image/jpeg;v=b2"))
	// This is a bug, though one which won't occur in practice:
	assert.False(t, CanSatisfy(`application/signed-exchange;x="a,b";v="b2"`))

	assert.True(t, CanSatisfy(`application/signed-exchange;v=b2`))
	assert.True(t, CanSatisfy(`application/signed-exchange;v="b2"`))
	assert.True(t, CanSatisfy(`application/signed-exchange;v=b2;q=0.8`))
	assert.True(t, CanSatisfy(`application/signed-exchange;v=b1,application/signed-exchange;v=b2`))
	assert.True(t, CanSatisfy(`application/signed-exchange;x="v=b1";v="b2"`))
	assert.True(t, CanSatisfy("*/*, application/signed-exchange;v=b2"))
	assert.True(t, CanSatisfy("*/* \t,\t application/signed-exchange;v=b2"))
	// This is the same bug:
	assert.True(t, CanSatisfy(`application/signed-exchange;x="y,application/signed-exchange;v=b2,z";v=b1`))
}
