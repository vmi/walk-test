package main

import (
	. "github.com/lxn/walk/declarative"
)

func main() {
	size := Size{Width: 400, Height: 50}
	mw := MainWindow{
		Title:  "lxn Walk",
		Size:   size,
		Layout: VBox{},
		Children: []Widget{
			Label{Text: "Text1"},
			Label{Text: "Text2"},
			Label{Text: "Text3"},
			Label{Text: "Text4"},
			Label{Text: "Text5"},
			Label{Text: "Text6"},
			Label{Text: "Text7"},
			Label{Text: "Text8"},
			Label{Text: "Text9"},
			Label{Text: "Text10"},
		},
	}
	mw.Run()
}
