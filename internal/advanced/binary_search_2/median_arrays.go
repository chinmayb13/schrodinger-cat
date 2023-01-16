package binarysearch2

import "math"

/*
Problem Description
There are two sorted arrays A and B of sizes N and M respectively.

Find the median of the two sorted arrays ( The median of the array formed by merging both the arrays ).

NOTE:
The overall run time complexity should be O(log(m+n)).
IF the number of elements in the merged array is even, then the median is the average of (n/2)th and (n/2+1)th element. For example, if the array is [1 2 3 4], the median is (2 + 3) / 2.0 = 2.5.


Problem Constraints
1 <= N + M <= 2*106

Input Format
The first argument is an integer array A of size N.
The second argument is an integer array B of size M.

Output Format
Return a decimal value denoting the median of two sorted arrays.

Example Input
Input 1:
A = [1, 4, 5]
B = [2, 3]

Input 2:
A = [1, 2, 3]
B = [4]

Example Output
Output 1:
3.0

Output 2:
2.5

Example Explanation
Explanation 1:
The median of both the sorted arrays will be 3.0.

Explanation 2:
The median of both the sorted arrays will be (2+3)/2 = 2.5.
*/

func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	arrSize1 := len(nums1)
	arrSize2 := len(nums2)

	if arrSize1 > arrSize2 {
		return FindMedianSortedArrays(nums2, nums1)
	}

	low, high := 0, arrSize1
	for low <= high {
		//cut1 determines how many elements to choose for left half from small array
		cut1 := low + (high-low)>>1
		//cut2 determines how many elements to choose for left half from larger array
		cut2 := (arrSize1+arrSize2+1)>>1 - cut1

		//left1 is last element of smaller array in left half
		left1 := getLeftPart(cut1, nums1)
		//left2 is last element of larger array in left half
		left2 := getLeftPart(cut2, nums2)
		//right1 is first element of smaller array in right half
		right1 := getRightPart(cut1+1, nums1)
		//right2 is first element of larger array in right half
		right2 := getRightPart(cut2+1, nums2)

		if left1 <= right2 && left2 <= right1 {
			if (arrSize1+arrSize2)&1 == 0 {
				return float64(max(left1, left2)+min(right1, right2)) / 2.0
			}
			return float64(max(left1, left2))
		} else if left1 > right2 {
			high = cut1 - 1
		} else {
			low = cut1 + 1
		}
	}
	return -1
}

func FindMedianSortedArraysAlt(nums1 []int, nums2 []int) float64 {

	var ans1 int
	nums1Size := len(nums1)
	nums2Size := len(nums2)
	least, maximum := getMinMax(nums1, nums2)
	low, high := least, maximum
	totalSize := nums1Size + nums2Size
	expectedCountLE := (totalSize + 1) / 2
	for low <= high {
		mid := low + (high-low)/2
		countLE := countLessThanEqualTo(mid, nums1, nums2)
		if countLE >= expectedCountLE {
			ans1 = mid
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	if (totalSize & 1) > 0 {
		return float64(ans1)
	}
	var ans2 int
	low, high = ans1, maximum
	for low <= high {
		mid := low + (high-low)/2
		countLE := countLessThanEqualTo(mid, nums1, nums2)
		if countLE >= (expectedCountLE + 1) {
			ans2 = mid
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return float64(ans1+ans2) / 2.0

}

func max(inp1, inp2 int) int {
	if inp1 > inp2 {
		return inp1
	}
	return inp2
}

func min(inp1, inp2 int) int {
	if inp1 < inp2 {
		return inp1
	}
	return inp2
}

func countLessThanEqualTo(val int, nums ...[]int) int {
	ans := 0
	for i := range nums {
		if len(nums[i]) > 0 {
			idx := -1
			low, high := 0, len(nums[i])-1
			for low <= high {
				mid := low + (high-low)/2
				if nums[i][mid] <= val {
					idx = mid
					low = mid + 1
				} else {
					high = mid - 1
				}
			}
			ans += idx + 1
		}
	}
	return ans
}

func getMinMax(inp ...[]int) (minVal, maxVal int) {
	minVal = 1e9 + 1
	maxVal = 0
	for i := range inp {
		if len(inp[i]) > 0 {
			minVal = min(minVal, inp[i][0])
			maxVal = max(maxVal, inp[i][len(inp[i])-1])
		}
	}
	return
}

func getLeftPart(cut int, arr []int) int {
	ans := math.MinInt
	if cut > 0 {
		return arr[cut-1]
	}
	return ans
}

func getRightPart(cut int, arr []int) int {
	ans := math.MaxInt
	if cut <= len(arr) {
		return arr[cut-1]
	}
	return ans
}
