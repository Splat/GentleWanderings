package main

import (
	"GentleWanderings/lib"
	"GentleWanderings/lib/printer"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	game := lib.NewGame()
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("╔═══════════════════════════════════════════════════════════╗")
	fmt.Println("║           Welcome to Gentle Wanderings                    ║")
	fmt.Println("║         A Cozy Map-Making Adventure                        ║")
	fmt.Println("╚═══════════════════════════════════════════════════════════╝")
	fmt.Println()

	currentTile := game.GetTile(game.CurrentX, game.CurrentY)
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
			fmt.Println("║" + printer.CenterText("Journey Summary", 60) + "║")
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

			newTile := game.GetTile(game.CurrentX, game.CurrentY)
			fmt.Println()
			fmt.Println("╔════════════════════════════════════════════════════════════╗")
			fmt.Println("║" + printer.CenterText(newTile.Theme, 60) + "║")
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
