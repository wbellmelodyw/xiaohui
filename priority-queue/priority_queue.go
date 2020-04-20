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
	for childrenIndex > 0 && temp > q.queue[childrenIndex] {
		q.queue[parentIndex] = q.queue[childrenIndex]
		parentIndex = childrenIndex
		childrenIndex = (parentIndex - 1) / 2
	}
	q.queue[childrenIndex] = temp
}

func downAdjust() {

}
