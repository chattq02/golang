package basic

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAddOne(t *testing.T) {
	// var (
	// 	input = 1
	// 	output = 2
	// )
	// actual := AddOne(1)
	// if actual != output {
    //     t.Errorf("AddOne(%d), input %d, actual = %d",input, output, actual)
    // }
	assert.Equal(t, AddOne(2), 4, "AddOne(2) should be 3")
}


// TestRequire chặn câu lệnh phía dưới khi fail
func TestRequire(t *testing.T) {
	require.Equal(t, 2, 3)
	fmt.Println("Not executing")
}


// TestAssert ko chặn câu lệnh phía dưới khi fail
func TestAssert(t *testing.T) {
	require.Equal(t, 2, 3)
	fmt.Println("executing")
}