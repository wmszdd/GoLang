package main

import (
	"fmt"
	"math"
)

const s string = "constant"

func main() {
	fmt.Println(s)

	const n = 500000000

	const d = 3e20 / n
	fmt.Println(d)

	fmt.Println(int64(d))

	//这里的 math.Sin 函数需要一个 float64 的参数，n 会自动确定类型
	fmt.Println(math.Sin(n))

}
