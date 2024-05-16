package ui

type Item struct {
	Title string
	Page  Page
	Exec  func()
}
