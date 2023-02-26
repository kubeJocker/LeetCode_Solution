package _355_design_twitter

import "container/heap"

type Twitter struct {
	users map[int]*user
	time  int
}

type user struct {
	userId    int
	followee  map[int]*user
	tweetList *tweetText
}

type tweetText struct {
	tweetId int
	time    int
	next    *tweetText
}

type tweetTextHeap []*tweetText

func (h tweetTextHeap) Len() int           { return len(h) }
func (h tweetTextHeap) Less(i, j int) bool { return h[i].time > h[j].time }
func (h tweetTextHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *tweetTextHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(*tweetText))
}

func (h *tweetTextHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// Constructor /** Initialize your data structure here. */
func Constructor() Twitter {
	t := Twitter{
		users: make(map[int]*user),
	}
	return t
}

// PostTweet /** Compose a new tweet. */
func (t *Twitter) PostTweet(userId int, tweetId int) {
	text := tweetText{
		tweetId: tweetId,
		time:    t.getTime(),
	}
	u := t.getUser(userId)
	text.next = u.tweetList
	u.tweetList = &text
}

// GetNewsFeed /** Retrieve the 10 most recent tweet ids in the user's news feed. Each item in the news feed must be posted by users who the user followed or by the user herself. Tweets must be ordered from most recent to least recent. */
func (t *Twitter) GetNewsFeed(userId int) []int {
	u := t.getUser(userId)

	texts := make([]*tweetText, 0, 1+len(u.followee))
	if u.tweetList != nil {
		texts = append(texts, u.tweetList)
	}
	for _, v := range u.followee {
		if v.tweetList != nil {
			texts = append(texts, v.tweetList)
		}
	}

	res := make([]int, 0, 10)

	textHeap := tweetTextHeap(texts)
	h := &textHeap
	heap.Init(h)
	for len(res) < 10 && h.Len() > 0 {
		text := heap.Pop(h).(*tweetText)
		res = append(res, text.tweetId)
		if text.next != nil {
			heap.Push(h, text.next)
		}
	}
	return res
}

// Follow /** Follower follows a followee. If the operation is invalid, it should be a no-op. */
func (t *Twitter) Follow(followerId int, followeeId int) {
	if followerId == followeeId {
		return
	}
	u1 := t.getUser(followerId)
	u2 := t.getUser(followeeId)
	u1.followee[followeeId] = u2
}

// Unfollow /** Follower unfollows a followee. If the operation is invalid, it should be a no-op. */
func (t *Twitter) Unfollow(followerId int, followeeId int) {
	if followerId == followeeId {
		return
	}
	u1 := t.getUser(followerId)
	_ = t.getUser(followeeId)
	delete(u1.followee, followeeId)
}

func (t *Twitter) getUser(userId int) *user {
	u, ok := t.users[userId]
	if !ok {
		u = &user{
			userId:   userId,
			followee: make(map[int]*user),
		}
		t.users[userId] = u
	}
	return u
}

func (t *Twitter) getTime() int {
	t.time++
	return t.time
}
