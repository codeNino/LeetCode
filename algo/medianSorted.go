package algo

func mergeSortedArrays(nums1 []int, nums2 []int) []int {
	m, n := len(nums1), len(nums2)
	merged := make([]int, 0, m+n)

	i, j := 0, 0

	for i < m && j < n {
		if nums1[i] <= nums2[j] {
			merged = append(merged, nums1[i])
			i++
		} else {
			merged = append(merged, nums2[j])
			j++
		}
	}

	// Append remaining elements from nums1, if any
	for i < m {
		merged = append(merged, nums1[i])
		i++
	}

	// Append remaining elements from nums2, if any
	for j < n {
		merged = append(merged, nums2[j])
		j++
	}

	return merged
}

func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	merged := mergeSortedArrays(nums1, nums2)

	q := len(merged) / 2
	m := len(merged) % 2

	if m != 0 {
		return float64(merged[q])
	} else {
		return float64(merged[q-1]+merged[q]) / float64(2)
	}
}
