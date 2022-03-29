
package find
import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestFindWhenWork(T *testing.T) {
	assert := assert.New(T);
	
	a := []int {1, 3  ,4};
	b := Find(5, 10, 3, 4, 1);
	assert.Equal(a, b, "Need fo find [1,3,4]");
	
	a = []int(nil);
	b = Find(5, 10);
	assert.Equal(a, b, "Need to find nothing");
	
	a = []int {2};
	b = Find(5, 2);
	assert.Equal(a, b, "need to find [2]");
}