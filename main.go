package main

import (
	"github.com/Gelio/go-js-diagram/pkg/components"
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
)

func main() {
	vecty.SetTitle("JS Diagrams in Go")
	vecty.AddStylesheet("styles/styles.css")
	vecty.RenderBody(&pageView{})
}

type pageView struct {
	vecty.Core
}

func (p *pageView) Render() vecty.ComponentOrHTML {
	return elem.Body(
		&components.Box{},
	)
}
