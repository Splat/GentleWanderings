package printer

import "fmt"

// ShowJournal displays the journal entry in a formatted box to the console.
// needs to take on more of the printing
func ShowJournal() {
	journal := fmt.Sprintf(`
╔════════════════════════════════════════════════════════════╗
║%s║
╚════════════════════════════════════════════════════════════╝
`, CenterText("Journal", 60))

	PrintToConsole(journal)
}

func ShowCurrentLocation() {
	locationInfo := fmt.Sprintf(`
╔════════════════════════════════════════════════════════════╗
║%s║
╚════════════════════════════════════════════════════════════╝
	`, CenterText("Current Location", 60))

	PrintToConsole(locationInfo)
}

func ShowStatistics() {
	statistics := fmt.Sprintf(`
╔════════════════════════════════════════════════════════════╗
║%s║
╚════════════════════════════════════════════════════════════╝
	`, CenterText("Statistics", 60))

	PrintToConsole(statistics)
}

func ShowInventory() {
	inventory := fmt.Sprintf(`
╔════════════════════════════════════════════════════════════╗
║%s║
╚════════════════════════════════════════════════════════════╝
`, CenterText("Collection", 60))

	PrintToConsole(inventory)
}
