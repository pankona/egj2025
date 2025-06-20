package main

// LoadStage1 creates stage 1 - Tutorial: Simple symmetric layout
// This stage teaches basic movement and jumping
// Grid layout: 40x30 cells (800x600 pixels with 20px cells)
func LoadStage1() *Stage {
	return &Stage{
		Platforms: []Platform{
			// Ground platform (full width, 3 cells high at bottom)
			CreateGridGroundPlatform(),

			// Simple symmetric platforms for tutorial
			// Left side platform (7 cells wide, 1 cell high, at position 7,22)
			// Original: (150, 450, 100, 20) -> Grid: (7, 22, 5, 1)
			CreateGridPlatform(7, 22, 5, 1),
			// Right side platform (mirror at position 27,22)
			// Original: (550, 450, 100, 20) -> Grid: (27, 22, 5, 1)
			CreateGridPlatform(27, 22, 5, 1),

			// Higher symmetric platforms
			// Original: (200, 350, 100, 20) -> Grid: (10, 17, 5, 1)
			CreateGridPlatform(10, 17, 5, 1),
			// Original: (500, 350, 100, 20) -> Grid: (25, 17, 5, 1)
			CreateGridPlatform(25, 17, 5, 1),

			// Goal platforms - symmetric at bottom center
			// Original: (320, ScreenHeight-70, 60, 20) -> Grid: (16, 26, 3, 1)
			CreateGridGoalPlatform(16, 26, 3, 1),
			// Original: (420, ScreenHeight-70, 60, 20) -> Grid: (21, 26, 3, 1)
			CreateGridGoalPlatform(21, 26, 3, 1),
		},
	}
}
