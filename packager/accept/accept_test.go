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
	assert.False(t, CanSatisfy(`application/signed-exchange;v="b1,b2"`))
	assert.False(t, CanSatisfy(`application/signed-exchange;x="y,application/signed-exchange;v=b3,z";v=b1`))

	assert.True(t, CanSatisfy(`application/signed-exchange;v=b3`))
	assert.True(t, CanSatisfy(`application/signed-exchange;v="b3"`))
	assert.True(t, CanSatisfy(`application/signed-exchange;v="b2,b3,b4"`))
	assert.True(t, CanSatisfy(`application/signed-exchange;v=b3;q=0.8`))
	assert.True(t, CanSatisfy(`application/signed-exchange;v=b2;q=0.9,application/signed-exchange;v="b3,b4";q=0.8`))
	assert.True(t, CanSatisfy(`application/signed-exchange;v=b1,application/signed-exchange;v=b3`))
	assert.True(t, CanSatisfy(`application/signed-exchange;x="v=b1";v="b3"`))
	assert.True(t, CanSatisfy("*/*, application/signed-exchange;v=b3"))
	assert.True(t, CanSatisfy("*/* \t,\t application/signed-exchange;v=b3"))
	assert.True(t, CanSatisfy(`application/signed-exchange;x="a,b";v="b3"`))
}
