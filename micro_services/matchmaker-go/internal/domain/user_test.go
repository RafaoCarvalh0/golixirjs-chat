package domain

import (
	"fmt"
	"sync"
	"testing"
)

const checkMark = "\u2713"
const ballotX = "\u2717"

type randomMatcherMock struct{}

func (rm randomMatcherMock) randomMatch(queue *UserQueue) (Match, error) {
	return Match{User: User{ID: ""}, UserPair: User{ID: ""}}, fmt.Errorf("error")
}

func Test_CreateMatch(t *testing.T) {
	userCaller := User{ID: "bar"}
	queue := NewUserQueue()

	for _, id := range []string{"foo", "baz"} {
		queue.userMap[id] = User{ID: id}
	}

	t.Log("Given a user, a queue and a RandomMatcher.")
	{
		testDescription := "\t\tReturns a new match with user caller, a user pair that is not the caller, 'match created' status and err nil when there are users in the queue."
		newMatch, status, err := CreateMatch(&userCaller, &queue, &userCaller)
		if newMatch.User.ID != userCaller.ID ||
			(newMatch.UserPair.ID != "foo" && newMatch.UserPair.ID != "baz") ||
			status != "match created" ||
			err != nil {
			t.Fatal(testDescription, newMatch, status, err, ballotX)
		}
		t.Log(testDescription, checkMark)

		testDescription = "\t\tReturns 'waiting for a pair...' status whe there's only the user caller in queue."
		queue = NewUserQueue()
		queue.userMap[userCaller.ID] = User{ID: userCaller.ID}
		newMatch, status, err = CreateMatch(&userCaller, &queue, &userCaller)
		if status != "waiting for a pair..." || err != nil {
			t.Fatal(testDescription,
				newMatch, status, err, ballotX)
		}
		t.Log(testDescription, checkMark)

		testDescription = "\t\tReturns 'waiting for a pair...' status with unexpected error message when creating a random match fails"
		queue = NewUserQueue()
		queue.userMap[userCaller.ID] = User{ID: userCaller.ID}
		queue.userMap["foo"] = User{ID: "FOO"}
		newMatch, status, err = CreateMatch(&userCaller, &queue, randomMatcherMock{})
		if newMatch.User.ID != "" ||
			newMatch.UserPair.ID != "" ||
			status != "waiting for a pair..." ||
			err.Error() != "unexpected error while trying to match users" {
			t.Fatal(testDescription,
				newMatch, status, err, ballotX)
		}
		t.Log(testDescription, checkMark)
	}

}

func Test_ConcurrentCreateMatch(t *testing.T) {
	var wg sync.WaitGroup
	var statusWaitingCount int
	var queueLimitExceeded bool
	var usersMatchedCount int

	userIDs := []string{"u1", "u2", "u3", "u4", "u5", "u6", "u7", "u8", "u9", "u10", "u11"}

	queue := NewUserQueue()

	wg.Add(len(userIDs))
	for _, id := range userIDs {
		go func(userID string) {
			defer wg.Done()

			user := User{ID: userID}
			match, status, _ := user.CreateMatch(&queue)

			if match.User.ID != "" && match.UserPair.ID != "" {
				usersMatchedCount += 2
			}

			if status == "waiting for a pair..." {
				statusWaitingCount++
			}

			if len(queue.userMap) > 2 {
				queueLimitExceeded = true
			}

		}(id)
	}

	wg.Wait()

	t.Log("Given many users queued at the same time.")
	{
		testDescription := "\t\tIf an uneven number of calls happen, one user should remain in the queue."
		if len(queue.userMap) == 0 {
			t.Fatal(testDescription, statusWaitingCount, ballotX)
		}
		t.Log(testDescription, checkMark)

		testDescription = "\t\tAn even number of users will always match."
		if usersMatchedCount%2 != 0 {
			t.Fatal(testDescription, usersMatchedCount, ballotX)
		}
		t.Log(testDescription, checkMark)

		testDescription = "\t\tThe queue should not have more than 2 users at once."
		if queueLimitExceeded {
			t.Fatal(testDescription, queueLimitExceeded, ballotX)
		}
		t.Log(testDescription, checkMark)

		testDescription = "\t\tIf an uneven number o calls happened, half plus one of them should return a 'waiting for pair...' status."
		if statusWaitingCount != (len(userIDs)/2)+1 {
			t.Fatal(testDescription, statusWaitingCount, ballotX)
		}
		t.Log(testDescription, checkMark)
	}
}

func Test_addUserToQueue(t *testing.T) {
	user := User{ID: "bar"}
	queue := NewUserQueue()

	t.Log("Given a user and a user queue.")
	{
		testDescription := "\t\tAdds user to the provided queue when not present."
		user.addUserToQueue(&queue)
		_, userInQueue := queue.userMap[user.ID]
		if !userInQueue {
			t.Fatal(testDescription,
				ballotX, queue.userMap)
		}
		t.Log(testDescription, checkMark)

		testDescription = "\t\tDoes nothing whe the user is already present."
		user.addUserToQueue(&queue)
		_, userInQueue = queue.userMap[user.ID]
		if !userInQueue {
			t.Fatal(testDescription,
				ballotX, queue.userMap)
		}
		t.Log(testDescription, checkMark)
	}
}

func Test_removeUsersFromQueue(t *testing.T) {
	userToRemove1 := User{ID: "foo"}
	userToRemove2 := User{ID: "bar"}
	userToRemove3 := User{ID: "baz"}

	queue := UserQueue{
		mu: &sync.Mutex{},
		userMap: map[string]User{
			"foo":  userToRemove1,
			"u001": {ID: "u001"},
			"u002": {ID: "u002"},
			"u003": {ID: "u003"},
			"u004": userToRemove2,
			"u005": {ID: "u005"},
			"u006": {ID: "u006"},
			"u007": {ID: "u007"},
			"u008": {ID: "u008"},
			"u009": userToRemove3,
		},
	}

	t.Log("Given a user queue and a list of users.")
	{
		testDescription := "\t\tRemoves provided users from queue"
		removeUsersFromQueue(&queue, &userToRemove1, &userToRemove2, &userToRemove3)
		_, foundUser1 := queue.userMap[userToRemove1.ID]
		_, foundUser2 := queue.userMap[userToRemove2.ID]
		_, foundUser3 := queue.userMap[userToRemove3.ID]

		if foundUser1 || foundUser2 || foundUser3 {
			t.Fatal(testDescription,
				ballotX, foundUser1, foundUser2, foundUser3)

		}
		t.Log(testDescription, checkMark)

		testDescription = "\t\tDoes nothing to the queue when provided a list of users that were already removed"
		originalQueue := queue.userMap
		removeUsersFromQueue(&queue, &userToRemove1, &userToRemove1, &userToRemove1)
		if len(originalQueue) != len(queue.userMap) {
			t.Fatal(testDescription,
				ballotX)

		}
		t.Log(testDescription, checkMark)
	}
}

func Test_getRandomUserFromQueue(t *testing.T) {
	queue := NewUserQueue()

	t.Log("Given a user and a user queue.")
	{
		testDescription := "\t\tReturns a random user that is not the caller"
		userCaller := User{ID: "foo"}
		queue.userMap[userCaller.ID] = userCaller
		randomUser := User{ID: "bar"}
		queue.userMap[randomUser.ID] = randomUser

		for range 10 {
			returnedUser, _ := userCaller.getRandomUserFromQueue(&queue)

			if returnedUser == userCaller {
				t.Fatal(testDescription,
					ballotX, queue.userMap)
			}
		}
		t.Log(testDescription, checkMark)

		testDescription = "\t\tReturns a random user from the queue"
		queue = NewUserQueue()
		queue.userMap = map[string]User{
			"foo":  userCaller,
			"u001": {ID: "u001"},
			"u002": {ID: "u002"},
			"u003": {ID: "u003"},
			"u004": {ID: "u004"},
			"u005": {ID: "u005"},
			"u006": {ID: "u006"},
			"u007": {ID: "u007"},
			"u008": {ID: "u008"},
			"u009": {ID: "u009"},
		}
		returnedUser, _ := userCaller.getRandomUserFromQueue(&queue)
		if _, userIsInQueue := queue.userMap[returnedUser.ID]; !userIsInQueue {
			t.Fatal(testDescription,
				ballotX, queue.userMap)
		}
		t.Log(testDescription, checkMark)
	}

}
