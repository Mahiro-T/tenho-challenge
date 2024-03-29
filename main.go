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
		elem.Heading1(vecty.Text("天和チャレンジ")),
		elem.Div(elem.Heading1(vecty.Text(p.mahjongTehai)),
			elem.Button(
				vecty.Markup(event.Click(func(_ *vecty.Event) {
					p.mahjongTehai = dealingTiles()
					vecty.Rerender(p)
				})),
				vecty.Text("配牌"),
			),
		))
}
func main() {
	vecty.SetTitle("天和チャレンジ")
	vecty.RenderBody(new(page))
}
