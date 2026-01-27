package lib

// Item represents a collectible object
type Item struct {
	Name        string
	Description string
	Category    string // keepsake, treasure, curiosity
	FoundAt     string
	FoundDay    int
}
