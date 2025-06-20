package main

import (
	"bytes"
	"fmt"
	"image/color"
	"io"
	"log"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/wav"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	ScreenWidth  = 800
	ScreenHeight = 600

	// Physics constants
	SPEED         = 1.0
	GRAVITY       = 0.35
	JUMP_STRENGTH = 13.0

	// Unit constants
	UnitSize = 20

	// Grid system constants
	CellSize   = 20                      // Each grid cell is 20x20 pixels (same as UnitSize)
	GridWidth  = ScreenWidth / CellSize  // 40 cells wide
	GridHeight = ScreenHeight / CellSize // 30 cells high

	// UI constants
	StageTextX = 10
	StageTextY = 30
)

var (
	// UI colors
	WhiteColor = color.RGBA{255, 255, 255, 255}
)

type GameState int

const (
	StatePlaying GameState = iota
	StateGameOver
	StateCleared
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
	IsGoal              bool // Mark this platform as a goal zone
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
	Position GridPosition
	Size     GridSize
	IsGoal   bool
}

type Stage struct {
	Platforms []Platform
}

// BGMManager manages background music playback
type BGMManager struct {
	audioContext *audio.Context
	player       *audio.Player
	isPlaying    bool
	isPaused     bool
}

// NewBGMManager creates a new BGM manager
func NewBGMManager() *BGMManager {
	audioContext := audio.NewContext(44100)
	return &BGMManager{
		audioContext: audioContext,
		isPlaying:    false,
		isPaused:     false,
	}
}

// LoadBGM loads BGM from a WAV data stream
func (bgm *BGMManager) LoadBGM(wavData []byte) error {
	// Stop current BGM if playing
	bgm.Stop()

	// Create audio stream from WAV data
	wavStream := bytes.NewReader(wavData)
	decodedStream, err := wav.DecodeWithSampleRate(bgm.audioContext.SampleRate(), wavStream)
	if err != nil {
		return fmt.Errorf("failed to decode WAV: %w", err)
	}

	// Create infinite loop stream
	loopStream := audio.NewInfiniteLoop(decodedStream, decodedStream.Length())

	// Create player
	player, err := bgm.audioContext.NewPlayer(loopStream)
	if err != nil {
		return fmt.Errorf("failed to create audio player: %w", err)
	}

	bgm.player = player
	return nil
}

// Play starts playing the BGM
func (bgm *BGMManager) Play() {
	if bgm.player == nil {
		return
	}

	if bgm.isPaused {
		// Resume from pause
		bgm.player.Play()
		bgm.isPaused = false
		bgm.isPlaying = true
	} else if !bgm.isPlaying {
		// Start from beginning
		bgm.player.Rewind()
		bgm.player.Play()
		bgm.isPlaying = true
	}
}

// Pause pauses the BGM playback
func (bgm *BGMManager) Pause() {
	if bgm.player != nil && bgm.isPlaying {
		bgm.player.Pause()
		bgm.isPaused = true
		bgm.isPlaying = false
	}
}

// Stop stops the BGM playback
func (bgm *BGMManager) Stop() {
	if bgm.player != nil {
		bgm.player.Pause()
		bgm.player.Rewind()
		bgm.isPlaying = false
		bgm.isPaused = false
	}
}

// Resume resumes the BGM playback from pause
func (bgm *BGMManager) Resume() {
	bgm.Play()
}

// SetVolume sets the BGM volume (0.0 to 1.0)
func (bgm *BGMManager) SetVolume(volume float64) {
	if bgm.player != nil {
		bgm.player.SetVolume(volume)
	}
}

// IsPlaying returns whether BGM is currently playing
func (bgm *BGMManager) IsPlaying() bool {
	return bgm.isPlaying
}

// IsPaused returns whether BGM is currently paused
func (bgm *BGMManager) IsPaused() bool {
	return bgm.isPaused
}

// Close releases resources
func (bgm *BGMManager) Close() {
	if bgm.player != nil {
		bgm.player.Close()
	}
}

type Game struct {
	BlueUnit    *Unit
	RedUnit     *Unit
	Stage       *Stage
	State       GameState
	Font        *text.GoTextFace
	StageLoader *StageLoader
	BGM         *BGMManager
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
		X:      GridToPixelX(gridPlatform.Position.X),
		Y:      GridToPixelY(gridPlatform.Position.Y),
		Width:  GridToPixelSize(gridPlatform.Size.Width),
		Height: GridToPixelSize(gridPlatform.Size.Height),
		Color:  color,
		IsGoal: gridPlatform.IsGoal,
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

	// Apply horizontal movement only if not stopped
	if !u.Stopped {
		u.VX = SPEED * float64(u.Direction)
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

		// Check if unit is horizontally overlapping with platform
		horizontalOverlap := unitRight > platformLeft && unitLeft < platformRight

		// Landing on top of platform (falling down) - skip goal platforms
		if !platform.IsGoal && horizontalOverlap && u.VY > 0 && unitBottom > platformTop && unitTop < platformTop {
			u.Y = platformTop - UnitSize
			u.VY = 0
			u.OnGround = true
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

func (u *Unit) jump() {
	if u.OnGround {
		u.VY = -JUMP_STRENGTH
		u.OnGround = false
	}
}

func (g *Game) checkGameOver() bool {
	// Check if either unit fell off the screen
	return g.BlueUnit.Y > float64(ScreenHeight) || g.RedUnit.Y > float64(ScreenHeight)
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
	// Reset units to starting positions
	g.BlueUnit.X = 100
	g.BlueUnit.Y = 100
	g.BlueUnit.VX = SPEED
	g.BlueUnit.VY = 0
	g.BlueUnit.Direction = 1
	g.BlueUnit.OnGround = false
	g.BlueUnit.Stopped = false

	g.RedUnit.X = 600
	g.RedUnit.Y = 100
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
	} else {
		// No more stages, restart from stage 1
		g.StageLoader.ResetToFirstStage()
		g.resetGame()
	}
}

// BGM convenience methods for game states

// StartBGM starts the background music
func (g *Game) StartBGM() {
	if g.BGM != nil {
		g.BGM.Play()
	}
}

// StopBGM stops the background music
func (g *Game) StopBGM() {
	if g.BGM != nil {
		g.BGM.Stop()
	}
}

// PauseBGM pauses the background music
func (g *Game) PauseBGM() {
	if g.BGM != nil {
		g.BGM.Pause()
	}
}

// ResumeBGM resumes the background music
func (g *Game) ResumeBGM() {
	if g.BGM != nil {
		g.BGM.Resume()
	}
}

// SetBGMVolume sets the BGM volume (0.0 to 1.0)
func (g *Game) SetBGMVolume(volume float64) {
	if g.BGM != nil {
		g.BGM.SetVolume(volume)
	}
}

// LoadBGMFromData loads BGM from WAV data
func (g *Game) LoadBGMFromData(wavData []byte) error {
	if g.BGM != nil {
		return g.BGM.LoadBGM(wavData)
	}
	return fmt.Errorf("BGM manager not initialized")
}

func (g *Game) Update() error {
	switch g.State {
	case StatePlaying:
		// Handle keyboard input
		// F key for blue unit jump
		if inpututil.IsKeyJustPressed(ebiten.KeyF) {
			g.BlueUnit.jump()
		}

		// J key for red unit jump
		if inpututil.IsKeyJustPressed(ebiten.KeyJ) {
			g.RedUnit.jump()
		}

		// BGM control keys for testing
		// M key to start/resume BGM
		if inpututil.IsKeyJustPressed(ebiten.KeyM) {
			if g.BGM != nil && !g.BGM.IsPlaying() {
				g.StartBGM()
			}
		}

		// N key to pause BGM
		if inpututil.IsKeyJustPressed(ebiten.KeyN) {
			if g.BGM != nil && g.BGM.IsPlaying() {
				g.PauseBGM()
			}
		}

		// B key to stop BGM
		if inpututil.IsKeyJustPressed(ebiten.KeyB) {
			if g.BGM != nil {
				g.StopBGM()
			}
		}

		// Handle touch input for gameplay
		touchIDs := inpututil.AppendJustPressedTouchIDs(nil)
		for _, id := range touchIDs {
			x, _ := ebiten.TouchPosition(id)
			// Left half of screen = F key (blue unit jump)
			if x < ScreenWidth/2 {
				g.BlueUnit.jump()
			} else {
				// Right half of screen = J key (red unit jump)
				g.RedUnit.jump()
			}
		}

		// Update physics for both units
		g.BlueUnit.updatePhysics(g.Stage)
		g.RedUnit.updatePhysics(g.Stage)

		// Check game state conditions
		if g.checkGameOver() {
			g.State = StateGameOver
		} else if g.checkCleared() {
			g.State = StateCleared
		}

	case StateGameOver:
		// Handle restart with space key
		if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
			g.resetGame()
		}

		// Handle touch input for retry - any touch triggers retry
		touchIDs := inpututil.AppendJustPressedTouchIDs(nil)
		if len(touchIDs) > 0 {
			g.resetGame()
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
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
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

	// Draw stage number in top-left corner during gameplay
	if g.State == StatePlaying {
		stageText := fmt.Sprintf("Stage %d", g.StageLoader.CurrentStageIndex)
		op := &text.DrawOptions{}
		op.GeoM.Translate(StageTextX, StageTextY)
		op.ColorScale.ScaleWithColor(WhiteColor)
		text.Draw(screen, stageText, g.Font, op)

		// Display BGM status in top-right corner
		var bgmStatus string
		if g.BGM != nil {
			if g.BGM.IsPlaying() {
				bgmStatus = "BGM: Playing"
			} else if g.BGM.IsPaused() {
				bgmStatus = "BGM: Paused"
			} else {
				bgmStatus = "BGM: Stopped"
			}
		} else {
			bgmStatus = "BGM: Not available"
		}

		bgmOp := &text.DrawOptions{}
		bgmOp.GeoM.Translate(float64(ScreenWidth-150), StageTextY)
		bgmOp.ColorScale.ScaleWithColor(WhiteColor)
		text.Draw(screen, bgmStatus, g.Font, bgmOp)

		// Display BGM controls hint
		controlsText := "BGM: M=Play, N=Pause, B=Stop"
		controlsOp := &text.DrawOptions{}
		controlsOp.GeoM.Translate(10, float64(ScreenHeight-20))
		controlsOp.ColorScale.ScaleWithColor(WhiteColor)
		text.Draw(screen, controlsText, g.Font, controlsOp)
	}

	// Draw game state text with background
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

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}

// generateSampleBGM creates a simple synthetic BGM for testing purposes
// This generates a simple sine wave pattern as a demonstration
func generateSampleBGM() []byte {
	// Create a simple WAV header + data for demonstration
	// This is a minimal implementation for testing
	sampleRate := 44100
	duration := 2      // 2 seconds
	frequency := 440.0 // A4 note

	// Calculate number of samples
	numSamples := sampleRate * duration

	// Create WAV header (44 bytes)
	header := make([]byte, 44)

	// "RIFF" chunk descriptor
	copy(header[0:4], "RIFF")
	// File size - 8 bytes (will be updated)
	header[4] = byte((36 + numSamples*2) & 0xff)
	header[5] = byte(((36 + numSamples*2) >> 8) & 0xff)
	header[6] = byte(((36 + numSamples*2) >> 16) & 0xff)
	header[7] = byte(((36 + numSamples*2) >> 24) & 0xff)
	// "WAVE" format
	copy(header[8:12], "WAVE")

	// "fmt " sub-chunk
	copy(header[12:16], "fmt ")
	// Sub-chunk size (16 for PCM)
	header[16] = 16
	header[17] = 0
	header[18] = 0
	header[19] = 0
	// Audio format (1 for PCM)
	header[20] = 1
	header[21] = 0
	// Number of channels (1 for mono)
	header[22] = 1
	header[23] = 0
	// Sample rate
	header[24] = byte(sampleRate & 0xff)
	header[25] = byte((sampleRate >> 8) & 0xff)
	header[26] = byte((sampleRate >> 16) & 0xff)
	header[27] = byte((sampleRate >> 24) & 0xff)
	// Byte rate (sample rate * channels * bits per sample / 8)
	byteRate := sampleRate * 1 * 16 / 8
	header[28] = byte(byteRate & 0xff)
	header[29] = byte((byteRate >> 8) & 0xff)
	header[30] = byte((byteRate >> 16) & 0xff)
	header[31] = byte((byteRate >> 24) & 0xff)
	// Block align (channels * bits per sample / 8)
	header[32] = 2
	header[33] = 0
	// Bits per sample
	header[34] = 16
	header[35] = 0

	// "data" sub-chunk
	copy(header[36:40], "data")
	// Sub-chunk size (number of samples * channels * bits per sample / 8)
	dataSize := numSamples * 2
	header[40] = byte(dataSize & 0xff)
	header[41] = byte((dataSize >> 8) & 0xff)
	header[42] = byte((dataSize >> 16) & 0xff)
	header[43] = byte((dataSize >> 24) & 0xff)

	// Generate audio data (simple sine wave)
	data := make([]byte, numSamples*2) // 16-bit samples
	for i := 0; i < numSamples; i++ {
		// Generate sine wave sample
		t := float64(i) / float64(sampleRate)
		sample := int16(32767 * 0.1 * math.Sin(2*math.Pi*frequency*t)) // Low volume

		// Convert to little-endian bytes
		data[i*2] = byte(sample & 0xff)
		data[i*2+1] = byte((sample >> 8) & 0xff)
	}

	// Combine header and data
	result := make([]byte, len(header)+len(data))
	copy(result, header)
	copy(result[len(header):], data)

	return result
}

func main() {
	ebiten.SetWindowSize(ScreenWidth, ScreenHeight)
	ebiten.SetWindowTitle("UNION JUMPERS")
	ebiten.SetWindowResizable(true)

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

	// Create BGM manager
	bgmManager := NewBGMManager()

	// Load sample BGM for demonstration
	sampleBGMData := generateSampleBGM()
	if err := bgmManager.LoadBGM(sampleBGMData); err != nil {
		log.Printf("Warning: Failed to load sample BGM: %v", err)
	} else {
		// Set a reasonable volume for the demo
		bgmManager.SetVolume(0.3)
	}

	game := &Game{
		BlueUnit: &Unit{
			X:         100,
			Y:         100,
			VX:        SPEED,
			VY:        0,
			Direction: 1,
			Color:     color.RGBA{0, 100, 255, 255}, // Blue
			OnGround:  false,
			Stopped:   false,
		},
		RedUnit: &Unit{
			X:         600,
			Y:         100,
			VX:        -SPEED,
			VY:        0,
			Direction: -1,
			Color:     color.RGBA{255, 100, 100, 255}, // Red
			OnGround:  false,
			Stopped:   false,
		},
		Stage:       stageLoader.GetCurrentStage(), // Load first stage
		State:       StatePlaying,
		Font:        font,
		StageLoader: stageLoader,
		BGM:         bgmManager,
	}
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
