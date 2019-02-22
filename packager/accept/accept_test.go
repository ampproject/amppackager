package accept

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanSatisfy(t *testing.T) {
	assert.False(t, CanSatisfy(""))
	assert.False(t, CanSatisfy("*/*"))
	assert.False(t, CanSatisfy("image/jpeg;v=b3"))
	assert.False(t, CanSatisfy(`application/signed-exchange;v=b2`))
	// This is a bug that will be triggered when a UA starts supporting multiple SXG versions:
	assert.False(t, CanSatisfy(`application/signed-exchange;x="a,b";v="b3"`))

	assert.True(t, CanSatisfy(`application/signed-exchange;v=b3`))
	assert.True(t, CanSatisfy(`application/signed-exchange;v="b3"`))
	assert.True(t, CanSatisfy(`application/signed-exchange;v=b3;q=0.8`))
	assert.True(t, CanSatisfy(`application/signed-exchange;v=b1,application/signed-exchange;v=b3`))
	assert.True(t, CanSatisfy(`application/signed-exchange;x="v=b1";v="b3"`))
	assert.True(t, CanSatisfy("*/*, application/signed-exchange;v=b3"))
	assert.True(t, CanSatisfy("*/* \t,\t application/signed-exchange;v=b3"))
	// This is the same bug, though one which won't occur in practice:
	assert.True(t, CanSatisfy(`application/signed-exchange;x="y,application/signed-exchange;v=b3,z";v=b1`))
}
