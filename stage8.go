package main

// LoadStage8 creates stage 8 - Generated from ASCII art
// Grid layout: 40x31 cells (800x620 pixels with 20px cells)
func LoadStage8() *Stage {
	return &Stage{
		Platforms: []Platform{

			// Regular platform at (0, 0) size 40x1
			CreateGridPlatform(0, 0, 40, 1),

			// Regular platform at (0, 1) size 1x30
			CreateGridPlatform(0, 1, 1, 30),

			// Regular platform at (19, 1) size 2x6
			CreateGridPlatform(19, 1, 2, 6),

			// Regular platform at (39, 1) size 1x30
			CreateGridPlatform(39, 1, 1, 30),

			// Regular platform at (7, 5) size 2x2
			CreateGridPlatform(7, 5, 2, 2),

			// Regular platform at (24, 5) size 5x2
			CreateGridPlatform(24, 5, 5, 2),

			// Regular platform at (9, 6) size 12x1
			CreateGridPlatform(9, 6, 12, 1),

			// Regular platform at (29, 6) size 11x1
			CreateGridPlatform(29, 6, 11, 1),

			// Regular platform at (1, 7) size 1x1
			CreateGridPlatform(1, 7, 1, 1),

			// Regular platform at (2, 8) size 1x1
			CreateGridPlatform(2, 8, 1, 1),

			// Regular platform at (3, 9) size 1x1
			CreateGridPlatform(3, 9, 1, 1),

			// Regular platform at (4, 10) size 1x1
			CreateGridPlatform(4, 10, 1, 1),

			// Regular platform at (5, 11) size 4x2
			CreateGridPlatform(5, 11, 4, 2),

			// Regular platform at (14, 11) size 2x2
			CreateGridPlatform(14, 11, 2, 2),

			// Regular platform at (24, 11) size 2x2
			CreateGridPlatform(24, 11, 2, 2),

			// Regular platform at (31, 11) size 4x2
			CreateGridPlatform(31, 11, 4, 2),

			// Regular platform at (11, 12) size 18x1
			CreateGridPlatform(11, 12, 18, 1),

			// Regular platform at (5, 17) size 8x2
			CreateGridPlatform(5, 17, 8, 2),

			// Regular platform at (15, 17) size 7x2
			CreateGridPlatform(15, 17, 7, 2),

			// Regular platform at (24, 17) size 4x2
			CreateGridPlatform(24, 17, 4, 2),

			// Regular platform at (30, 17) size 10x2
			CreateGridPlatform(30, 17, 10, 2),

			// Regular platform at (1, 23) size 4x2
			CreateGridPlatform(1, 23, 4, 2),

			// Regular platform at (11, 23) size 2x2
			CreateGridPlatform(11, 23, 2, 2),

			// Regular platform at (15, 23) size 5x2
			CreateGridPlatform(15, 23, 5, 2),

			// Regular platform at (22, 23) size 7x2
			CreateGridPlatform(22, 23, 7, 2),

			// Regular platform at (31, 23) size 2x2
			CreateGridPlatform(31, 23, 2, 2),

			// Regular platform at (5, 24) size 2x1
			CreateGridPlatform(5, 24, 2, 1),

			// Regular platform at (9, 24) size 11x1
			CreateGridPlatform(9, 24, 11, 1),

			// Regular platform at (29, 24) size 6x1
			CreateGridPlatform(29, 24, 6, 1),

			// Regular platform at (1, 29) size 4x2
			CreateGridPlatform(1, 29, 4, 2),

			// Regular platform at (7, 29) size 1x2
			CreateGridPlatform(7, 29, 1, 2),

			// Regular platform at (16, 29) size 1x2
			CreateGridPlatform(16, 29, 1, 2),

			// Regular platform at (19, 29) size 1x2
			CreateGridPlatform(19, 29, 1, 2),

			// Regular platform at (24, 29) size 1x2
			CreateGridPlatform(24, 29, 1, 2),

			// Regular platform at (27, 29) size 1x2
			CreateGridPlatform(27, 29, 1, 2),

			// Regular platform at (32, 29) size 1x2
			CreateGridPlatform(32, 29, 1, 2),

			// Regular platform at (35, 29) size 5x2
			CreateGridPlatform(35, 29, 5, 2),

			// Regular platform at (8, 30) size 9x1
			CreateGridPlatform(8, 30, 9, 1),

			// Regular platform at (20, 30) size 5x1
			CreateGridPlatform(20, 30, 5, 1),

			// Regular platform at (28, 30) size 5x1
			CreateGridPlatform(28, 30, 5, 1),

			// Goal platform at (1, 27) size 2x2
			CreateGridGoalPlatform(1, 27, 2, 2),

			// Speed-up platform at (9, 5) size 10x1
			CreateGridSpeedUpPlatform(9, 5, 10, 1),

			// Speed-up platform at (11, 11) size 3x1
			CreateGridSpeedUpPlatform(11, 11, 3, 1),

			// Speed-up platform at (16, 11) size 3x1
			CreateGridSpeedUpPlatform(16, 11, 3, 1),

			// Speed-up platform at (21, 11) size 3x1
			CreateGridSpeedUpPlatform(21, 11, 3, 1),

			// Speed-up platform at (26, 11) size 3x1
			CreateGridSpeedUpPlatform(26, 11, 3, 1),

			// Speed-up platform at (13, 23) size 2x1
			CreateGridSpeedUpPlatform(13, 23, 2, 1),

			// Speed-up platform at (29, 23) size 2x1
			CreateGridSpeedUpPlatform(29, 23, 2, 1),

			// Speed-up platform at (8, 29) size 8x1
			CreateGridSpeedUpPlatform(8, 29, 8, 1),

			// Speed-up platform at (20, 29) size 4x1
			CreateGridSpeedUpPlatform(20, 29, 4, 1),

			// Speed-up platform at (28, 29) size 4x1
			CreateGridSpeedUpPlatform(28, 29, 4, 1),

			// Speed-down platform at (29, 5) size 10x1
			CreateGridSpeedDownPlatform(29, 5, 10, 1),

			// Speed-down platform at (19, 11) size 2x1
			CreateGridSpeedDownPlatform(19, 11, 2, 1),

			// Speed-down platform at (5, 23) size 2x1
			CreateGridSpeedDownPlatform(5, 23, 2, 1),

			// Speed-down platform at (9, 23) size 2x1
			CreateGridSpeedDownPlatform(9, 23, 2, 1),

			// Speed-down platform at (33, 23) size 2x1
			CreateGridSpeedDownPlatform(33, 23, 2, 1),
		},
		Spikes: []Spike{

			// Spike at (9, 12)
			CreateGridSpike(9, 12),

			// Spike at (10, 12)
			CreateGridSpike(10, 12),

			// Spike at (29, 12)
			CreateGridSpike(29, 12),

			// Spike at (30, 12)
			CreateGridSpike(30, 12),

			// Spike at (13, 18)
			CreateGridSpike(13, 18),

			// Spike at (14, 18)
			CreateGridSpike(14, 18),

			// Spike at (22, 18)
			CreateGridSpike(22, 18),

			// Spike at (23, 18)
			CreateGridSpike(23, 18),

			// Spike at (28, 18)
			CreateGridSpike(28, 18),

			// Spike at (29, 18)
			CreateGridSpike(29, 18),

			// Spike at (7, 24)
			CreateGridSpike(7, 24),

			// Spike at (8, 24)
			CreateGridSpike(8, 24),

			// Spike at (20, 24)
			CreateGridSpike(20, 24),

			// Spike at (21, 24)
			CreateGridSpike(21, 24),

			// Spike at (5, 30)
			CreateGridSpike(5, 30),

			// Spike at (6, 30)
			CreateGridSpike(6, 30),

			// Spike at (17, 30)
			CreateGridSpike(17, 30),

			// Spike at (18, 30)
			CreateGridSpike(18, 30),

			// Spike at (25, 30)
			CreateGridSpike(25, 30),

			// Spike at (26, 30)
			CreateGridSpike(26, 30),

			// Spike at (33, 30)
			CreateGridSpike(33, 30),

			// Spike at (34, 30)
			CreateGridSpike(34, 30),
		},
	}
}

// GetStage8StartPositions returns the starting positions for stage 8
func GetStage8StartPositions() (blueX, blueY, redX, redY float64) {
	// Convert grid coordinates to pixel coordinates
	return 360, 80, 760, 80
}
