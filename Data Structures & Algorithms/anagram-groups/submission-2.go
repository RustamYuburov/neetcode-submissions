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
    // hash = make(map[string][]int)

    // for _, s := range strs {
    //     charInd = [30]rune{}

    //     for _, char := range s {
    //         charInd[char - 'a'] = char
    //     }

    //     hash[charInd] = append(hash[charInd], s)
    // }

    // arr = make([]string)

    // for _, st := range hash {
    //     arr = append(arr, st)
    // }

    // return arr
}
