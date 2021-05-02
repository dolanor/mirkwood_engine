package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
)

type TextDrawer struct {
	screen *ebiten.Image
	fonts  *fontAssets
}

func (d *TextDrawer) Title(msg string, x, y int) {
	text.Draw(d.screen, msg, d.fonts.mplusTitleFont, x, y, color.White)
}

func (d *TextDrawer) Body(msg string, x, y int) {
	text.Draw(d.screen, msg, d.fonts.mplusNormalFont, x, y, color.White)
}

func (d *TextDrawer) Small(msg string, x, y int) {
	text.Draw(d.screen, msg, d.fonts.mplusSmallFont, x, y, color.White)
}

func (d *TextDrawer) Mini(msg string, x, y int) {
	text.Draw(d.screen, msg, d.fonts.mplusMiniFont, x, y, color.White)
}

func (d *TextDrawer) Notif(msg string, x, y int) {
	text.Draw(d.screen, msg, d.fonts.mplusNotificationFont, x, y, color.White)
}
