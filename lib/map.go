package lib

import (
	"fmt"
	"strings"
)

type Tile struct {
	X           int
	Y           int
	Theme       string
	Description string
	Discovery   string
	Visited     bool
	Item        *Item // Optional item found at this location
}

type Direction struct {
	Name string
	DX   int
	DY   int
}

type LocationOption struct {
	Theme       string
	Description string
}

func (g *Game) tileKey(x, y int) string {
	return fmt.Sprintf("%d,%d", x, y)
}

func (g *Game) GetTile(x, y int) *Tile {
	return g.Map[g.tileKey(x, y)]
}

func (g *Game) GetAdjacentDirections() []Direction {
	directions := []Direction{
		{Name: "North", DX: 0, DY: 1},
		{Name: "South", DX: 0, DY: -1},
		{Name: "East", DX: 1, DY: 0},
		{Name: "West", DX: -1, DY: 0},
	}

	available := []Direction{}
	for _, dir := range directions {
		newX, newY := g.CurrentX+dir.DX, g.CurrentY+dir.DY
		if g.GetTile(newX, newY) == nil {
			available = append(available, dir)
		}
	}

	return available
}

// GenerateLocationOptions creates 3 themed location options for the player
func (g *Game) GenerateLocationOptions() []LocationOption {
	themes := []string{
		"Mushroom Circle", "Mossy Stones", "Babbling Brook", "Wildflower Meadow",
		"Hollow Tree", "Crystal Pool", "Foggy Hollow", "Sunlit Glade",
		"Berry Thicket", "Stone Circle", "Whispering Willows", "Hidden Grotto",
		"Autumn Vale", "Morning Mist", "Starlit Clearing", "Gentle Waterfall",
	}

	descriptors := [][]string{
		{"ancient", "forgotten", "peaceful", "mysterious", "enchanted"},
		{"soft light dances across", "shadows play among", "gentle sounds echo from", "a strange calm pervades"},
		{"You feel drawn here", "Something calls to you", "A sense of wonder fills you", "Time seems to slow"},
	}

	options := []LocationOption{}
	usedThemes := make(map[string]bool)

	for i := 0; i < 3; i++ {
		var theme string
		for {
			theme = themes[g.rand.Intn(len(themes))]
			if !usedThemes[theme] {
				usedThemes[theme] = true
				break
			}
		}

		adj := descriptors[0][g.rand.Intn(len(descriptors[0]))]
		detail := descriptors[1][g.rand.Intn(len(descriptors[1]))]
		feeling := descriptors[2][g.rand.Intn(len(descriptors[2]))]

		options = append(options, LocationOption{
			Theme:       theme,
			Description: fmt.Sprintf("An %s %s where %s the space. %s.", adj, strings.ToLower(theme), detail, feeling),
		})
	}

	return options
}

// ShowMap displays an enhanced map with box-drawing characters
func (g *Game) ShowMap() {
	minX, maxX := 0, 0
	minY, maxY := 0, 0

	for _, tile := range g.Map {
		if tile.X < minX {
			minX = tile.X
		}
		if tile.X > maxX {
			maxX = tile.X
		}
		if tile.Y < minY {
			minY = tile.Y
		}
		if tile.Y > maxY {
			maxY = tile.Y
		}
	}

	// Add padding
	minX--
	maxX++
	minY--
	maxY++

	width := maxX - minX + 1

	fmt.Println()
	fmt.Println("╔" + strings.Repeat("═", width*4-1) + "╗")
	fmt.Println("║" + CenterText("Your Map", width*4-1) + "║")
	fmt.Println("╠" + strings.Repeat("═", width*4-1) + "╣")

	for y := maxY; y >= minY; y-- {
		line := "║ "
		for x := minX; x <= maxX; x++ {
			tile := g.GetTile(x, y)
			if tile != nil {
				if x == g.CurrentX && y == g.CurrentY {
					line += "📍"
				} else if tile.Item != nil {
					line += "🎁"
				} else {
					line += "■ "
				}
			} else {
				// Check if there's an adjacent explored tile
				hasAdjacent := false
				for _, d := range []struct{ dx, dy int }{{0, 1}, {0, -1}, {1, 0}, {-1, 0}} {
					if g.GetTile(x+d.dx, y+d.dy) != nil {
						hasAdjacent = true
						break
					}
				}
				if hasAdjacent {
					line += "· "
				} else {
					line += "  "
				}
			}
			line += " "
		}
		fmt.Println(line + "║")
	}

	fmt.Println("╚" + strings.Repeat("═", width*4-1) + "╝")
	fmt.Println()
	fmt.Println("Legend: 📍 You  ■ Explored  🎁 Has Item  · Unexplored")
	fmt.Println()
}

// ShowDetailedMap shows the map with location names
func (g *Game) ShowDetailedMap() {
	minX, maxX := 0, 0
	minY, maxY := 0, 0

	for _, tile := range g.Map {
		if tile.X < minX {
			minX = tile.X
		}
		if tile.X > maxX {
			maxX = tile.X
		}
		if tile.Y < minY {
			minY = tile.Y
		}
		if tile.Y > maxY {
			maxY = tile.Y
		}
	}

	fmt.Println()
	fmt.Println("╔════════════════════════════════════════════════════════════╗")
	fmt.Println("║" + CenterText("Detailed Map", 60) + "║")
	fmt.Println("╚════════════════════════════════════════════════════════════╝")
	fmt.Println()

	// List all locations
	locations := []*Tile{}
	for _, tile := range g.Map {
		locations = append(locations, tile)
	}

	// Sort by turn order (Y then X for visual consistency)
	for i := 0; i < len(locations); i++ {
		for j := i + 1; j < len(locations); j++ {
			if locations[i].Y < locations[j].Y || (locations[i].Y == locations[j].Y && locations[i].X > locations[j].X) {
				locations[i], locations[j] = locations[j], locations[i]
			}
		}
	}

	for _, tile := range locations {
		marker := "■"
		if tile.X == g.CurrentX && tile.Y == g.CurrentY {
			marker = "📍"
		}

		pos := fmt.Sprintf("(%d,%d)", tile.X, tile.Y)
		fmt.Printf("%s %s %s\n", marker, tile.Theme, pos)
		if tile.Item != nil {
			fmt.Printf("   🎁 Contains: %s\n", tile.Item.Name)
		}
	}
	fmt.Println()
}
