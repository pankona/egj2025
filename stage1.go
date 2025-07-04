package main

// LoadStage1 creates stage 1 - Generated from ASCII art
// Grid layout: 40x31 cells (800x620 pixels with 20px cells)
func LoadStage1() *Stage {
	return &Stage{
		Platforms: []Platform{

			// Regular platform at (0, 0) size 40x1
			CreateGridPlatform(0, 0, 40, 1),

			// Regular platform at (0, 1) size 1x30
			CreateGridPlatform(0, 1, 1, 30),

			// Regular platform at (39, 1) size 1x30
			CreateGridPlatform(39, 1, 1, 30),

			// Regular platform at (5, 28) size 1x3
			CreateGridPlatform(5, 28, 1, 3),

			// Regular platform at (34, 28) size 1x3
			CreateGridPlatform(34, 28, 1, 3),

			// Regular platform at (1, 29) size 39x2
			CreateGridPlatform(1, 29, 39, 2),

			// Goal platform at (19, 27) size 2x2
			CreateGridGoalPlatform(19, 27, 2, 2),
		},
		Spikes: []Spike{},
	}
}

// GetStage1StartPositions returns the starting positions for stage 1
func GetStage1StartPositions() (blueX, blueY, redX, redY float64) {
	// Convert grid coordinates to pixel coordinates
	return 20, 560, 760, 560
}
