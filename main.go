package main

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"github.com/hexops/vecty/event"
	"math/rand"
	"sort"
	"time"
)

type page struct {
	vecty.Core
	mahjongTehai string
}

func RandomMahjongTile() string {
	rand.NewSource(time.Now().UnixNano())
	r := rand.Intn(0x1F022-0x1F000) + 0x1F000
	return string(rune(r))
}

func dealingTiles() string {
	mahjongTehai := ""
	for i := 0; i < 14; i++ {
		mahjongTehai += RandomMahjongTile()
	}

	tehaiRunes := []rune(mahjongTehai)

	sort.Slice(tehaiRunes, func(i, j int) bool {
		return tehaiRunes[i] > tehaiRunes[j]
	})

	sortedTehai := string(tehaiRunes)
	return sortedTehai
}

func (p *page) Render() vecty.ComponentOrHTML {
	if p.mahjongTehai == "" {
		p.mahjongTehai = "🀫🀫🀫🀫🀫🀫🀫🀫🀫🀫🀫🀫🀫🀫"
	}
	return elem.Body(
		vecty.Markup(
			vecty.Style("display", "flex"),
			vecty.Style("flex-direction", "column"),
			vecty.Style("justify-content", "center"),
			vecty.Style("align-items", "center"),
		),

		elem.Div(
			vecty.Markup(
				vecty.Style("font-family", `"Sawarabi Mincho", serif`),
				vecty.Style("font-weight", "400"),
				vecty.Style("font-style", "normal"),
				vecty.Style("font-size", "7vw"),
				vecty.Style("margin-top", "2%"),
				vecty.Style("margin-bottom", "2%"),
			),
			vecty.Text("天和チャレンジ"),
		),

		elem.Div(
			vecty.Markup(
				vecty.Style("font-family", `AppleSymbols`),
				vecty.Style("-webkit-user-select", "text"),
				vecty.Style("user-select", "text"),
				vecty.Style("font-size", "7vw"),
				vecty.Style("margin-top", "2%"),
				vecty.Style("margin-bottom", "2%"),
			),
			vecty.Text(p.mahjongTehai),
		),

		elem.Div(
			vecty.Text("下のボタンをクリックすると配牌されます"),
		),

		elem.Div(elem.Button(
			vecty.Markup(
				vecty.Class("button"),
				event.Click(func(_ *vecty.Event) {
					p.mahjongTehai = dealingTiles()
					vecty.Rerender(p)
				}),
			),
			vecty.Text("配牌"),
		),
		),
	)
}
func main() {
	vecty.SetTitle("天和チャレンジ")
	vecty.RenderBody(new(page))
}
