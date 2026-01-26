package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

// Tile represents a discovered location on the map
type Tile struct {
	X           int
	Y           int
	Theme       string
	Description string
	Discovery   string
	Visited     bool
	Item        *Item // Optional item found at this location
}

// Item represents a collectible object
type Item struct {
	Name        string
	Description string
	Category    string // keepsake, treasure, curiosity
	FoundAt     string
	FoundDay    int
}

// Game holds the game state
type Game struct {
	Map        map[string]*Tile
	CurrentX   int
	CurrentY   int
	TurnCount  int
	JournalLog []string
	Inventory  []*Item
	rand       *rand.Rand
}

// NewGame initializes a new game
func NewGame() *Game {
	source := rand.NewSource(time.Now().UnixNano())
	g := &Game{
		Map:        make(map[string]*Tile),
		CurrentX:   0,
		CurrentY:   0,
		TurnCount:  1,
		JournalLog: []string{},
		Inventory:  []*Item{},
		rand:       rand.New(source),
	}

	// Create starting tile
	startTile := &Tile{
		X:           0,
		Y:           0,
		Theme:       "Quiet Grove",
		Description: "A peaceful clearing surrounded by ancient trees, dappled sunlight filtering through the leaves.",
		Discovery:   "You begin your journey here, where the world feels safe and full of possibility.",
		Visited:     true,
	}
	g.Map[g.tileKey(0, 0)] = startTile
	g.JournalLog = append(g.JournalLog, "Day 1: "+startTile.Discovery)

	return g
}

func (g *Game) tileKey(x, y int) string {
	return fmt.Sprintf("%d,%d", x, y)
}

func (g *Game) getTile(x, y int) *Tile {
	return g.Map[g.tileKey(x, y)]
}

// GetAdjacentDirections returns available directions to explore
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
		if g.getTile(newX, newY) == nil {
			available = append(available, dir)
		}
	}

	return available
}

type Direction struct {
	Name string
	DX   int
	DY   int
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

type LocationOption struct {
	Theme       string
	Description string
}

// GenerateDiscovery creates a discovery event for the new location
func (g *Game) GenerateDiscovery(theme string) string {
	discoveries := []string{
		"You discover %s and feel a deep connection to this place.",
		"As you arrive at %s, you notice details you hadn't expected.",
		"The %s reveals itself slowly, inviting you to linger.",
		"%s feels like it has been waiting for you.",
		"You find yourself drawn deeper into %s.",
	}

	template := discoveries[g.rand.Intn(len(discoveries))]
	return fmt.Sprintf(template, strings.ToLower(theme))
}

// GenerateItem creates a random item (or returns nil if no item)
func (g *Game) GenerateItem(theme string, turnCount int) *Item {
	// 60% chance to find an item
	if g.rand.Float32() > 0.6 {
		return nil
	}

	itemsByCategory := map[string][]string{
		"keepsake": {
			"Smooth River Stone", "Pressed Flower", "Acorn Cap", "Bird Feather",
			"Seashell Fragment", "Dried Leaf", "Pinecone", "Lucky Pebble",
			"Glass Bead", "Carved Twig", "Moss Sample", "Butterfly Wing",
		},
		"treasure": {
			"Ancient Coin", "Crystal Shard", "Silver Locket", "Brass Key",
			"Jade Figurine", "Pearl", "Golden Ring", "Copper Medallion",
			"Gemstone", "Amber", "Moonstone", "Opal",
		},
		"curiosity": {
			"Strange Map Fragment", "Mysterious Note", "Odd Compass", "Faded Photograph",
			"Old Journal Page", "Weathered Letter", "Riddle Scroll", "Poetry Fragment",
			"Sheet Music", "Recipe Card", "Star Chart", "Encrypted Message",
		},
	}

	categories := []string{"keepsake", "treasure", "curiosity"}
	category := categories[g.rand.Intn(len(categories))]
	items := itemsByCategory[category]
	itemName := items[g.rand.Intn(len(items))]

	descriptions := map[string][]string{
		"keepsake": {
			"A simple treasure that reminds you of this moment.",
			"Something small but meaningful.",
			"A gentle reminder of your journey.",
			"It feels right to carry this with you.",
		},
		"treasure": {
			"It glimmers softly in your hand, valuable yet mysterious.",
			"Worth keeping safe - who knows its story?",
			"A prize from your wanderings.",
			"Something precious, left behind long ago.",
		},
		"curiosity": {
			"This raises more questions than it answers.",
			"You sense there's a story here, waiting to unfold.",
			"Strange and intriguing - you must learn more.",
			"A puzzle piece from someone else's tale.",
		},
	}

	descList := descriptions[category]
	desc := descList[g.rand.Intn(len(descList))]

	return &Item{
		Name:        itemName,
		Description: desc,
		Category:    category,
		FoundAt:     theme,
		FoundDay:    turnCount,
	}
}

// Explore creates a new tile in the given direction
func (g *Game) Explore(dir Direction, option LocationOption) *Item {
	newX := g.CurrentX + dir.DX
	newY := g.CurrentY + dir.DY

	discovery := g.GenerateDiscovery(option.Theme)
	item := g.GenerateItem(option.Theme, g.TurnCount)

	newTile := &Tile{
		X:           newX,
		Y:           newY,
		Theme:       option.Theme,
		Description: option.Description,
		Discovery:   discovery,
		Visited:     true,
		Item:        item,
	}

	g.Map[g.tileKey(newX, newY)] = newTile
	g.CurrentX = newX
	g.CurrentY = newY
	g.TurnCount++

	logEntry := fmt.Sprintf("Day %d: %s", g.TurnCount, discovery)
	g.JournalLog = append(g.JournalLog, logEntry)

	if item != nil {
		g.Inventory = append(g.Inventory, item)
		logEntry := fmt.Sprintf("  → Found: %s", item.Name)
		g.JournalLog = append(g.JournalLog, logEntry)
	}

	return item
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
	fmt.Println("║" + centerText("Your Map", width*4-1) + "║")
	fmt.Println("╠" + strings.Repeat("═", width*4-1) + "╣")

	for y := maxY; y >= minY; y-- {
		line := "║ "
		for x := minX; x <= maxX; x++ {
			tile := g.getTile(x, y)
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
					if g.getTile(x+d.dx, y+d.dy) != nil {
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
	fmt.Println("║" + centerText("Detailed Map", 60) + "║")
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

func centerText(text string, width int) string {
	if len(text) >= width {
		return text
	}
	leftPad := (width - len(text)) / 2
	rightPad := width - len(text) - leftPad
	return strings.Repeat(" ", leftPad) + text + strings.Repeat(" ", rightPad)
}

// ShowInventory displays the player's collected items
func (g *Game) ShowInventory() {
	fmt.Println()
	fmt.Println("╔════════════════════════════════════════════════════════════╗")
	fmt.Println("║" + centerText("Your Collection", 60) + "║")
	fmt.Println("╚════════════════════════════════════════════════════════════╝")

	if len(g.Inventory) == 0 {
		fmt.Println("\nYour pack is empty. Perhaps you'll find something as you wander...")
		fmt.Println()
		return
	}

	// Group by category
	categories := map[string][]*Item{
		"keepsake":  {},
		"treasure":  {},
		"curiosity": {},
	}

	for _, item := range g.Inventory {
		categories[item.Category] = append(categories[item.Category], item)
	}

	categoryNames := map[string]string{
		"keepsake":  "🍃 Keepsakes",
		"treasure":  "💎 Treasures",
		"curiosity": "❓ Curiosities",
	}

	categoryOrder := []string{"keepsake", "treasure", "curiosity"}

	for _, cat := range categoryOrder {
		items := categories[cat]
		if len(items) == 0 {
			continue
		}

		fmt.Printf("\n%s (%d)\n", categoryNames[cat], len(items))
		fmt.Println(strings.Repeat("─", 60))

		for i, item := range items {
			fmt.Printf("%d. %s\n", i+1, item.Name)
			fmt.Printf("   %s\n", item.Description)
			fmt.Printf("   Found at %s on Day %d\n", item.FoundAt, item.FoundDay)
			if i < len(items)-1 {
				fmt.Println()
			}
		}
	}

	fmt.Printf("\n%s Total items collected: %d\n", strings.Repeat("─", 60), len(g.Inventory))
	fmt.Println()
}

// ShowMenu displays the main game menu
func (g *Game) ShowMenu(scanner *bufio.Scanner) {
	for {
		fmt.Println()
		fmt.Println("╔════════════════════════════════════════════════════════════╗")
		fmt.Println("║" + centerText("Menu", 60) + "║")
		fmt.Println("╠════════════════════════════════════════════════════════════╣")
		fmt.Println("║  1. View Map                                               ║")
		fmt.Println("║  2. Detailed Map (with locations)                          ║")
		fmt.Println("║  3. View Inventory                                         ║")
		fmt.Println("║  4. Read Journal                                           ║")
		fmt.Println("║  5. Current Location Info                                  ║")
		fmt.Println("║  6. Game Statistics                                        ║")
		fmt.Println("║  7. Return to Journey                                      ║")
		fmt.Println("╚════════════════════════════════════════════════════════════╝")

		fmt.Print("\nChoose an option (1-7): ")
		if !scanner.Scan() {
			return
		}

		choice := strings.TrimSpace(scanner.Text())

		switch choice {
		case "1":
			g.ShowMap()

		case "2":
			g.ShowDetailedMap()

		case "3":
			g.ShowInventory()

		case "4":
			g.ShowJournal()

		case "5":
			g.ShowCurrentLocation()

		case "6":
			g.ShowStatistics()

		case "7":
			fmt.Println("\nReturning to your journey...")
			return

		default:
			fmt.Println("\nInvalid choice. Please try again.")
		}

		fmt.Print("\nPress Enter to continue...")
		scanner.Scan()
	}
}

// ShowJournal displays the journey log
func (g *Game) ShowJournal() {
	fmt.Println()
	fmt.Println("╔════════════════════════════════════════════════════════════╗")
	fmt.Println("║" + centerText("Your Journey", 60) + "║")
	fmt.Println("╚════════════════════════════════════════════════════════════╝")
	fmt.Println()

	for _, entry := range g.JournalLog {
		if strings.HasPrefix(entry, "  →") {
			// Item entries
			fmt.Println(entry)
		} else {
			// Day entries
			fmt.Println(entry)
		}
	}
	fmt.Println()
}

// ShowCurrentLocation displays detailed info about current location
func (g *Game) ShowCurrentLocation() {
	tile := g.getTile(g.CurrentX, g.CurrentY)
	if tile == nil {
		return
	}

	fmt.Println()
	fmt.Println("╔════════════════════════════════════════════════════════════╗")
	fmt.Println("║" + centerText("Current Location", 60) + "║")
	fmt.Println("╚════════════════════════════════════════════════════════════╝")
	fmt.Println()
	fmt.Printf("🌿 %s\n", tile.Theme)
	fmt.Printf("📍 Position: (%d, %d)\n\n", tile.X, tile.Y)
	fmt.Printf("%s\n\n", tile.Description)

	if tile.Item != nil {
		fmt.Printf("🎁 You found: %s\n", tile.Item.Name)
		fmt.Printf("   %s\n", tile.Item.Description)
	} else {
		fmt.Println("This location holds no items, just peaceful presence.")
	}
	fmt.Println()
}

// ShowStatistics displays game statistics
func (g *Game) ShowStatistics() {
	fmt.Println()
	fmt.Println("╔════════════════════════════════════════════════════════════╗")
	fmt.Println("║" + centerText("Statistics", 60) + "║")
	fmt.Println("╚════════════════════════════════════════════════════════════╝")
	fmt.Println()

	fmt.Printf("🗓️  Days Traveled: %d\n", g.TurnCount)
	fmt.Printf("🗺️  Locations Discovered: %d\n", len(g.Map))
	fmt.Printf("🎒 Items Collected: %d\n", len(g.Inventory))

	// Count by category
	categories := map[string]int{}
	for _, item := range g.Inventory {
		categories[item.Category]++
	}

	if len(g.Inventory) > 0 {
		fmt.Println("\nCollection breakdown:")
		if categories["keepsake"] > 0 {
			fmt.Printf("  🍃 Keepsakes: %d\n", categories["keepsake"])
		}
		if categories["treasure"] > 0 {
			fmt.Printf("  💎 Treasures: %d\n", categories["treasure"])
		}
		if categories["curiosity"] > 0 {
			fmt.Printf("  ❓ Curiosities: %d\n", categories["curiosity"])
		}
	}

	// Calculate exploration extent
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

	width := maxX - minX + 1
	height := maxY - minY + 1

	fmt.Printf("\n🧭 Map Dimensions: %d × %d\n", width, height)
	fmt.Printf("📏 Furthest North: %d, South: %d, East: %d, West: %d\n", maxY, minY, maxX, minX)
	fmt.Println()
}

func main() {
	game := NewGame()
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("╔═══════════════════════════════════════════════════════════╗")
	fmt.Println("║           Welcome to Gentle Wanderings                    ║")
	fmt.Println("║         A Cozy Map-Making Adventure                        ║")
	fmt.Println("╚═══════════════════════════════════════════════════════════╝")
	fmt.Println()

	currentTile := game.getTile(game.CurrentX, game.CurrentY)
	fmt.Printf("🌿 %s\n", currentTile.Theme)
	fmt.Printf("%s\n", currentTile.Description)
	fmt.Printf("\n%s\n", currentTile.Discovery)

	for {
		fmt.Println("\n" + strings.Repeat("─", 60))

		// Show available directions
		directions := game.GetAdjacentDirections()
		if len(directions) == 0 {
			fmt.Println("\nYou have explored all directions from here!")
			fmt.Println("Commands: [menu] | [m]ap | [i]nventory | [j]ournal | [q]uit")
		} else {
			fmt.Println("\nWhere would you like to wander?")
			for i, dir := range directions {
				fmt.Printf("  %d. Explore %s\n", i+1, dir.Name)
			}
			fmt.Println("\nOther: [menu] | [m]ap | [i]nventory | [j]ournal | [q]uit")
		}

		fmt.Print("\n> ")
		if !scanner.Scan() {
			break
		}

		input := strings.ToLower(strings.TrimSpace(scanner.Text()))

		switch input {
		case "menu":
			game.ShowMenu(scanner)

		case "m", "map":
			game.ShowMap()

		case "i", "inv", "inventory":
			game.ShowInventory()

		case "j", "journal":
			game.ShowJournal()

		case "q", "quit":
			fmt.Println()
			fmt.Println("╔════════════════════════════════════════════════════════════╗")
			fmt.Println("║" + centerText("Journey Summary", 60) + "║")
			fmt.Println("╚════════════════════════════════════════════════════════════╝")
			fmt.Printf("\n🗓️  Days traveled: %d\n", game.TurnCount)
			fmt.Printf("🗺️  Locations discovered: %d\n", len(game.Map))
			fmt.Printf("🎒 Items collected: %d\n\n", len(game.Inventory))
			fmt.Println("Thank you for wandering with us. Until next time... 🌙✨")
			fmt.Println()
			return

		default:
			// Try to parse as a direction number
			choice, err := strconv.Atoi(input)
			if err != nil || choice < 1 || choice > len(directions) {
				fmt.Println("Invalid choice. Please try again.")
				continue
			}

			selectedDir := directions[choice-1]

			// Generate 3 location options
			options := game.GenerateLocationOptions()

			fmt.Printf("\n✨ As you head %s, three paths reveal themselves:\n\n", selectedDir.Name)
			for i, opt := range options {
				fmt.Printf("%d. %s\n   %s\n\n", i+1, opt.Theme, opt.Description)
			}

			fmt.Print("Which path calls to you? (1-3): ")
			if !scanner.Scan() {
				break
			}

			optChoice, err := strconv.Atoi(strings.TrimSpace(scanner.Text()))
			if err != nil || optChoice < 1 || optChoice > 3 {
				fmt.Println("Let's try that again...")
				continue
			}

			selectedOption := options[optChoice-1]
			foundItem := game.Explore(selectedDir, selectedOption)

			newTile := game.getTile(game.CurrentX, game.CurrentY)
			fmt.Println()
			fmt.Println("╔════════════════════════════════════════════════════════════╗")
			fmt.Println("║" + centerText(newTile.Theme, 60) + "║")
			fmt.Println("╚════════════════════════════════════════════════════════════╝")
			fmt.Printf("\n%s\n", newTile.Description)
			fmt.Printf("\n%s\n", newTile.Discovery)

			if foundItem != nil {
				fmt.Println()
				fmt.Println(strings.Repeat("─", 60))
				fmt.Printf("\n✨ You found something! ✨\n\n")
				fmt.Printf("🎁 %s\n", foundItem.Name)
				fmt.Printf("   %s\n", foundItem.Description)
				fmt.Println()
			}
		}
	}
}
