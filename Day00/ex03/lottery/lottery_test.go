package lottery

import (
	"testing"
	"github.com/stretchr/testify/assert"
  )
  

  func TestAnswer(t *testing.T) {
  
	for i:=0; i < 1000 ; i++ { 
		number := Answer();
		assert.Condition(t, func() bool {if number >= 0 || number <= 99 {
			return true; }
		return false;},  "out of bound")
	}
  
  }