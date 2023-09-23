package database

import (
	"testing"
	"unsafe"
)

func TestS_StructMem(t *testing.T) {
	t.Log(unsafe.Alignof(Contents{}))
}
