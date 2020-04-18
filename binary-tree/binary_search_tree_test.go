package binary_tree

import (
"container/list"
"fmt"
"testing"
)

func Test_Recursive(t *testing.T){
	nums := list.New()
	nums.PushBack(3)
	nums.PushBack(2)
	nums.PushBack(9)
	nums.PushBack(nil)
	nums.PushBack(nil)
	nums.PushBack(10)
	nums.PushBack(nil)
	nums.PushBack(nil)
	nums.PushBack(8)
	nums.PushBack(nil)
	nums.PushBack(4)
	node := createBinaryTree(nums)
	fmt.Println("递归前序遍历:")
	PreOrderTravel(node)
	fmt.Println()
	fmt.Println("递归中序遍历:")
	inOrderTravel(node)
	fmt.Println()
	fmt.Println("递归后序遍历:")
	postOrderTravel(node)
	fmt.Println()
	fmt.Println("stack前序遍历:")
	preOrderTravelByStack(node)
	fmt.Println()
	fmt.Println("stack中序遍历:")
	inOrderTravelByStack(node)
	fmt.Println()
	fmt.Println("stack后序遍历:")//深度优先
	postOrderTravelByStack(node)
	fmt.Println()
	fmt.Println("层序遍历:")//广度有限
	levelOrderTravel(node)
	t.Log("a")
}

func createBinaryTree(list *list.List) *Node{
	node := new(Node)

	if  list.Len()<= 0{
		return node
	}
	data := list.Front()
	list.Remove(data)

	if data.Value != nil {
		node.value = data.Value
		node.leftChild = createBinaryTree(list)
		node.rightChild = createBinaryTree(list)
		//if data.Value == 2{
		//	fmt.Print(list)
		//}
	}
	return node
}