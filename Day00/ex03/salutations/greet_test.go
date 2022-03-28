package salutations

import (
	"testing"
	"github.com/stretchr/testify/assert"
  )
  
  func TestFrench(t *testing.T) {
  
	var a string = Francais();
	var b string = "Bienvenue sur GoLang!\n"
  
	assert.Equal(t, a, b, "The two words should be the same.")
  }

  func TestEnglish(t *testing.T) {
  
	var a string = English();
	var b string = "Welcome to GoLang\n"
  
	assert.Equal(t, a, b, "The two words should be the same.")
  }

  func TestChinese(t *testing.T) {
  
	var a string = Chinois();
	var b string = "欢迎来到 Golang!\n"
  
	assert.Equal(t, a, b, "The two words should be the same.")
  }