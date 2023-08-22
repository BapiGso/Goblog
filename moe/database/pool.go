package database

import (
	"sync"
)

// S
// 本包中 smoe_ 前缀是sql语句存储
// 本包中 struct_ 前缀是Data中的结构体的方法和少量关联查询拆成多次查询的sql
type S struct {
	PostArr   []Contents
	PageArr   []Contents
	MediaArr  []Contents
	CommArr   []Comments
	UserInfo  User
	AccessLog []Access
}

// Free 清空s结构体存储的data，然后返还到池中
// todo 记得清空数据，因为sqlx的select方法是append，而不是clear后scan
func (s *S) Free() {
	//s.Data = nil
	ss.Put(s)
}

// qpu对象池
var ss = sync.Pool{
	New: func() any {
		return new(S)
	},
}

// NewQPU 创建新的query progress unit
func NewQPU() *S {
	p := ss.Get().(*S)
	return p
}
