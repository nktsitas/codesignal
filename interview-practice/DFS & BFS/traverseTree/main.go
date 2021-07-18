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

type cursor struct {
	node       *Tree
	parents    []*Tree
	directions []string
}

type treeStack []*Tree

func (ts treeStack) Push(t *Tree) treeStack {
	ts = append(ts, t)
	return ts
}

func (ts treeStack) Pop() *Tree {
	if len(ts) == 0 {
		return nil
	}

	toReturn := ts[len(ts)-1]

	ts = append(ts[:len(ts)-1], ts[len(ts):]...)

	return toReturn
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

func traverseTreeDFS(t *Tree) []int {
	toReturn := make([]int, 0)
	nodesMap := make(map[int][]*Tree)

	if t == nil {
		return []int{}
	}

	current := cursor{
		node:       t,
		parents:    make([]*Tree, 0),
		directions: make([]string, 0),
	}

	level := 0
	skipLeftCheck := false
	wentDown := true

	for current.node != nil {
		if len(current.directions) > 0 && !wentDown {
			// fmt.Println("why not - came from: ", current.directions[len(current.directions)-1])
			if current.directions[len(current.directions)-1] == "right" {
				// fmt.Println("up! premature")
				level--
				if len(current.parents) > 0 {
					current.node = current.parents[len(current.parents)-1]

					// pop
					// fmt.Printf("before: %v\n", len(current.parents))
					current.parents = append(current.parents[:len(current.parents)-1], current.parents[len(current.parents):]...)
					current.directions = append(current.directions[:len(current.directions)-1], current.directions[len(current.directions):]...)

					// fmt.Printf("after: %v\n", len(current.parents))
					// fmt.Printf("parents: %v\n", current.parents)
					// fmt.Printf("directions: %v\n", current.directions)
				} else {
					current.node = nil
				}

				continue
			} else {
				skipLeftCheck = true
				current.directions = append(current.directions[:len(current.directions)-1], current.directions[len(current.directions):]...)
			}
		}

		if nodesMap[level] == nil {
			nodesMap[level] = make([]*Tree, 0)
		}

		if wentDown {
			// fmt.Println("I'm: ", current.node.Value.(int))
			nodesMap[level] = append(nodesMap[level], current.node)
		}

		if current.node.Left != nil && !skipLeftCheck {
			// fmt.Println("left!")
			level++
			current.parents = append(current.parents, current.node)
			current.directions = append(current.directions, "left")
			current.node = current.node.Left
			wentDown = true
			continue
		}

		if current.node.Right != nil {
			// fmt.Println("right!")
			level++
			current.parents = append(current.parents, current.node)
			current.directions = append(current.directions, "right")
			current.node = current.node.Right
			skipLeftCheck = false
			wentDown = true
			continue
		}

		// fmt.Println("up!")
		wentDown = false
		level--
		if len(current.parents) > 0 {
			current.node = current.parents[len(current.parents)-1]

			// pop
			// fmt.Printf("before: %v\n", len(current.parents))
			current.parents = append(current.parents[:len(current.parents)-1], current.parents[len(current.parents):]...)
			// fmt.Printf("after: %v\n", len(current.parents))
			// fmt.Printf("parents: %v\n", current.parents)
			// fmt.Printf("directions: %v\n", current.directions)
		} else {
			current.node = nil
		}

	}

	for i := 0; i < len(nodesMap); i++ {
		iterLevelNodes := nodesMap[i]

		fmt.Printf("height: %v\n", i)
		fmt.Printf("level Nodes: %v\n", iterLevelNodes)

		for _, iterNode := range iterLevelNodes {
			fmt.Println(iterNode.Value)
			// toReturn = append(toReturn, iterNode.Value.(int))
			toReturn = append(toReturn, iterNode.Value)
		}

	}

	return toReturn
}

func traverseTreeBFS(t *Tree) []int {
	queue := make(treeQueue, 0)

	if t == nil {
		return []int{}
	}

	var current *Tree
	queue = queue.Enqueue(t)
	toReturn := make([]int, 0)

	for {

		current, queue = queue.Dequeue()
		// queue.Print()
		// fmt.Println("queue: ", queue)

		if current == nil {
			break
		}

		// fmt.Println("I'm: ", current.Value)
		toReturn = append(toReturn, current.Value)
		// fmt.Println("queue: ", queue)
		// queue.Print()

		if current.Left != nil {

			// fmt.Println("my left: ", current.Left.Value)
			queue = queue.Enqueue(current.Left)
			// queue.Print()
		}

		if current.Right != nil {
			// fmt.Println("my right: ", current.Right.Value)
			queue = queue.Enqueue(current.Right)
			// queue.Print()
		}

		// current, queue := queue.Dequeue()
		// queue.Print()
		// // fmt.Println("queue: ", queue)

		// fmt.Println("I'm: ", current.Value)
		// // fmt.Println("queue: ", queue)
		// queue.Print()

	}

	return toReturn
}

func main() {
	test1 := `{"value":1,"left":{"value":2,"left":null,"right":{"value":3,"left":null,"right":null}},"right":{"value":4,"left":{"value":5,"left":null,"right":null},"right":null}}`

	var tree *Tree
	json.Unmarshal([]byte(test1), &tree)

	fmt.Println(tree)

	// fmt.Println(traverseTreeDFS(tree))
	fmt.Println(traverseTreeBFS(tree))

}
