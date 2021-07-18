package main

import (
	"encoding/json"
	"fmt"
)

type Tree struct {
	Value int
	Left  *Tree
	Right *Tree
}

type treeQueue []*Tree

func (ts treeQueue) Enqueue(t *Tree) treeQueue {
	ts = append(ts, t)
	return ts
}

func (ts treeQueue) Dequeue() (*Tree, treeQueue) {
	if len(ts) == 0 {
		return nil, ts
	}

	toReturn := ts[0]

	ts = append(ts[:0], ts[1:]...)

	return toReturn, ts
}

func (ts treeQueue) Max() *Tree {
	if len(ts) == 0 {
		return nil
	}

	max := ts[0]

	for _, iterTree := range ts {
		if iterTree.Value > max.Value {
			max = iterTree
		}
	}

	return max
}

func largestValuesInTreeRows(t *Tree) []int {
	queue := make(treeQueue, 0)

	if t == nil {
		return []int{}
	}

	queue = queue.Enqueue(t)
	toReturn := make([]int, 0)

	for len(queue) > 0 {
		toReturn = append(toReturn, queue.Max().Value)

		maxQueue := make(treeQueue, 0)
		for len(queue) > 0 {
			t, queue = queue.Dequeue()

			if t.Left != nil {
				maxQueue = maxQueue.Enqueue(t.Left)
			}
			if t.Right != nil {
				maxQueue = maxQueue.Enqueue(t.Right)
			}
		}

		queue = maxQueue
	}

	return toReturn
}

func main() {
	test1 := `{"value":1,"left":{"value":2,"left":null,"right":{"value":3,"left":null,"right":null}},"right":{"value":4,"left":{"value":5,"left":null,"right":null},"right":null}}`

	var tree *Tree
	json.Unmarshal([]byte(test1), &tree)

	fmt.Println(tree)

	// fmt.Println(traverseTreeDFS(tree))
	fmt.Println(largestValuesInTreeRows(tree))
}
