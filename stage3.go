package main

// LoadStage3 creates stage 3 - Tutorial: Advanced symmetric layout
// This stage teaches timing and coordination with symmetric design
// Grid layout: 40x30 cells (800x600 pixels with 20px cells)
func LoadStage3() *Stage {
	return &Stage{
		Platforms: []Platform{
			// Ground platform (full width, 3 cells high at bottom)
			CreateGridGroundPlatform(),

			// Starting platforms (symmetric)
			// Original: (50, 450, 100, 20) -> Grid: (2, 22, 5, 1)
			CreateGridPlatform(2, 22, 5, 1),
			// Original: (650, 450, 100, 20) -> Grid: (32, 22, 5, 1)
			CreateGridPlatform(32, 22, 5, 1),

			// Lower symmetric platforms with gaps
			// Original: (180, 420, 60, 15) -> Grid: (9, 21, 3, 1)
			CreateGridPlatform(9, 21, 3, 1),
			// Original: (560, 420, 60, 15) -> Grid: (28, 21, 3, 1)
			CreateGridPlatform(28, 21, 3, 1),

			// Mid platforms (staggered but symmetric)
			// Original: (120, 360, 70, 15) -> Grid: (6, 18, 3, 1)
			CreateGridPlatform(6, 18, 3, 1),
			// Original: (610, 360, 70, 15) -> Grid: (30, 18, 3, 1)
			CreateGridPlatform(30, 18, 3, 1),

			// Original: (220, 300, 70, 15) -> Grid: (11, 15, 3, 1)
			CreateGridPlatform(11, 15, 3, 1),
			// Original: (510, 300, 70, 15) -> Grid: (25, 15, 3, 1)
			CreateGridPlatform(25, 15, 3, 1),

			// Higher platforms
			// Original: (160, 240, 80, 15) -> Grid: (8, 12, 4, 1)
			CreateGridPlatform(8, 12, 4, 1),
			// Original: (560, 240, 80, 15) -> Grid: (28, 12, 4, 1)
			CreateGridPlatform(28, 12, 4, 1),

			// Central bridge area
			// Original: (270, 200, 60, 15) -> Grid: (13, 10, 3, 1)
			CreateGridPlatform(13, 10, 3, 1),
			// Original: (470, 200, 60, 15) -> Grid: (23, 10, 3, 1)
			CreateGridPlatform(23, 10, 3, 1),

			// Final approach platforms
			// Original: (340, 160, 50, 15) -> Grid: (17, 8, 2, 1)
			CreateGridPlatform(17, 8, 2, 1),
			// Original: (410, 160, 50, 15) -> Grid: (20, 8, 2, 1)
			CreateGridPlatform(20, 8, 2, 1),

			// Goal platform - single central goal
			// Original: (370, 120, 60, 20) -> Grid: (18, 6, 3, 1)
			CreateGridGoalPlatform(18, 6, 3, 1),
		},
	}
}
