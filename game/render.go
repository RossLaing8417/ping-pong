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

	pos := g.playerLeft.Position
	for y := pos.TL.Y; y <= pos.BR.Y; y += 1 {
		buffer = append(buffer, DrawCommand{
			Coord: Coord{pos.TL.X, y},
			Data:  ' ',
			Style: style,
		})
	}

	pos = g.playerRight.Position
	for y := pos.TL.Y; y <= pos.BR.Y; y += 1 {
		buffer = append(buffer, DrawCommand{
			Coord: Coord{pos.TL.X, y},
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
