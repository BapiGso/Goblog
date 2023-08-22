package blog

import (
	"github.com/BapiGso/SMOE/moe/query"
	"strconv"
)

// IsNum 首页返回1，不是数字返回err调用404，其他为对应页数
func isNum(numstr string) (int, error) {
	if numstr == "" {
		return 1, nil
	}
	num, err := strconv.ParseUint(numstr, 10, 64)
	if err != nil {
		return 0, err
	}
	return int(num), nil
}

func sortComms(data []query.Comments) [][]query.Comments {
	var final [][]query.Comments
	parentMap := make(map[uint32]int)
	for _, v := range data {
		//父评论新建一个组，因为按时间排序肯定比子评论先
		if v.Parent == 0 {
			//初始化tmp的同时就把v加入切片
			tmp := []query.Comments{v}
			final = append(final, tmp)
			parentMap[v.Coid] = len(final) - 1
		} else { //如果是子评论，找自己属于哪个父评论组
			index := parentMap[v.Parent]
			final[index] = append(final[index], v)
			parentMap[v.Coid] = index
		}
	}
	//fmt.Println(parentIndexMap)
	return final
}
