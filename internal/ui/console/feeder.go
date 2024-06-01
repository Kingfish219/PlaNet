package console

import (
	"github.com/Kingfish219/PlaNet/internal/ui"
	"github.com/Kingfish219/PlaNet/internal/ui/console/pages"
)

func FeedUI(console *ConsoleUI) *ui.Page {
	return pages.Main(console.dnsRepository)
}
