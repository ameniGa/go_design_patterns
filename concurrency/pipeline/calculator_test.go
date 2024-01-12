package pipeline

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLaunchPipeline(t *testing.T) {
	tt := [][]int{
		{3, 14},
		{5, 55},
	}

	for _, tc := range tt {
		res := LaunchPipeline(tc[0])
		assert.Equal(t, tc[1], res)
	}

}
