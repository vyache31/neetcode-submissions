func hasDuplicate(nums []int) bool {
    cntNums := make(map[int]int, len(nums))
    for _, num := range nums {
        cntNums[num]++
        if cntNums[num] > 1 {
            return true
        }
    }
    return false
}
