package ui

type Page struct {
	Title    string
	Actions  []Action
	Initiate func()
}
