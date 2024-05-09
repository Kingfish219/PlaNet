package console

type Action struct {
	Title string
	Exec  func()
}
