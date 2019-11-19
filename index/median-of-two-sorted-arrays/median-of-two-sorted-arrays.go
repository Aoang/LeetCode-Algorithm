package main

import "math"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	l := len(nums1) + len(nums2)
	k := l / 2
	if l%2 == 0 {
		return (find(nums1, 0, nums2, 0, k+1) + find(nums1, 0, nums2, 0, k)) / 2.0
	}
	return find(nums1, 0, nums2, 0, k+1)
}

func find(nums1 []int, n int, nums2 []int, m int, k int) float64 {
	if n >= len(nums1) {
		return float64(nums2[m+k-1])
	} else if m >= len(nums2) {
		return float64(nums1[n+k-1])
	}
	if k == 1 {
		a := math.Min(float64(nums1[n]), float64(nums2[m]))
		return a
	}

	p1 := n + k/2 - 1
	p2 := m + k - k/2 - 1
	var mid1, mid2 int
	if p1 >= len(nums1) {
		mid2 = nums2[p2]
		mid1 = mid2 + 1
	} else if p2 >= len(nums2) {
		mid1 = nums1[p1]
		mid2 = mid1 + 1
	} else {
		mid1 = nums1[p1]
		mid2 = nums2[p2]
	}

	if mid1 < mid2 {
		k = k - (p1 + 1 - n)
		if k == 0 {
			return float64(mid2)
		}
		return find(nums1, p1+1, nums2, m, k)
	} else {
		k = k - (p2 + 1 - m)
		if k == 0 {
			return float64(mid1)
		}
		return find(nums1, n, nums2, p2+1, k)
	}
}
