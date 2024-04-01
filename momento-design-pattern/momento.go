package main

type momento struct {
	state string
}

func (m *momento) getSavedState() string {
	return m.state
}
