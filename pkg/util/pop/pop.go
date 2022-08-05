package main

import (
	"fmt"
)

func main() {
	// numA0 := []int{1, 7, 11}
	// numB0 := []int{2, 4, 6}
	// k0 := 3
	// fmt.Println(kSmallestPairs(numA0, numB0, k0))
	// numA1 := []int{1, 1, 2}
	// numB1 := []int{1, 2, 3}
	// k1 := 2
	// fmt.Println(kSmallestPairs(numA1, numB1, k1))
	numA2 := []int{1, 2}
	numB2 := []int{3}
	k2 := 3
	fmt.Println(kSmallestPairs2(numA2, numB2, k2))
	numA1 := []int{1, 1, 2}
	numB1 := []int{1, 2, 3}
	k1 := 10
	fmt.Println(kSmallestPairs2(numA1, numB1, k1))
}

type block struct {
	sum  int
	x    int
	y    int
	next *block
}

func (head *block) insert(new *block) {
	p := head
	prevP := head
	for {
		if p.sum > new.sum {
			prevP.next = new
			new.next = p
			break
		}
		prevP = p
		p = p.next
		// p, prevP = p.next, p
		if p == nil {
			prevP.next = new
			break
		}
	}
}

func (head *block) pop() {
	*head = &head.next
}

func kSmallestPairs2(nums1 []int, nums2 []int, k int) [][]int {
	res := [][]int{}
	// temp := math.MinInt
	if len(nums1) > k {
		nums1 = nums1[:k]
	}
	if len(nums2) > k {
		nums2 = nums2[:k]
	}
	//head := new(block)
	//head.sum = nums1[0] + nums2[0]
	//head.x = 0
	//head.y = 0
	head := &block{nums1[0] + nums2[0], 0, 0, nil}
	valid := make(map[[2]int]bool)
	valid[[2]int{0, 0}] = true

	for i := 0; i < k; i++ {
		if head != nil {
			temp := []int{nums1[head.x], nums2[head.y]}
			res = append(res, temp)
		} else {
			break
		}

		if _, ok := valid[[2]int{head.x + 1, head.y}]; !ok && head.x+1 < k && head.x+1 < len(nums1) {
			new1 := &block{nums1[head.x+1] + nums2[head.y], head.x + 1, head.y, nil}
			head.insert(new1)
			valid[[2]int{head.x + 1, head.y}] = true
		}
		if _, ok := valid[[2]int{head.x, head.y + 1}]; !ok && head.y+1 < k && head.y+1 < len(nums2) {
			new2 := &block{nums1[head.x] + nums2[head.y+1], head.x, head.y + 1, nil}
			head.insert(new2)
			valid[[2]int{head.x, head.y + 1}] = true
		}
		head.pop()
		//head = head.next
	}
	return res

}
