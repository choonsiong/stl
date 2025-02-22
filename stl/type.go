package stl

// STL 文件中的三角面元
type TriangleFace struct {
	N [3]float32 `json:"n" xml:"n" bson:"n"`
	A [3]float32 `json:"a" xml:"a" bson:"a"`
	B [3]float32 `json:"b" xml:"b" bson:"b"`
	C [3]float32 `json:"c" xml:"c" bson:"c"`
}

// 识别的 STL 文件采用二进制编码格式
// 二进制编码格式的 STL 文件使用固定的字节数给出三角面元的几何信息
// 文件起始的 80 字节是文件头，用于存储文件名
// 紧接着使用 4 字节的整数描述模型的三角面元个数
// 后面逐个给出每个三角面元的集合信息
// 每个三角面元占用固定的 50 字节，依次是：
// 3 个 4 字节浮点数 （三角面元的法线矢量）
// 3 个 4 字节浮点数 （三角面元第 1 个顶点的坐标）
// 3 个 4 字节浮点数 （三角面元第 2 个顶点的坐标）
// 3 个 4 字节浮点数 （三角面元第 3 个顶点的坐标）
// 三角面元的最后 2 个字节用来描述三角面元的属性信息
// 一个完整的二进制编码格式的 STL 文件的大小为三角面元数 * 50 + 84 字节
type ModelSTL struct {
	Name              string         `json:"name" bson:"name"`
	FaceNum           int32          `json:"face_num" bson:"face_num"`
	TriangleFaceArray []TriangleFace `json:"triangle_face_array" bson:"triangle_face_array"`
}
