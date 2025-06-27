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

// GetCurrentStageStartPositions returns the starting positions for the current stage
func (sl *StageLoader) GetCurrentStageStartPositions() (blueX, blueY, redX, redY float64) {
	switch sl.CurrentStageIndex {
	case 1:
		return GetStage1StartPositions()
	case 2:
		return GetStage2StartPositions()
	case 3:
		return GetStage3StartPositions()
	case 4:
		return GetStage4StartPositions()
	case 5:
		return GetStage5StartPositions()
	case 6:
		return GetStage6StartPositions()
	case 7:
		return GetStage7StartPositions()
	case 8:
		return GetStage8StartPositions()
	case 9:
		return GetStage9StartPositions()
	case 10:
		return GetStage10StartPositions()
	default:
		// Default to stage 1 positions if invalid index
		return GetStage1StartPositions()
	}
}

// Common platform colors and definitions
var (
	GroundColor   = color.RGBA{100, 100, 100, 255} // Gray for ground
	PlatformColor = color.RGBA{150, 150, 150, 255} // Light gray for platforms
	GoalColor     = color.RGBA{255, 255, 0, 255}   // Yellow for goal platforms
	SpikeColor    = color.RGBA{255, 0, 0, 255}     // Red for spikes
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

// Grid-based helper functions for platform creation

// CreateGridGroundPlatform creates a ground platform using grid coordinates
// Default: full width, 2.5 cells high at bottom
func CreateGridGroundPlatform() Platform {
	gridPlatform := GridPlatform{
		Position: GridPosition{X: 0, Y: GridHeight - 3}, // 3 cells from bottom (2.5 rounded up)
		Size:     GridSize{Width: GridWidth, Height: 3},
		IsGoal:   false,
	}
	return GridPlatformToPlatform(gridPlatform, GroundColor)
}

// CreateGridPlatform creates a platform using grid coordinates
func CreateGridPlatform(x, y, width, height int) Platform {
	gridPlatform := GridPlatform{
		Position: GridPosition{X: x, Y: y},
		Size:     GridSize{Width: width, Height: height},
		IsGoal:   false,
	}
	return GridPlatformToPlatform(gridPlatform, PlatformColor)
}

// CreateGridGoalPlatform creates a goal platform using grid coordinates
func CreateGridGoalPlatform(x, y, width, height int) Platform {
	gridPlatform := GridPlatform{
		Position: GridPosition{X: x, Y: y},
		Size:     GridSize{Width: width, Height: height},
		IsGoal:   true,
	}
	return GridPlatformToPlatform(gridPlatform, GoalColor)
}

// CreateGridSpike creates a spike using grid coordinates
func CreateGridSpike(x, y int) Spike {
	return Spike{
		X:     GridToPixelX(x),
		Y:     GridToPixelY(y),
		Color: SpikeColor,
	}
}
