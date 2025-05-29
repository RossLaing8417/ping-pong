package game

import "github.com/gdamore/tcell/v2"

type DrawCommand struct {
	Coord
	Data  rune
	Style tcell.Style
}

func (g *Game) GetDrawCommands(style tcell.Style) []DrawCommand {
	buffer := g.buffer

	for x := g.arena.TL.X; x <= g.arena.BR.X; x += 1 {
		buffer = append(buffer, DrawCommand{
			Coord: Coord{x, g.arena.TL.Y},
			Data:  ' ',
			Style: style,
		})
		buffer = append(buffer, DrawCommand{
			Coord: Coord{x, g.arena.BR.Y},
			Data:  ' ',
			Style: style,
		})
	}

	for y := g.arena.TL.Y; y <= g.arena.BR.Y; y += 1 {
		buffer = append(buffer, DrawCommand{
			Coord: Coord{g.arena.TL.X, y},
			Data:  ' ',
			Style: style,
		})
		buffer = append(buffer, DrawCommand{
			Coord: Coord{g.arena.BR.X, y},
			Data:  ' ',
			Style: style,
		})
	}

	player := g.playerLeft
	for y := player.Top; y <= player.Bottom; y += 1 {
		buffer = append(buffer, DrawCommand{
			Coord: Coord{player.X, y},
			Data:  ' ',
			Style: style,
		})
	}

	player = g.playerRight
	for y := player.Top; y <= player.Bottom; y += 1 {
		buffer = append(buffer, DrawCommand{
			Coord: Coord{player.X, y},
			Data:  ' ',
			Style: style,
		})
	}

	buffer = append(buffer, DrawCommand{
		Coord: Coord{g.puck.Position.X, g.puck.Position.Y},
		Data:  ' ',
		Style: style,
	})

	return buffer
}
