package main

import "encoding/hex"

func main()  {
	byte_data := []byte("01040000000271cb")
	// 将 byte 装换为 16进制的字符串
	hex_string_data := hex.EncodeToString(byte_data)
	// byte 转 16进制 的结果
	println(hex_string_data)

	/* ====== 分割线 ====== */

	// 将 16进制的字符串 转换 byte
	hex_data, _ := hex.DecodeString(hex_string_data)
	// 将 byte 转换 为字符串 输出结果
	println(string(hex_data))
}
