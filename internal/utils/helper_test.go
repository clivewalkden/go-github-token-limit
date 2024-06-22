package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCenterString_WithNormalString(t *testing.T) {
	result := CenterString("test", 10)
	assert.Equal(t, "   test   ", result)
}

func TestCenterString_WithEmptyString(t *testing.T) {
	result := CenterString("", 10)
	assert.Equal(t, "          ", result)
}

func TestCenterString_WithCenteredString(t *testing.T) {
	result := CenterString("   test   ", 10)
	assert.Equal(t, "   test   ", result)
}

func TestCenterString_WithLongString(t *testing.T) {
	result := CenterString("this is a long string", 10)
	assert.Equal(t, "this is a long string", result)
}

func TestCenterString_WithOddWidth(t *testing.T) {
	result := CenterString("test", 11)
	assert.Equal(t, "   test    ", result)
}

func TestObscureToken_WithShortToken(t *testing.T) {
	result := ObscureToken("test")
	assert.Equal(t, "test", result)
}

func TestObscureToken_WithLongToken(t *testing.T) {
	result := ObscureToken("thisisaverylongtokenthisisaverylongtoken")
	assert.Equal(t, "thisisaveryl...oken", result)
}
