package main

import (
	"bytes"
	"image/color"
	"testing"

	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

func createTestFont() *text.GoTextFace {
	fontSource, err := text.NewGoTextFaceSource(bytes.NewReader(fonts.MPlus1pRegular_ttf))
	if err != nil {
		return nil
	}
	return &text.GoTextFace{
		Source: fontSource,
		Size:   24,
	}
}

func TestGameLogic(t *testing.T) {
	// Create a test stage with a goal platform
	stage := &Stage{
		Platforms: []Platform{
			// Ground platform
			{X: 0, Y: 550, Width: 800, Height: 50, Color: color.RGBA{100, 100, 100, 255}, IsGoal: false},
			// Goal platform
			{X: 350, Y: 530, Width: 100, Height: 20, Color: color.RGBA{255, 255, 0, 255}, IsGoal: true},
		},
	}

	// Create units
	blueUnit := &Unit{
		X:         100,
		Y:         100,
		VX:        SPEED,
		VY:        0,
		Direction: 1,
		Color:     color.RGBA{0, 100, 255, 255},
		OnGround:  false,
		Stopped:   false,
	}

	redUnit := &Unit{
		X:         600,
		Y:         100,
		VX:        -SPEED,
		VY:        0,
		Direction: -1,
		Color:     color.RGBA{255, 100, 100, 255},
		OnGround:  false,
		Stopped:   false,
	}

	game := &Game{
		BlueUnit: blueUnit,
		RedUnit:  redUnit,
		Stage:    stage,
		State:    StatePlaying,
		Font:     createTestFont(),
	}

	t.Run("両キャラがゴールプラットフォームに到達していない場合はクリアしない", func(t *testing.T) {
		// Reset units
		game.BlueUnit.X = 100
		game.BlueUnit.Y = 530
		game.BlueUnit.OnGround = true
		game.BlueUnit.Stopped = false
		game.RedUnit.X = 600
		game.RedUnit.Y = 530
		game.RedUnit.OnGround = true
		game.RedUnit.Stopped = false

		if game.checkCleared() {
			t.Error("両キャラがゴールプラットフォームにいない場合はクリアしてはいけない")
		}
	})

	t.Run("片方のキャラのみがゴールプラットフォームに到達した場合はクリアしない", func(t *testing.T) {
		// Reset and position units - only blue is in goal
		game.BlueUnit.X = 375
		game.BlueUnit.Y = 535
		game.BlueUnit.OnGround = true
		game.BlueUnit.Stopped = true
		game.RedUnit.X = 600
		game.RedUnit.Y = 530
		game.RedUnit.OnGround = true
		game.RedUnit.Stopped = false

		if game.checkCleared() {
			t.Error("片方のキャラのみがゴールプラットフォームにいる場合はクリアしてはいけない")
		}
	})

	t.Run("両キャラがゴールプラットフォームに到達した場合はクリアする", func(t *testing.T) {
		// Position both units in goal platform
		game.BlueUnit.X = 360
		game.BlueUnit.Y = 535
		game.BlueUnit.OnGround = true
		game.BlueUnit.Stopped = true
		game.RedUnit.X = 380
		game.RedUnit.Y = 535
		game.RedUnit.OnGround = true
		game.RedUnit.Stopped = true

		if !game.checkCleared() {
			t.Error("両キャラがゴールプラットフォームにいる場合はクリアすべき")
		}
	})

	t.Run("キャラがゴールプラットフォームにいても地面についていない場合はクリアしない", func(t *testing.T) {
		// Position units in goal area but not on ground
		game.BlueUnit.X = 360
		game.BlueUnit.Y = 500
		game.BlueUnit.OnGround = false
		game.BlueUnit.Stopped = false
		game.RedUnit.X = 380
		game.RedUnit.Y = 500
		game.RedUnit.OnGround = false
		game.RedUnit.Stopped = false

		if game.checkCleared() {
			t.Error("キャラが地面についていない場合はクリアしてはいけない")
		}
	})
}

func TestUnitGoalDetection(t *testing.T) {
	// Test if units stop correctly when fully inside goal platform
	stage := &Stage{
		Platforms: []Platform{
			// Ground platform
			{X: 0, Y: 550, Width: 800, Height: 50, Color: color.RGBA{100, 100, 100, 255}, IsGoal: false},
			// Goal platform
			{X: 350, Y: 530, Width: 100, Height: 20, Color: color.RGBA{255, 255, 0, 255}, IsGoal: true},
		},
	}

	t.Run("キャラがゴールプラットフォームに完全に入ると停止状態になる", func(t *testing.T) {
		unit := &Unit{
			X:         360, // Inside goal platform (350-450)
			Y:         535, // Inside goal platform (530-550)
			VX:        SPEED,
			VY:        0,
			Direction: 1,
			Color:     color.RGBA{0, 100, 255, 255},
			OnGround:  true,
			Stopped:   false,
		}

		unit.updatePhysics(stage)

		if !unit.Stopped {
			t.Error("キャラがゴールプラットフォームに完全に入った場合は停止すべき")
		}
	})

	t.Run("キャラがゴールプラットフォームに部分的に重なっても停止しない", func(t *testing.T) {
		unit := &Unit{
			X:         340, // Partially outside goal platform (unit goes from 340-360, goal is 350-450)
			Y:         535,
			VX:        SPEED,
			VY:        0,
			Direction: 1,
			Color:     color.RGBA{0, 100, 255, 255},
			OnGround:  true,
			Stopped:   false,
		}

		unit.updatePhysics(stage)

		if unit.Stopped {
			t.Error("キャラがゴールプラットフォームに部分的にしか入っていない場合は停止してはいけない")
		}
	})

	t.Run("キャラがゴールプラットフォーム外にいる場合は停止しない", func(t *testing.T) {
		unit := &Unit{
			X:         100, // Outside goal platform
			Y:         530,
			VX:        SPEED,
			VY:        0,
			Direction: 1,
			Color:     color.RGBA{0, 100, 255, 255},
			OnGround:  true,
			Stopped:   false,
		}

		unit.updatePhysics(stage)

		if unit.Stopped {
			t.Error("キャラがゴールプラットフォーム外にいる場合は停止してはいけない")
		}
	})
}

func TestGameStateCleared(t *testing.T) {
	// Create a test stage with a goal platform
	stage := &Stage{
		Platforms: []Platform{
			// Ground platform
			{X: 0, Y: 550, Width: 800, Height: 50, Color: color.RGBA{100, 100, 100, 255}, IsGoal: false},
			// Goal platform
			{X: 350, Y: 530, Width: 100, Height: 20, Color: color.RGBA{255, 255, 0, 255}, IsGoal: true},
		},
	}

	// Create units positioned in goal
	blueUnit := &Unit{
		X:         360,
		Y:         535,
		VX:        0,
		VY:        0,
		Direction: 1,
		Color:     color.RGBA{0, 100, 255, 255},
		OnGround:  true,
		Stopped:   true,
	}

	redUnit := &Unit{
		X:         380,
		Y:         535,
		VX:        0,
		VY:        0,
		Direction: -1,
		Color:     color.RGBA{255, 100, 100, 255},
		OnGround:  true,
		Stopped:   true,
	}

	game := &Game{
		BlueUnit: blueUnit,
		RedUnit:  redUnit,
		Stage:    stage,
		State:    StatePlaying,
		Font:     createTestFont(),
	}

	t.Run("ゲームクリア状態でUpdate呼び出しが正常に動作する", func(t *testing.T) {
		// ゲームをクリア状態にする
		if game.checkCleared() {
			game.State = StateCleared
		}

		// Updateが正常に実行されることを確認
		err := game.Update()
		if err != nil {
			t.Errorf("Update()でエラーが発生: %v", err)
		}

		if game.State != StateCleared {
			t.Error("ゲーム状態がStateCleared状態を維持していない")
		}
	})
}