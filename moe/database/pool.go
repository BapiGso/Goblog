package database

import (
	"sync"
)

// QPU Query Processing Unit 输出一些需要查询结果的sql处理器
type QPU struct {
	User     User
	Contents []Contents
	Comments []Comments
	Access   []Access
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
	return q
}

// FreeQPU 清空s结构体存储的data，然后返还到池中
// 记得清空数据，因为sqlx的select方法是append，而不是clear后scan
func FreeQPU(q *QPU) {
	clear(q.Contents)
	clear(q.Comments)
	clear(q.Access)
	qpuPool.Put(q)
}
