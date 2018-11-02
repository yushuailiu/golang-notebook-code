/*
demo1 包提供了一个计算相关的函数和常量。

第二段落的包注释~
链接测试：https://github.com/yushuailiu
 */
package demo0

// 数学 π 的值
const Pi = 3.1415926

// 数学常量
const (
	Ln2    = 0.693147180559945309417232121458176568075500134360255254120680009
	Log2E  = 1 / Ln2
	Ln10   = 2.30258509299404568401799145468436420760110148862877297603332790
	Log10E = 1 / Ln10
)

// Sum 用来计算一个 int 列表内所有元素的和
func Sum(items []int) int {
	sum := 0
	for i :=0 ; i < len(items) ; i++ {
		sum += items[i]
	}
	return sum
}