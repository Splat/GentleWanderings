package printer

import "fmt"

func ShowMenu() {
	menu := fmt.Sprintf(`
╔════════════════════════════════════════════════════════════╗
║%s║
╠════════════════════════════════════════════════════════════╣
║  1. View Map                                               ║
║  2. Detailed Map (with locations)                          ║
║  3. View Inventory                                         ║
║  4. Read Journal                                           ║
║  5. Current Location Info                                  ║
║  6. Game Statistics                                        ║
║  7. Return to Journey                                      ║
╚════════════════════════════════════════════════════════════╝

Choose an option (1-7): `, CenterText("Menu", 60))

	PrintToConsole(menu)
}
