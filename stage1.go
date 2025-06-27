package main

// LoadStage1 creates stage 1 - Generated from ASCII art
// Grid layout: 40x30 cells (800x600 pixels with 20px cells)
func LoadStage1() *Stage {
	return &Stage{
		Platforms: []Platform{

			// Regular platform at (0, 0) size 40x1
			CreateGridPlatform(0, 0, 40, 1),

			// Regular platform at (4, 25) size 1x5
			CreateGridPlatform(4, 25, 1, 5),

			// Regular platform at (35, 25) size 1x5
			CreateGridPlatform(35, 25, 1, 5),

			// Regular platform at (0, 27) size 14x3
			CreateGridPlatform(0, 27, 14, 3),

			// Regular platform at (15, 27) size 10x3
			CreateGridPlatform(15, 27, 10, 3),

			// Regular platform at (26, 27) size 14x3
			CreateGridPlatform(26, 27, 14, 3),

			// Goal platform at (19, 26) size 2x1
			CreateGridGoalPlatform(19, 26, 2, 1),

		},
		Spikes: []Spike{

			// Spike at (14, 29)
			CreateGridSpike(14, 29),

			// Spike at (25, 29)
			CreateGridSpike(25, 29),

		},
	}
}

// GetStage1StartPositions returns the starting positions for stage 1
func GetStage1StartPositions() (blueX, blueY, redX, redY float64) {
	// Convert grid coordinates to pixel coordinates
	return 0, 520, 780, 520
}
