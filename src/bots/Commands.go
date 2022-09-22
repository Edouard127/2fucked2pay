package bots

import (
	"github.com/edouard127/mc-go-1.12.2/PathFinding"
	"github.com/edouard127/mc-go-1.12.2/data/entities"
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
		case "walk":
			g.WalkStraight(20)
		case "obsi":
			highway := ObsidianHighway(g, 100, 2)
			BuildHighway(g, highway)
		}
	}
}

func ObsidianHighway(g *Game, length int, width int) []Vector3 {
	// This function is inspired by HighwayTools from Lambda
	var path []Vector3
	dir := g.GetPlayer().GetFacing()
	for i := 0; i < length; i++ {
		// Make a floor of obsidian of width size and make the number of blocks equal to each side of the player position
		for j := -width; j <= width; j++ {
			for k := -width; k <= width; k++ {
				switch dir {
				case entities.DNorth:
					path = append(path, Vector3{X: float64(j), Y: 0, Z: float64(k)})
				case entities.DSouth:
					path = append(path, Vector3{X: float64(-j), Y: 0, Z: float64(-k)})
				case entities.DWest:
					path = append(path, Vector3{X: float64(k), Y: 0, Z: float64(-j)})
				case entities.DEast:
					path = append(path, Vector3{X: float64(-k), Y: 0, Z: float64(j)})
				}
			}
		}
	}
	return path
}

func BuildHighway(g *Game, path []Vector3) {
	go func() {
		for {
			select {
			case e := <-g.Events:
				switch e.(type) {
				case TickEvent:
					if len(path) == 0 {
						g.Chat("Done")
						return
					}
					pos := g.GetPlayer().Position.Add(path[0])
					g.LookAt(pos)
					g.SetPosition(pos)
					path = path[1:] // Fix multiple blocks being placed at the same time
				}
			}
		}
	}()
}
