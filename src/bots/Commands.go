package bots

import (
	"github.com/edouard127/mc-go-1.12.2/PathFinding"
	. "github.com/edouard127/mc-go-1.12.2/maths"
	. "github.com/edouard127/mc-go-1.12.2/struct"
	"strconv"
	"strings"
)

func ParseCommands(g *Game, s string) {
	if strings.HasPrefix(string(s[0]), "!") && g.Server.Addr == "127.0.0.1" /* <-- Avoid people using commands */ {
		s = strings.TrimPrefix(s, "!")
		s = strings.ToLower(s)
		split := strings.Split(s, " ")
		switch split[0] {
		case "goto":
			if len(split) == 4 {
				pos := g.GetPlayer().Position
				x, _ := strconv.ParseFloat(split[1], 64)
				y, _ := strconv.ParseFloat(split[2], 64)
				z, _ := strconv.ParseFloat(split[3], 64)
				start := PathFinding.Node{
					Position: pos,
				}
				end := PathFinding.Node{
					Position: Vector3{
						X: x,
						Y: y,
						Z: z,
					},
				}
				path := PathFinding.NewAStar(&start, &end)
				calc := PathFinding.Compute(path, g)

				if calc.PathFound {
					g.Chat("Path found with " + strconv.Itoa(len(calc.Path.Nodes)) + " nodes")
					for _, node := range calc.Path.Nodes {
						g.Chat("X: " + strconv.FormatFloat(node.Position.X, 'f', 6, 64) + " Y: " + strconv.FormatFloat(node.Position.Y, 'f', 6, 64) + " Z: " + strconv.FormatFloat(node.Position.Z, 'f', 6, 64))
					}

				} else {
					g.Chat("Path not found")
				}
			}
		case "closest":
			g.LookYawPitch(0, 90)
		case "eat":
			g.Eat()
		case "swing":
			g.SwingHand(true)
			g.SwingHand(false)
		case "jump":
			g.TweenJump()
		case "slot":
			for i, item := range g.GetPlayer().Inventory {
				if item.Count == 64 {
					g.Chat("Slot " + strconv.Itoa(i) + " is full")
				}
			}
		}
	}
}
