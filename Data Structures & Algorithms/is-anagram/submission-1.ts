class Solution {
    /**
     * @param {string} s
     * @param {string} t
     * @return {boolean}
     */
    isAnagram(s: string, t: string): boolean {
		if (s.length !== t.length) {
			return false;
		}

		const hashMap: Record<string, number> = {}

		for (let char of s) {
			if (!hashMap[char]) {
				hashMap[char] = 0;
			}

			hashMap[char] += 1;
		}

		for (let char2 of t) {
			if (!hashMap[char2]) {
				return false;
			}

			hashMap[char2] -= 1;
		}

		for (const [key, value] of Object.entries(hashMap)) {
			if (value > 0) {
				return false;
			};
		}

		return true
	}
}
