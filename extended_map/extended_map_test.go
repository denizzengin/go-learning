package extended_map

import (
	"fmt"
	"testing"
	"github.com/magiconair/properties/assert"
)


func Test_Map(t *testing.T) {

	expected := []int{2, 3, 4}
	arr := []int{1, 2, 3}
	CalculateAdd(arr)
	fmt.Println(expected)
	fmt.Println(arr)	
	assert.Equal(t, arr, expected)
}
