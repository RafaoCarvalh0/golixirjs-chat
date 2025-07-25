package domain

import (
	"testing"
)

const checkMark = "\u2713"
const ballotX = "\u2717"

func Test_alreadyQueued(t *testing.T) {
	existingUser := User{ID: "bar"}
	queue := map[string]User{}

	queue[existingUser.ID] = existingUser

	ids := []string{"foo", "baz"}
	for _, id := range ids {
		queue[id] = User{ID: id}
	}

	t.Log("Given a user and a user queue.")
	{
		testDescription := "\t\tReturns true when the user caller is present in the queue."
		userInQueue := existingUser.alreadyQueued(queue)
		if !userInQueue {
			t.Fatal(testDescription,
				ballotX, userInQueue)
		}
		t.Log(testDescription, checkMark)

		testDescription = "\t\tReturns false when the user caller is not in the queue."
		emptyQueue := map[string]User{}
		userInQueue = existingUser.alreadyQueued(emptyQueue)
		if userInQueue {
			t.Fatal(testDescription,
				ballotX, userInQueue)
		}
		t.Log(testDescription, checkMark)
	}

}

func Test_addUserToQueue(t *testing.T) {
	user := User{ID: "bar"}
	queue := map[string]User{}

	t.Log("Given a user and a user queue.")
	{
		testDescription := "\t\tAdds user to the provided queue when not present."
		user.addUserToQueue(queue)
		_, userInQueue := queue[user.ID]
		if !userInQueue {
			t.Fatal(testDescription,
				ballotX, queue)
		}
		t.Log(testDescription, checkMark)

		testDescription = "\t\tDoes nothing whe the user is already present."
		user.addUserToQueue(queue)
		_, userInQueue = queue[user.ID]
		if !userInQueue {
			t.Fatal(testDescription,
				ballotX, queue)
		}
		t.Log(testDescription, checkMark)
	}
}

func Test_removeUsersFromQueue(t *testing.T) {
	userToRemove1 := User{ID: "foo"}
	userToRemove2 := User{ID: "bar"}
	userToRemove3 := User{ID: "baz"}

	queue := map[string]User{
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
	}

	t.Log("Given a user queue and a list of users.")
	{
		testDescription := "\t\tRemoves provided users from queue"
		removeUsersFromQueue(queue, &userToRemove1, &userToRemove2, &userToRemove3)
		_, foundUser1 := queue[userToRemove1.ID]
		_, foundUser2 := queue[userToRemove2.ID]
		_, foundUser3 := queue[userToRemove3.ID]

		if foundUser1 || foundUser2 || foundUser3 {
			t.Fatal(testDescription,
				ballotX, foundUser1, foundUser2, foundUser3)

		}
		t.Log(testDescription, checkMark)

		testDescription = "\t\tDoes nothing to the queue when provided a list of users that were already removed"
		originalQueue := queue
		removeUsersFromQueue(queue, &userToRemove1, &userToRemove1, &userToRemove1)
		if len(originalQueue) != len(queue) {
			t.Fatal(testDescription,
				ballotX)

		}
		t.Log(testDescription, checkMark)
	}
}

func Test_getRandomUserFromQueue(t *testing.T) {
	queue := map[string]User{}

	t.Log("Given a user and a user queue.")
	{
		testDescription := "\t\tReturns a random user that is not the caller"
		userCaller := User{ID: "foo"}
		queue[userCaller.ID] = userCaller
		randomUser := User{ID: "bar"}
		queue[randomUser.ID] = randomUser

		for range 10 {
			returnedUser, _ := userCaller.getRandomUserFromQueue(queue)

			if returnedUser == userCaller {
				t.Fatal(testDescription,
					ballotX, queue)
			}
		}
		t.Log(testDescription, checkMark)

		testDescription = "\t\tReturns a random user from the queue"
		queue = map[string]User{
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
		returnedUser, _ := userCaller.getRandomUserFromQueue(queue)
		if _, userIsInQueue := queue[returnedUser.ID]; !userIsInQueue {
			t.Fatal(testDescription,
				ballotX, queue)
		}
		t.Log(testDescription, checkMark)
	}

}
