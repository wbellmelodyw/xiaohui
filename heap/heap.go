package heap

//二叉树
//上一层  是下一层-1*2
/**
	(*
   *  *)
  (* * *) *
	最小堆上浮
*/
func upAdjust(a *[]int) {
	array := *a
	children := len(array) - 1
	parent := (children - 1) / 2
	temp := array[children]                    // 不需要交换 最后赋值就可以
	for children > 0 && array[parent] > temp { //满足这两个条件才交换
		array[children] = array[parent]
		children = parent
		parent = (children - 1) / 2
	}
	array[children] = temp
}

/**
	 0
   (1  2)
  (3 4) 5 *
	index * 2+1
*/
func downAdjust(a *[]int, parentIndex int) {
	array := *a
	aLen := len(array)
	childrenIndex := parentIndex*2 + 1
	temp := array[parentIndex]
	for childrenIndex < aLen {
		//取最小的孩子节点做置换
		if childrenIndex+1 < aLen && array[childrenIndex] > array[childrenIndex+1] {
			childrenIndex++
		}
		//父节点最小就不需要下沉了
		if temp < array[childrenIndex] {
			break
		}
		//置换1-赋值
		array[parentIndex] = array[childrenIndex]
		//置换2-更新parent
		parentIndex = childrenIndex
		//3-update children
		childrenIndex = parentIndex*2 + 1
	}
	array[parentIndex] = temp
}

func buildHeap(a *[]int) {
	//从最后一个父亲节点开始下沉
	array := *a
	for i := (len(array) - 2) / 2; i >= 0; i-- { //-1-1
		downAdjust(a, i)
	}
}
