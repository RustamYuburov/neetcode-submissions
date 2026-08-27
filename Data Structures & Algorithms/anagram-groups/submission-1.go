func groupAnagrams(strs []string) [][]string {
	hashMap := make(map[[26]int][]string)

    for _, s := range strs {
        count := [26]int{}

        for _, c := range s {
            count[c-'a'] += 1
        }

        hashMap[count] = append(hashMap[count], s)
    }

    var result [][]string
    for _, group := range hashMap {
        result = append(result, group)
    }

    return result
}
