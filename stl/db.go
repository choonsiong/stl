package stl

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

// go get -u go.mongodb.org/mongo-driver
// go get -u go.mongodb.org/mongo-driver/mongo
// go get -u go.mongodb.org/mongo-driver/mongo/options
// go get -u go.mongodb.org/mongo-driver/bson

// 输入参数 stlName: 需要在 MongoDB 中查询 STL 网格数据的名称
// 输入参数 user, pwd, ip, port: 连接 MongoDB 的用户名，密码，IP 地址，端口号
// 输入参数 database, collection: 将 STL 数据存入 MongoDB 时指定的数据库和集合的名称
// 输入参数 timeout: 连接 MongoDB 的超时时间，单位为秒
// 返回 ModelSTL 类型的数据 modelSTL，包含 JSON 形式的 STL 网格数据
// 以及函数中间过程可能产生的错误 err
func QuerySTLMongo(stlName, user, pwd, ip string, port int, database, collection string, timeout int64) (modelSTL ModelSTL, err error) {
	// 设置连接 MongoDB 的用户名，密码，IP 地址，端口号
	clientOptions := options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%s@%s:%d", user, pwd, ip, port))

	// 连接 MongoDB 需要传入一个上下文对象 context，这里使用 context.WithTimeout，超时秒数为 timeout 之所以需要传入 context 对象，
	// 是因为当顶级 Goroutine 退出时，可以方便地终结所有子 Goroutine
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(timeout))
	// 当程序退出时，清空 context 对象 ctx
	defer cancel()

	// 连接到 MongoDB，得到数据库客户端对象 client
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		// 在发生连接错误时，打印错误并返回
		log.Println("stl/db.go/QuerySTLMongo, connect to mongodb error:", err)
		return
	}

	// 使用 MongoDB 客户端对象 client 对数据库执行 ping 指令，如果有错误，则说明连接存在问题
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Println("stl/db.go/QuerySTLMongo, ping mongodb fatal error:", err)
		return
	}

	// 在函数退出时，关闭客户端对象 client
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			log.Println("stl/db.go/QuerySTLMongo, disconnected to MongoDB error:", err)
			return
		}
	}()

	// 根据输入的 database 和 collection 名称，获取存储 STL 数据的 MongoDB 集合
	coll := client.Database(database).Collection(collection)

	// 查询条件变量 filterM 用来检索 name 字段为 stlName 的数据库记录
	filterM := bson.M{"name": stlName}

	// 根据查询条件变量 filterM 查找一条数据库记录，并将结果反序列化为 ModelSTL 类型的数据 modelSTL
	err = coll.FindOne(ctx, filterM).Decode(&modelSTL)
	if err != nil {
		log.Printf("stl/db.go/QuerySTLMongo, find and decode modelSTL(name=%s) error: %s\n", stlName, err)
		return
	}

	return
}

// 将 STL 文件中的数据存入 MongoDB 中
// 输入参数 modelSTL: ModelSTL 类型的 STL 三角面元信息
// 输入参数 user, pwd, ip, port: 连接 MongoDB 的用户名，密码，IP 地址，端口号
// 输入参数 database, collection: 将 STL 数据存入 MongoDB 时指定的数据库和集合的名称
// 输入参数 timeout: 连接 MongoDB 的超时时间，单位为秒
// 返回函数中间过程可能产生的错误 err
func SaveSTLMongo(modelSTL ModelSTL, user, pwd, ip string, port int, database, collection string, timeout int64) (err error) {
	// 设置连接 MongoDB 的用户名，密码，IP 地址，端口号
	clientOptions := options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%s@%s:%d", user, pwd, ip, port))

	// 连接 MongoDB 需要传入一个上下文对象 context，这里使用 context.WithTimeout，超时秒数为 timeout 之所以需要传入 context 对象，
	// 是因为当顶级 Goroutine 退出时，可以方便地终结所有子 Goroutine
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(timeout))
	// 当程序退出时，清空 context 对象 ctx
	defer cancel()

	// 连接到 MongoDB，得到数据库客户端对象 client
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		// 在发生连接错误时，打印错误并返回
		log.Println("stl/db.go/SaveSTLMongo, connect to mongodb error:", err)
		return
	}

	// 使用 MongoDB 客户端对象 client 对数据库执行 ping 指令，如果有错误，则说明连接存在问题
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Println("stl/db.go/SaveSTLMongo, ping mongodb fatal error:", err)
		return
	}

	// 在函数退出时，关闭客户端对象 client
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			log.Println("stl/db.go/SaveSTLMongo, disconnected to MongoDB error:", err)
			return
		}
	}()

	// 根据输入的 database 和 collection 名称，获取存储 STL 数据的 MongoDB 集合
	coll := client.Database(database).Collection(collection)

	// 查询条件变量 filterM 用来检索是否已经有名称为 modelSTL.Name 的记录，若有，则删除
	filterM := bson.M{"name": modelSTL.Name}
	docCount, err := coll.CountDocuments(ctx, filterM)

	// 删除全部的 "name" 字段为 modelSTL.Name 的数据库记录
	if docCount > 0 {
		_, err = coll.DeleteMany(ctx, filterM)
		if err != nil {
			log.Println("stl/db.go/SaveSTLMongo, delete documents error:", err)
			return
		}

	}

	// 插入 modelSTL 结构体数据，可以根据 ModelSTL 数据类型的 json tag 值来设置相应的 document 字段
	_, err = coll.InsertOne(ctx, modelSTL)
	if err != nil {
		log.Println("stl/db.go/SaveSTLMongo, insert document error:", err)
		return
	}

	return
}
