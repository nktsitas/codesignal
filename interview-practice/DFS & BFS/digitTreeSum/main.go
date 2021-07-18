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

type treeStack []*Tree

func (ts treeStack) Push(t *Tree) treeStack {
	ts = append(ts, t)
	return ts
}

func (ts treeStack) Pop() (*Tree, treeStack) {
	if len(ts) == 0 {
		return nil, ts
	}

	toReturn := ts[len(ts)-1]

	ts = append(ts[:len(ts)-1], ts[len(ts):]...)

	return toReturn, ts
}

func (ts treeStack) Top() *Tree {
	if len(ts) == 0 {
		return nil
	}

	return ts[len(ts)-1]
}

func (ts treeStack) RegisterNum() int64 {
	multiplier := 1
	toReturn := 0

	var val *Tree

	for len(ts) > 0 {
		val, ts = ts.Pop()
		toReturn += multiplier * val.Value
		multiplier *= 10
	}

	return int64(toReturn)
}

func (ts treeStack) Print() {
	for _, iterTree := range ts {
		fmt.Printf("%v ", iterTree.Value)
	}
	fmt.Println()
}

// ---

type stringStack []string

func (ss stringStack) Push(s string) stringStack {
	ss = append(ss, s)
	return ss
}

func (ss stringStack) Pop() (string, stringStack) {
	if len(ss) == 0 {
		return "", ss
	}

	toReturn := ss[len(ss)-1]

	ss = append(ss[:len(ss)-1], ss[len(ss):]...)

	return toReturn, ss
}

func (ss stringStack) Top() string {
	if len(ss) == 0 {
		return ""
	}

	return ss[len(ss)-1]
}

func digitTreeSum(t *Tree) int64 {
	if t == nil {
		return 0
	}

	stack := make(treeStack, 0)
	directionsStack := make(stringStack, 0)

	toReturn := int64(0)
	wentUp := false

	stack = stack.Push(t)

	for len(stack) > 0 {
		if stack.Top().Left != nil && !wentUp {
			stack = stack.Push(stack.Top().Left)
			directionsStack = directionsStack.Push("left")
			wentUp = false
			continue
		}

		if stack.Top().Right != nil {
			stack = stack.Push(stack.Top().Right)
			directionsStack = directionsStack.Push("right")
			wentUp = false
			continue
		}

		if !wentUp {
			toReturn += stack.RegisterNum()
			stack.Print()
		}

		_, stack = stack.Pop()
		var direction string
		direction, directionsStack = directionsStack.Pop()
		for direction == "right" {
			_, stack = stack.Pop()
			direction, directionsStack = directionsStack.Pop()
		}

		wentUp = true
	}

	return toReturn
}

func main() {
	// test1 := `{"value":1,"left":{"value":2,"left":null,"right":{"value":3,"left":null,"right":null}},"right":{"value":4,"left":{"value":5,"left":null,"right":null},"right":null}}`
	test2 := `{"value":1,"left":{"value":3,"left":{"value":8,"left":null,"right":null},"right":{"value":7,"left":{"value":5,"left":null,"right":null},"right":{"value":4,"left":{"value":9,"left":null,"right":null},"right":{"value":2,"left":null,"right":null}}}},"right":{"value":4,"left":{"value":9,"left":{"value":8,"left":null,"right":null},"right":null},"right":{"value":5,"left":{"value":7,"left":null,"right":null},"right":{"value":5,"left":null,"right":null}}}}`
	// test3 := `{"value":1,"left":{"value":3,"left":{"value":5,"left":{"value":6,"left":{"value":9,"left":null,"right":null},"right":null},"right":null},"right":null},"right":{"value":5,"left":null}}`

	var tree *Tree
	// json.Unmarshal([]byte(test1), &tree)
	json.Unmarshal([]byte(test2), &tree)
	// json.Unmarshal([]byte(test3), &tree)

	fmt.Println(tree)

	fmt.Println(digitTreeSum(tree))
}
