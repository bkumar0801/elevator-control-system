package queue

import (
	"container/heap"
	"fmt"
	"sort"
	"strings"
)

/*
Goal ...
*/
type Goal struct {
	Floor    int
	Priority float64
	Index    int
}

/*
PriorityQueue ...
*/
type PriorityQueue []*Goal

/*
NewPriorityQueue ...
*/
func NewPriorityQueue() *PriorityQueue {
	pq := new(PriorityQueue)
	heap.Init(pq)
	return pq
}

/*
NewGoal ...
*/
func NewGoal(goalFloor int, prio float64) *Goal {
	return &Goal{Floor: goalFloor, Priority: prio}
}

/*
PushGoal ...
*/
func (pq *PriorityQueue) PushGoal(goal *Goal) {
	heap.Push(pq, goal)
	pq.update(goal, goal.Floor, goal.Priority)
	sort.Sort(pq)
}

/*
Len ...
*/
func (pq PriorityQueue) Len() int {
	return len(pq)
}

/*
Less ...
*/
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Priority < pq[j].Priority
}

/*
Swap ...
*/
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index = i
	pq[j].Index = j
}

/*
Push ...
*/
func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Goal)
	item.Index = n
	*pq = append(*pq, item)
}

/*
update ...
*/
func (pq *PriorityQueue) update(goal *Goal, value int, priority float64) {
	goal.Floor = value
	goal.Priority = priority
	heap.Fix(pq, goal.Index)
}

/*
Pop ...
*/
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.Index = -1
	*pq = old[0 : n-1]
	return item
}

/*
Peek ...
*/
func (pq PriorityQueue) Peek() *Goal {
	n := len(pq)
	goal := pq[n-1]
	return goal
}

/*
Find ...
*/
func (pq PriorityQueue) Find(value int) *Goal {
	for i := 0; i < pq.Len(); i++ {
		goal := pq[i]
		if goal.Floor == value {
			return goal
		}
	}
	return nil
}

/*
PrintGoals ...
*/
func (pq PriorityQueue) PrintGoals() string {

	resultString := make([]string, pq.Len()+1)
	if pq.Len() == 0 {
		resultString[0] = "No pending request"
	}
	for i := 0; i < pq.Len(); i++ {
		goal := pq[pq.Len()-1-i]
		resultString[i] = fmt.Sprintf("\n\t\t\tGoal #%v (floor: %v priority: %v)", i, goal.Floor, goal.Priority)
	}
	return strings.Join(resultString, " ")
}
