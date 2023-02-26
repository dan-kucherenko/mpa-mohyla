package sorter

import "fmt"

func MergeSort(arr []int) {
	hlprArr := make([]int, len(arr))
	step := 0
	notSorted := true
	for notSorted {
		step++
		if step%2 != 0 {
			notSorted = mergeStep(arr, hlprArr)
			fmt.Println("Step", step, arr)
		} else {
			notSorted = mergeStep(hlprArr, arr)
			fmt.Println("Step", step, hlprArr)
		}
	}
	if step%2 == 0 {
		copyArr(hlprArr, arr, 0, len(arr)-1)
	}
}

func merge(arr, hlprArr []int, left, mid, right int) {
	i := left
	j := mid + 1
	for k := left; k <= right; k++ {
		if i > mid {
			hlprArr[k] = arr[j]
			j++
		} else {
			if j > right {
				hlprArr[k] = arr[i]
				i++
			} else {
				if arr[i] < arr[j] {
					hlprArr[k] = arr[i]
					i++
				} else {
					hlprArr[k] = arr[j]
					j++
				}
			}
		}
	}
}

func findPair(arr []int, left int, mid, right *int) bool {
	if left <= len(arr)-1 {
		*mid = left
		for *mid < len(arr)-1 && arr[*mid] <= arr[*mid+1] {
			*mid++
		}
		if *mid == len(arr)-1 {
			*right = *mid
			return false
		} else {
			*right = *mid + 1
			for *right < len(arr)-1 && arr[*right] <= arr[*right+1] {
				*right++
			}
			return true
		}
	}
	return false
}

func mergeStep(arr, hlprArr []int) bool {
	var left, mid, right int
	for findPair(arr, left, &mid, &right) {
		merge(arr, hlprArr, left, mid, right)
		left = right + 1
	}
	if left == 0 && right == len(arr)-1 {
		return false
	} else {
		if left <= len(arr)-1 {
			copyArr(arr, hlprArr, left, len(arr)-1)
		}
		return true
	}
}

func copyArr(arr, hlprArr []int, left, right int) {
	for i := left; i <= right; i++ {
		hlprArr[i] = arr[i]
	}
}
