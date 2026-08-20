class Solution {
    /**
     * @param {number[]} nums
     * @return {number[]}
     */
    getConcatenation(nums: number[]): number[] {
        const ans = [...nums];
        const len = nums.length;
        for (let i = 0; i < len; i++) {
            ans[i + len] = nums[i]
        }
        return ans
    }
}
