package deepcopy
import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestDeepCopy(T * testing.T) {
	assert := assert.New(T);

	a:= []int{1, 3, 4};
	b, err:= DeepCopy(a, 0, 2);

	assert.Equal(a, b, "Need to be equal");

	a[0] = 0;
	assert.NotEqual(a, b, "Need to be inequal");
	assert.Equal(err, nil , "Err need to be nil");

	b, err = DeepCopy(a, 0, 1);
	a = []int{0, 3};
	assert.Equal(a, b, "Need to be equal")

	b, err = DeepCopy(a, 1, 1);
	a = []int{3};
	assert.Equal(a, b, "Need to be equal")


	_, err = DeepCopy(a, 3, 10);
	assert.NotEqual(err, nil, "need to find an error");

	_, err = DeepCopy(a, 2, 0);
	assert.NotEqual(err, nil, "need to find an error");
}