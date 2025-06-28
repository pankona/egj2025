package main

import (
	"bytes"
	"fmt"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	ScreenWidth  = 800
	ScreenHeight = 620

	// Physics constants
	SPEED         = 1.5 // Increased for better responsiveness
	GRAVITY       = 0.35
	JUMP_STRENGTH = 5.9 // Allows jumping over 2 platforms but not 3

	// Unit constants
	UnitSize = 20

	// Grid system constants
	CellSize   = 20                      // Each grid cell is 20x20 pixels (same as UnitSize)
	GridWidth  = ScreenWidth / CellSize  // 40 cells wide
	GridHeight = ScreenHeight / CellSize // 30 cells high

	// UI constants
	StageTextX = 10
	StageTextY = 30

	// Audio constants
	SampleRate = 44100
)

var (
	// UI colors
	WhiteColor = color.RGBA{255, 255, 255, 255}

	// Debug mode flag (initialized based on platform)
	DebugMode bool
)

type GameState int

const (
	StateTitle GameState = iota
	StatePlaying
	StateGameOver
	StateCleared
	StateAllCleared
)

type Unit struct {
	X, Y      float64
	VX, VY    float64
	Direction int // 1 for right, -1 for left
	Color     color.Color
	OnGround  bool
	Stopped   bool // Whether the unit has stopped at the goal
}

type Platform struct {
	X, Y, Width, Height float64
	Color               color.Color
	IsGoal              bool    // Mark this platform as a goal zone
	SpeedModifier       float64 // Speed multiplier when standing on this platform (1.0 = normal, >1.0 = faster, <1.0 = slower)
}

type Spike struct {
	X, Y  float64
	Color color.Color
}

// GridPosition represents a position in the grid coordinate system
type GridPosition struct {
	X, Y int
}

// GridSize represents the size in grid coordinates
type GridSize struct {
	Width, Height int
}

// GridPlatform represents a platform in grid coordinates
type GridPlatform struct {
	Position      GridPosition
	Size          GridSize
	IsGoal        bool
	SpeedModifier float64 // Speed multiplier when standing on this platform
}

type Stage struct {
	Platforms []Platform
	Spikes    []Spike
}

type Game struct {
	BlueUnit      *Unit
	RedUnit       *Unit
	Stage         *Stage
	State         GameState
	Font          *text.GoTextFace
	StageLoader   *StageLoader
	SoundManager  *SoundManager
	BlinkCounter  int  // Counter for blinking text animation
	BlinkVisible  bool // Whether blinking text is currently visible
}

// Grid coordinate conversion functions

// GridToPixelX converts grid X coordinate to pixel X coordinate
func GridToPixelX(gridX int) float64 {
	return float64(gridX * CellSize)
}

// GridToPixelY converts grid Y coordinate to pixel Y coordinate
func GridToPixelY(gridY int) float64 {
	return float64(gridY * CellSize)
}

// GridToPixelSize converts grid size to pixel size
func GridToPixelSize(gridSize int) float64 {
	return float64(gridSize * CellSize)
}

// PixelToGridX converts pixel X coordinate to grid X coordinate
func PixelToGridX(pixelX float64) int {
	return int(pixelX / CellSize)
}

// PixelToGridY converts pixel Y coordinate to grid Y coordinate
func PixelToGridY(pixelY float64) int {
	return int(pixelY / CellSize)
}

// GridPlatformToPlatform converts a GridPlatform to a Platform with pixel coordinates
func GridPlatformToPlatform(gridPlatform GridPlatform, color color.Color) Platform {
	return Platform{
		X:             GridToPixelX(gridPlatform.Position.X),
		Y:             GridToPixelY(gridPlatform.Position.Y),
		Width:         GridToPixelSize(gridPlatform.Size.Width),
		Height:        GridToPixelSize(gridPlatform.Size.Height),
		Color:         color,
		IsGoal:        gridPlatform.IsGoal,
		SpeedModifier: gridPlatform.SpeedModifier,
	}
}

func (u *Unit) checkCollisionWithPlatform(platform Platform) bool {
	unitLeft := u.X
	unitRight := u.X + UnitSize
	unitTop := u.Y
	unitBottom := u.Y + UnitSize

	platformLeft := platform.X
	platformRight := platform.X + platform.Width
	platformTop := platform.Y
	platformBottom := platform.Y + platform.Height

	return unitRight > platformLeft && unitLeft < platformRight &&
		unitBottom > platformTop && unitTop < platformBottom
}

func (u *Unit) updatePhysics(stage *Stage) {
	// Apply gravity
	u.VY += GRAVITY

	// Calculate current speed modifier based on platforms the unit is standing on
	speedModifier := 1.0
	if u.OnGround {
		for _, platform := range stage.Platforms {
			// Check if unit is standing on this platform
			unitLeft := u.X
			unitRight := u.X + UnitSize
			unitBottom := u.Y + UnitSize

			platformLeft := platform.X
			platformRight := platform.X + platform.Width
			platformTop := platform.Y

			// Check if unit is on top of platform (standing on it)
			if unitRight > platformLeft && unitLeft < platformRight &&
				unitBottom >= platformTop && unitBottom <= platformTop+5 { // Small tolerance for "on platform"
				speedModifier = platform.SpeedModifier
				break // Use the first matching platform's speed modifier
			}
		}
	}

	// Apply horizontal movement only if not stopped
	if !u.Stopped {
		u.VX = SPEED * float64(u.Direction) * speedModifier
		// Update horizontal position
		u.X += u.VX
	} else {
		u.VX = 0
	}

	// Wall collision (screen boundaries) - only if not stopped
	if !u.Stopped {
		if u.X <= 0 {
			u.X = 0
			u.Direction = 1 // Move right
		} else if u.X >= float64(ScreenWidth-UnitSize) {
			u.X = float64(ScreenWidth - UnitSize)
			u.Direction = -1 // Move left
		}
	}

	// Update vertical position
	u.Y += u.VY

	// Platform collision detection
	u.OnGround = false
	for _, platform := range stage.Platforms {
		unitLeft := u.X
		unitRight := u.X + UnitSize
		unitTop := u.Y
		unitBottom := u.Y + UnitSize

		platformLeft := platform.X
		platformRight := platform.X + platform.Width
		platformTop := platform.Y
		platformBottom := platform.Y + platform.Height

		// Check if unit is horizontally overlapping with platform
		horizontalOverlap := unitRight > platformLeft && unitLeft < platformRight
		// Check if unit is vertically overlapping with platform
		verticalOverlap := unitBottom > platformTop && unitTop < platformBottom

		// Landing on top of platform (falling down) - skip goal platforms
		if !platform.IsGoal && horizontalOverlap && u.VY > 0 && unitBottom > platformTop && unitTop < platformTop {
			u.Y = platformTop - UnitSize
			u.VY = 0
			u.OnGround = true
		}

		// Horizontal collision detection - skip goal platforms
		// Check horizontal collision for all non-goal platforms
		if !platform.IsGoal && verticalOverlap && !u.Stopped {
			// Only check horizontal collision if unit is not on top of this platform
			isOnTopOfPlatform := u.OnGround && unitBottom >= platformTop && unitBottom <= platformTop+5

			if !isOnTopOfPlatform {
				// Check collision from left side (moving right)
				if u.Direction > 0 && unitRight > platformLeft && unitLeft < platformLeft {
					u.X = platformLeft - UnitSize
					u.Direction = -1 // Reverse direction to left
				}
				// Check collision from right side (moving left)
				if u.Direction < 0 && unitLeft < platformRight && unitRight > platformRight {
					u.X = platformRight
					u.Direction = 1 // Reverse direction to right
				}
			}
		}
	}

	// Prevent falling through bottom of screen
	if u.Y > float64(ScreenHeight) {
		u.Y = float64(ScreenHeight - UnitSize)
		u.OnGround = true
		u.VY = 0
	}

	// Check if unit is completely inside goal platform area (for stopping and clearing)
	if u.OnGround {
		for _, platform := range stage.Platforms {
			if platform.IsGoal {
				unitLeft := u.X
				unitRight := u.X + UnitSize
				unitTop := u.Y
				unitBottom := u.Y + UnitSize
				platformLeft := platform.X
				platformRight := platform.X + platform.Width
				platformTop := platform.Y
				platformBottom := platform.Y + platform.Height

				// Check if unit is completely inside the goal platform
				if unitLeft >= platformLeft && unitRight <= platformRight &&
					unitTop >= platformTop && unitBottom <= platformBottom {
					u.Stopped = true
					break
				}
			}
		}
	}
}

func (u *Unit) jump(soundManager *SoundManager) {
	if u.OnGround {
		u.VY = -JUMP_STRENGTH
		u.OnGround = false
		soundManager.PlayJumpSound()
	}
}

func (g *Game) checkGameOver() bool {
	// Check if either unit fell off the screen
	if g.BlueUnit.Y > float64(ScreenHeight) || g.RedUnit.Y > float64(ScreenHeight) {
		return true
	}

	// Check if either unit touched a spike
	for _, spike := range g.Stage.Spikes {
		if g.checkUnitSpikeCollision(g.BlueUnit, spike) || g.checkUnitSpikeCollision(g.RedUnit, spike) {
			return true
		}
	}

	return false
}

func (g *Game) checkUnitSpikeCollision(unit *Unit, spike Spike) bool {
	unitLeft := unit.X
	unitRight := unit.X + UnitSize
	unitTop := unit.Y
	unitBottom := unit.Y + UnitSize

	spikeLeft := spike.X
	spikeRight := spike.X + CellSize
	spikeTop := spike.Y
	spikeBottom := spike.Y + CellSize

	// Check if unit and spike overlap
	return unitRight > spikeLeft && unitLeft < spikeRight &&
		unitBottom > spikeTop && unitTop < spikeBottom
}

func (g *Game) checkCleared() bool {
	// Check if both units are on goal platforms
	blueOnGoal := false
	redOnGoal := false

	for _, platform := range g.Stage.Platforms {
		if platform.IsGoal {
			if g.BlueUnit.checkCollisionWithPlatform(platform) && g.BlueUnit.OnGround {
				blueOnGoal = true
			}
			if g.RedUnit.checkCollisionWithPlatform(platform) && g.RedUnit.OnGround {
				redOnGoal = true
			}
		}
	}

	return blueOnGoal && redOnGoal
}

func (g *Game) resetGame() {
	// Get starting positions from current stage
	blueX, blueY, redX, redY := g.StageLoader.GetCurrentStageStartPositions()

	// Reset units to stage-specific starting positions
	g.BlueUnit.X = blueX
	g.BlueUnit.Y = blueY
	g.BlueUnit.VX = SPEED
	g.BlueUnit.VY = 0
	g.BlueUnit.Direction = 1
	g.BlueUnit.OnGround = false
	g.BlueUnit.Stopped = false

	g.RedUnit.X = redX
	g.RedUnit.Y = redY
	g.RedUnit.VX = -SPEED
	g.RedUnit.VY = 0
	g.RedUnit.Direction = -1
	g.RedUnit.OnGround = false
	g.RedUnit.Stopped = false

	// Reload current stage
	g.Stage = g.StageLoader.GetCurrentStage()
	g.State = StatePlaying
}

func (g *Game) advanceToNextStageOrRestart() {
	if g.StageLoader.NextStage() {
		// Advanced to next stage, reset game with new stage
		g.resetGame()
		// Restart BGM when advancing to next stage
		g.SoundManager.StartBGM()
	} else {
		// No more stages, go to all cleared state
		g.State = StateAllCleared
		g.SoundManager.StopBGM()
	}
}

func (g *Game) Update() error {
	// Update blinking animation for title and all cleared screens
	g.BlinkCounter++
	if g.BlinkCounter >= 30 { // Blink every 30 frames (0.5 seconds at 60 FPS)
		g.BlinkVisible = !g.BlinkVisible
		g.BlinkCounter = 0
	}

	// Ensure BGM is playing only during gameplay (NewInfiniteLoop handles the looping automatically)
	if g.State == StatePlaying && g.SoundManager.bgmPlayer != nil && !g.SoundManager.bgmPlayer.IsPlaying() {
		g.SoundManager.StartBGM()
	}

	switch g.State {
	case StateTitle:
		// Handle any key to start game
		// Check keyboard input
		keys := inpututil.AppendPressedKeys(nil)
		if len(keys) > 0 {
			g.State = StatePlaying
			g.SoundManager.StartBGM()
		}

		// Handle touch input
		touchIDs := inpututil.AppendJustPressedTouchIDs(nil)
		if len(touchIDs) > 0 {
			g.State = StatePlaying
			g.SoundManager.StartBGM()
		}

	case StatePlaying:
		// Handle keyboard input
		// F key for blue unit jump
		if inpututil.IsKeyJustPressed(ebiten.KeyF) {
			g.BlueUnit.jump(g.SoundManager)
		}

		// J key for red unit jump
		if inpututil.IsKeyJustPressed(ebiten.KeyJ) {
			g.RedUnit.jump(g.SoundManager)
		}

		// Handle touch input for gameplay
		touchIDs := inpututil.AppendJustPressedTouchIDs(nil)
		for _, id := range touchIDs {
			x, _ := ebiten.TouchPosition(id)
			// Left half of screen = F key (blue unit jump)
			if x < ScreenWidth/2 {
				g.BlueUnit.jump(g.SoundManager)
			} else {
				// Right half of screen = J key (red unit jump)
				g.RedUnit.jump(g.SoundManager)
			}
		}

		// Update physics for both units
		g.BlueUnit.updatePhysics(g.Stage)
		g.RedUnit.updatePhysics(g.Stage)

		// Check game state conditions
		if g.checkGameOver() {
			g.SoundManager.PlayDeadSound()
			g.State = StateGameOver
		} else if g.checkCleared() {
			g.SoundManager.StopBGM()
			g.SoundManager.PlayClearSound()
			g.State = StateCleared
		}

	case StateGameOver:
		// Handle restart with space key
		if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
			g.resetGame()
			g.SoundManager.StartBGM()
		}

		// Handle touch input for retry - any touch triggers retry
		touchIDs := inpututil.AppendJustPressedTouchIDs(nil)
		if len(touchIDs) > 0 {
			g.resetGame()
			g.SoundManager.StartBGM()
		}

	case StateCleared:
		// Handle next stage with space key
		if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
			g.advanceToNextStageOrRestart()
		}

		// Handle touch input for next stage - any touch advances
		touchIDs := inpututil.AppendJustPressedTouchIDs(nil)
		if len(touchIDs) > 0 {
			g.advanceToNextStageOrRestart()
		}

	case StateAllCleared:
		// Handle any key to restart from stage 1
		// Check keyboard input
		keys := inpututil.AppendPressedKeys(nil)
		if len(keys) > 0 {
			g.StageLoader.ResetToFirstStage()
			g.resetGame()
			g.SoundManager.StartBGM()
		}

		// Handle touch input
		touchIDs := inpututil.AppendJustPressedTouchIDs(nil)
		if len(touchIDs) > 0 {
			g.StageLoader.ResetToFirstStage()
			g.resetGame()
			g.SoundManager.StartBGM()
		}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	switch g.State {
	case StateTitle:
		// TODO: Add background image for title screen
		// Draw semi-transparent background for now
		vector.DrawFilledRect(screen, 0, 0, ScreenWidth, ScreenHeight, color.RGBA{20, 30, 50, 255}, false)

		// Draw title
		titleOp := &text.DrawOptions{}
		titleOp.GeoM.Translate(float64(ScreenWidth/2-120), float64(ScreenHeight/2-80))
		titleOp.ColorScale.ScaleWithColor(WhiteColor)
		text.Draw(screen, "UNION JUMPERS", g.Font, titleOp)

		// Draw blinking "Press any key to start" text
		if g.BlinkVisible {
			startOp := &text.DrawOptions{}
			startOp.GeoM.Translate(float64(ScreenWidth/2-130), float64(ScreenHeight/2-20))
			startOp.ColorScale.ScaleWithColor(WhiteColor)
			text.Draw(screen, "Press any key to start", g.Font, startOp)
		}

	case StateAllCleared:
		// TODO: Add background image for all cleared screen
		// Draw semi-transparent background for now
		vector.DrawFilledRect(screen, 0, 0, ScreenWidth, ScreenHeight, color.RGBA{50, 20, 50, 255}, false)

		// Draw congratulations message
		congratsOp := &text.DrawOptions{}
		congratsOp.GeoM.Translate(float64(ScreenWidth/2-140), float64(ScreenHeight/2-80))
		congratsOp.ColorScale.ScaleWithColor(color.RGBA{255, 255, 100, 255}) // Golden color
		text.Draw(screen, "Congratulations!", g.Font, congratsOp)

		// Draw completion message
		completeOp := &text.DrawOptions{}
		completeOp.GeoM.Translate(float64(ScreenWidth/2-120), float64(ScreenHeight/2-40))
		completeOp.ColorScale.ScaleWithColor(WhiteColor)
		text.Draw(screen, "All stages cleared!", g.Font, completeOp)

		// Draw blinking restart message
		if g.BlinkVisible {
			restartOp := &text.DrawOptions{}
			restartOp.GeoM.Translate(float64(ScreenWidth/2-120), float64(ScreenHeight/2+20))
			restartOp.ColorScale.ScaleWithColor(WhiteColor)
			text.Draw(screen, "Press any key to restart", g.Font, restartOp)
		}

	default:
		// Draw gameplay elements (StatePlaying, StateGameOver, StateCleared)
		// Draw platforms
		for _, platform := range g.Stage.Platforms {
			platformColor := platform.Color
			// Highlight goal platforms
			if platform.IsGoal {
				platformColor = color.RGBA{255, 255, 0, 255} // Yellow for goal
			}
			vector.DrawFilledRect(screen, float32(platform.X), float32(platform.Y), float32(platform.Width), float32(platform.Height), platformColor, false)
		}

		// Draw blue unit
		vector.DrawFilledRect(screen, float32(g.BlueUnit.X), float32(g.BlueUnit.Y), UnitSize, UnitSize, g.BlueUnit.Color, false)

		// Draw red unit
		vector.DrawFilledRect(screen, float32(g.RedUnit.X), float32(g.RedUnit.Y), UnitSize, UnitSize, g.RedUnit.Color, false)

		// Draw spikes as upward triangles
		for _, spike := range g.Stage.Spikes {
			// Define triangle vertices (upward pointing)
			x := float32(spike.X)
			y := float32(spike.Y)
			size := float32(CellSize)

			// Use vector.DrawFilledRect to create a simple spike representation
			// Draw as a smaller rectangle at the center to represent spike
			spikeSize := size * 0.8
			offset := (size - spikeSize) / 2
			vector.DrawFilledRect(screen, x+offset, y+offset, spikeSize, spikeSize, spike.Color, false)
		}

		// Draw stage number in top-left corner during gameplay
		if g.State == StatePlaying {
			stageText := fmt.Sprintf("Stage %d", g.StageLoader.CurrentStageIndex)
			op := &text.DrawOptions{}
			op.GeoM.Translate(StageTextX, StageTextY)
			op.ColorScale.ScaleWithColor(WhiteColor)
			text.Draw(screen, stageText, g.Font, op)
		}

		// Draw game state overlay text with background
		switch g.State {
		case StateGameOver:
			// Draw semi-transparent background
			vector.DrawFilledRect(screen, 0, 0, ScreenWidth, ScreenHeight, color.RGBA{0, 0, 0, 150}, false)

			// Draw first line
			op1 := &text.DrawOptions{}
			op1.GeoM.Translate(float64(ScreenWidth/2-80), float64(ScreenHeight/2-30))
			op1.ColorScale.ScaleWithColor(WhiteColor)
			text.Draw(screen, "GAME OVER", g.Font, op1)

			// Draw second line
			op2 := &text.DrawOptions{}
			op2.GeoM.Translate(float64(ScreenWidth/2-120), float64(ScreenHeight/2+10))
			op2.ColorScale.ScaleWithColor(WhiteColor)
			text.Draw(screen, "Press SPACE to retry", g.Font, op2)

		case StateCleared:
			// Draw semi-transparent background
			vector.DrawFilledRect(screen, 0, 0, ScreenWidth, ScreenHeight, color.RGBA{0, 0, 0, 150}, false)

			// Draw first line
			op1 := &text.DrawOptions{}
			op1.GeoM.Translate(float64(ScreenWidth/2-100), float64(ScreenHeight/2-30))
			op1.ColorScale.ScaleWithColor(WhiteColor)
			text.Draw(screen, "STAGE CLEARED!", g.Font, op1)

			// Draw second line
			op2 := &text.DrawOptions{}
			if g.StageLoader.CurrentStageIndex < g.StageLoader.TotalStages {
				op2.GeoM.Translate(float64(ScreenWidth/2-140), float64(ScreenHeight/2+10))
				op2.ColorScale.ScaleWithColor(WhiteColor)
				text.Draw(screen, "Press SPACE for next stage", g.Font, op2)
			} else {
				op2.GeoM.Translate(float64(ScreenWidth/2-120), float64(ScreenHeight/2+10))
				op2.ColorScale.ScaleWithColor(WhiteColor)
				text.Draw(screen, "Press SPACE to restart", g.Font, op2)
			}
		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}

func main() {
	// Initialize debug mode based on platform
	initDebugMode()

	// Log debug mode status
	if DebugMode {
		log.Println("Debug mode: ENABLED (starting from stage 0)")
	} else {
		log.Println("Debug mode: DISABLED (starting from stage 1)")
	}

	ebiten.SetWindowSize(ScreenWidth, ScreenHeight)
	ebiten.SetWindowTitle("UNION JUMPERS")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	// Initialize font
	fontSource, err := text.NewGoTextFaceSource(bytes.NewReader(fonts.MPlus1pRegular_ttf))
	if err != nil {
		log.Fatal(err)
	}
	font := &text.GoTextFace{
		Source: fontSource,
		Size:   24,
	}

	// Create stage loader
	stageLoader := NewStageLoader()

	// Create sound manager
	soundManager := NewSoundManager()

	// Get starting positions for the first stage
	blueX, blueY, redX, redY := stageLoader.GetCurrentStageStartPositions()

	game := &Game{
		BlueUnit: &Unit{
			X:         blueX,
			Y:         blueY,
			VX:        SPEED,
			VY:        0,
			Direction: 1,
			Color:     color.RGBA{0, 100, 255, 255}, // Blue
			OnGround:  false,
			Stopped:   false,
		},
		RedUnit: &Unit{
			X:         redX,
			Y:         redY,
			VX:        -SPEED,
			VY:        0,
			Direction: -1,
			Color:     color.RGBA{255, 100, 100, 255}, // Red
			OnGround:  false,
			Stopped:   false,
		},
		Stage:        stageLoader.GetCurrentStage(), // Load first stage
		State:        StateTitle,                    // Start with title screen
		Font:         font,
		StageLoader:  stageLoader,
		SoundManager: soundManager,
		BlinkCounter: 0,
		BlinkVisible: true,
	}
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
