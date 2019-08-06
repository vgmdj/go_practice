package Top_K_Frequent_Elements

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTopKFrequentElements(t *testing.T) {
	ast := assert.New(t)

	ast.Equal([]int{1, 3}, topKFrequent([]int{1, 1, 1, 2, 3, 3}, 2))
}
