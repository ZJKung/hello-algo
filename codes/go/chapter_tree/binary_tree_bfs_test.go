// File: binary_tree_bfs_test.go
// Created Time: 2022-11-26
// Author: Reanon (793584285@qq.com)

package chapter_tree

import (
	"fmt"
	"testing"

	. "github.com/krahets/hello-algo/pkg"
)

func TestHierOrder(t *testing.T) {
	/* 初始化二叉树 */
	// 这里借助了一个从数组直接生成二叉树的函数
	root := ArrToTree([]any{1, 2, 3, 4, 5, 6, 7})
	fmt.Println("\n初始化二叉树: ")
	PrintTree(root)

	// 层序遍历
	nums := hierOrder(root)
	fmt.Println("\n层序遍历的结点打印序列 =", nums)
}
