func twoSum(nums []int, target int) []int {
    indexes := make(map[int]int)

   for i, num := range nums {
        val := target - num

        if index, ok := indexes[val]; ok {
            return []int{index, i}
        }

        indexes[num] = i
   }

   return []int{0, 0}
}
