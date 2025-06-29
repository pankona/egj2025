# UNION JUMPERS

A two-handed puzzle action game where you control two characters simultaneously with your left and right hands.

Created for [Ebitengine Game Jam 2025](https://itch.io/jam/ebitengine-game-jam-2025).

## 🎮 Play Online

[Play UNION JUMPERS on GitHub Pages](https://pankona.github.io/egj2025/)

## 📖 About the Game

UNION JUMPERS is a cooperative puzzle platformer where players must guide two characters (blue and red) to the goal simultaneously. The unique twist is that each character is controlled by a different hand - creating a challenging coordination experience.

### Features

- **Dual Character Control**: Control blue character with left hand (F key) and red character with right hand (J key)
- **10 Stages**: Progressively challenging levels from tutorial to expert
- **Stage Gimmicks**:
  - Spikes (red triangles) - instant game over on contact
  - Speed-up platforms (green) - increases movement speed
  - Speed-down platforms (orange) - decreases movement speed
- **Cross-Platform**: Works on PC browsers and mobile devices

### Controls

**PC (Keyboard)**:

- `F` key: Jump (Blue character / Left hand)
- `J` key: Jump (Red character / Right hand)
- `Space`: Retry/Next stage

**Mobile/Tablet**:

- Tap left half of screen: Jump (Blue character)
- Tap right half of screen: Jump (Red character)

## 🛠️ Development

### Prerequisites

- Go 1.24 or later

### Building

```bash
# Build for WebAssembly
make build-wasm

# Run local development server (port 8080)
make serve-wasm
```

## 📁 Project Structure

```
egj2025/
├── main.go              # Main game logic
├── sound.go             # Sound system
├── assets.go            # Embedded assets
├── stage_loader.go      # Stage management
├── stage*.go            # Generated stage data
├── stage*.txt           # Stage definitions (40x31 ASCII grid)
├── assets/              # Game assets (images, sounds)
├── web/                 # Web files for WASM build
└── cmd/stagegen/        # Stage generation tool
```

## 🎵 Credits

### Development

- **Author**: pankona
- **AI Assistant**: Claude (Code assistance)

### Assets

- **Sound Effects**:
  - [Springin' Sound Stock](https://www.springin.org/sound-stock/)
  - [ポケットサウンド](https://pocket-se.info/)
  - [魔王魂](https://maou.audio)
  - [効果音ラボ](https://soundeffect-lab.info)
  - [Howling-Indicator](https://howlingindicator.net)
- **Font**: M+ Font (embedded in Ebitengine)

## 📝 License

MIT

## 🙏 Acknowledgments

Special thanks to:

- [Ebitengine](https://ebitengine.org/) community for the awesome game engine
- Ebitengine Game Jam 2025 organizers
- All playtesters and feedback providers
