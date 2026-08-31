func longestCommonPrefix(strs []string) string {

    var output strings.Builder

    for i := 0; i < len(strs[0]); i++ {
        for _, s := range strs {
            if i == len(s) || s[i] != strs[0][i] {
                return output.String()
            }
        }
        output.WriteString(string(strs[0][i]))
    }

    return output.String()
}
