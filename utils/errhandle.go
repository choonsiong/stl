// 用于错误处理的函数集
package utils

import "log"

// 可恢复的错误处理函数，阻止程序退出
// 用于虽然发生致命错误，但不让整个程序退出的情景
func RecoverHandler(info string) {
	// 内置的 recover() 函数用于处理发生 panic 级别错误的 Goroutine 信息
	if err := recover(); err != nil {
		log.Println("\""+info+"\", recover error occurred:", err)
	}
}

// 用于一般性的错误处理
func ErrorHandler(err error, info string) {
	if err != nil {
		log.Println("\""+info+"\", common error occurred:", err)
	}
}

// 当发生导致整个程序不能再继续正常工作的严重错误时（如启动 HTTP 服务器失败），强制退出程序
func PanicHandler(err error, info string) {
	if err != nil {
		// 打印错误信息并强制退出程序
		log.Fatalln(info, "exited, panic error occurred:", err)
	}
}
