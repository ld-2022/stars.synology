package path

import (
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"
)

func ExistFile(str string) bool {
	if _, err := os.Stat(str); err == nil {
		return true
	}
	return false
}

func GetCurrentRunDir() string {
	return getCurrentAbPath()
}
func GetCurrentRunFile() string {
	return filepath.Join(getCurrentAbPath(), filepath.Base(os.Args[0]))
}
func GetCurrentRunPath(filename string) string {
	return filepath.Join(GetCurrentRunDir(), filename)
}

// 最终方案-全兼容
func getCurrentAbPath() string {
	dir := getCurrentAbPathByExecutable()
	//fmt.Println("getCurrentAbPathByExecutable:", dir)
	//fmt.Println("getTmpDir():", getTmpDir())
	//if strings.Contains(dir, getTmpDir()) {
	//	return getCurrentAbPathByCaller()
	//}
	return dir
}

// 获取系统临时目录，兼容go run
func getTmpDir() string {
	dir := os.Getenv("TEMP")
	if dir == "" {
		dir = os.Getenv("TMP")
	}
	res, _ := filepath.EvalSymlinks(dir)
	return res
}

// 获取当前执行文件绝对路径
func getCurrentAbPathByExecutable() string {
	exePath, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	res, _ := filepath.EvalSymlinks(filepath.Dir(exePath))
	return res
}

// 获取当前执行文件绝对路径（go run）
func getCurrentAbPathByCaller() string {
	var abPath string
	_, filename, _, ok := runtime.Caller(0)
	if ok {
		abPath = path.Dir(filename)
	}
	fmt.Println("getCurrentAbPathByCaller:", abPath)
	return abPath
}
