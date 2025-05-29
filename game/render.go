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

func (g *Game) GetDrawCommands(baseStyle, altStyle tcell.Style) []DrawCommand {
	buffer := g.buffer

	for x := g.arena.TL.X; x <= g.arena.BR.X; x += 1 {
		buffer = append(buffer, DrawCommand{
			Coord: Coord{x, g.arena.TL.Y},
			Data:  ' ',
			Style: altStyle,
		})
		buffer = append(buffer, DrawCommand{
			Coord: Coord{x, g.arena.BR.Y},
			Data:  ' ',
			Style: altStyle,
		})
	}

	for y := g.arena.TL.Y; y <= g.arena.BR.Y; y += 1 {
		buffer = append(buffer, DrawCommand{
			Coord: Coord{g.arena.TL.X, y},
			Data:  ' ',
			Style: altStyle,
		})
		buffer = append(buffer, DrawCommand{
			Coord: Coord{g.arena.BR.X, y},
			Data:  ' ',
			Style: altStyle,
		})
	}

	player := g.playerLeft
	for y := player.Top; y <= player.Bottom; y += 1 {
		buffer = append(buffer, DrawCommand{
			Coord: Coord{player.X, y},
			Data:  ' ',
			Style: altStyle,
		})
	}

	player = g.playerRight
	for y := player.Top; y <= player.Bottom; y += 1 {
		buffer = append(buffer, DrawCommand{
			Coord: Coord{player.X, y},
			Data:  ' ',
			Style: altStyle,
		})
	}

	x := (g.arena.BR.X - 1 - g.arena.TL.X) / 4
	y := g.arena.TL.Y + 2

	buffer = append(buffer, DrawCommand{
		Coord: Coord{x, y},
		Data:  rune(fmt.Sprintf("%d", g.playerLeft.Score)[0]),
		Style: baseStyle,
	})

	buffer = append(buffer, DrawCommand{
		Coord: Coord{x * 3, y},
		Data:  rune(fmt.Sprintf("%d", g.playerRight.Score)[0]),
		Style: baseStyle,
	})

	buffer = append(buffer, DrawCommand{
		Coord: Coord{g.puck.Position.X, g.puck.Position.Y},
		Data:  ' ',
		Style: altStyle,
	})

	return buffer
}
