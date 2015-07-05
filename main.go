package main

import "golang.org/x/tour/tree"
import "fmt"
import "sort"

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
  walk(t, ch)
  close(ch)
}

func walk(t *tree.Tree, ch chan int) {
  if t != nil {
    ch <- t.Value
    walk(t.Left, ch)
    walk(t.Right, ch)
  }
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	t1ch := make(chan int)
	t2ch := make(chan int)
	t1val := []int{}
	t2val := []int{}

	go Walk(t1, t1ch)
	go Walk(t2, t2ch)

	for i := range t1ch {
		t1val = append(t1val, i)
	}
	for i := range t2ch {
		t2val = append(t2val, i)
	}

	sort.Ints(t1val)
	sort.Ints(t2val)
	for i, _ := range t1val {
		if t1val[i] != t2val[i] {
			return false
		}
	}
	return true
}

func main() {
	k := 1
	t := tree.New(k)
	ch := make(chan int)
	go Walk(t, ch)
	for i := range ch {
		fmt.Println(i)
	}

	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))
}
