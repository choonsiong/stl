package main

import (
	"example.com/stl/handler"
	"example.com/stl/utils"
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	log.Println("main() started at:", time.Now().Format("2006-01-02 15:04:05"))

	// 创建一个名称为 mux 的 http multiplexer （多路复用器）
	// 不使用 http.DefaultServeMux
	mux := http.NewServeMux()

	// 定义服务器的 IP 地址
	// 当使用 127.0.0.1 时，服务器程序只能被本机访问
	// 当使用 0.0.0.0 时，服务器程序可被所在网路上的任意计算机访问
	ip := "0.0.0.0"

	// 定义服务器程序的运行端口，端口范围为 0-65535，建议使用大于 8080 的非活动端口
	port := "8080"

	// 当用户访问 / 或 /ping 两个路由地址时
	// 调用 handler 包中的 Ping() 函数进行处理
	mux.HandleFunc("/", handler.Ping)
	mux.HandleFunc("/ping", handler.Ping)
	// 以此类推...
	mux.HandleFunc("/get-stl-list", handler.GetSTLList)
	mux.HandleFunc("/save-stl-mongo", handler.SaveSTLMongo)
	mux.HandleFunc("/query-stl-mongo", handler.QuerySTLMongo)

	// 使用 Go 语言的标准库的 http 包创建 server 对象
	server := &http.Server{
		Addr:              ip + ":" + port,
		Handler:           mux,
		TLSConfig:         nil,
		ReadTimeout:       time.Second * 150,
		ReadHeaderTimeout: 0,
		WriteTimeout:      time.Second * 600,
		IdleTimeout:       0,
		MaxHeaderBytes:    0,
		TLSNextProto:      nil,
		ConnState:         nil,
		ErrorLog:          nil,
		BaseContext:       nil,
		ConnContext:       nil,
	}
	log.Println("server started at:", ip+":"+port)
	err := server.ListenAndServe()
	if err != nil {
		// 定义一个字符串型变量，用来存储错误信息
		//waitStr := ""

		// 当 server 对象在调用 ListenAndServe() 方法时发生错误，应强制退出程序
		// 使用 fmt.Scanf() 函数会将程序运行卡住，在按 Enter 键后，程序继续进行
		// 主要目的是防止程序退出命令行窗口一闪而过
		//_, _ = fmt.Scanf("panic error %s occurred, press \"Enter\" key to exit...\n", &waitStr)

		fmt.Println("panic error occurred, press \"Enter\" key to exit...")
		fmt.Scanln() // Waits for the user to press Enter

		// 使用 utils 包中定义的 panic 错误处理函数，程序强制退出
		utils.PanicHandler(err, "main.go")
	}
}
