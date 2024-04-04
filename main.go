package main

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"github.com/hexops/vecty/event"
	"math/rand"
	"sort"
	"strings"
	"time"
)

type page struct {
	vecty.Core
	mahjongTehai string
}

func intSliceFromRange(min, max int) []int {
	s := make([]int, max-min+1)
	for i := range s {
		s[i] = min + i
	}
	return s
}

func setOfMahjongTiles() []int {
	const (
		beginOfMahjongTiles = 0x1F000
		endOfMahjongTiles   = 0x1F021
	)
	return intSliceFromRange(beginOfMahjongTiles, endOfMahjongTiles)
}

func repeatSlice[T any](size int, v []T) []T {
	s := make([]T, 0, size*len(v))
	for i := 0; i < size; i++ {
		s = append(s, v...)
	}
	return s
}

func mahjongTiles() []int {
	return repeatSlice(4, setOfMahjongTiles())
}

func orderOfMahjongTileTypes(tile int) int {
	type tileRange struct{ begin, end, order int }
	const orderOfUnknown = -1
	tileRanges := []tileRange{
		{begin: 0x1F000, end: 0x1F003, order: 3}, // Wind Tiles
		{begin: 0x1F004, end: 0x1F006, order: 4}, // Dragon Tiles
		{begin: 0x1F007, end: 0x1F00F, order: 0}, // Character Tiles
		{begin: 0x1F010, end: 0x1F018, order: 1}, // Bamboo Tiles
		{begin: 0x1F019, end: 0x1F021, order: 2}, // Circle Tiles
	}
	for _, r := range tileRanges {
		if r.begin <= tile && tile <= r.end {
			return r.order
		}
	}
	return orderOfUnknown
}

const countOfDealingTiles = 14

func dealingTiles() string {
	tiles := mahjongTiles()
	rand.Shuffle(len(tiles), func(i, j int) { tiles[i], tiles[j] = tiles[j], tiles[i] })
	dealingTiles := tiles[:countOfDealingTiles]
	sort.Ints(dealingTiles)
	sort.SliceStable(dealingTiles, func(i, j int) bool {
		return orderOfMahjongTileTypes(dealingTiles[i]) < orderOfMahjongTileTypes(dealingTiles[j])
	})
	dealingTilesStr := make([]string, len(dealingTiles))
	for i, tile := range dealingTiles {
		dealingTilesStr[i] = string(rune(tile))
	}
	return strings.Join(dealingTilesStr, "")
}

func (p *page) Render() vecty.ComponentOrHTML {
	var twitterButton vecty.ComponentOrHTML

	if p.mahjongTehai == "" {
		const mahjongTailBack = "\U0001F02B"
		p.mahjongTehai = strings.Repeat(mahjongTailBack, countOfDealingTiles)
	} else {
		twitterButton = elem.Div(
			elem.Anchor(
				vecty.Markup(
					vecty.Property("href", "https://twitter.com/intent/tweet?text="+p.mahjongTehai+"%0Ahttps://xn--0tr30i.xn--08j2a8s0b5d8jq891d.com/"+"&hashtags=天和チャレンジ"),
					vecty.Property("target", "_blank"),
					vecty.Property("rel", `"noopener noreferrer"`),
					vecty.Class("button"),
					vecty.Style("margin-top", "6%"),
					vecty.Style("margin-bottom", "3%"),
				),
				vecty.Text("𝕏に投稿する"),
			),
		)
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

		elem.Div(
			elem.Div(
				vecty.Markup(
					vecty.Style("font-family", `"Sawarabi Mincho", serif`),
					vecty.Style("font-weight", "400"),
					vecty.Style("font-style", "normal"),
				),

				elem.Button(
					vecty.Markup(
						vecty.Class("button"),
						vecty.Style("margin-top", "6%"),
						vecty.Style("margin-bottom", "3%"),

						event.Click(
							func(_ *vecty.Event) {
								p.mahjongTehai = dealingTiles()
								vecty.Rerender(p)
							},
						),
					),
					vecty.Text("配牌"),
				),
			)),
		twitterButton,
	)
}
func main() {
	rand.Seed(time.Now().UnixNano())
	vecty.SetTitle("天和チャレンジ")
	vecty.RenderBody(new(page))
}
