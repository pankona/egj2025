package main

// LoadStage3 creates stage 3 - Generated from ASCII art
// Grid layout: 40x31 cells (800x620 pixels with 20px cells)
func LoadStage3() *Stage {
	return &Stage{
		Platforms: []Platform{

			// Regular platform at (0, 0) size 40x1
			CreateGridPlatform(0, 0, 40, 1),

			// Regular platform at (0, 1) size 1x30
			CreateGridPlatform(0, 1, 1, 30),

			// Regular platform at (39, 1) size 1x30
			CreateGridPlatform(39, 1, 1, 30),

			// Regular platform at (19, 17) size 2x8
			CreateGridPlatform(19, 17, 2, 8),

			// Regular platform at (1, 19) size 5x1
			CreateGridPlatform(1, 19, 5, 1),

			// Regular platform at (34, 19) size 6x1
			CreateGridPlatform(34, 19, 6, 1),

			// Regular platform at (14, 24) size 12x1
			CreateGridPlatform(14, 24, 12, 1),

			// Regular platform at (1, 29) size 10x2
			CreateGridPlatform(1, 29, 10, 2),

			// Regular platform at (13, 29) size 14x2
			CreateGridPlatform(13, 29, 14, 2),

			// Regular platform at (29, 29) size 11x2
			CreateGridPlatform(29, 29, 11, 2),

			// Goal platform at (19, 27) size 2x2
			CreateGridGoalPlatform(19, 27, 2, 2),

			// Speed-up platform at (3, 24) size 11x1
			CreateGridSpeedUpPlatform(3, 24, 11, 1),

			// Speed-up platform at (26, 24) size 11x1
			CreateGridSpeedUpPlatform(26, 24, 11, 1),

			// Speed-down platform at (6, 19) size 11x1
			CreateGridSpeedDownPlatform(6, 19, 11, 1),

			// Speed-down platform at (23, 19) size 11x1
			CreateGridSpeedDownPlatform(23, 19, 11, 1),
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

// GetStage3StartPositions returns the starting positions for stage 3
func GetStage3StartPositions() (blueX, blueY, redX, redY float64) {
	// Convert grid coordinates to pixel coordinates
	return 20, 360, 760, 360
}
