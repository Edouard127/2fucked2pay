package bots

import (
	"fmt"
	"github.com/edouard127/mc-go-1.12.2/PathFinding"
	. "github.com/edouard127/mc-go-1.12.2/struct"
	"strconv"
	"strings"
)

func ParseCommands(g *Game, s string) {
	if strings.HasPrefix(string(s[0]), "!") {
		s = strings.TrimPrefix(s, "!")
		s = strings.ToLower(s)
		split := strings.Split(s, " ")
		switch split[0] {
		case "goto":
			if len(split) == 4 {
				fmt.Println(g.GetPlayer().X, g.GetPlayer().Y, g.GetPlayer().Z)
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
			e := g.ClosestEntity(20)
			if e != nil {
				g.Chat("Closest entity: " + string((e).ID))
				g.Attack(e)
			} else {
				g.Chat("No entities found")
			}
		case "swing":
			g.SwingHand(true)
			g.SwingHand(false)
		case "jump":
			g.TweenJump()
		}
	}
}
