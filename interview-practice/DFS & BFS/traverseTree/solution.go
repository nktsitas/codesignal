package main

type Tree struct {
	Value int
	Left  *Tree
	Right *Tree
}

// ---

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

func traverseTree(t *Tree) []int {
	if t == nil {
		return []int{}
	}

	// initialize
	var current *Tree
	queue := make(treeQueue, 0)
	queue = queue.Enqueue(t)
	toReturn := make([]int, 0)

	for {
		current, queue = queue.Dequeue()

		if current == nil {
			break
		}

		toReturn = append(toReturn, current.Value)

		if current.Left != nil {
			queue = queue.Enqueue(current.Left)
		}

		if current.Right != nil {
			queue = queue.Enqueue(current.Right)
		}
	}

	return toReturn
}
