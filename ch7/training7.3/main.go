package main

import "fmt"

type tree struct {
	value int
	left, right *tree
}

func (t *tree) String() string {
	var values []int
	var f func (tr *tree)
	f = func (tr *tree) {
		if tr.left != nil {
			f(tr.left)
		}
		values = append(values, tr.value)
		if tr.right != nil {
			f(tr.right)
		}
	}
	f(t)

	return fmt.Sprintf("%v", values)
}

func Sort(values []int) {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	appendValues(values[:0], root)
}

func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func add(t *tree, value int) *tree {
	if t == nil {
		t = new (tree)
		t.value = value
		return t
	}

	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}

func main() {
	values := []int{2, 8, 5, 6, 1, 9, 7}
	Sort(values)
	fmt.Println(values)
}