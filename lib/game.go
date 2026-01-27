package lib

import (
	"GentleWanderings/lib/printer"
	"bufio"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

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

// ShowInventory displays the player's collected items
func (g *Game) ShowInventory() {
	printer.ShowInventory()

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
		printer.ShowMenu() // prints out the menu in the lib printer
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
	printer.ShowJournal()

	for _, entry := range g.JournalLog {
		if strings.HasPrefix(entry, "  →") {
			fmt.Println(entry) // Item entries
		} else {
			fmt.Println(entry) // Day entries
		}
	}
	fmt.Println()
}

// ShowCurrentLocation displays detailed info about current location
func (g *Game) ShowCurrentLocation() {
	tile := g.GetTile(g.CurrentX, g.CurrentY)
	if tile == nil {
		return
	}

	printer.ShowCurrentLocation()

	// TODO: Move printing function to the printer
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
	printer.ShowStatistics()

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
