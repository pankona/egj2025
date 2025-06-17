package main

import "image/color"

// StageLoader manages stage loading and progression
type StageLoader struct {
	CurrentStageIndex int
	TotalStages       int
}

// NewStageLoader creates a new stage loader
func NewStageLoader() *StageLoader {
	return &StageLoader{
		CurrentStageIndex: 1, // Start from stage 1
		// NOTE: TotalStages is hardcoded and must be updated manually when adding stages
		// TODO: Consider dynamic stage counting for better scalability
		TotalStages: 10,
	}
}

// LoadStage loads the stage by index
// NOTE: Using switch statement for stage loading. For better scalability with many stages,
// consider using a map[int]func() *Stage approach for dynamic stage registration.
func (sl *StageLoader) LoadStage(stageIndex int) *Stage {
	switch stageIndex {
	case 1:
		return LoadStage1()
	case 2:
		return LoadStage2()
	case 3:
		return LoadStage3()
	case 4:
		return LoadStage4()
	case 5:
		return LoadStage5()
	case 6:
		return LoadStage6()
	case 7:
		return LoadStage7()
	case 8:
		return LoadStage8()
	case 9:
		return LoadStage9()
	case 10:
		return LoadStage10()
	default:
		// Default to stage 1 if invalid index
		return LoadStage1()
	}
}

// GetCurrentStage returns the current stage
func (sl *StageLoader) GetCurrentStage() *Stage {
	return sl.LoadStage(sl.CurrentStageIndex)
}

// NextStage advances to the next stage
func (sl *StageLoader) NextStage() bool {
	if sl.CurrentStageIndex < sl.TotalStages {
		sl.CurrentStageIndex++
		return true
	}
	return false
}

// PreviousStage goes back to the previous stage
func (sl *StageLoader) PreviousStage() bool {
	if sl.CurrentStageIndex > 1 {
		sl.CurrentStageIndex--
		return true
	}
	return false
}

// ResetToFirstStage resets to stage 1
func (sl *StageLoader) ResetToFirstStage() {
	sl.CurrentStageIndex = 1
}

// Common platform colors and definitions
var (
	GroundColor   = color.RGBA{100, 100, 100, 255} // Gray for ground
	PlatformColor = color.RGBA{150, 150, 150, 255} // Light gray for platforms
	GoalColor     = color.RGBA{255, 255, 0, 255}   // Yellow for goal platforms
)

// Helper functions for common platform types
func CreateGroundPlatform() Platform {
	return Platform{
		X:      0,
		Y:      float64(ScreenHeight - 50),
		Width:  float64(ScreenWidth),
		Height: 50,
		Color:  GroundColor,
		IsGoal: false,
	}
}

func CreatePlatform(x, y, width, height float64) Platform {
	return Platform{
		X:      x,
		Y:      y,
		Width:  width,
		Height: height,
		Color:  PlatformColor,
		IsGoal: false,
	}
}

func CreateGoalPlatform(x, y, width, height float64) Platform {
	return Platform{
		X:      x,
		Y:      y,
		Width:  width,
		Height: height,
		Color:  GoalColor,
		IsGoal: true,
	}
}