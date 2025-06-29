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

			// Regular platform at (5, 6) size 7x1
			CreateGridPlatform(5, 6, 7, 1),

			// Regular platform at (22, 6) size 9x1
			CreateGridPlatform(22, 6, 9, 1),

			// Regular platform at (4, 7) size 1x1
			CreateGridPlatform(4, 7, 1, 1),

			// Regular platform at (12, 7) size 3x1
			CreateGridPlatform(12, 7, 3, 1),

			// Regular platform at (15, 8) size 3x1
			CreateGridPlatform(15, 8, 3, 1),

			// Regular platform at (38, 8) size 2x1
			CreateGridPlatform(38, 8, 2, 1),

			// Regular platform at (18, 9) size 3x1
			CreateGridPlatform(18, 9, 3, 1),

			// Regular platform at (37, 9) size 1x1
			CreateGridPlatform(37, 9, 1, 1),

			// Regular platform at (36, 10) size 1x1
			CreateGridPlatform(36, 10, 1, 1),

			// Regular platform at (21, 11) size 6x2
			CreateGridPlatform(21, 11, 6, 2),

			// Regular platform at (29, 11) size 1x2
			CreateGridPlatform(29, 11, 1, 2),

			// Regular platform at (34, 11) size 2x2
			CreateGridPlatform(34, 11, 2, 2),

			// Regular platform at (30, 12) size 6x1
			CreateGridPlatform(30, 12, 6, 1),

			// Regular platform at (1, 13) size 1x1
			CreateGridPlatform(1, 13, 1, 1),

			// Regular platform at (2, 14) size 1x1
			CreateGridPlatform(2, 14, 1, 1),

			// Regular platform at (3, 15) size 1x1
			CreateGridPlatform(3, 15, 1, 1),

			// Regular platform at (4, 16) size 1x1
			CreateGridPlatform(4, 16, 1, 1),

			// Regular platform at (5, 17) size 8x2
			CreateGridPlatform(5, 17, 8, 2),

			// Regular platform at (15, 17) size 1x2
			CreateGridPlatform(15, 17, 1, 2),

			// Regular platform at (21, 17) size 1x2
			CreateGridPlatform(21, 17, 1, 2),

			// Regular platform at (24, 17) size 1x2
			CreateGridPlatform(24, 17, 1, 2),

			// Regular platform at (27, 17) size 1x2
			CreateGridPlatform(27, 17, 1, 2),

			// Regular platform at (30, 17) size 3x2
			CreateGridPlatform(30, 17, 3, 2),

			// Regular platform at (16, 18) size 6x1
			CreateGridPlatform(16, 18, 6, 1),

			// Regular platform at (25, 18) size 3x1
			CreateGridPlatform(25, 18, 3, 1),

			// Regular platform at (33, 18) size 1x1
			CreateGridPlatform(33, 18, 1, 1),

			// Regular platform at (37, 18) size 3x1
			CreateGridPlatform(37, 18, 3, 1),

			// Regular platform at (38, 21) size 2x4
			CreateGridPlatform(38, 21, 2, 4),

			// Regular platform at (37, 22) size 3x3
			CreateGridPlatform(37, 22, 3, 3),

			// Regular platform at (5, 23) size 2x2
			CreateGridPlatform(5, 23, 2, 2),

			// Regular platform at (9, 23) size 1x2
			CreateGridPlatform(9, 23, 1, 2),

			// Regular platform at (13, 23) size 7x2
			CreateGridPlatform(13, 23, 7, 2),

			// Regular platform at (22, 23) size 4x2
			CreateGridPlatform(22, 23, 4, 2),

			// Regular platform at (28, 23) size 1x2
			CreateGridPlatform(28, 23, 1, 2),

			// Regular platform at (36, 23) size 4x2
			CreateGridPlatform(36, 23, 4, 2),

			// Regular platform at (10, 24) size 10x1
			CreateGridPlatform(10, 24, 10, 1),

			// Regular platform at (29, 24) size 11x1
			CreateGridPlatform(29, 24, 11, 1),

			// Regular platform at (1, 27) size 1x4
			CreateGridPlatform(1, 27, 1, 4),

			// Regular platform at (2, 28) size 1x3
			CreateGridPlatform(2, 28, 1, 3),

			// Regular platform at (3, 29) size 3x2
			CreateGridPlatform(3, 29, 3, 2),

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

			// Speed-up platform at (22, 5) size 9x1
			CreateGridSpeedUpPlatform(22, 5, 9, 1),

			// Speed-up platform at (30, 11) size 4x1
			CreateGridSpeedUpPlatform(30, 11, 4, 1),

			// Speed-up platform at (29, 23) size 7x1
			CreateGridSpeedUpPlatform(29, 23, 7, 1),

			// Speed-up platform at (9, 29) size 4x1
			CreateGridSpeedUpPlatform(9, 29, 4, 1),

			// Speed-up platform at (17, 29) size 4x1
			CreateGridSpeedUpPlatform(17, 29, 4, 1),

			// Speed-up platform at (25, 29) size 4x1
			CreateGridSpeedUpPlatform(25, 29, 4, 1),

			// Speed-up platform at (33, 29) size 4x1
			CreateGridSpeedUpPlatform(33, 29, 4, 1),

			// Speed-down platform at (16, 17) size 5x1
			CreateGridSpeedDownPlatform(16, 17, 5, 1),

			// Speed-down platform at (25, 17) size 2x1
			CreateGridSpeedDownPlatform(25, 17, 2, 1),

			// Speed-down platform at (10, 23) size 3x1
			CreateGridSpeedDownPlatform(10, 23, 3, 1),
		},
		Spikes: []Spike{

			// Spike at (17, 12)
			CreateGridSpike(17, 12),

			// Spike at (18, 12)
			CreateGridSpike(18, 12),

			// Spike at (19, 12)
			CreateGridSpike(19, 12),

			// Spike at (20, 12)
			CreateGridSpike(20, 12),

			// Spike at (27, 12)
			CreateGridSpike(27, 12),

			// Spike at (28, 12)
			CreateGridSpike(28, 12),

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

			// Spike at (26, 24)
			CreateGridSpike(26, 24),

			// Spike at (27, 24)
			CreateGridSpike(27, 24),

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
	return 500, 80, 580, 80
}
