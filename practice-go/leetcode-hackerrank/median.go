package main 

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    result := []int{}
    i, j := 0, 0

    // 1. Merge until one array is empty
    for i < len(nums1) && j < len(nums2) {
        if nums1[i] < nums2[j] {
            result = append(result, nums1[i])
            i++
        } else {
            result = append(result, nums2[j])
            j++
        }
    }

    // 2. Append remaining elements (crucial!)
    result = append(result, nums1[i:]...)
    result = append(result, nums2[j:]...)

    // 3. Calculate Median
    n := len(result)
    if n%2 == 0 {
        // Median is the average of the two middle elements
        return float64(result[n/2-1]+result[n/2]) / 2.0
    }
    // Median is the single middle element
    return float64(result[n/2])
}
