package stl

import (
	"encoding/binary"
	"example.com/stl/utils"
	"fmt"
	"log"
	"math"
	"os"
	"path/filepath"
)

// 输入 STL 文件名和所在的相对路径，返回 ModelSTL 类型的结构体和错误讯息
func ReadSTLFile(filename, dirPath string) (stl ModelSTL, err error) {
	info := "stl/stl.go/ReadSTLFile"
	defer func() {
		if err := recover(); err != nil {
			log.Printf("%s, read stl file error occurred: %s\n", info, err)
		}
	}()

	// 利用标准函数库函数 filepath.Join() 拼接 STL 文件的绝对路径
	f, err := os.Open(filepath.Join(utils.ExecutableDir, dirPath, filename))
	if err != nil {
		log.Printf("%s, open stl %s file error occurred: %s\n", info, filename, err)
		return
	}
	defer func() {
		err = f.Close()
		if err != nil {
			log.Println(info, fmt.Sprintf("%s, close file error occurred: %s\n", info, err))
		}
		return
	}()

	stl.Name = filename

	b := make([]byte, 4, 4)
	n, err := f.ReadAt(b, 80) // first 80 bytes is 模型名称, next 4 bytes 是三角面元数量
	if err != nil {
		log.Println(info, fmt.Sprintf("%s, read the stl model face nums failed at %d byte\n", info, 80+n))
		return
	}

	// 在 PC 上按照字节的小端序 LittleEndian 读取文件
	stl.FaceNum = int32(binary.LittleEndian.Uint32(b))
	stl.TriangleFaceArray = make([]TriangleFace, stl.FaceNum, stl.FaceNum)

	b = make([]byte, 50*stl.FaceNum, 50*stl.FaceNum) // 接下来的每 50 bytes 是一个面元 (N-12,A-12,B-12,C-12,属性-2)
	n, err = f.ReadAt(b, 84)
	if err != nil {
		log.Println(info, fmt.Sprintf("%s, read the stl face data failed at %d byte\n", info, 80+4+n))
		return
	}

	var i int32

	for i = 0; i < stl.FaceNum; i++ {
		offset := i * 50
		stl.TriangleFaceArray[i].N = [3]float32{
			math.Float32frombits(binary.LittleEndian.Uint32(b[0+offset : 4+offset])),
			math.Float32frombits(binary.LittleEndian.Uint32(b[4+offset : 8+offset])),
			math.Float32frombits(binary.LittleEndian.Uint32(b[8+offset : 12+offset])),
		}
		stl.TriangleFaceArray[i].A = [3]float32{
			math.Float32frombits(binary.LittleEndian.Uint32(b[12+offset : 16+offset])),
			math.Float32frombits(binary.LittleEndian.Uint32(b[16+offset : 20+offset])),
			math.Float32frombits(binary.LittleEndian.Uint32(b[20+offset : 24+offset])),
		}
		stl.TriangleFaceArray[i].B = [3]float32{
			math.Float32frombits(binary.LittleEndian.Uint32(b[24+offset : 28+offset])),
			math.Float32frombits(binary.LittleEndian.Uint32(b[28+offset : 32+offset])),
			math.Float32frombits(binary.LittleEndian.Uint32(b[32+offset : 36+offset])),
		}
		stl.TriangleFaceArray[i].C = [3]float32{
			math.Float32frombits(binary.LittleEndian.Uint32(b[36+offset : 40+offset])),
			math.Float32frombits(binary.LittleEndian.Uint32(b[40+offset : 44+offset])),
			math.Float32frombits(binary.LittleEndian.Uint32(b[44+offset : 48+offset])),
		}
	}

	return
}
