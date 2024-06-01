package pages

import (
	"os"

	"github.com/Kingfish219/PlaNet/internal/ui"
)

func Exit() *ui.Page {
	return &ui.Page{
		Key:   "c_exit",
		Title: "Goodbye...",
		Initiate: func() {
			os.Exit(3)
		},
	}
}
