package main

type originator struct {
	state string
}

func (e *originator) createMomento() *momento {
	return &momento{state: e.state}
}

func (e *originator) restoreMomento(m *momento) {
	e.state = m.getSavedState()
}

func (e *originator) setState(state string) {
	e.state = state
}

func (e *originator) getState() string {
	return e.state
}
