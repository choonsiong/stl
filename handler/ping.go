package handler

import (
	"example.com/stl/utils"
	"fmt"
	"log"
	"net/http"
	"time"
)

// 用于测试服务器程序的可连通性，给客户端返回一个字符串 “Pong” + 服务器时间
func Ping(w http.ResponseWriter, r *http.Request) {
	// 当发生致命错误时，调用 utils 包中的 RecoverHandler() 函数进行故障恢复，防止意外退出
	defer utils.RecoverHandler("handler.Ping")
	defer func() {
		// 关闭客户端的 HTTP 请求体 r.Body
		err := r.Body.Close()
		if err != nil {
			log.Println("handler \"Ping\" close the request body error occurred:", err)
			return
		}
	}()

	// 查看客户端请求的方法
	fmt.Println("handler \"Ping\": client request method is:", r.Method)
	// w 为响应写入器
	// w 返回的数据必须为字节切片格式，所以使用 []byte() 对返回的字符串进行强制的数据类型转换
	num, err := w.Write([]byte("Pong! " + time.Now().Format("2006-01-02 15:04:05")))
	// 如果 w 在返回响应时发生错误，则调用 utils 包的一般错误处理函数进行处理
	utils.ErrorHandler(err, "handler.Ping write response")
	// 在服务器端输出返回给客户端的字节数
	log.Println("Ping handler write", num, "bytes")
}
