package components

import (
	"strconv"
	"syscall/js"

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
	mouseMoveFunc                    js.Func
	Children                         func() vecty.ComponentOrHTML `vecty:"slot"`
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
				event.MouseDown(b.onMouseDown),
				vecty.MarkupIf(b.listening,
					event.MouseUp(b.stopListeningForMouseMove),
				),
			),
			b.Children(),
		),
	)
}

func (b *Box) onMouseDown(v *vecty.Event) {
	b.stopListeningForMouseMove(v)

	b.listening = true
	b.dragStartBoxX = b.x
	b.dragStartBoxY = b.y
	b.dragStartMouseX = v.Value.Get("pageX").Int()
	b.dragStartMouseY = v.Value.Get("pageY").Int()

	b.mouseMoveFunc = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		e := args[0]

		b.x = e.Get("pageX").Int() - b.dragStartMouseX + b.dragStartBoxX
		b.y = e.Get("pageY").Int() - b.dragStartMouseY + b.dragStartBoxY
		vecty.Rerender(b)

		return nil
	})

	js.Global().Call("addEventListener", "mousemove", b.mouseMoveFunc)
	vecty.Rerender(b)
}

func (b *Box) stopListeningForMouseMove(v *vecty.Event) {
	b.listening = false

	js.Global().Call("removeEventListener", "mousemove", b.mouseMoveFunc)
	b.mouseMoveFunc.Release()

	vecty.Rerender(b)
}
