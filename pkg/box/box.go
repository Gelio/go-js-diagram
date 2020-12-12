package box

import (
	"fmt"
	"strconv"

	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"github.com/hexops/vecty/event"
)

type Box struct {
	vecty.Core
	x, y                             int
	listening                        bool
	dragStartBoxX, dragStartBoxY     int
	dragStartMouseX, dragStartMouseY int
}

func (b *Box) Render() vecty.ComponentOrHTML {
	return elem.Div(
		vecty.Markup(
			vecty.Style("position", "relative"),
			vecty.Style("background-color", "#ffc0c0"),
			vecty.Style("width", "100%"),
			vecty.Style("height", "100%"),
		),
		elem.Div(
			vecty.Markup(
				vecty.Style("position", "absolute"),
				vecty.Style("top", strconv.Itoa(b.y)+"px"),
				vecty.Style("left", strconv.Itoa(b.x)+"px"),
				vecty.Style("cursor", "grab"),
				vecty.Style("user-select", "none"),
				event.MouseDown(func(v *vecty.Event) {
					b.listening = true
					b.dragStartBoxX = b.x
					b.dragStartBoxY = b.y
					b.dragStartMouseX = v.Value.Get("pageX").Int()
					b.dragStartMouseY = v.Value.Get("pageY").Int()
					fmt.Println("MouseDown")
					vecty.Rerender(b)
				}),
				vecty.MarkupIf(b.listening,
					event.MouseUp(func(v *vecty.Event) {
						b.listening = false
						fmt.Println("MOuseUp")
						vecty.Rerender(b)
					}),
					event.MouseLeave(func(v *vecty.Event) {
						b.listening = false
						fmt.Println("MouseLeave")
						vecty.Rerender(b)
					}),
					event.MouseMove(func(v *vecty.Event) {
						if !b.listening {
							return
						}

						b.x = v.Value.Get("pageX").Int() - b.dragStartMouseX + b.dragStartBoxX
						b.y = v.Value.Get("pageY").Int() - b.dragStartMouseY + b.dragStartBoxY
						fmt.Println("Moving", b.x, b.y)
						vecty.Rerender(b)
					}),
				),
			),
			elem.Paragraph(
				vecty.Text("Drag me, senpai"),
			),
		),
	)
}
