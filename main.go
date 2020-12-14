package main

import (
	"bufio"
	"fmt"
	"net/http"

	"github.com/Gelio/go-js-diagram/pkg/components"
	"github.com/Gelio/go-js-diagram/pkg/geolocation"
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

func (p *pageView) Mount() {
	coordsChan, err := geolocation.GetLocation()
	if err != nil {
		fmt.Println("Error when getting geolocation:", err)
		return
	}

	go func() {
		coords := <-coordsChan
		if coords.Err != nil {
			fmt.Println("Error when getting geolocation:", coords.Err)
			return
		}

		fmt.Println("Got geolocation:", coords.Latitude, coords.Longitude)
	}()

	go func() {
		resp, err := http.DefaultClient.Get("https://openlibrary.org/books/OL7353617M.json")
		if err != nil {
			fmt.Println("Cannot get response", err)
			return
		}

		r := bufio.NewReader(resp.Body)
		line, _, err := r.ReadLine()
		if err != nil {
			fmt.Println("Cannot read first line", err)
			return
		}

		fmt.Println("Got response:", string(line))
	}()
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
