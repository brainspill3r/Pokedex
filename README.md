# 🎮 Pokédex CLI 

A command-line Pokédex application built with Go that lets you explore the Pokémon world, catch Pokémon, and build your personal collection! 

## ✨ Features

🗺️ **Explore the World**
- Navigate through 20+ location areas with `map` and `mapb` commands
- Discover Pokémon in different areas with the `explore` command

🎯 **Catch Pokémon** 
- Throw Pokéballs at wild Pokémon with realistic catch mechanics
- Difficulty based on Pokémon's base experience level
- Build your personal Pokédex collection

📊 **Detailed Pokémon Info**
- View stats (HP, Attack, Defense, Special Attack, Special Defense, Speed)
- Check types, height, weight, and base experience
- Only inspect Pokémon you've caught!

⚡ **Performance Optimized**
- Built-in caching system for lightning-fast responses
- Thread-safe cache with automatic cleanup
- No repeated API calls for the same data

## 🚀 Installation

### Prerequisites
- Go 1.22 or higher installed on your system

### Setup
```bash
# Clone or download the project
git clone <your-repo-url>
cd pokedexcli

# Initialize Go module (if not already done)
go mod init pokedexcli

# Build the project
go build .

# Run the application
./pokedexcli
```

Or run directly with:
```bash
go run .
```

## 🎮 Usage

Once you start the application, you'll see the Pokédex prompt:

```
Pokedex > 
```

Type `help` to see all available commands!

### 📝 Commands

| Command | Description | Example |
|---------|-------------|---------|
| `help` | 📖 Show all available commands | `help` |
| `map` | 🗺️ Display next 20 location areas | `map` |
| `mapb` | ⬅️ Display previous 20 location areas | `mapb` |
| `explore <area>` | 🔍 Explore a location to find Pokémon | `explore canalave-city-area` |
| `catch <pokemon>` | ⚽ Attempt to catch a Pokémon | `catch pikachu` |
| `inspect <pokemon>` | 👁️ View details of caught Pokémon | `inspect pikachu` |
| `pokedex` | 📚 List all your caught Pokémon | `pokedex` |
| `exit` | 🚪 Exit the application | `exit` |

## 🎯 Example Session

```bash
Pokedex > map
canalave-city-area
eterna-city-area
pastoria-city-area
# ... more locations

Pokedex > explore pastoria-city-area
Exploring pastoria-city-area...
Found Pokemon:
 - tentacool
 - tentacruel
 - magikarp
 - gyarados

Pokedex > catch magikarp
Throwing a Pokeball at magikarp...
magikarp was caught!
You may now inspect it with the inspect command.

Pokedex > inspect magikarp
Name: magikarp
Height: 9
Weight: 100
Stats:
  -hp: 20
  -attack: 10
  -defense: 55
  -special-attack: 15
  -special-defense: 20
  -speed: 80
Types:
  - water

Pokedex > pokedex
Your Pokedex:
 - magikarp
```

## 🏗️ Project Structure

```
pokedexcli/
├── main.go              # 🎮 Main application and command handlers
├── types.go             # 📋 Data structures and type definitions
├── pokeapi.go          # 🌐 PokeAPI integration and HTTP client
├── pokecache.go        # ⚡ Caching system implementation
├── pokecache_test.go   # 🧪 Cache tests
├── repl.go             # 🔧 Input processing utilities
├── repl_test.go        # 🧪 Input processing tests
└── go.mod              # 📦 Go module definition
```

## 🛠️ Technical Details

### Built With
- **Go 1.22+** - Core language
- **PokeAPI** - Pokemon data source (https://pokeapi.co/)
- **Native Go packages**: `net/http`, `encoding/json`, `math/rand`, `sync`, `time`

### Architecture Features
- **🏗️ Command Pattern** - Extensible command system
- **💾 Thread-Safe Caching** - Automatic cache expiration and cleanup
- **🎲 Randomized Mechanics** - Realistic catch probability system
- **📊 RESTful API Integration** - Full PokeAPI integration with error handling

### Catch Mechanics
The catch system uses a probability formula based on Pokémon difficulty:
```go
catchChance = max(10, 70 - baseExperience/5)
```
- 🟢 **Easy Pokémon** (low base exp): ~60-70% catch rate
- 🟡 **Medium Pokémon** (100-150 base exp): ~40-50% catch rate  
- 🔴 **Hard Pokémon** (200+ base exp): ~10-30% catch rate

## 🧪 Testing

Run the test suite:
```bash
# Run all tests
go test ./...

# Run with verbose output
go test -v ./...

# Run specific tests
go test -v . -run TestCleanInput
```

## 🤝 Contributing

1. Fork the project
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📄 License

This project is for educational purposes. Pokémon and related characters are trademarks of Nintendo, Game Freak, and Creatures.

