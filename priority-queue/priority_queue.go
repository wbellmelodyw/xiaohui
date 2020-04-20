package priority_queue

type priorityQueue struct {
	queue []int
}

func (q *priorityQueue) enQueue(element int) {
	q.queue = append(q.queue, element)
	q.upAdjust()
}

/**
优先队列 从小到达 入队上浮 出队下降
*/
func (q *priorityQueue) upAdjust() {
	childrenIndex := len(q.queue) - 1
	parentIndex := (childrenIndex - 1) / 2
	temp := q.queue[childrenIndex]
	for childrenIndex > 0 && temp < q.queue[parentIndex] {
		q.queue[childrenIndex] = q.queue[parentIndex]
		childrenIndex = parentIndex
		parentIndex = parentIndex / 2
	}
	q.queue[childrenIndex] = temp
}

func (q *priorityQueue) deQueue() int {
	head := q.queue[0]
	length := len(q.queue)
	q.queue[0] = q.queue[length-1] //保证最大的开始下沉
	q.queue = q.queue[:length-1]
	q.downAdjust()
	return head
}

func (q *priorityQueue) downAdjust() {
	length := len(q.queue)
	parentIndex := 0
	childrenIndex := parentIndex*2 + 1
	temp := q.queue[parentIndex]
	for childrenIndex < length {
		if childrenIndex+1 < length && q.queue[childrenIndex] > q.queue[childrenIndex+1] {
			childrenIndex++
		}

		if temp < q.queue[childrenIndex] {
			break
		}
		q.queue[parentIndex] = q.queue[childrenIndex]
		parentIndex = childrenIndex
		childrenIndex = parentIndex*2 + 1
	}
	q.queue[parentIndex] = temp
}
