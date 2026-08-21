class Solution {
    /**
     * @param {number[]} nums
     * @param {number} target
     * @return {number[]}
     */
    twoSum(nums: number[], target: number): number[] {
        // for (let i = 0; i < nums.length; i++) {
        //     for (let j = i + 1; i < nums.length; j++) {
        //         if (nums[i] + nums[j] === target) {
        //             return [i, j]
        //         }
        //     }
        // }

        const hash = {}

        for (let i = 0; i < nums.length; i++) {
            const secondNum = target - nums[i]
            
            if (typeof hash[secondNum] === 'number') {
                return [hash[secondNum], i]
            }
                
            hash[nums[i]] = i
        }
    }
}
