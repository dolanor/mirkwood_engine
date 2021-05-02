/* Started 12 April 2021
(C) Myu-Unix, 2021 - MIT Licensed - Assets used with fair use in mind, don't sue me */

package main

import (
	"embed"
	"fmt"
	"image/color"
	"log"
	"math"
	"math/rand"
	"strconv"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text"
)

type adventurer struct {
	name           string
	class          string
	race           string
	item1          string // Usually the "main" weapon
	item2          string
	item3          string
	item4          string
	item5          string
	posx           float64
	posy           float64
	hp_max         string
	STR            string
	DEX            string
	CON            string
	INT            string
	WIS            string
	CHA            string
	alignment      string
	ac_armor_class string
}

type enemy struct {
	name           string
	race           string
	hp_max         string
	ac_armor_class string
	item1          string // Usually the enemy weapon
	item2          string
	item3          string
	item4          string
	posx           float64
	posy           float64
	alive          bool
}

var (
	//go:embed fonts
	//go:embed images
	assetFS embed.FS

	MyConfig       adventurer // TEST
	player         [2]adventurer
	npc            [4]enemy
	keyStates      = map[ebiten.Key]int{}
	cmd_run        []byte
	engine_version = "Mirkwood Engine 0.7.0 (Prototype)"
	engine_text    = "Written in Go + Ebiten // Not all those who wander are lost"
)

type config struct {
	fullscreen        bool
	hidden            bool
	showInventory     bool
	link              bool
	dm                bool
	debug             bool
	splash            bool
	header_posx       float64
	notification_posx float64
	rand              *rand.Rand
}

type state struct {
	playerSelected int
	enemySelected  int
	round          int
	d20            int
	d4             int
	d6             int
	d8             int
}

func newConfig() config {
	randSrc := rand.NewSource(time.Now().UnixNano())
	return config{
		fullscreen:        true,
		hidden:            true,
		showInventory:     true,
		dm:                false,
		debug:             false,
		link:              false,
		splash:            true,
		header_posx:       0,
		notification_posx: 1920,
		rand:              rand.New(randSrc),
	}
}

func newState() state {
	return state{
		playerSelected: 0,
		enemySelected:  0,
		round:          0,
		d20:            20,
		d4:             4,
		d6:             6,
		d8:             8,
	}
}

func init() {
	// TODO : Replace by json file config
	player[0] = adventurer{name: "Myu", class: "Level 1 Ranger", race: "Elf", item1: "Elven Shortbow +1 (45m/1d6)", item2: "Elvish Dagger +1 (1d4)", item3: "Leather Armor (AC11)", item4: "Lembas (5)", item5: "Camping supplies", posx: 630, posy: 210, hp_max: "15 HP", STR: "STR 12", DEX: "DEX 14", CON: "CON 13", INT: "INT 12", WIS: "WIS 13", CHA: "CHA 10", alignment: "Chaotic good", ac_armor_class: "AC 13"}
	player[1] = adventurer{name: "Dolph", class: "Level 1 Druid", race: "Elf", item1: "Staff of Adornment (1d6 - 1d8)", item2: "Rope", item3: "Healing Herbs", posx: 560, posy: 280, hp_max: "12 HP", STR: "STR 8", DEX: "DEX 10", CON: "CON 7", INT: "INT 15", WIS: "WIS 14", CHA: "CHA 12", alignment: "Lawful good", ac_armor_class: "AC 10"}
	npc[0] = enemy{name: "Ghaz", race: "Level 1 Goblin", posx: 1200, posy: 700, hp_max: "8 HP", ac_armor_class: "AC 5", item1: "Club (1d4)", alive: true}
	npc[1] = enemy{name: "Dhurg", race: "Level 2 Goblin Warg Rider", posx: 1100, posy: 750, hp_max: "10 HP", ac_armor_class: "AC 7", item1: "Hand-Axe (1d6)", alive: true}
	npc[2] = enemy{name: "Dorg", race: "Level 1 Skeleton Archer", posx: 1150, posy: 800, hp_max: "5 HP", ac_armor_class: "AC 6", item1: "Longbow (1d6)", alive: true}
	npc[3] = enemy{name: "Dakh", race: "Level 1 Skeleton", posx: 1150, posy: 700, hp_max: "6 HP", ac_armor_class: "AC 5", item1: "Hand-Axe (1d6)", alive: true}
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Images options
	opAdventurer1 := &ebiten.DrawImageOptions{}
	opAdventurer2 := &ebiten.DrawImageOptions{}
	opEnemy1 := &ebiten.DrawImageOptions{}
	opEnemy2 := &ebiten.DrawImageOptions{}
	opEnemy3 := &ebiten.DrawImageOptions{}
	opEnemy4 := &ebiten.DrawImageOptions{}
	opBackground := &ebiten.DrawImageOptions{}
	opInventory := &ebiten.DrawImageOptions{}
	opHeader := &ebiten.DrawImageOptions{}
	opDice20 := &ebiten.DrawImageOptions{}
	opDice4 := &ebiten.DrawImageOptions{}
	opDice6 := &ebiten.DrawImageOptions{}
	opDice8 := &ebiten.DrawImageOptions{}
	opHide := &ebiten.DrawImageOptions{}
	opNotification := &ebiten.DrawImageOptions{}
	opAdventurer1.GeoM.Translate(player[0].posx, player[0].posy)
	opAdventurer2.GeoM.Translate(player[1].posx, player[1].posy)
	opEnemy1.GeoM.Translate(npc[0].posx, npc[0].posy)
	opEnemy2.GeoM.Translate(npc[1].posx, npc[1].posy)
	opEnemy3.GeoM.Translate(npc[2].posx, npc[2].posy)
	opEnemy4.GeoM.Translate(npc[3].posx, npc[3].posy)
	opHeader.GeoM.Translate(g.config.header_posx, 32)
	opInventory.GeoM.Translate(g.config.header_posx, 220)
	opDice20.GeoM.Translate(16, 120)
	opDice4.GeoM.Translate(16, 230)
	opDice6.GeoM.Translate(16, 340)
	opDice8.GeoM.Translate(16, 450)
	opHide.GeoM.Translate(1041, 629)
	opNotification.GeoM.Translate(g.config.notification_posx, 16)
	drawer := TextDrawer{screen: screen, fonts: g.assets.fonts}

	// Draw images
	if g.config.splash { // This shows the splashscreen
		screen.DrawImage(g.assets.images.splashImage, opBackground)
		drawer.Title("~ Into Mirkwood ~", 730, 400)
		drawer.Body("A short tabletop tutorial campaign", 725, 575)
		drawer.Body("Myu & Dolph <3", 865, 650)
		drawer.Small("Press 'p' to start", 1700, 1000)
		return
	}

	// Map background handler
	screen.DrawImage(g.assets.images.background1Image, opBackground)
	if g.config.notification_posx > 32 {
		g.config.notification_posx -= 128
	}
	screen.DrawImage(g.assets.images.notificationImage, opNotification)

	// Draw a line between selected player and target (if alive)
	if g.config.link {
		if npc[g.state.enemySelected].alive {
			ebitenutil.DrawLine(screen, player[g.state.playerSelected].posx+16, player[g.state.playerSelected].posy+32, npc[g.state.enemySelected].posx+16, npc[g.state.enemySelected].posy+32, color.RGBA{255, 128, 0, 255})
			ebitenutil.DrawLine(screen, player[g.state.playerSelected].posx+17, player[g.state.playerSelected].posy+33, npc[g.state.enemySelected].posx+17, npc[g.state.enemySelected].posy+33, color.RGBA{255, 128, 0, 255})
			a := int(npc[g.state.enemySelected].posx) - int(player[g.state.playerSelected].posx)
			b := int(npc[g.state.enemySelected].posy) - int(player[g.state.playerSelected].posy)
			// Rough distance in "ft" from pixels
			distance := math.Sqrt(float64((a*a))+float64((b*b))) / 10
			drawer.Small(string(strconv.Itoa(int(distance))), int(distance*5+player[g.state.playerSelected].posx), int(distance*5+player[g.state.playerSelected].posy))
			drawer.Small("ft", int(distance*5+player[g.state.playerSelected].posx+30), int(distance*5+player[g.state.playerSelected].posy))
		}
	}
	// Drawing dices and values
	screen.DrawImage(g.assets.images.dice20Image, opDice20)
	screen.DrawImage(g.assets.images.dice4Image, opDice4)
	screen.DrawImage(g.assets.images.dice6Image, opDice6)
	screen.DrawImage(g.assets.images.dice8Image, opDice8)
	drawer.Body(string(strconv.Itoa(g.state.d20)), 140, 200)
	drawer.Body(string(strconv.Itoa(g.state.d4)), 140, 300)
	drawer.Body(string(strconv.Itoa(g.state.d6)), 140, 400)
	drawer.Body(string(strconv.Itoa(g.state.d8)), 140, 500)
	// Drawing adventurers/players
	screen.DrawImage(g.assets.images.adventurer1Image, opAdventurer1)
	screen.DrawImage(g.assets.images.adventurer2Image, opAdventurer2)
	// Player "token" data
	drawer.Small(string(player[g.state.playerSelected].name), int(player[g.state.playerSelected].posx+48), int(player[g.state.playerSelected].posy))
	// TEST - JSON gathered
	//text.Draw(screen, string(MyConfig.name), mplusSmallFont, int(player[g.state.playerSelected].posx+48), int(player[g.state.playerSelected].posy), color.White)
	drawer.Mini(string(player[g.state.playerSelected].hp_max), int(player[g.state.playerSelected].posx+64), int(player[g.state.playerSelected].posy+18))
	drawer.Mini(string(player[g.state.playerSelected].ac_armor_class), int(player[g.state.playerSelected].posx+72), int(player[g.state.playerSelected].posy+32))
	drawer.Mini(string(player[g.state.playerSelected].item1), int(player[g.state.playerSelected].posx+72), int(player[g.state.playerSelected].posy+46))

	if g.config.debug {
		drawer.Body(engine_version, 40, 960)
		drawer.Mini(engine_text, 40, 982)
		drawer.Small("PLAYER : ", 32, 560)
		drawer.Small("ENEMY : ", 32, 600)
		drawer.Small("ROUND : ", 32, 640)
		drawer.Small(strconv.Itoa(g.state.playerSelected+1), 156, 560)
		drawer.Small(strconv.Itoa(g.state.enemySelected+1), 156, 600)
		drawer.Small(strconv.Itoa(g.state.round), 156, 640)
	}

	// If NPC is alive, draw it
	if npc[0].alive {
		screen.DrawImage(g.assets.images.enemy1Image, opEnemy1)
	}
	if npc[1].alive {
		screen.DrawImage(g.assets.images.enemy2Image, opEnemy2)
	}
	if npc[2].alive {
		screen.DrawImage(g.assets.images.enemy3Image, opEnemy3)
	}
	if npc[3].alive {
		screen.DrawImage(g.assets.images.enemy4Image, opEnemy4)
	}

	// INVENTORY CARD
	if g.config.showInventory {
		// Show header animation
		if g.config.header_posx < 1450 {
			g.config.header_posx += 290
		}
		// Show player header image
		if g.state.playerSelected == 0 {
			screen.DrawImage(g.assets.images.header1Image, opHeader)
		} else {
			screen.DrawImage(g.assets.images.header2Image, opHeader)
		}
		screen.DrawImage(g.assets.images.inventoryImage, opInventory)

		drawer.Body(string(player[g.state.playerSelected].name), 1480, 82)
		drawer.Small(string(player[g.state.playerSelected].class), 1480, 114)
		drawer.Small(string(player[g.state.playerSelected].hp_max), 1480, 146)
		drawer.Small(string(player[g.state.playerSelected].ac_armor_class), 1540, 146)
		drawer.Small(string(player[g.state.playerSelected].alignment), 1490, 178)
		drawer.Body("-- INVENTORY --", 1500, 232)
		//text.Draw(screen, "Range 3-18", mplusMiniFont, 1720, 50, color.White)
		drawer.Mini(string(player[g.state.playerSelected].STR), 1770, 70)
		drawer.Mini(string(player[g.state.playerSelected].DEX), 1770, 90)
		drawer.Mini(string(player[g.state.playerSelected].CON), 1770, 110)
		drawer.Mini(string(player[g.state.playerSelected].INT), 1770, 130)
		drawer.Mini(string(player[g.state.playerSelected].WIS), 1770, 150)
		drawer.Mini(string(player[g.state.playerSelected].CHA), 1770, 170)
		drawer.Small(string(player[g.state.playerSelected].item1), 1532, 270)
		drawer.Small(string(player[g.state.playerSelected].item2), 1532, 310)
		drawer.Small(string(player[g.state.playerSelected].item3), 1532, 350)
		drawer.Small(string(player[g.state.playerSelected].item4), 1532, 390)
		drawer.Small(string(player[g.state.playerSelected].item5), 1532, 430)
		//text.Draw(screen, string(player[g.state.playerSelected].item6), mplusSmallFont, 1532, 470, color.White)
	} // INVENTORY CARD END

	// Show/hide enemy data
	if npc[g.state.enemySelected].alive {
		drawer.Small(string(npc[g.state.enemySelected].race), int(npc[g.state.enemySelected].posx+48), int(npc[g.state.enemySelected].posy-10))
		drawer.Mini(string(npc[g.state.enemySelected].hp_max), int(npc[g.state.enemySelected].posx+64), int(npc[g.state.enemySelected].posy+18))
		drawer.Mini(string(npc[g.state.enemySelected].ac_armor_class), int(npc[g.state.enemySelected].posx+72), int(npc[g.state.enemySelected].posy+32))
		drawer.Mini(string(npc[g.state.enemySelected].item1), int(npc[g.state.enemySelected].posx+72), int(npc[g.state.enemySelected].posy+46))
	}

	// "For of war"/hidden roof for map 1
	if g.config.hidden {
		screen.DrawImage(g.assets.images.hideImage, opHide)
	}

	// Notification for round
	if g.state.round == 0 {
		drawer.Notif("Setting the scene !", 72, 72)
		drawer.Notif("DM explains the scene and/or what happens next.", 72, 94)
	} else if g.state.round == 1 {
		drawer.Notif("Movement - Up to your speed", 72, 72)
		drawer.Notif("Interaction - i.e opening a door, sheathing a weapon", 72, 94)
	} else if g.state.round == 2 {
		drawer.Notif("Action - Attack, Dash, Improvise, Hide, Search, ...", 72, 72)
		drawer.Notif("Combat resolution", 72, 94)
	}

	// DM cheat sheet
	if g.config.dm {
		screen.DrawImage(g.assets.images.dmImage, opBackground)
		drawer.Small("--- SUPERSIMPLIFIED COMBAT RULES (WIP) ---", 32, 32)
		drawer.Small("Is anyone surprised ? If you surprise an enemy, you'll have an additional turn.", 32, 82)
		drawer.Small("Everyone rolls initiative (1d20 + initiative modifier) and the one with highest start first", 32, 132)
		drawer.Small("On your turn, you can move a distance up to your speed and take 1 action", 32, 182)
		drawer.Small("To attack, roll a d20 and add weapons modifiers and check that against AC value", 32, 214)
		drawer.Small("Then to it, roll the dice from you weapon (i.e 1d6)", 32, 246)
		drawer.Small("--- SKILL CHECKS/SAVINGS THROWS (1d20) ---", 32, 320)
		drawer.Small("DM can ask for a skill check before a player can process with an action. This is resolved with a D20 roll +/- modifiers", 32, 370)
		drawer.Small("DM can ask for a saving throw based on abilities. Must resolve the difficulty (DC) set by the DM or else fail", 32, 420)
		drawer.Small("DC : 5 = very easy / 10 = Easy / 15 = Moderate / 20 = Hard / 25 = Very Hard", 32, 470)
		drawer.Small("--- MIRKWOOD ENGINE KEYBOARD SHORTCUTS ---", 32, 520)
		drawer.Small("P to switch player - R to roll a dice - Z/S/Q/D to move player - K to quit - up/down/left/right to move ennemies - e to switch enemies", 32, 570)
		drawer.Small("K to quit - up/down/left/right to move ennemies - e to switch enemies - L to link", 32, 620)
		drawer.Small("I - Show inventory/character panel - U DM info - KP1/KP2/KP3/KP4 to 'kill' enemies 1/2/3/4 - N for next round - G debug info", 32, 670)
		text.Draw(screen, "PRESS 'U' to open/close this panel :)", g.assets.fonts.mplusLargeFont, 500, 900, color.White)
	}
	return
}

func (g *Game) Update() error {
	// Handle keypress and set states
	g.handleState()

	return nil
}

type Game struct {
	config *config
	state  *state
	assets *assets
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 1920, 1080
}

func main() {
	cfg := newConfig()
	state := newState()

	a, err := loadAssets()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(engine_version)

	// TEST
	err = readConfigPlayer1()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf(string(MyConfig.name))

	ebiten.SetFullscreen(true)
	ebiten.SetWindowSize(1920, 1080)
	ebiten.SetWindowTitle(engine_version)
	game := Game{
		config: &cfg,
		state:  &state,
		assets: a,
	}
	if err := ebiten.RunGame(&game); err != nil {
		log.Fatal(err)
	}
}
