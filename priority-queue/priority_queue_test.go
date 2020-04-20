package priority_queue

import (
	"fmt"
	"testing"
)

func TestPriorityQueue(t *testing.T) {
	q := new(priorityQueue)
	q.enQueue(5)
	q.enQueue(3)
	q.enQueue(10)
	q.enQueue(2)
	q.enQueue(7)
	fmt.Println(q)
	fmt.Println(q.deQueue())
	fmt.Println(q)
	fmt.Println(q.deQueue())
	fmt.Println(q)
	fmt.Println(q.deQueue())
	fmt.Println(q)
}
