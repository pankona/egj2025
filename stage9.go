package main

// LoadStage9 creates stage 9 - Character Swap: Complex role reversal
// This stage features multiple character swaps and complex asymmetric design
func LoadStage9() *Stage {
	return &Stage{
		Platforms: []Platform{
			// Ground platform
			CreateGroundPlatform(),

			// Complex starting area with immediate role confusion
			CreatePlatform(400, 480, 80, 15), // Central start
			CreatePlatform(200, 460, 60, 15), // Left option
			CreatePlatform(600, 460, 60, 15), // Right option

			// First swap area - characters must take opposite paths
			CreatePlatform(120, 420, 50, 15),
			CreatePlatform(680, 420, 50, 15),

			CreatePlatform(180, 380, 40, 15),
			CreatePlatform(620, 380, 40, 15),

			// Central maze section
			CreatePlatform(280, 360, 30, 15),
			CreatePlatform(520, 360, 30, 15),

			CreatePlatform(320, 320, 35, 15),
			CreatePlatform(480, 320, 35, 15),

			CreatePlatform(260, 280, 40, 15),
			CreatePlatform(540, 280, 40, 15),

			// Second swap area - more complex crossing
			CreatePlatform(360, 240, 25, 15),
			CreatePlatform(440, 240, 25, 15),
			CreatePlatform(400, 220, 30, 15),

			// Asymmetric final approach
			CreatePlatform(180, 200, 80, 15),
			CreatePlatform(580, 180, 60, 15),

			CreatePlatform(100, 160, 50, 15),
			CreatePlatform(650, 140, 70, 15),

			CreatePlatform(200, 120, 40, 15),
			CreatePlatform(560, 100, 45, 15),

			// Multiple goals requiring strategic positioning
			CreateGoalPlatform(50, 80, 40, 15),  // Far left
			CreateGoalPlatform(380, 60, 40, 15), // Center
			CreateGoalPlatform(710, 80, 40, 15), // Far right
		},
	}
}
