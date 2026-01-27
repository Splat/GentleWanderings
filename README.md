# Gentle Wanderings 🌿

A cozy 2D map-making game inspired by Solo-journalling and map-making RPGS, built in Go. Explore procedurally generated locations, make choices about where to wander, and build your own unique map through gentle discovery.

## Features

- **Procedural Generation**: Each direction offers 3 unique location options with themed descriptions
- **Map Building**: Your world grows organically as you explore in cardinal directions
- **Journal System**: Track your journey through automatically generated log entries
- **Cozy Atmosphere**: Peaceful themes like Mushroom Circles, Babbling Brooks, and Sunlit Glades
- **Simple ASCII Map**: Visual representation of your explored world

## How to Run

Make sure you have Go installed (version 1.21 or higher):

```bash
go version
```

Then run the game:

```bash
cd GentleWanderings
go run main.go
```

Or build an executable:

```bash
go build -o gentle-wanderings main.go
./gentle-wanderings
```

## How to Play

1. **Start**: You begin in a Quiet Grove
2. **Explore**: Choose a cardinal direction (North, South, East, West)
3. **Choose**: Pick from 3 procedurally generated location options
4. **Discover**: Each new location is added to your map and journal
5. **Continue**: Keep exploring to build your unique world

### Commands

- **1-4**: Choose a direction to explore
- **1-3**: Choose which location option to visit
- **m** or **map**: View your current map (@ shows your position)
- **j** or **journal**: Read your journey log
- **q** or **quit**: End your session

## Code Structure

- **Game struct**: Holds all game state (map, position, journal)
- **Tile struct**: Represents each discovered location
- **Procedural Generation**: Random but themed location creation
- **Turn-based**: Each exploration is a new "day" in your journey

## Future Enhancement Ideas

### Graphics & UI
- Add a graphical tileset with sprites
- Implement a proper 2D rendering engine (Ebiten, Pixel, or SDL)
- Create beautiful hand-drawn location illustrations
- Add smooth camera movement and transitions

### Gameplay Depth
- **Oracle Tables**: Add prompt tables like table fables (What do you find? What happens next?)
- **Themes & Biomes**: Different regions with unique generation rules
- **Items & Memories**: Collect keepsakes from locations
- **Characters**: Meet other wanderers with their own stories
- **Seasons**: Time passing that changes the world's appearance
- **Questions**: Player answers questions that shape the world

### Persistence
- Save/load game state to JSON
- Multiple save slots
- Export map as image
- Share seeds for reproducible worlds

### Procedural Variety
- More location types and themes
- Weather and time-of-day variations
- Rare "special" locations with unique mechanics
- Connected location chains (forest → deep forest → ancient grove)

### Polish
- Background music and ambient sounds
- Achievements for exploration milestones
- Color-coded terminal output
- Animated ASCII art transitions
- Story generator based on your path

## TODO Additions

To evolve gameplay, consider adding:

1. **Prompt Cards**: "When you arrive, you notice..." with multiple choice outcomes
2. **Domains**: Specialized areas (The Depths, The Wilds, The Ruins)
3. **Landmarks**: Special locations that tell bigger stories
4. **Journaling Prompts**: Questions that help players narrate their experience
5. **Danger/Delight**: Balance peaceful with mysterious/challenging moments
6. **Rendering**: Begin building a library to autogenerate the map
7. **Generated Dialogue**: Make the journalling entries from input with generated AI
8. **Confrontation**: Currnetly there are no consequences or decision making beyond the map.

## Example Play Session

```
🌿 Quiet Grove
A peaceful clearing surrounded by ancient trees, dappled sunlight filtering through the leaves.

You begin your journey here, where the world feels safe and full of possibility.

Where would you like to wander?
  1. Explore North
  2. Explore South
  3. Explore East
  4. Explore West

> 1

✨ As you head North, three paths reveal themselves:

1. Whispering Willows
   A mysterious whispering willows where soft light dances across the space. Something calls to you.

2. Foggy Hollow
   An ancient foggy hollow where gentle sounds echo from the space. You feel drawn here.

3. Morning Mist
   A peaceful morning mist where shadows play among the space. Time seems to slow.

Which path calls to you? (1-3): 2

🌿 Foggy Hollow
An ancient foggy hollow where gentle sounds echo from the space. You feel drawn here.

As you arrive at foggy hollow, you notice details you hadn't expected.
```

## Technical Notes

- Pure Go with no external dependencies
- Uses `math/rand` for procedural generation
- Terminal-based interface using standard input/output
- Map stored as hashmap for efficient sparse grid

## License

MIT - Feel free to use, modify, and expand upon this project!

---

Happy wandering! 🌙✨