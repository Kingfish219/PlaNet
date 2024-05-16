package ui

type Page struct {
	Title    string
	Items    []Item
	Initiate func()
}
