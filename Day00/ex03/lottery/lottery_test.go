package lottery

import (
	"testing"
	"github.com/stretchr/testify/assert"
  )
  
  func Testnumbers(t *testing.T) {
  
	var a string = "eeee";
	var b string = "Bienvenue sur GoLang!\n"
  
	assert.Equal(t, a, b, "The two words should be the same.")
  }