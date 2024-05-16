package ui

type Item struct {
	Key   string
	Title string
	Page  Page
	Exec  func()
}
