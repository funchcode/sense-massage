package main

import "fmt"

type Point struct {
	array                   []int
	total, startIdx, endIdx int
}

func (p *Point) hasNext() bool {
	if len(p.array) > p.endIdx {
		if len(p.array) == p.endIdx+1 {
			p.endIdx++
		}
		return true
	}
	return false
}

func (p *Point) gap() int {
	return p.startIdx - p.endIdx
}

func (p *Point) getTotal() int {
	return p.total
}

func (p *Point) moveStartIdx() int {
	if len(p.array) > p.startIdx+1 {
		p.startIdx++
		if p.startIdx != 0 {
			p.total -= p.array[p.startIdx-1]
		}
		if p.startIdx == 0 || p.startIdx != p.endIdx {
			p.total += p.array[p.startIdx]
		}
		if p.startIdx > p.endIdx {
			p.endIdx++
			p.total += p.array[p.endIdx]
		}
	}
	fmt.Println("t = ", p.total)
	return p.startIdx
}

func (p *Point) moveEndIdx() int {
	if len(p.array) > p.endIdx {
		p.endIdx++
		p.total += p.array[p.endIdx]
	}
	fmt.Println("t = ", p.total)
	return p.endIdx
}

func solution1788702(sequence []int, k int) []int {

	optimalGap := -1

	point := &Point{
		array:    sequence,
		total:    0,
		startIdx: 0,
		endIdx:   0,
	}

	for point.hasNext() {
		if point.total == k {
			if optimalGap == -1 || point.gap() < optimalGap {
				optimalGap = point.gap()
			}
			point.moveEndIdx()
		} else if point.total > k {
			point.moveStartIdx()
		} else {
			point.moveEndIdx()
		}
	}

	return []int{}
}

func solution178870(sequence []int, k int) []int {

	optimalGap := -1
	t := 0
	startIdx := 0
	endIdx := 0

	fmt.Println("k = ", k)

	for _, seq := range sequence {

		if startIdx == 0 || startIdx != endIdx {
			t += seq
		}

		for endIdx < len(sequence) {

			if startIdx != endIdx {
				t += sequence[endIdx]
			}

			if startIdx > endIdx {
				endIdx++
			}

			fmt.Println("startIdx = ", startIdx)
			fmt.Println("endIdx = ", endIdx)
			fmt.Println("t = ", t)

			if t == k {
				gap := endIdx - startIdx

				if optimalGap == -1 || gap < optimalGap {
					fmt.Println(startIdx)
					fmt.Println(endIdx)
					optimalGap = gap
					break
				}

			} else if t < k {
				endIdx++
			} else if t > k {
				break
			}

		}

		startIdx++

		t -= seq

	}

	return []int{}
}

func main() {

	sequence01 := []int{1, 1, 1, 2, 3, 4, 5}
	k01 := 5

	solution178870(sequence01, k01)

}
