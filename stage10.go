package main

// LoadStage10 creates stage 10 - Generated from ASCII art
// Grid layout: 40x30 cells (800x600 pixels with 20px cells)
func LoadStage10() *Stage {
	return &Stage{
		Platforms: []Platform{
			// Ground platform (full width, 3 cells high at bottom)
			CreateGridGroundPlatform(),

			// Regular platform at (9, 17) size 5x1
			CreateGridPlatform(9, 17, 5, 1),

			// Regular platform at (24, 17) size 5x1
			CreateGridPlatform(24, 17, 5, 1),

			// Regular platform at (7, 22) size 5x1
			CreateGridPlatform(7, 22, 5, 1),

			// Regular platform at (24, 22) size 5x1
			CreateGridPlatform(24, 22, 5, 1),

			// Regular platform at (0, 28) size 40x2
			CreateGridPlatform(0, 28, 40, 2),

			// Goal platform at (4, 26) size 3x1
			CreateGridGoalPlatform(4, 26, 3, 1),

			// Goal platform at (23, 26) size 3x1
			CreateGridGoalPlatform(23, 26, 3, 1),
		},
	}
}

// GetStage10StartPositions returns the starting positions for stage 10
func GetStage10StartPositions() (blueX, blueY, redX, redY float64) {
	// Convert grid coordinates to pixel coordinates
	return 0, 540, 720, 540
}
