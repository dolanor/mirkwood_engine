package main

import (
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func (g *Game) handleState() {
	// Move selected player
	if inpututil.IsKeyJustPressed(ebiten.KeyW) {
		g.state.player[g.state.playerSelected].posy -= 70
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyS) {
		g.state.player[g.state.playerSelected].posy += 70
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyA) {
		g.state.player[g.state.playerSelected].posx -= 70
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyD) {
		g.state.player[g.state.playerSelected].posx += 70
	}
	// Move selected enemy
	if inpututil.IsKeyJustPressed(ebiten.KeyUp) {
		g.state.npc[g.state.enemySelected].posy -= 70
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyDown) {
		g.state.npc[g.state.enemySelected].posy += 70
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyLeft) {
		g.state.npc[g.state.enemySelected].posx -= 70
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyRight) {
		g.state.npc[g.state.enemySelected].posx += 70
	}
	// Toogle fullscreen
	if inpututil.IsKeyJustPressed(ebiten.KeyF) {
		g.config.fullscreen = !g.config.fullscreen
		ebiten.SetFullscreen(g.config.fullscreen)
	}
	// Player choice
	if inpututil.IsKeyJustPressed(ebiten.KeyP) {
		g.config.splash = false
		g.config.header_posx = 0
		go click_sound()
		if g.state.playerSelected < 1 {
			g.state.playerSelected += 1
		} else {
			g.state.playerSelected = 0
		}
	}
	// DM screen
	if inpututil.IsKeyJustPressed(ebiten.KeyU) {
		go click_sound()
		g.config.dm = !g.config.dm
	}
	// Link/Measure
	if inpututil.IsKeyJustPressed(ebiten.KeyL) {
		go click_sound()
		g.config.link = !g.config.link
	}
	// Select enemy
	if inpututil.IsKeyJustPressed(ebiten.KeyE) {
		go click_sound()
		if g.state.enemySelected < 3 {
			g.state.enemySelected += 1
		} else {
			g.state.enemySelected = 0
		}
	}
	// Show some debug info
	if inpututil.IsKeyJustPressed(ebiten.KeyG) {
		go click_sound()
		g.config.debug = !g.config.debug
	}
	// Toogle inventory
	if inpututil.IsKeyJustPressed(ebiten.KeyI) {
		go click_sound()
		g.config.header_posx = 0
		g.config.showInventory = !g.config.showInventory
	}
	// Quit
	if inpututil.IsKeyJustPressed(ebiten.KeyK) {
		os.Exit(0)
	}

	// Hidden area on the map
	if inpututil.IsKeyJustPressed(ebiten.KeyH) {
		g.config.hidden = !g.config.hidden
	}
	// Kill enemy (temporary)
	if inpututil.IsKeyJustPressed(ebiten.KeyKP1) {
		g.state.npc[0].alive = false
		g.config.link = false
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyKP2) {
		g.state.npc[1].alive = false
		g.config.link = false
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyKP3) {
		g.state.npc[2].alive = false
		g.config.link = false
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyKP4) {
		g.state.npc[3].alive = false
		g.config.link = false
	}
	// Remove health from enemies WIP
	/*if inpututil.IsKeyJustPressed(ebiten.KeyMinus) {
	  g.state.npc[g.state.enemySelected].hp_max = strconv(g.state.npc[g.state.enemySelected].hp_max) -= true
	} */

	// Change game round
	if inpututil.IsKeyJustPressed(ebiten.KeyN) {
		g.config.notification_posx = 1920
		go click_sound()
		if g.state.round < 2 {
			g.state.round += 1
		} else {
			g.state.round = 0
		}
	}
	// Dices be rollin'
	if inpututil.IsKeyJustPressed(ebiten.KeyR) { // roll dices
		go dice_sound()
		g.state.d20 = g.config.rand.Intn(20) + 1
		g.state.d4 = g.config.rand.Intn(4) + 1
		g.state.d6 = g.config.rand.Intn(6) + 1
		g.state.d8 = g.config.rand.Intn(8) + 1
	}

	// Next map - Disabled
	/*if inpututil.IsKeyJustPressed(ebiten.KeyN) {
	  go click_sound()
	  g.config.header_posx = 0
	  STATE_MAP=2
	  g.state.npc[0].alive=1
	  g.state.npc[1].alive=1
	  g.state.npc[2].alive=1
	  g.state.npc[3].alive=1
	  g.state.player[0].posx = 240
	  g.state.player[0].posy = 250
	  g.state.player[1].posx = 180
	  g.state.player[1].posy = 340
	  g.state.npc[0].posx = 1240
	  g.state.npc[0].posy = 650
	  g.state.npc[1].posx = 1180
	  g.state.npc[1].posy = 840
	  g.state.npc[2].posx = 1240
	  g.state.npc[2].posy = 250
	  g.state.npc[3].posx = 1580
	  g.state.npc[3].posy = 340
	} */
}
