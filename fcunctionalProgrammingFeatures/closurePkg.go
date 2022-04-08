package main

import "fmt"

/*
	函数编程特性：闭包
		特征：闭包有多种特征，但其最主要的特征就是在 函数体内部又再次实现了一个方法。
		功能：闭包最主要的功能就是 利用作用域。
		缺点：闭包写法容易导致奇怪的BUG,代码可读性低，非必要不要使用闭包
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}
	res := LeverOrder(root)
	fmt.Println(res)
}

func LeverOrder(root *TreeNode) [][]int {
	var res [][]int = make([][]int, 0)

	// 定义闭包 变量 BackTrack
	var BackTrack func(*TreeNode, int)
	BackTrack = func(root *TreeNode, depth int) {
		if root == nil {
			return
		}
		if depth == len(res) {
			res = append(res, []int{})
		}
		res[depth] = append(res[depth], root.Val)
		BackTrack(root.Left, depth+1)
		BackTrack(root.Right, depth+1)
	}
	BackTrack(root, 0)
	return res
}
