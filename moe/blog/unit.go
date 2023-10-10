package blog

import (
	"SMOE/moe/database"
	"strconv"
)

// validateNum 首页返回1，不是数字返回err调用404，其他为对应页数
func validateNum(numStr string) (int, error) {
	if numStr == "" {
		return 1, nil
	}
	num, err := strconv.ParseUint(numStr, 10, 64)
	if err != nil {
		return 0, err
	}
	return int(num), nil
}

func sortComms(data []database.Comments) [][]database.Comments {
	var final [][]database.Comments
	parentMap := make(map[uint32]int)
	for _, v := range data {
		//父评论新建一个组，因为按时间排序肯定比子评论先
		if v.Parent == 0 {
			//初始化tmp的同时就把v加入切片
			tmp := []database.Comments{v}
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
