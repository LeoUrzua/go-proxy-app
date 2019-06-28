package middleware

import (
	"fmt"
	"github.com/kataras/iris"
	"sort"
)

type QueueNode struct {
	Domain   string
	Weight   int
	Priority int
}

var Queue []*QueueNode

type Repository interface {
	Read() []*QueueNode
}

func (q *QueueNode) Read() []*QueueNode {
	return MockQueue()
}

func MockQueue() []*QueueNode {
	return []*QueueNode{
		{
			Domain:   "alpha",
			Weight:   5,
			Priority: 5,
		},
		{
			Domain:   "beta",
			Weight:   5,
			Priority: 5,
		},
		{
			Domain:   "omega",
			Weight:   5,
			Priority: 1,
		},
	}
}

func InitQueue() {
	Queue = append(Queue, &QueueNode{})
}

func Handler(c iris.Context) {
	_ = c.GetHeader("domain")
	var repo Repository
	repo = &QueueNode{}

	fmt.Println(repo.Read())

	for _, row := range repo.Read() {
		Queue = append(Queue, row)
	}

	sortQueue()
	c.Next()
}

func sortQueue() {
	sort.Slice(Queue, func(i, j int) bool {
		return Queue[i].Weight*Queue[i].Priority < Queue[i].Weight*Queue[i].Priority
	})
}
