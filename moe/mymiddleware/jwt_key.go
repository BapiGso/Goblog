package mymiddleware

import "strconv"

// JWTKey 生成一个随机[]byte
// var JWTKey = []byte(strconv.Itoa(rand.Int()))
var JWTKey = []byte(strconv.Itoa(123))
