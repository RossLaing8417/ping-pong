package game

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
)

type DrawCommand struct {
	Coord
	Data  rune
	Style tcell.Style
}

func (g *Game) GetDrawCommands() []DrawCommand {
	buffer := g.buffer

	for x := g.arena.TL.X; x <= g.arena.BR.X; x += 1 {
		buffer = append(buffer, DrawCommand{
			Coord: Coord{x, g.arena.TL.Y},
			Data:  ' ',
			Style: g.arena.Style,
		})
		buffer = append(buffer, DrawCommand{
			Coord: Coord{x, g.arena.BR.Y},
			Data:  ' ',
			Style: g.arena.Style,
		})
	}

	for y := g.arena.TL.Y; y <= g.arena.BR.Y; y += 1 {
		buffer = append(buffer, DrawCommand{
			Coord: Coord{g.arena.TL.X, y},
			Data:  ' ',
			Style: g.arena.Style,
		})
		buffer = append(buffer, DrawCommand{
			Coord: Coord{g.arena.BR.X, y},
			Data:  ' ',
			Style: g.arena.Style,
		})
	}

	player := g.playerLeft
	for y := player.Top; y <= player.Bottom; y += 1 {
		buffer = append(buffer, DrawCommand{
			Coord: Coord{player.X, y},
			Data:  ' ',
			Style: player.Style,
		})
	}

	player = g.playerRight
	for y := player.Top; y <= player.Bottom; y += 1 {
		buffer = append(buffer, DrawCommand{
			Coord: Coord{player.X, y},
			Data:  ' ',
			Style: player.Style,
		})
	}

	x := (g.arena.BR.X - 1 - g.arena.TL.X) / 4
	y := g.arena.TL.Y + 2

	buffer = append(buffer, DrawCommand{
		Coord: Coord{x, y},
		Data:  rune(fmt.Sprintf("%d", g.playerLeft.Score)[0]),
		Style: g.baseStyle,
	})

	buffer = append(buffer, DrawCommand{
		Coord: Coord{x * 3, y},
		Data:  rune(fmt.Sprintf("%d", g.playerRight.Score)[0]),
		Style: g.baseStyle,
	})

	buffer = append(buffer, DrawCommand{
		Coord: Coord{g.puck.Position.X, g.puck.Position.Y},
		Data:  ' ',
		Style: g.puck.Style,
	})

	return buffer
}
