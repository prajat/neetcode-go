package main

import "sort"

/*
There are n cars traveling to the same destination on a one-lane highway.
You are given two arrays of integers position and speed, both of length n.
position[i] is the position of the ith car (in miles)
speed[i] is the speed of the ith car (in miles per hour)
The destination is at position target miles.
A car can not pass another car ahead of it. It can only catch up to another car and then drive at the same speed as the car ahead of it.
A car fleet is a non-empty set of cars driving at the same position and same speed. A single car is also considered a car fleet.
If a car catches up to a car fleet the moment the fleet reaches the destination, then the car is considered to be part of the fleet.
Return the number of different car fleets that will arrive at the destination.
*/

/*
Sort cars by position descending, closest to the target first.
Compute time each car needs to reach the target.
If a car’s time is greater than the fleet ahead, it starts a new fleet.
Otherwise, it merges into the fleet ahead.
*/
func carFleet(target int, position, speed []int) int {
	n := len(position)
	type car struct {
		pos  int
		time float64
	}

	cars := make([]car, n)

	for i := 0; i < n; i++ {
		cars[i] = car{
			pos:  position[i],
			time: float64(target-position[i]) / float64(speed[i]),
		}
	}

	sort.Slice(cars, func(i, j int) bool {
		return cars[i].pos > cars[j].pos
	})

	fleets := 0
	maxTime := 0.0

	for i := 0; i < n; i++ {
		if cars[i].time > maxTime {
			fleets++
			maxTime = cars[i].time
		}
	}
	return fleets
}
