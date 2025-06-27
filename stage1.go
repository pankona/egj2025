package main

// LoadStage1 creates stage 1 - Generated from ASCII art
// Grid layout: 40x30 cells (800x600 pixels with 20px cells)
func LoadStage1() *Stage {
	return &Stage{
		Platforms: []Platform{

			// Regular platform at (0, 0) size 40x1
			CreateGridPlatform(0, 0, 40, 1),

			// Regular platform at (0, 1) size 1x29
			CreateGridPlatform(0, 1, 1, 29),

			// Regular platform at (19, 1) size 2x8
			CreateGridPlatform(19, 1, 2, 8),

			// Regular platform at (39, 1) size 1x29
			CreateGridPlatform(39, 1, 1, 29),

			// Regular platform at (1, 4) size 16x1
			CreateGridPlatform(1, 4, 16, 1),

			// Regular platform at (23, 4) size 17x1
			CreateGridPlatform(23, 4, 17, 1),

			// Regular platform at (3, 8) size 5x1
			CreateGridPlatform(3, 8, 5, 1),

			// Regular platform at (9, 8) size 22x1
			CreateGridPlatform(9, 8, 22, 1),

			// Regular platform at (32, 8) size 5x1
			CreateGridPlatform(32, 8, 5, 1),

			// Regular platform at (4, 26) size 1x4
			CreateGridPlatform(4, 26, 1, 4),

			// Regular platform at (35, 26) size 1x4
			CreateGridPlatform(35, 26, 1, 4),

			// Regular platform at (1, 28) size 13x2
			CreateGridPlatform(1, 28, 13, 2),

			// Regular platform at (15, 28) size 10x2
			CreateGridPlatform(15, 28, 10, 2),

			// Regular platform at (26, 28) size 14x2
			CreateGridPlatform(26, 28, 14, 2),

			// Goal platform at (19, 27) size 2x1
			CreateGridGoalPlatform(19, 27, 2, 1),

		},
		Spikes: []Spike{

			// Spike at (8, 8)
			CreateGridSpike(8, 8),

			// Spike at (31, 8)
			CreateGridSpike(31, 8),

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
	return 20, 20, 760, 20
}
