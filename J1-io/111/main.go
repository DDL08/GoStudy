package main

import (
	"io"
	"os"
	"strings"
)

func ReadFrom(reader io.Reader, num int) ([]byte, error) {
	//输入两个参数，传入一个实现了 io.Reader 接口的对象，一个数字
	// 返回两个参数（字节+报错
	p := make([]byte, num)
	/*make(...): Go 的内建函数，用于创建切片（slice）、映射（map）或通道（chan）
	// 创建一个能装 num 的字节的切片
	*/
	n, err := reader.Read(p)
	//读取到数据就返回数据，没读取到就返回错误（err）
	if n > 0 {
		return p[:n], nil
		//从切片 p 的第 0 个元素开始，取到第 n-1 个元素（不包含第 n 个），然后第二个参数是nil空，表示没错
	}
	return p, err
}

func main() {
	/*
							data, err := ReadFrom(os.Stdin, 11)
							//表示输入11个字节的数字
							if err != nil {
								fmt.Println("报错", err)
								return
							} else {
								fmt.Println("数据", string(data))
							}
		-------------------------------------------------------------------------------------------------

					reader := strings.NewReader("Go语言中文网")
					p := make([]byte, 6)
					n, err := reader.ReadAt(p, 2)
					//从第二个字节开始读6个字节，一个字符是1个字节，一个汉字是3个字节
					//ReadAt是
					if err != nil {
						panic(err)
					}
					fmt.Printf("%s, %d\n", p, n)

		-------------------------------------------------------------------------------------------------
		file, err := os.Open("C:\\Users\\LEGION\\Desktop\\ezRce\\num.txt")
			if err != nil {
				panic(err)
			}
			defer file.Close()
			//延迟执行，直到遇到return或者
			writer := bufio.NewWriter(os.Stdout)
			writer.ReadFrom(file)
			//将文件写入writer的缓冲区
			writer.Flush()
			//将缓冲区中的数据强制写入到标准输出。
		-------------------------------------------------------------------------------------------------
		reader := strings.NewReader("DDL咚咚")
		//Seek 设置下一次 Read 或 Write 的偏移量
		reader.Seek(-6, io.SeekEnd)
		r, _, _ := reader.ReadRune()
		fmt.Printf("%c\n", r)
		//注意：当文件 studygolang.txt 不存在或找不到时，file.Close() 会 panic，因为 file 是 nil。因此，应该将 defer file.Close() 放在错误检查之后。
	*/

	// var buffer bytes.Buffer

	// // 向缓冲区写入字节
	// buffer.WriteByte('a')
	// buffer.WriteByte('b')
	// buffer.WriteByte('c')
	// // 将缓冲区内容转换为字符串并打印
	// //newCh, _ := buffer.ReadByte()
	// //fmt.Println(newCh)         // 输出: 97
	// fmt.Println(buffer.String()) // 输出: abc

	// reader := io.TeeReader(strings.NewReader("Go语言中文网"), os.Stdout)
	// reader.Read(make([]byte, 20))

	io.CopyN(os.Stdout, strings.NewReader("Go语言中文网"), 8)
}

/*
# 样例1
go run main.go
12345678900
数据 12345678900

go run main.go
111111111111111111111111111111
数据 11111111111
-------------------------------------------------------------------------------------------------
# 样例2
输出↓
语言, 6

-------------------------------------------------------------------------------------------------
# 样例2
*/
