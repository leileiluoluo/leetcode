package main

import "fmt"

type TimeMap struct {
	timestamps map[string][]int
	values     map[int]string
}

func Constructor() TimeMap {
	return TimeMap{make(map[string][]int), make(map[int]string)}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	this.timestamps[key] = append(this.timestamps[key], timestamp)
	this.values[timestamp] = value
}

func (this *TimeMap) Get(key string, timestamp int) string {
	if timestamps, ok := this.timestamps[key]; ok {
		for i := len(timestamps) - 1; i >= 0; i-- {
			if timestamp >= timestamps[i] {
				return this.values[timestamps[i]]
			}
		}
	}
	return ""
}

func main() {
	obj := Constructor()
	obj.Set("foo", "bar", 1)
	fmt.Println(obj.Get("foo", 1))
	fmt.Println(obj.Get("foo", 3))
	obj.Set("foo", "bar2", 4)
	fmt.Println(obj.Get("foo", 4))
	fmt.Println(obj.Get("foo", 5))
}
