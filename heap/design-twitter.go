package main

import "container/heap"

/*
Implement a simplified version of Twitter which allows users to post tweets, follow/unfollow each other, and view the 10 most recent tweets within their own news feed.
Users and tweets are uniquely identified by their IDs (integers).
Implement the following methods:
Twitter() Initializes the twitter object.
void postTweet(int userId, int tweetId) Publish a new tweet with ID tweetId by the user userId. You may assume that each tweetId is unique.
List<Integer> getNewsFeed(int userId) Fetches at most the 10 most recent tweet IDs in the user's news feed. Each item must be posted by users who the user is following or by the user themself. Tweets IDs should be ordered from most recent to least recent.
void follow(int followerId, int followeeId) The user with ID followerId follows the user with ID followeeId.
void unfollow(int followerId, int followeeId) The user with ID followerId unfollows the user with ID followeeId.
*/

type Tweet struct {
	tweetId int
	time    int
}

type MaxHeap []Tweet

func (h MaxHeap) Len() int {
	return len(h)
}

func (h MaxHeap) Less(i, j int) bool {
	return h[i].time > h[j].time
}

func (h MaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MaxHeap) Push(x any) {
	*h = append(*h, x.(Tweet))
}

func (h *MaxHeap) Pop() any {
	old := *h
	n := len(old)

	x := old[n-1]
	*h = old[:n-1]
	return x
}

type Twitter struct {
	time      int
	followMap map[int]map[int]bool
	tweetMap  map[int][]Tweet
}

func Constructor() Twitter {
	return Twitter{
		followMap: make(map[int]map[int]bool),
		tweetMap:  make(map[int][]Tweet),
	}
}

func (this *Twitter) PostTweet(userId, tweetId int) {
	this.time++
	this.tweetMap[userId] = append(this.tweetMap[userId], Tweet{
		tweetId: tweetId,
		time:    this.time,
	})
}

func (this *Twitter) GetNewsFeed(userId int) []int {
	h := &MaxHeap{}
	heap.Init(h)

	// include the user themself
	users := make(map[int]bool)
	users[userId] = true

	// include followers
	for followee := range this.followMap[userId] {
		users[followee] = true
	}

	// push all tweets from relevant users
	for uid := range users {
		for _, tweet := range this.tweetMap[uid] {
			heap.Push(h, tweet)
		}
	}

	result := make([]int, 0, 10)
	for h.Len() > 0 && len(result) < 10 {
		result = append(result, heap.Pop(h).(Tweet).tweetId)
	}

	return result
}

func (this *Twitter) Follow(followerId, followeeId int) {
	if followeeId == followerId {
		return
	}

	if this.followMap[followerId] == nil {
		this.followMap[followerId] = make(map[int]bool)
	}
	this.followMap[followerId][followeeId] = true
}

func (this *Twitter) Unfollow(followerId, followeeId int) {
	if this.followMap[followerId] != nil {
		delete(this.followMap[followerId], followeeId)
	}
}
