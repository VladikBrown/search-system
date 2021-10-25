package search

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBuildDocumentService(t *testing.T) {
	s := BuildDocumentService([]DocumentArg{
		{
			Name:  "first",
			Words: []string{"new", "york", "times"},
		},
		{
			Name:  "second",
			Words: []string{"new", "york", "post"},
		},
		{
			Name:  "third",
			Words: []string{"los", "angeles", "times"},
		},
	})
	res := s.Search([]string{"new", "new", "times"})
	require.Equal(t, 3, len(res))
	assert.Equal(t, "0", res[0].Doc.id)
	assert.Equal(t, 774, int(res[0].SimilarityRate*1000))
	assert.Equal(t, 292, int(res[1].SimilarityRate*1000))
	assert.Equal(t, 112, int(res[2].SimilarityRate*1000))
	assert.Equal(t, "1", res[1].Doc.id)
	assert.Equal(t, "2", res[2].Doc.id)
}
