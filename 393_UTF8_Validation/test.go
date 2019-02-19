package main

import "fmt"

func validUtf8(data []int) bool {
	if 0 == len(data) {
		return true
	}
	switch {
	case 0x00 == 0x80&data[0]:
		return validUtf8(data[1:])
	case 0xC0 == 0xE0&data[0] &&
		len(data) > 1 && 0x80 == 0xC0&data[1]:
		return validUtf8(data[2:])
	case 0xE0 == 0xF0&data[0] &&
		len(data) > 2 && 0x80 == 0xC0&data[1] &&
		0x80 == 0xC0&data[2]:
		return validUtf8(data[3:])
	case 0xF0 == 0xF8&data[0] &&
		len(data) > 3 && 0x80 == 0xC0&data[1] &&
		0x80 == 0xC0&data[2] && 0x80 == 0xC0&data[3]:
		return validUtf8(data[4:])
	}
	return false
}

func main() {
	data := []int{240, 162, 138, 147, 145}
	fmt.Println(validUtf8(data))
}
