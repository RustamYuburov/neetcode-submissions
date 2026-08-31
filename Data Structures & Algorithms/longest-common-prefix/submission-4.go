func longestCommonPrefix(strs []string) string {
    var output strings.Builder

    base := strs[0]

    for i, _ := range base {
        for _, s := range strs {
            if i == len(s) || s[i] != base[i] {
                return output.String()
            }
        }
        output.WriteString(string(base[i]))
    }

    return output.String()
}
