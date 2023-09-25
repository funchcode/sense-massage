package main

import "fmt"

func solution118667(queue1, queue2 []int) int {

	var queueSize = len(queue1)
	var totalQ1 = 0
	var totalQ2 = 0
	var q1OffSet = 0
	var q2OffSet = 0

	for _, q1 := range queue1 {
		totalQ1 += q1
	}

	for _, q2 := range queue2 {
		totalQ2 += q2
	}

	for i := 0; i < queueSize*4; i++ {

		if totalQ1 == totalQ2 {
			return i
		}

		if totalQ1 > totalQ2 {
			if q1OffSet/queueSize%2 == 0 {
				totalQ2 += queue1[q1OffSet%queueSize]
				totalQ1 -= queue1[q1OffSet%queueSize]
				q1OffSet++
			} else {
				totalQ2 += queue2[q1OffSet%queueSize]
				totalQ1 -= queue2[q1OffSet%queueSize]
				q1OffSet++
			}
		} else {
			if q2OffSet/queueSize%2 == 0 {
				totalQ1 += queue2[q2OffSet%queueSize]
				totalQ2 -= queue2[q2OffSet%queueSize]
				q2OffSet++
			} else {
				totalQ1 += queue1[q2OffSet%queueSize]
				totalQ2 -= queue1[q2OffSet%queueSize]
				q2OffSet++
			}
		}
	}

	return -1
}

func main() {
	//queue1 := []int{3, 2, 7, 2}
	//queue2 := []int{4, 6, 5, 1}

	//queue1 := []int{1, 1}
	//queue2 := []int{1, 5}

	queue1 := []int{101, 100}
	queue2 := []int{102, 103}

	//queue1 := []int{3, 3, 3, 3}
	//queue2 := []int{3, 3, 21, 3}

	//queue1 := []int{1, 2, 1, 2}
	//queue2 := []int{1, 10, 1, 2}

	fmt.Println(solution118667(queue1, queue2))
}

/**
201, 205
101-100/102-103

303, 103
101-100-102/103

202, 204
100-102/103-101

305, 101
100-102-103/101

205, 201
102-103/101-100

103, 303
103/101-100-102

204, 202
103-101/100-102
*/
