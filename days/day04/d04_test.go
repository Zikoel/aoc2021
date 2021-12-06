package day04

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func BoardDecoding(t *testing.T) {
	assert := assert.New(t)

	boardStr := `22 13 17 11  0
8  2 23  4 24
21  9 14 16  7
6 10  3 18  5
1 12 20 15 19`

	b := decodeBoard(boardStr)

	assert.Equal(5, len(b.rows))
	assert.Equal(5, len(b.rows[0].unmarked))
	assert.Equal(22, b.rows[0].unmarked[0])
	assert.Equal(22, b.rows[1].unmarked[1])
}
