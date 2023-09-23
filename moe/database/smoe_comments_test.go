package database

import (
	"testing"
)

func BenchmarkS_SortComments(b *testing.B) {
	qpu := NewQPU()
	qpu.CommentsWithCid(48)
	for i := 0; i < b.N; i++ {
		qpu.SortComments()
	}
}
