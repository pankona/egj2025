package main

import (
	"bytes"
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
	ScreenHeight = 600

	// Physics constants
	SPEED         = 1.0
	GRAVITY       = 0.35
	JUMP_STRENGTH = 13.0

	// Unit constants
	UnitSize = 20
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

type Stage struct {
	Platforms []Platform
}

type Game struct {
	BlueUnit *Unit
	RedUnit  *Unit
	Stage    *Stage
	State    GameState
	Font     *text.GoTextFace
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

	g.State = StatePlaying
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

	case StateGameOver, StateCleared:
		// Handle restart with space key
		if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
			g.resetGame()
		}

		// Handle touch input for retry - any touch triggers retry
		touchIDs := inpututil.AppendJustPressedTouchIDs(nil)
		if len(touchIDs) > 0 {
			g.resetGame()
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

	// Draw game state text with background
	switch g.State {
	case StateGameOver:
		// Draw semi-transparent background
		vector.DrawFilledRect(screen, 0, 0, ScreenWidth, ScreenHeight, color.RGBA{0, 0, 0, 150}, false)
		
		// Draw first line
		op1 := &text.DrawOptions{}
		op1.GeoM.Translate(float64(ScreenWidth/2-80), float64(ScreenHeight/2-30))
		op1.ColorScale.ScaleWithColor(color.RGBA{255, 255, 255, 255})
		text.Draw(screen, "GAME OVER", g.Font, op1)
		
		// Draw second line
		op2 := &text.DrawOptions{}
		op2.GeoM.Translate(float64(ScreenWidth/2-120), float64(ScreenHeight/2+10))
		op2.ColorScale.ScaleWithColor(color.RGBA{255, 255, 255, 255})
		text.Draw(screen, "Press SPACE to retry", g.Font, op2)

	case StateCleared:
		// Draw semi-transparent background
		vector.DrawFilledRect(screen, 0, 0, ScreenWidth, ScreenHeight, color.RGBA{0, 0, 0, 150}, false)
		
		// Draw first line
		op1 := &text.DrawOptions{}
		op1.GeoM.Translate(float64(ScreenWidth/2-100), float64(ScreenHeight/2-30))
		op1.ColorScale.ScaleWithColor(color.RGBA{255, 255, 255, 255})
		text.Draw(screen, "STAGE CLEARED!", g.Font, op1)
		
		// Draw second line
		op2 := &text.DrawOptions{}
		op2.GeoM.Translate(float64(ScreenWidth/2-120), float64(ScreenHeight/2+10))
		op2.ColorScale.ScaleWithColor(color.RGBA{255, 255, 255, 255})
		text.Draw(screen, "Press SPACE to retry", g.Font, op2)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
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
		Stage: &Stage{
			Platforms: []Platform{
				// Ground platform
				{X: 0, Y: float64(ScreenHeight - 50), Width: float64(ScreenWidth), Height: 50, Color: color.RGBA{100, 100, 100, 255}, IsGoal: false},
				// Some platforms for testing
				{X: 200, Y: 400, Width: 150, Height: 20, Color: color.RGBA{150, 150, 150, 255}, IsGoal: false},
				{X: 450, Y: 300, Width: 150, Height: 20, Color: color.RGBA{150, 150, 150, 255}, IsGoal: false},
				{X: 100, Y: 200, Width: 100, Height: 20, Color: color.RGBA{150, 150, 150, 255}, IsGoal: false},
				{X: 550, Y: 200, Width: 100, Height: 20, Color: color.RGBA{150, 150, 150, 255}, IsGoal: false},
				// Goal platform (at the bottom center)
				{X: 350, Y: float64(ScreenHeight - 70), Width: 100, Height: 20, Color: color.RGBA{255, 255, 0, 255}, IsGoal: true},
			},
		},
		State: StatePlaying,
		Font:  font,
	}
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
