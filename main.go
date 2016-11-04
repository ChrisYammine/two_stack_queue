package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type stack []int

func (s stack) Push(x int) stack {
	return append(s, x)
}

func (s stack) Pop() (stack, int) {
	l := len(s)
	if l == 0 {
		return s, 0
	}
	return s[:l-1], s[l-1]
}

type Queue struct {
	newest stack
	oldest stack
}

func NewQueue() *Queue {
	return &Queue{make(stack, 0), make(stack, 0)}
}

// Pushes to the 'newest' stack
func (q *Queue) Enqueue(x int) {
	q.newest = q.newest.Push(x)
}

// Removes oldest element
func (q *Queue) Dequeue() {
	q.MaybeShiftStacks()
	s, _ := q.oldest.Pop()
	q.oldest = s
}

// Peeks oldest element
func (q *Queue) Peek() {
	q.MaybeShiftStacks()
	_, x := q.oldest.Pop()
	fmt.Println(x)
}

// Rebalance stacks
func (q *Queue) MaybeShiftStacks() {
	if len(q.oldest) == 0 {
		for len(q.newest) > 0 {
			new, x := q.newest.Pop()
			q.newest = new
			q.oldest = q.oldest.Push(x)
		}
	}
}

func main() {
	queue := NewQueue()
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	n_queries, _ := strconv.Atoi(scanner.Text())

	for i := 0; i < n_queries; i++ {
		scanner.Scan()
		cmd := scanner.Text()
		if len(cmd) > 1 {
			values := strings.Split(cmd, " ")
			x, _ := strconv.Atoi(values[1])
			queue.Enqueue(x)
		}
		switch cmd {
		case "2":
			queue.Dequeue()
		case "3":
			queue.Peek()
		}
	}
}
