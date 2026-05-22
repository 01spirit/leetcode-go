package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// 1. 创建快速读写器
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush() // 程序结束时刷新缓冲区

	// 2. 读一个整数 n（表示测试用例数或数组长度）
	var n int
	fmt.Fscan(in, &n)

	// 3. 读 n 个整数到切片
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}
	fmt.Println("n: ", n)
	fmt.Println("arr: ", a)

	// 4. 处理逻辑...
	ans := solve(a)

	// 5. 输出结果
	fmt.Fprintln(out, ans)
}

func init() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)

	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}

	fmt.Fprintln(out, "n: ", n)
}

func solve(a []int) int {
	// 你的算法逻辑
	return 0
}
