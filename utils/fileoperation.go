// 对文件进行操作的函数集
package utils

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

var ExecutableDir string

// init 函数会在 utils 包中的其它函数之前运行
// 当其它的包引入 utils 包后，utils 包中的 init() 函数会先于其它函数运行
func init() {
	InitExecutableDir()
}

// 获取程序编译后的可执行文件所在文件夹的名称
func InitExecutableDir() {
	// 获取程序编译后的可执行文件路径 executableFilePath，包含可执行文件的名称
	executableFilePath, err := os.Executable()
	if err != nil {
		log.Println("\"GetFileList\" get executable file dir path error:", err)
		return
	}
	// 先使用 Go 语言标准库中的 filepath.Dir() 函数去除可执行文件的名称
	// 再使用 Go 语言标准库中的 filepath.EvalSymLinks() 函数去除符号链接
	// ExecutableDir 变量存储的就是 dirPath 所在的绝对路径
	ExecutableDir, err = filepath.EvalSymlinks(filepath.Dir(executableFilePath))
	if err != nil {
		fmt.Println("eval symlinks error:", err)
		PanicHandler(err, "utils/GetFileList")
	}
	fmt.Println("ExecutableDir:", ExecutableDir)
}

// 传入一个相对于程序可执行文件的相对目录，返回目录下的所有 STL 文件名的字符串切片
func GetFileList(dirPath string) (fileList []string, err error) {
	// 使用 Go 语言标准库中的 os.ReadDir() 函数获取路径下的所有文件
	files, err := os.ReadDir(filepath.Join(ExecutableDir, dirPath))
	if err != nil {
		log.Println("\"GetFileList\" read dir of \"", dirPath, "\" error:", err)
		return
	}
	// 获取文件夹下的文件数量
	// 下面假设该文件夹下既有子文件夹，也有扩展名不是 STL 的文件
	fileListLen := len(files)
	// 根据文件数量， 创建一个字符串数组，用于存放 STL 文件名
	fileList = make([]string, fileListLen, fileListLen)
	// 初始化一个 uint32 类型变量 numFilesSTL，用于存放确切的 STL 文件的数量
	var numFileSTL uint32 = 0
	for _, f := range files {
		// 使用标准库提供的 IsDir() 函数判断 f 是文件还是文件夹
		if !f.IsDir() {
			// 如果 f 为文件
			// 获取文件全名的后 4 位，即文件的扩展名，将其存储在变量 extName 中
			extName := f.Name()[len(f.Name())-4 : len(f.Name())]
			// 如果文件的扩展名 extName 为 ".STL" 或 ".stl"
			isSTL := (extName == ".STL") || (extName == ".stl")
			if isSTL {
				// 则 numFileSTL 加 1
				numFileSTL++
				// 将 STL 文件名放入 fileList 中
				fileList[numFileSTL-1] = f.Name()
			}
		}
	}
	// 只返回前 numFileSTL 个有效的 STL 文件名
	fileList = fileList[:numFileSTL]

	return
}
