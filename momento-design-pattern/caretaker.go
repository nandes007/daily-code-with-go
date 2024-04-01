package main

type caretaker struct {
	momentoArray []*momento
}

func (c *caretaker) addMomento(m *momento) {
	c.momentoArray = append(c.momentoArray, m)
}

func (c *caretaker) getMomento(index int) *momento {
	return c.momentoArray[index]
}
