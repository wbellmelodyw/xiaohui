package binary_tree

import (
	"fmt"
	"github.com/golang-collections/collections/queue"
	stack "github.com/golang-collections/collections/stack"
)


type Node struct {
	value interface{}
	leftChild *Node
	rightChild *Node
}

/***************************递归遍历***************************/
//前序遍历
func PreOrderTravel(node *Node){
	if node.value == nil{
		return
	}
	fmt.Println(node.value)
	PreOrderTravel(node.leftChild)
	PreOrderTravel(node.rightChild)
}
//中序遍历
func inOrderTravel(node *Node){
	if node.value == nil{
		return
	}
	inOrderTravel(node.leftChild)
	fmt.Println(node.value)
	inOrderTravel(node.rightChild)
}
//后序遍历
func postOrderTravel(node *Node){
	if node.value == nil{
		return
	}
	postOrderTravel(node.leftChild)
	postOrderTravel(node.rightChild)
	fmt.Println(node.value)
}

/***************************递归遍历***************************/

/***************************递归遍历stack***************************/
func preOrderTravelByStack(node *Node){
	stack := stack.New()
	for stack.Len()!=0 || node.value != nil{
		for node.value != nil{
			fmt.Println(node.value)
			stack.Push(node)
			node = node.leftChild
		}
		if stack.Len() >0{
			last := stack.Pop().(*Node)
			node = last.rightChild
		}
	}
}


func inOrderTravelByStack(node *Node){
	stack := stack.New()
	for stack.Len()!=0 || node.value != nil{
		for node.value != nil{
			stack.Push(node)
			node = node.leftChild
		}
		if stack.Len() >0{
			last := stack.Pop().(*Node)
			fmt.Println(last.value)
			node = last.rightChild
		}
	}
}

func postOrderTravelByStack(node *Node){
	stack := stack.New()
	cur := node
	pre := new(Node)
	for stack.Len() >0 || cur.value != nil{

		for cur.value != nil {
			stack.Push(cur)
			cur = cur.leftChild
		}
		//
		cur = stack.Peek().(*Node)
		if cur.rightChild.value == nil || cur.rightChild.value == pre.value{
			stack.Pop()
			fmt.Println(cur.value)
			pre = cur
			//pop就要初始化了
			cur = new(Node)
		}else{
			cur = cur.rightChild
		}
	}
}

/***************************递归遍历stack***************************/


//层序列遍历(队列实现)
func levelOrderTravel(node *Node){
	q := queue.New()
	q.Enqueue(node)
	for q.Len() > 0{
		n := q.Dequeue().(*Node)
		fmt.Println(n.value)
		if n.leftChild.value != nil{
			q.Enqueue(n.leftChild)
		}
		if n.rightChild.value != nil{
			q.Enqueue(n.rightChild)
		}
	}
}