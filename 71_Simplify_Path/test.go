package main

import "fmt"

func simplifyPath(path string) string {
	if len(path) < 2 {
		return path
	}
	canonicalPath := path[:1]
	for i := 1; i < len(path); {
		switch path[i-1 : i+1] {
		case "//", "./":
		case "/.":
			if i+1 < len(path) &&
				'/' != path[i+1] && '.' != path[i+1] {
				canonicalPath += "."
			}
		case "..":
			// valid directory name
			if i+1 < len(path) &&
				'/' != path[i+1] {
				canonicalPath += ".."
				i++
				for ; i < len(path) && '/' != path[i]; i++ {
					canonicalPath += string(path[i])
				}
				continue
			}

			// move directory up a level
			if len(canonicalPath) > 1 &&
				'/' == canonicalPath[len(canonicalPath)-1] {
				j := len(canonicalPath) - 2
				for ; j > 0; j-- {
					if '/' == canonicalPath[j] {
						break
					}
				}
				canonicalPath = canonicalPath[:j+1]
			}
		default:
			canonicalPath += string(path[i])
		}
		i++
	}

	// trim last '/'
	if len(canonicalPath) > 1 &&
		'/' == canonicalPath[len(canonicalPath)-1] {
		canonicalPath = canonicalPath[:len(canonicalPath)-1]
	}
	return canonicalPath
}

func main() {
	path := "/.aa/.hidden/.."
	fmt.Println(simplifyPath(path))
}
