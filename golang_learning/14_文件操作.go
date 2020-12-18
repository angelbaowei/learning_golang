package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	f, err := os.Open("./README.md")  // Open() 只读形式打开
	if err != nil {
		fmt.Println("打开失败")
		return
	}

	var resSlice1 []byte
	var tmpSlice1 = make([]byte, 32)
	for {
		count1, err := f.Read(tmpSlice1)  // 1 Read读到[]byte
		if err == io.EOF {
			fmt.Println("读取完毕")
			break
		}
		if err != nil {
			fmt.Println("读取失败")
			return
		}
		fmt.Printf("读到 %v 个字节\n", count1)
		//fmt.Println(string(tmpSlice1))
		resSlice1 = append(resSlice1, tmpSlice1[:count1]...)
	}
	fmt.Println(string(resSlice1))
	f.Close()


	f, err = os.Open("./README.md")
	if err != nil {
		fmt.Println("打开失败")
		return
	}
	var resString2 string
	reader := bufio.NewReader(f)  // 2 通过bufio读
	for {
		str, err := reader.ReadString('\n')   // 读一行
		if err == io.EOF {
			resString2 += str   // 使用bufio最后也要加上去
			fmt.Println("读取完毕")
			break
		}
		//fmt.Println(str)
		resString2 += str
	}
	fmt.Println(resString2)
	f.Close()


	resSlice2, err := ioutil.ReadFile("./README.md")  // 3 通过ioutil打开和读取  一次读取全部  适合小文件  前两种是基于流的读取 适合大文件
	if err != nil {
		fmt.Println("打开或读取失败")
		return
	}
	fmt.Println(string(resSlice2))


	f, err = os.OpenFile("./test_file_op.txt", os.O_CREATE | os.O_RDWR, 0666)  // 创建并读写
	f.WriteString("写入string\n")
	f.Write(resSlice2)  // 写入 []byte  1 通过Write写入
	f.Close()

	f, err = os.OpenFile("./test_file_op.txt", os.O_RDWR | os.O_TRUNC, 0666)  // 创建并读写  O_TRUNC表示写入前清空文件  不加O_TRUNC就是按顺序覆盖写入
	writer := bufio.NewWriter(f)   // 2 通过bufio写入
	writer.WriteString("bufio写入string\n")  // 这句是写入到缓存了
	writer.Flush()  // 清空缓存 即将缓存内容写入到文件
	f.Close()

	err = ioutil.WriteFile("./test_file_op.txt", []byte("ioutil写入[]byte\n"), 0666)  // 3 通过ioutil写入
	if err != nil {
		fmt.Println("写入失败")
		return
	}


	// 创建目录
	err = os.Mkdir("./test_dir_op", 0666)
	if err != nil {
		fmt.Println("目录已经存在")
	}

	err = os.MkdirAll("./test_dir_op/dir1/dir2", 0666)  // 递归创建
	if err != nil {
		fmt.Println(err)
	}

	err = os.Remove("./test_dir_op/dir1/dir2")  // 删除一个文件或目录
	err = os.RemoveAll("./test_dir_op")  // 删除整个父目录

	err = os.Rename("./test_file_op.txt", "rename_test_file_op.txt")  // 重命名


}