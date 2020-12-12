package main

import (
	"fmt"

	"github.com/Gelio/go-js-diagram/pkg/components"
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"github.com/hexops/vecty/event"
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
		&components.Box{
			Children: func() vecty.ComponentOrHTML {
				return elem.Div(
					elem.Paragraph(
						vecty.Text("Drag me, senpai"),
					),
					&counter{Text: "Count"},
				)
			},
		},
	)
}

type counter struct {
	vecty.Core
	count int
	Text  string `vecty:"prop"`
}

func (c *counter) Render() vecty.ComponentOrHTML {
	return elem.Button(
		vecty.Markup(
			event.Click(func(v *vecty.Event) {
				c.count++
				vecty.Rerender(c)
			}),
		),
		vecty.Text(fmt.Sprintf("%s: %d", c.Text, c.count)),
	)
}
