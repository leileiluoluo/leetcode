package main

import "fmt"

func compareVersion(version1 string, version2 string) int {
	i, j := 0, 0
	v1, v2 := 0, 0
	for i < len(version1) || j < len(version2) {
		for ; i < len(version1); i++ {
			if '.' == version1[i] {
				i++
				break
			}
			if 0 == v1 && '0' == version1[i] {
				continue
			}
			v1 = 10*v1 + int(version1[i]-'0')
		}

		for ; j < len(version2); j++ {
			if '.' == version2[j] {
				j++
				break
			}
			if 0 == v2 && '0' == version2[j] {
				continue
			}
			v2 = 10*v2 + int(version2[j]-'0')
		}

		if v1 != v2 {
			break
		}
		v1, v2 = 0, 0
	}

	if v1 > v2 {
		return 1
	}
	if v1 < v2 {
		return -1
	}
	return 0
}

func main() {
	version1, version2 := "1.0", "1.1"
	fmt.Println(compareVersion(version1, version2))
}
