package printer

import (
	"fmt"
	"strings"
)

func ShowMap(width int) {
	mapRender := fmt.Sprintf(`
╔%s╗
║%s║
╠%s╣
`, strings.Repeat("═", width*4-1),
		CenterText("Map", width*4-1),
		strings.Repeat("═", width*4-1))

	PrintToConsole(mapRender)
}
