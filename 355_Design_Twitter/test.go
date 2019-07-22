package main

import (
	"fmt"
	"time"
)

type Twitter struct {
	tweets map[int][]tweet
	feeds  map[int][]tweet
	fans   map[int][]int
}

type tweet struct {
	id   int
	time time.Time
}

func Constructor() Twitter {
	tweets := make(map[int][]tweet)
	feeds := make(map[int][]tweet)
	fans := make(map[int][]int)

	return Twitter{tweets, feeds, fans}
}

func (this *Twitter) PostTweet(userId int, tweetId int) {
	time.Sleep(time.Nanosecond)

	now := time.Now()
	newTweet := tweet{tweetId, now}
	this.tweets[userId] = append(this.tweets[userId], newTweet)
	for _, followerId := range this.fans[userId] {
		this.feeds[followerId] = append(this.feeds[followerId], newTweet)
	}
	this.feeds[userId] = append(this.feeds[userId], tweet{tweetId, now})
}

func (this *Twitter) GetNewsFeed(userId int) []int {
	var feedIds []int
	feeds := this.feeds[userId]
	count := 0
	for i := len(feeds) - 1; i >= 0; i-- {
		if count >= 10 {
			break
		}
		feedIds = append(feedIds, feeds[i].id)
		count++
	}
	return feedIds
}

func (this *Twitter) Follow(followerId int, followeeId int) {
	if followerId == followeeId {
		return
	}

	found := false
	for _, item := range this.fans[followeeId] {
		if item == followerId {
			found = true
			break
		}
	}
	if !found {
		this.fans[followeeId] = append(this.fans[followeeId], followerId)
		this.feeds[followerId] = merge(this.feeds[followerId], this.tweets[followeeId])
	}
}

func merge(left, right []tweet) []tweet {
	var r []tweet
	if 0 == len(left) ||
		0 == len(right) {
		return append(left, right...)
	}
	i, j := 0, 0
	for i < len(left) && j < len(right) {
		for i < len(left) && left[i].time.Before(right[j].time) {
			r = append(r, left[i])
			i++
		}
		for j < len(right) && i < len(left) &&
			right[j].time.Before(left[i].time) {
			r = append(r, right[j])
			j++
		}
	}
	for i < len(left) {
		r = append(r, left[i])
		i++
	}
	for j < len(right) {
		r = append(r, right[j])
		j++
	}
	return r
}

func (this *Twitter) Unfollow(followerId int, followeeId int) {
	if followerId == followeeId {
		return
	}

	for i := 0; i < len(this.fans[followeeId]); i++ {
		item := this.fans[followeeId][i]
		if item == followerId {
			this.fans[followeeId] = append(this.fans[followeeId][:i], this.fans[followeeId][i+1:]...)

			for _, tweet := range this.tweets[followeeId] {
				for i, item := range this.feeds[followerId] {
					if item.id == tweet.id {
						this.feeds[followerId] = append(this.feeds[followerId][:i], this.feeds[followerId][i+1:]...)
						break
					}
				}
			}
			break
		}
	}
}

func main() {
	obj := Constructor()
	obj.PostTweet(1, 4)
	obj.PostTweet(2, 5)
	obj.Unfollow(1, 2)
	obj.Follow(1, 2)
	fmt.Println(obj.GetNewsFeed(1))
}
