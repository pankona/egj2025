package main

// LoadStage0 creates stage 0 - Generated from ASCII art
// Grid layout: 40x31 cells (800x620 pixels with 20px cells)
func LoadStage0() *Stage {
	return &Stage{
		Platforms: []Platform{

			// Regular platform at (0, 0) size 40x1
			CreateGridPlatform(0, 0, 40, 1),

			// Regular platform at (0, 1) size 1x30
			CreateGridPlatform(0, 1, 1, 30),

			// Regular platform at (39, 1) size 1x30
			CreateGridPlatform(39, 1, 1, 30),

			// Regular platform at (23, 6) size 17x1
			CreateGridPlatform(23, 6, 17, 1),

			// Regular platform at (5, 11) size 4x2
			CreateGridPlatform(5, 11, 4, 2),

			// Regular platform at (30, 11) size 1x2
			CreateGridPlatform(30, 11, 1, 2),

			// Regular platform at (11, 12) size 20x1
			CreateGridPlatform(11, 12, 20, 1),

			// Regular platform at (1, 13) size 1x1
			CreateGridPlatform(1, 13, 1, 1),

			// Regular platform at (2, 14) size 1x1
			CreateGridPlatform(2, 14, 1, 1),

			// Regular platform at (3, 15) size 1x1
			CreateGridPlatform(3, 15, 1, 1),

			// Regular platform at (4, 16) size 1x1
			CreateGridPlatform(4, 16, 1, 1),

			// Regular platform at (5, 17) size 2x2
			CreateGridPlatform(5, 17, 2, 2),

			// Regular platform at (12, 17) size 1x2
			CreateGridPlatform(12, 17, 1, 2),

			// Regular platform at (15, 17) size 7x2
			CreateGridPlatform(15, 17, 7, 2),

			// Regular platform at (24, 17) size 1x2
			CreateGridPlatform(24, 17, 1, 2),

			// Regular platform at (27, 17) size 1x2
			CreateGridPlatform(27, 17, 1, 2),

			// Regular platform at (30, 17) size 1x2
			CreateGridPlatform(30, 17, 1, 2),

			// Regular platform at (34, 17) size 1x1
			CreateGridPlatform(34, 17, 1, 1),

			// Regular platform at (7, 18) size 6x1
			CreateGridPlatform(7, 18, 6, 1),

			// Regular platform at (25, 18) size 3x1
			CreateGridPlatform(25, 18, 3, 1),

			// Regular platform at (31, 18) size 3x1
			CreateGridPlatform(31, 18, 3, 1),

			// Regular platform at (37, 18) size 3x3
			CreateGridPlatform(37, 18, 3, 3),

			// Regular platform at (13, 22) size 1x3
			CreateGridPlatform(13, 22, 1, 3),

			// Regular platform at (26, 22) size 1x3
			CreateGridPlatform(26, 22, 1, 3),

			// Regular platform at (5, 23) size 2x2
			CreateGridPlatform(5, 23, 2, 2),

			// Regular platform at (9, 23) size 11x2
			CreateGridPlatform(9, 23, 11, 2),

			// Regular platform at (22, 23) size 15x2
			CreateGridPlatform(22, 23, 15, 2),

			// Regular platform at (1, 29) size 5x2
			CreateGridPlatform(1, 29, 5, 2),

			// Regular platform at (8, 29) size 1x2
			CreateGridPlatform(8, 29, 1, 2),

			// Regular platform at (13, 29) size 1x2
			CreateGridPlatform(13, 29, 1, 2),

			// Regular platform at (16, 29) size 1x2
			CreateGridPlatform(16, 29, 1, 2),

			// Regular platform at (21, 29) size 1x2
			CreateGridPlatform(21, 29, 1, 2),

			// Regular platform at (24, 29) size 1x2
			CreateGridPlatform(24, 29, 1, 2),

			// Regular platform at (29, 29) size 1x2
			CreateGridPlatform(29, 29, 1, 2),

			// Regular platform at (32, 29) size 1x2
			CreateGridPlatform(32, 29, 1, 2),

			// Regular platform at (37, 29) size 3x2
			CreateGridPlatform(37, 29, 3, 2),

			// Regular platform at (9, 30) size 5x1
			CreateGridPlatform(9, 30, 5, 1),

			// Regular platform at (17, 30) size 5x1
			CreateGridPlatform(17, 30, 5, 1),

			// Regular platform at (25, 30) size 5x1
			CreateGridPlatform(25, 30, 5, 1),

			// Regular platform at (33, 30) size 7x1
			CreateGridPlatform(33, 30, 7, 1),

			// Goal platform at (37, 27) size 2x2
			CreateGridGoalPlatform(37, 27, 2, 2),

			// Speed-up platform at (23, 5) size 16x1
			CreateGridSpeedUpPlatform(23, 5, 16, 1),

			// Speed-up platform at (11, 11) size 19x1
			CreateGridSpeedUpPlatform(11, 11, 19, 1),

			// Speed-up platform at (25, 17) size 2x1
			CreateGridSpeedUpPlatform(25, 17, 2, 1),

			// Speed-up platform at (31, 17) size 3x1
			CreateGridSpeedUpPlatform(31, 17, 3, 1),

			// Speed-up platform at (9, 29) size 4x1
			CreateGridSpeedUpPlatform(9, 29, 4, 1),

			// Speed-up platform at (17, 29) size 4x1
			CreateGridSpeedUpPlatform(17, 29, 4, 1),

			// Speed-up platform at (25, 29) size 4x1
			CreateGridSpeedUpPlatform(25, 29, 4, 1),

			// Speed-up platform at (33, 29) size 4x1
			CreateGridSpeedUpPlatform(33, 29, 4, 1),

			// Speed-down platform at (7, 17) size 5x1
			CreateGridSpeedDownPlatform(7, 17, 5, 1),
		},
		Spikes: []Spike{

			// Spike at (9, 12)
			CreateGridSpike(9, 12),

			// Spike at (10, 12)
			CreateGridSpike(10, 12),

			// Spike at (37, 17)
			CreateGridSpike(37, 17),

			// Spike at (38, 17)
			CreateGridSpike(38, 17),

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

			// Spike at (37, 24)
			CreateGridSpike(37, 24),

			// Spike at (38, 24)
			CreateGridSpike(38, 24),

			// Spike at (6, 30)
			CreateGridSpike(6, 30),

			// Spike at (7, 30)
			CreateGridSpike(7, 30),

			// Spike at (14, 30)
			CreateGridSpike(14, 30),

			// Spike at (15, 30)
			CreateGridSpike(15, 30),

			// Spike at (22, 30)
			CreateGridSpike(22, 30),

			// Spike at (23, 30)
			CreateGridSpike(23, 30),

			// Spike at (30, 30)
			CreateGridSpike(30, 30),

			// Spike at (31, 30)
			CreateGridSpike(31, 30),
		},
	}
}

// GetStage0StartPositions returns the starting positions for stage 0
func GetStage0StartPositions() (blueX, blueY, redX, redY float64) {
	// Convert grid coordinates to pixel coordinates
	return 760, 80, 620, 80
}
