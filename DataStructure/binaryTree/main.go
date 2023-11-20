package main

import "fmt"

type Hero struct {
	No    int
	name  string
	left  *Hero
	right *Hero
}

// 前序遍历
// 先输出 root 节点, 然后输出左子树, 然后输出右子树
func PreOrder(node *Hero) {
	if node != nil {
		fmt.Printf("no = %d name=%s\n", node.No, node.name)
		PreOrder(node.left)
		PreOrder(node.right)
	}
}

// 中序遍历 先输出root左子树, 再输出 root节点, 最后输出 root 的右子树
func InfixOrder(node *Hero) {
	if node != nil {
		PreOrder(node.left)
		fmt.Printf("no = %d name=%s\n", node.No, node.name)
		PreOrder(node.right)
	}
}

// 后序遍历
func PostOrder(node *Hero) {
	if node != nil {
		PreOrder(node.left)
		PreOrder(node.right)
		fmt.Printf("no = %d name=%s\n", node.No, node.name)
	}
}

func main() {
	// 构建一个二叉树
	root := &Hero{
		No:   1,
		name: "宋江",
	}
	left := &Hero{
		No:   2,
		name: "吴勇",
	}
	right1 := &Hero{
		No:   3,
		name: "卢俊义",
	}
	right2 := &Hero{
		No:   4,
		name: "林冲",
	}

	root.left = left
	root.right = right1
	right1.right = right2

	PreOrder(root)

}
