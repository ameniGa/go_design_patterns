package proxy

import "errors"

type User struct {
	ID int32
}

// UserFinder represents the Subject
type UserFinder interface {
	FindUser(id int32) (User, error)
}

// UserList is the real Subject
type UserList []User

func (ul *UserList) FindUser(id int32) (User, error) {
	for _, user := range *ul {
		if user.ID == id {
			return user, nil
		}
	}
	return User{}, errors.New("user not found")
}

// UserListProxy is the Proxy
// has the real subject
type UserListProxy struct {
	SomeDB                    UserList
	StackCache                UserList
	StackCapacity             int
	DidDidLastSearchUsedCache bool
}

func (up *UserListProxy) FindUser(id int32) (User, error) {
	user, err := up.StackCache.FindUser(id)
	if err == nil {
		up.DidDidLastSearchUsedCache = true
		return user, nil
	}

	user, err = up.SomeDB.FindUser(id)
	if err != nil {
		return User{}, err
	}
	up.addUserToStack(user)
	up.DidDidLastSearchUsedCache = false
	return user, nil
}

func (up *UserListProxy) addUserToStack(user User) {
	if len(up.StackCache) >= up.StackCapacity {
		up.StackCache = append(up.StackCache[1:], user)
		return
	}
	up.StackCache = append(up.StackCache, user)
}
