package proxy

import (
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func Test_UserListProxy(t *testing.T) {
	someDB := UserList{}
	random := rand.NewSource(2342342)
	for i := 0; i < 1000000; i++ {
		n := random.Int63()
		someDB = append(someDB, User{ID: int32(n)})
	}

	proxy := UserListProxy{
		SomeDB:        someDB,
		StackCache:    UserList{},
		StackCapacity: 2,
	}

	knownIDs := [3]int32{someDB[3].ID, someDB[4].ID, someDB[5].ID}

	t.Run("find user -empty cache", func(t *testing.T) {
		user, err := proxy.FindUser(knownIDs[0])
		assert.NoError(t, err)
		assert.Equal(t, someDB[3].ID, user.ID)
		assert.Equal(t, 1, len(proxy.StackCache))
		assert.False(t, proxy.DidDidLastSearchUsedCache)
	})

	t.Run("find user - user in cache", func(t *testing.T) {
		user, err := proxy.FindUser(knownIDs[0])
		assert.NoError(t, err)
		assert.Equal(t, someDB[3].ID, user.ID)
		assert.Equal(t, 1, len(proxy.StackCache))
		assert.True(t, proxy.DidDidLastSearchUsedCache)
	})

	t.Run("find user - cache is full", func(t *testing.T) {
		user1, err := proxy.FindUser(knownIDs[0])
		assert.NoError(t, err)
		user2, err := proxy.FindUser(knownIDs[1])
		assert.False(t, proxy.DidDidLastSearchUsedCache)
		user3, err := proxy.FindUser(knownIDs[2])
		assert.False(t, proxy.DidDidLastSearchUsedCache)

		assert.Equal(t, 2, len(proxy.StackCache))

		for _, user := range proxy.StackCache {
			if user.ID == user1.ID {
				t.Error("user1 should be gone from cache")
			}
		}

		for _, user := range proxy.StackCache {
			if user != user2 && user != user3 {
				t.Error("a non expected user found in cache")
			}
		}
	})
}
