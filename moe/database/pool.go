package database

import (
	"sync"
)

// QPU
// 本包中 smoe_ 前缀是sql语句存储
// 本包中 struct_ 前缀是Data中的结构体的方法和少量关联查询拆成多次查询的sql
type QPU struct {
	HaveNext  int
	PostArr   []Contents
	PageArr   []Contents
	MediaArr  []Contents
	CommArr   []Comments
	UserInfo  User
	AccessLog []Access
}

// Free 清空s结构体存储的data，然后返还到池中
// todo 记得清空数据，因为sqlx的select方法是append，而不是clear后scan
func (q *QPU) Free() {
	//clear(s.PostArr)
	qpuPool.Put(q)
}

// qpu对象池
var qpuPool = sync.Pool{
	New: func() any {
		return new(QPU)
	},
}

// NewQPU 创建新的query progress unit
func NewQPU() *QPU {
	q := qpuPool.Get().(*QPU)
	qpuPool.Put(q)
	return q
}
