

# Feature Showcase 🎨

This document showcases the enhanced CLI rendering features of Gentle Wanderings.

Enhanced Map Display
The map uses Unicode box-drawing characters for a polished look:

╔═══════════════════════════╗
║         Your Map          ║
╠═══════════════════════════╣
║  ·  ·  ■  ·  ║
║  ·  ■  📍 ■  ║
║  🎁 ■  ■  ·  ║
║  ·  ·  ·  ·  ║
╚═══════════════════════════╝

Legend: 📍 You  ■ Explored  🎁 Has Item  · Unexplored
Features:

📍 Shows your current position
■ Shows explored tiles
🎁 Shows tiles that contain items
· Shows unexplored adjacent areas
Automatically scales to your map size
Padded border keeps everything centered
Detailed Map View
See all your discoveries with coordinates and items:

╔════════════════════════════════════════════════════════════╗
║                        Detailed Map                        ║
╚════════════════════════════════════════════════════════════╝

■ Foggy Hollow (0,1)
🎁 Contains: Smooth River Stone
📍 Morning Mist (1,1)
■ Quiet Grove (0,0)
■ Babbling Brook (1,0)
🎁 Contains: Ancient Coin
■ Mushroom Circle (-1,0)
Inventory System
Items are organized by category with detailed information:

╔════════════════════════════════════════════════════════════╗
║                     Your Collection                        ║
╚════════════════════════════════════════════════════════════╝

🍃 Keepsakes (2)
────────────────────────────────────────────────────────────
1. Smooth River Stone
   A simple treasure that reminds you of this moment.
   Found at Foggy Hollow on Day 2

2. Pressed Flower
   Something small but meaningful.
   Found at Wildflower Meadow on Day 5

💎 Treasures (1)
────────────────────────────────────────────────────────────
1. Ancient Coin
   It glimmers softly in your hand, valuable yet mysterious.
   Found at Crystal Pool on Day 3

❓ Curiosities (1)
────────────────────────────────────────────────────────────
1. Strange Map Fragment
   This raises more questions than it answers.
   Found at Stone Circle on Day 4

────────────────────────────────────────────────────────────
Total items collected: 4
Item Categories:

🍃 Keepsakes: Simple, meaningful treasures (stones, flowers, feathers)
💎 Treasures: Valuable finds (coins, gems, jewelry)
❓ Curiosities: Mysterious items with stories (maps, notes, scrolls)
Interactive Menu
A comprehensive menu system for accessing all features:

╔════════════════════════════════════════════════════════════╗
║                          Menu                              ║
╠════════════════════════════════════════════════════════════╣
║  1. View Map                                               ║
║  2. Detailed Map (with locations)                          ║
║  3. View Inventory                                         ║
║  4. Read Journal                                           ║
║  5. Current Location Info                                  ║
║  6. Game Statistics                                        ║
║  7. Return to Journey                                      ║
╚════════════════════════════════════════════════════════════╝
Location Discovery
When you discover a new location, it's presented beautifully:

╔════════════════════════════════════════════════════════════╗
║                        Foggy Hollow                        ║
╚════════════════════════════════════════════════════════════╝

An ancient foggy hollow where gentle sounds echo from the space.
You feel drawn here.

The foggy hollow reveals itself slowly, inviting you to linger.

────────────────────────────────────────────────────────────

✨ You found something! ✨

🎁 Smooth River Stone
A simple treasure that reminds you of this moment.
Statistics View
Track your progress with detailed statistics:

╔════════════════════════════════════════════════════════════╗
║                        Statistics                          ║
╚════════════════════════════════════════════════════════════╝

🗓️  Days Traveled: 7
🗺️  Locations Discovered: 8
🎒 Items Collected: 4

Collection breakdown:
🍃 Keepsakes: 2
💎 Treasures: 1
❓ Curiosities: 1

🧭 Map Dimensions: 3 × 3
📏 Furthest North: 1, South: -1, East: 1, West: -1
Journal
Your journey is automatically documented:

╔════════════════════════════════════════════════════════════╗
║                       Your Journey                         ║
╚════════════════════════════════════════════════════════════╝

Day 1: You begin your journey here, where the world feels safe and full of possibility.
Day 2: As you arrive at foggy hollow, you notice details you hadn't expected.
→ Found: Smooth River Stone
Day 3: You discover crystal pool and feel a deep connection to this place.
→ Found: Ancient Coin
Day 4: The stone circle reveals itself slowly, inviting you to linger.
→ Found: Strange Map Fragment
Current Location Info
Get detailed information about where you are:

╔════════════════════════════════════════════════════════════╗
║                     Current Location                       ║
╚════════════════════════════════════════════════════════════╝

🌿 Foggy Hollow
📍 Position: (0, 1)

An ancient foggy hollow where gentle sounds echo from the space.
You feel drawn here.

🎁 You found: Smooth River Stone
A simple treasure that reminds you of this moment.
Exit Summary
When you finish playing, see a summary of your adventure:

╔════════════════════════════════════════════════════════════╗
║                      Journey Summary                       ║
╚════════════════════════════════════════════════════════════╝

🗓️  Days traveled: 7
🗺️  Locations discovered: 8
🎒 Items collected: 4

Thank you for wandering with us. Until next time... 🌙✨
All of these features work in any terminal that supports Unicode and emoji! The game automatically handles map scaling and text centering to create a polished, cozy experience.

