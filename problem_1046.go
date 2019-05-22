// https://leetcode.com/problems/last-stone-weight/
package leetcode

import "container/heap"

type Heap []int

func (h Heap) Len() int {
    return len(h)
}

func (h Heap) Less(i, j int) bool {
    return h[i] > h[j]
}

func (h Heap) Swap(i, j int) {
    h[i], h[j] = h[j], h[i]
}

func (h *Heap) Push(x interface{}) {
    *h = append(*h, x.(int))
}

func (h *Heap) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n - 1]
    *h = old[:n - 1]
    return x
}

func Abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}

func lastStoneWeight(stones []int) int {
    h := &Heap{}
    heap.Init(h)

    for _, stone := range stones {
        heap.Push(h, stone)
    }

    for len(*h) > 1 {
        a := heap.Pop(h).(int)
        b := heap.Pop(h).(int)
        heap.Push(h, Abs(a - b))
    }
    
    return (*h)[0]
}
