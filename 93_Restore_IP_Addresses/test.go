func restoreIpAddresses(s string) []string {
    return restore(s, 4)
}

func restore(s string, require int) []string {
    if 1 == require {
        if len(s) > 1 && '0' == s[0] {
            return []string{}
        }
        if v, _ := strconv.Atoi(s); v < 256 {
            return []string{s}
        }
        return []string{}
    }

    var r []string
    for i := 1; i < 4 && i+require-1 <= len(s); i++ {
        prefix := s[:i]
        if v, _ := strconv.Atoi(prefix); v < 256 {
            for _, j := range restore(s[i:], require-1) {
                r = append(r, prefix+"."+j)
            }
        }
        if '0' == s[0] {
            break
        }
    }
    return r
}
