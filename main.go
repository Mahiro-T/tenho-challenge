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

func RandomMahjongTile() int {
	rand.NewSource(time.Now().UnixNano())
	r := rand.Intn(0x1F022-0x1F000) + 0x1F000
	return r
}

func dealingTiles() string {
	mahjongHonorsWinds := ""
	mahjongHonorsDragons := ""
	mahjongManzu := ""
	mahjongSozu := ""
	mahjongPinzu := ""
	for {
		mahjongTehaiTemp := RandomMahjongTile()
		if mahjongTehaiTemp <= 0x1F003 {
			if strings.Count(mahjongHonorsWinds, string(rune(mahjongTehaiTemp))) < 4 {
				mahjongHonorsWinds += string(rune(mahjongTehaiTemp))
			} else {
				continue
			}
		} else if mahjongTehaiTemp <= 0x1F006 {
			if strings.Count(mahjongHonorsDragons, string(rune(mahjongTehaiTemp))) < 4 {
				mahjongHonorsDragons += string(rune(mahjongTehaiTemp))
			} else {
				continue
			}
		} else if mahjongTehaiTemp <= 0x1F00F {
			if strings.Count(mahjongManzu, string(rune(mahjongTehaiTemp))) < 4 {
				mahjongManzu += string(rune(mahjongTehaiTemp))
			} else {
				continue
			}
		} else if mahjongTehaiTemp <= 0x1F018 {
			if strings.Count(mahjongSozu, string(rune(mahjongTehaiTemp))) < 4 {
				mahjongSozu += string(rune(mahjongTehaiTemp))
			} else {
				continue
			}
		} else if mahjongTehaiTemp <= 0x1F021 {
			if strings.Count(mahjongPinzu, string(rune(mahjongTehaiTemp))) < 4 {
				mahjongPinzu += string(rune(mahjongTehaiTemp))
			} else {
				continue
			}
		}
		if len(mahjongHonorsWinds)+len(mahjongHonorsDragons)+len(mahjongManzu)+len(mahjongPinzu)+len(mahjongSozu) > 4*14 {
			break
		}
	}

	tehaiHonorsWinds := []rune(mahjongHonorsWinds)
	sort.Slice(tehaiHonorsWinds, func(i, j int) bool {
		return tehaiHonorsWinds[i] < tehaiHonorsWinds[j]
	})
	tehaiHonorsDragons := []rune(mahjongHonorsDragons)
	sort.Slice(tehaiHonorsDragons, func(i, j int) bool {
		return tehaiHonorsDragons[i] < tehaiHonorsDragons[j]
	})
	tehaiManzu := []rune(mahjongManzu)
	sort.Slice(tehaiManzu, func(i, j int) bool {
		return tehaiManzu[i] < tehaiManzu[j]
	})
	tehaiSozu := []rune(mahjongSozu)
	sort.Slice(tehaiSozu, func(i, j int) bool {
		return tehaiSozu[i] < tehaiSozu[j]
	})
	tehaiPinzu := []rune(mahjongPinzu)
	sort.Slice(tehaiPinzu, func(i, j int) bool {
		return tehaiPinzu[i] < tehaiPinzu[j]
	})

	sortedTehai := string(tehaiManzu) + string(tehaiSozu) + string(tehaiPinzu) + string(tehaiHonorsWinds) + string(tehaiHonorsDragons)
	return sortedTehai
}

func (p *page) Render() vecty.ComponentOrHTML {
	var twitterButton vecty.ComponentOrHTML

	if p.mahjongTehai == "" {
		p.mahjongTehai = "🀫🀫🀫🀫🀫🀫🀫🀫🀫🀫🀫🀫🀫🀫"
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
	vecty.SetTitle("天和チャレンジ")
	vecty.RenderBody(new(page))
}
