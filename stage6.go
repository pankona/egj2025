package main

// LoadStage6 creates stage 6 - Generated from ASCII art
// Grid layout: 40x31 cells (800x620 pixels with 20px cells)
func LoadStage6() *Stage {
	return &Stage{
		Platforms: []Platform{

			// Regular platform at (0, 0) size 40x1
			CreateGridPlatform(0, 0, 40, 1),

			// Regular platform at (0, 1) size 1x30
			CreateGridPlatform(0, 1, 1, 30),

			// Regular platform at (39, 1) size 1x30
			CreateGridPlatform(39, 1, 1, 30),

			// Regular platform at (19, 9) size 2x16
			CreateGridPlatform(19, 9, 2, 16),

			// Regular platform at (3, 12) size 3x1
			CreateGridPlatform(3, 12, 3, 1),

			// Regular platform at (11, 12) size 18x1
			CreateGridPlatform(11, 12, 18, 1),

			// Regular platform at (34, 12) size 3x1
			CreateGridPlatform(34, 12, 3, 1),

			// Regular platform at (1, 17) size 4x2
			CreateGridPlatform(1, 17, 4, 2),

			// Regular platform at (7, 17) size 2x2
			CreateGridPlatform(7, 17, 2, 2),

			// Regular platform at (11, 17) size 2x2
			CreateGridPlatform(11, 17, 2, 2),

			// Regular platform at (15, 17) size 2x2
			CreateGridPlatform(15, 17, 2, 2),

			// Regular platform at (23, 17) size 2x2
			CreateGridPlatform(23, 17, 2, 2),

			// Regular platform at (27, 17) size 2x2
			CreateGridPlatform(27, 17, 2, 2),

			// Regular platform at (31, 17) size 2x2
			CreateGridPlatform(31, 17, 2, 2),

			// Regular platform at (35, 17) size 5x2
			CreateGridPlatform(35, 17, 5, 2),

			// Regular platform at (3, 23) size 2x2
			CreateGridPlatform(3, 23, 2, 2),

			// Regular platform at (7, 23) size 2x2
			CreateGridPlatform(7, 23, 2, 2),

			// Regular platform at (11, 23) size 2x2
			CreateGridPlatform(11, 23, 2, 2),

			// Regular platform at (15, 23) size 10x2
			CreateGridPlatform(15, 23, 10, 2),

			// Regular platform at (27, 23) size 2x2
			CreateGridPlatform(27, 23, 2, 2),

			// Regular platform at (31, 23) size 2x2
			CreateGridPlatform(31, 23, 2, 2),

			// Regular platform at (35, 23) size 2x2
			CreateGridPlatform(35, 23, 2, 2),

			// Regular platform at (1, 29) size 10x2
			CreateGridPlatform(1, 29, 10, 2),

			// Regular platform at (13, 29) size 14x2
			CreateGridPlatform(13, 29, 14, 2),

			// Regular platform at (29, 29) size 11x2
			CreateGridPlatform(29, 29, 11, 2),

			// Goal platform at (19, 27) size 2x2
			CreateGridGoalPlatform(19, 27, 2, 2),

			// Speed-up platform at (29, 12) size 5x1
			CreateGridSpeedUpPlatform(29, 12, 5, 1),

			// Speed-down platform at (6, 12) size 5x1
			CreateGridSpeedDownPlatform(6, 12, 5, 1),
		},
		Spikes: []Spike{

			// Spike at (5, 18)
			CreateGridSpike(5, 18),

			// Spike at (6, 18)
			CreateGridSpike(6, 18),

			// Spike at (9, 18)
			CreateGridSpike(9, 18),

			// Spike at (10, 18)
			CreateGridSpike(10, 18),

			// Spike at (13, 18)
			CreateGridSpike(13, 18),

			// Spike at (14, 18)
			CreateGridSpike(14, 18),

			// Spike at (25, 18)
			CreateGridSpike(25, 18),

			// Spike at (26, 18)
			CreateGridSpike(26, 18),

			// Spike at (29, 18)
			CreateGridSpike(29, 18),

			// Spike at (30, 18)
			CreateGridSpike(30, 18),

			// Spike at (33, 18)
			CreateGridSpike(33, 18),

			// Spike at (34, 18)
			CreateGridSpike(34, 18),

			// Spike at (5, 24)
			CreateGridSpike(5, 24),

			// Spike at (6, 24)
			CreateGridSpike(6, 24),

			// Spike at (9, 24)
			CreateGridSpike(9, 24),

			// Spike at (10, 24)
			CreateGridSpike(10, 24),

			// Spike at (13, 24)
			CreateGridSpike(13, 24),

			// Spike at (14, 24)
			CreateGridSpike(14, 24),

			// Spike at (25, 24)
			CreateGridSpike(25, 24),

			// Spike at (26, 24)
			CreateGridSpike(26, 24),

			// Spike at (29, 24)
			CreateGridSpike(29, 24),

			// Spike at (30, 24)
			CreateGridSpike(30, 24),

			// Spike at (33, 24)
			CreateGridSpike(33, 24),

			// Spike at (34, 24)
			CreateGridSpike(34, 24),

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

// GetStage6StartPositions returns the starting positions for stage 6
func GetStage6StartPositions() (blueX, blueY, redX, redY float64) {
	// Convert grid coordinates to pixel coordinates
	return 360, 220, 420, 220
}
