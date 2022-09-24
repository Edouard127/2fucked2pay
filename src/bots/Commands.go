package bots

import (
	"github.com/edouard127/mc-go-1.12.2/PathFinding"
	"github.com/edouard127/mc-go-1.12.2/data/entities"
	. "github.com/edouard127/mc-go-1.12.2/maths"
	. "github.com/edouard127/mc-go-1.12.2/struct"
	"strconv"
	"strings"
	"time"
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
		case "walk":
			g.WalkStraight(20)
		case "obsi":
			highway := ObsidianHighway(g, 20)
			BuildHighway(g, highway)
		}
	}
}

func ObsidianHighway(g *Game, length int) []Vector3 {
	// This function is inspired by HighwayTools from Lambda

	// Make an array of the path with initial value
	var path []Vector3
	dir := g.GetPlayer().GetFacing()
	for i := 0; i < length; i++ {
		switch dir {
		case entities.DSouth:
			path = append(path, Vector3{X: g.GetPlayer().Position.X, Y: g.GetPlayer().Position.Y - 1, Z: g.GetPlayer().Position.Z + float64(i)})
		case entities.DNorth:
			path = append(path, Vector3{X: g.GetPlayer().Position.X, Y: g.GetPlayer().Position.Y - 1, Z: g.GetPlayer().Position.Z - float64(i)})
		case entities.DWest:
			path = append(path, Vector3{X: g.GetPlayer().Position.X - float64(i), Y: g.GetPlayer().Position.Y - 1, Z: g.GetPlayer().Position.Z})
		case entities.DEast:
			path = append(path, Vector3{X: g.GetPlayer().Position.X + float64(i), Y: g.GetPlayer().Position.Y - 1, Z: g.GetPlayer().Position.Z})
		}
	}
	return path
}

func BuildHighway(g *Game, path []Vector3) {
	go func() {
		for _, pos := range path {
			g.LookAt(pos)
			g.WalkToVector(pos)
			time.Sleep(50 * time.Millisecond)
		}
	}()
}
