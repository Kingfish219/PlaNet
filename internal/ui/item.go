package ui

type Item struct {
	Key      string
	Title    string
	ShortKey string
	Page     *Page
	Exec     func()
}
