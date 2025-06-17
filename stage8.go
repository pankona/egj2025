package main

// LoadStage8 creates stage 8 - Character Swap: Introduction to role reversal
// This stage introduces the concept where characters might need to swap roles or positions
// Note: Character swapping logic needs to be implemented in the game engine
func LoadStage8() *Stage {
	return &Stage{
		Platforms: []Platform{
			// Ground platform
			CreateGroundPlatform(),

			// Starting positions - Blue starts on right, Red on left (swapped)
			CreatePlatform(650, 470, 100, 20), // Blue starts here (usually Red's side)
			CreatePlatform(50, 470, 100, 20),  // Red starts here (usually Blue's side)

			// Crossing paths that encourage role reversal
			CreatePlatform(580, 420, 60, 15),
			CreatePlatform(120, 420, 60, 15),

			CreatePlatform(500, 370, 70, 15),
			CreatePlatform(200, 370, 70, 15),

			CreatePlatform(420, 320, 80, 15),
			CreatePlatform(280, 320, 80, 15),

			// Central intersection - characters must cross paths
			CreatePlatform(340, 270, 40, 15),
			CreatePlatform(420, 270, 40, 15),

			// Continuation after crossover
			CreatePlatform(200, 220, 70, 15),
			CreatePlatform(500, 220, 70, 15),

			CreatePlatform(120, 170, 60, 15),
			CreatePlatform(580, 170, 60, 15),

			// Final convergence area
			CreatePlatform(300, 130, 80, 15),
			CreatePlatform(420, 130, 80, 15),

			// Goal platforms - positioned to require the swap
			CreateGoalPlatform(150, 90, 50, 15), // Blue needs to reach this (left side)
			CreateGoalPlatform(600, 90, 50, 15), // Red needs to reach this (right side)
		},
	}
}
