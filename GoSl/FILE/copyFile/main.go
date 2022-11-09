package main

import (
	"fmt"
	"io"
	"os"
)

//文件拷贝函数
func copyfile(dstName, srcName string) (written int64, err error) {
	//以只读方式打开源文件
	src, err := os.Open(srcName)
	if err != nil {
		fmt.Printf("open %s failed,err:%v\n", srcName, err)
		return
	}
	defer src.Close()

	//以只写，创建的方式打开目标文件
	dst, err := os.OpenFile(dstName, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open %s failed ,err:%v\n", dstName, err)
		return
	}
	defer dst.Close()

	return io.Copy(dst, src) //调用io.copy()拷贝内容
}
func main() {
	_, err := copyfile("dst.txt", "src.txt")
	if err != nil {
		fmt.Println("copy file failed,err:", err)
		return
	}
	fmt.Println("copy done!!!!!")

}
