func removeElement(nums []int, val int) int {
    var pointer int

    for _, n := range nums {
        if n != val {
            nums[pointer] = n
            pointer++
        }
    }

    return pointer
}
