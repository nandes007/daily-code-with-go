package main

type userIterator struct {
	index int
	users []*user
}

func (u *userIterator) hasNext() bool {
	return u.index < len(u.users)
}

func (u *userIterator) getNext() *user {
	if u.hasNext() {
		user := u.users[u.index]
		u.index++
		return user
	}
	return nil
}
