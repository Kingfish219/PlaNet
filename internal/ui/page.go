package ui

type Page struct {
	Key      string
	Title    string
	Items    []Item
	Initiate func()
}
