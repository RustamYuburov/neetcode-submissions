func twoSum(nums []int, target int) []int {
    indexHash := make(map[int]int)

    for i, num := range nums {
        diff := target - num

        if j, found := indexHash[diff]; found {
            return []int{j, i}
        }

        indexHash[num] = i
    }

    return []int{}
}
