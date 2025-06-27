package main

// LoadStage1 creates stage 1 - Generated from ASCII art
// Grid layout: 40x30 cells (800x600 pixels with 20px cells)
func LoadStage1() *Stage {
	return &Stage{
		Platforms: []Platform{

			// Regular platform at (0, 0) size 40x1
			CreateGridPlatform(0, 0, 40, 1),

			// Regular platform at (4, 27) size 1x3
			CreateGridPlatform(4, 27, 1, 3),

			// Regular platform at (35, 27) size 1x3
			CreateGridPlatform(35, 27, 1, 3),

			// Regular platform at (0, 28) size 40x2
			CreateGridPlatform(0, 28, 40, 2),

			// Goal platform at (19, 27) size 2x1
			CreateGridGoalPlatform(19, 27, 2, 1),

		},
	}
}

// GetStage1StartPositions returns the starting positions for stage 1
func GetStage1StartPositions() (blueX, blueY, redX, redY float64) {
	// Convert grid coordinates to pixel coordinates
	return 0, 540, 780, 540
}
