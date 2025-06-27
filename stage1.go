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

			// Regular platform at (19, 1) size 2x24
			CreateGridPlatform(19, 1, 2, 24),

			// Regular platform at (39, 1) size 1x30
			CreateGridPlatform(39, 1, 1, 30),

			// Regular platform at (1, 5) size 16x2
			CreateGridPlatform(1, 5, 16, 2),

			// Regular platform at (23, 5) size 17x2
			CreateGridPlatform(23, 5, 17, 2),

			// Regular platform at (3, 11) size 5x2
			CreateGridPlatform(3, 11, 5, 2),

			// Regular platform at (10, 11) size 20x2
			CreateGridPlatform(10, 11, 20, 2),

			// Regular platform at (32, 11) size 5x2
			CreateGridPlatform(32, 11, 5, 2),

			// Regular platform at (1, 17) size 3x2
			CreateGridPlatform(1, 17, 3, 2),

			// Regular platform at (14, 17) size 3x2
			CreateGridPlatform(14, 17, 3, 2),

			// Regular platform at (23, 17) size 3x2
			CreateGridPlatform(23, 17, 3, 2),

			// Regular platform at (36, 17) size 4x2
			CreateGridPlatform(36, 17, 4, 2),

			// Regular platform at (4, 18) size 13x1
			CreateGridPlatform(4, 18, 13, 1),

			// Regular platform at (26, 18) size 14x1
			CreateGridPlatform(26, 18, 14, 1),

			// Regular platform at (3, 23) size 3x2
			CreateGridPlatform(3, 23, 3, 2),

			// Regular platform at (16, 23) size 8x2
			CreateGridPlatform(16, 23, 8, 2),

			// Regular platform at (34, 23) size 3x2
			CreateGridPlatform(34, 23, 3, 2),

			// Regular platform at (6, 24) size 31x1
			CreateGridPlatform(6, 24, 31, 1),

			// Regular platform at (1, 29) size 12x2
			CreateGridPlatform(1, 29, 12, 2),

			// Regular platform at (15, 29) size 10x2
			CreateGridPlatform(15, 29, 10, 2),

			// Regular platform at (27, 29) size 13x2
			CreateGridPlatform(27, 29, 13, 2),

			// Goal platform at (19, 28) size 2x1
			CreateGridGoalPlatform(19, 28, 2, 1),

			// Speed-up platform at (6, 23) size 10x1
			CreateGridSpeedUpPlatform(6, 23, 10, 1),

			// Speed-up platform at (24, 23) size 10x1
			CreateGridSpeedUpPlatform(24, 23, 10, 1),

			// Speed-up platform at (4, 28) size 3x1
			CreateGridSpeedUpPlatform(4, 28, 3, 1),

			// Speed-down platform at (4, 17) size 10x1
			CreateGridSpeedDownPlatform(4, 17, 10, 1),

			// Speed-down platform at (26, 17) size 10x1
			CreateGridSpeedDownPlatform(26, 17, 10, 1),

			// Speed-down platform at (33, 28) size 3x1
			CreateGridSpeedDownPlatform(33, 28, 3, 1),
		},
		Spikes: []Spike{

			// Spike at (8, 12)
			CreateGridSpike(8, 12),

			// Spike at (9, 12)
			CreateGridSpike(9, 12),

			// Spike at (30, 12)
			CreateGridSpike(30, 12),

			// Spike at (31, 12)
			CreateGridSpike(31, 12),

			// Spike at (13, 30)
			CreateGridSpike(13, 30),

			// Spike at (14, 30)
			CreateGridSpike(14, 30),

			// Spike at (25, 30)
			CreateGridSpike(25, 30),

			// Spike at (26, 30)
			CreateGridSpike(26, 30),
		},
	}
}

// GetStage1StartPositions returns the starting positions for stage 1
func GetStage1StartPositions() (blueX, blueY, redX, redY float64) {
	// Convert grid coordinates to pixel coordinates
	return 20, 20, 760, 20
}
