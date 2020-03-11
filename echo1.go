//echo1 输出其命令行参数
package main	//包

import (	//导入库
	"fmt"
	"os"
)

func main() {	//主函数
	var s, sep string	//变量
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]	//首次使用sep,初始值自动为空""
		sep = " "
	}
	fmt.Println(s)
}

