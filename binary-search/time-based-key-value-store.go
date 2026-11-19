package main

import "sort"

/*Design a time-based key-value data structure that can store multiple values for the same key at different time stamps and retrieve the key's value at a certain timestamp.
Implement the TimeMap class:
TimeMap() Initializes the object of the data structure.
void set(String key, String value, int timestamp) Stores the key key with the value value at the given time timestamp.
String get(String key, int timestamp) Returns a value such that set was called previously, with timestamp_prev <= timestamp. If there are multiple such values, it returns the value associated with the largest timestamp_prev. If there are no values, it returns "". */

type entry struct {
	timestamp int
	value     string
}

type TimeMap struct {
	data map[string][]entry
}

func Constructor() TimeMap {
	return TimeMap{
		data: make(map[string][]entry),
	}
}

func (this *TimeMap) Set(key, value string, timestamp int) {
	this.data[key] = append(this.data[key], entry{timestamp: timestamp, value: value})
}

func (this *TimeMap) Get(key string, timestamp int) string {
	entries, exists := this.data[key]
	if !exists {
		return ""
	}

	// first index whose timestamp is greater than the requested one
	index := sort.Search(len(entries), func(i int) bool {
		return entries[i].timestamp > timestamp
	})

	// nothing at or before this timestamp
	if index == 0 {
		return ""
	}

	return entries[index-1].value
}
