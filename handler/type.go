package handler

// 当 HTTP 服务器无法正确返回结果时，返回 ResponseStatus
type ResponseStatus struct {
	// 使用结构体标签 tag，当对 Status 字段进行 json 或 xml 序列化时，字段名称为 status
	Status string `json:"status" xml:"status"`
}

// 返回 STL 文件列表的结构体类型
type STLFileList struct {
	STLList []string `json:"stl_list" xml:"stl_list"`
}

type STLFileName struct {
	Name string `json:"name" xml:"name"`
}
