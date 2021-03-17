package files

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func init() {
	os.Setenv("NB_CODE_LINES_EDGES", "5")
}

func TestGetSurroundingLineIndex(t *testing.T) {
	code := `	1;
2;
	3;
4;
5;
6;
	7;
8;
9;
	10;
	11;
12;
	13;
14;
15;`
	codeSampleExpected := `	3;
4;
5;
6;
	7;
8;
9;
	10;
	11;
12;
	13;`
	start := getSurroundingLineIndex(code, 24, true)
	end := getSurroundingLineIndex(code, 24, false)

	assert.Equal(t, 7, start)
	assert.Equal(t, 48, end)
	assert.Equal(t, codeSampleExpected, code[start:end])
}

func TestGetSurroundingLineIndexEdgeCases(t *testing.T) {
	code := `	1;
2;
	3;
	4;
5;`
	start := getSurroundingLineIndex(code, 8, true)
	end := getSurroundingLineIndex(code, 8, false)

	assert.Equal(t, 0, start)
	assert.Equal(t, 17, end)
	assert.Equal(t, code, code[start:end])
}
