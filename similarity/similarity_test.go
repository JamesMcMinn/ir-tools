package similarity

import (
	"testing"
)

func TestSimilarity(t *testing.T) {

	v1 := map[string]uint{"hello": 2}
	v2 := map[string]uint{"hello": 1}
	a := 1.0

	if x := Similarity(v1, v2); x != a {
		t.Errorf("Wrong Similarity Score: cosine(%s, %s) = %d, want %d.", v1, v2, x, a)
	}

}
