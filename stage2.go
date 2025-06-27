package main

// LoadStage2 creates stage 2 - Generated from ASCII art
// Grid layout: 40x31 cells (800x620 pixels with 20px cells)
func LoadStage2() *Stage {
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

			// Regular platform at (1, 29) size 10x2
			CreateGridPlatform(1, 29, 10, 2),

			// Regular platform at (13, 29) size 14x2
			CreateGridPlatform(13, 29, 14, 2),

			// Regular platform at (29, 29) size 11x2
			CreateGridPlatform(29, 29, 11, 2),

			// Goal platform at (19, 27) size 2x2
			CreateGridGoalPlatform(19, 27, 2, 2),
		},
		Spikes: []Spike{

			// Spike at (11, 30)
			CreateGridSpike(11, 30),

			// Spike at (12, 30)
			CreateGridSpike(12, 30),

			// Spike at (27, 30)
			CreateGridSpike(27, 30),

			// Spike at (28, 30)
			CreateGridSpike(28, 30),
		},
	}
}

// GetStage2StartPositions returns the starting positions for stage 2
func GetStage2StartPositions() (blueX, blueY, redX, redY float64) {
	// Convert grid coordinates to pixel coordinates
	return 20, 560, 760, 560
}
