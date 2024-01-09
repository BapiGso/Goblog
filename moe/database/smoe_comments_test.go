package database

import (
	"testing"
)

func BenchmarkS_SortComments(b *testing.B) {
	qpu := NewQPU()
	err := qpu.CommentsWithCid("publish", 48)
	if err != nil {
		return
	}
	for i := 0; i < b.N; i++ {
		qpu.SortComments()
	}
}
