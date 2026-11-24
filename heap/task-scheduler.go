package main

import "container/heap"

/*
You are given an array of CPU tasks tasks, where tasks[i] is an uppercase english character from A to Z. You are also given an integer n.
Each CPU cycle allows the completion of a single task, and tasks may be completed in any order.
The only constraint is that identical tasks must be separated by at least n CPU cycles, to cooldown the CPU.
Return the minimum number of CPU cycles required to complete all tasks.
*/

/*We start by calculating the frequency of each task and initialize a variable time to track the total processing time. The task frequencies are inserted into a Max-Heap. We also use a queue to store tasks along with the time they become available after the cooldown. At each step, if the Max-Heap is empty, we update time to match the next available task in the queue, covering idle time. Otherwise, we process the most frequent task from the heap, decrement its frequency, and if it's still valid, add it back to the queue with its next available time. If the task at the front of the queue becomes available, we pop it and reinsert it into the heap. */

type task struct {
	readyAt int
	freq    int
}

func leastInterval(tasks []byte, n int) int {
	frequency := make(map[byte]int, 26)
	for _, task := range tasks {
		frequency[task]++
	}

	h := &maxHeap{}
	heap.Init(h)

	queue := make([]task, 0)
	time := 0

	for _, v := range frequency {
		heap.Push(h, v)
	}

	for h.Len() > 0 || len(queue) > 0 {
		time++

		if h.Len() > 0 {
			freq := heap.Pop(h).(int) - 1
			if freq > 0 {
				queue = append(queue, task{readyAt: time + n, freq: freq})
			}
		}

		if len(queue) > 0 && queue[0].readyAt == time {
			heap.Push(h, queue[0].freq)
			queue = queue[1:]
		}
	}
	return time
}
