package main

// LoadStage4 creates stage 4 - Asymmetric: Introduction to different paths
// This stage starts introducing asymmetric design where each character has different routes
func LoadStage4() *Stage {
	return &Stage{
		Platforms: []Platform{
			// Ground platform
			CreateGroundPlatform(),

			// Left side (Blue) - easier path
			CreatePlatform(50, 450, 100, 20),
			CreatePlatform(180, 400, 80, 20),
			CreatePlatform(100, 320, 120, 20),
			CreatePlatform(250, 260, 90, 20),

			// Right side (Red) - different path
			CreatePlatform(650, 460, 100, 20),
			CreatePlatform(580, 380, 70, 20),
			CreatePlatform(620, 300, 100, 20),
			CreatePlatform(540, 220, 80, 20),

			// Central connection areas
			CreatePlatform(320, 350, 80, 20),
			CreatePlatform(420, 280, 70, 20),

			// Shared final approach
			CreatePlatform(360, 180, 80, 20),

			// Goal platforms - close together for coordination
			CreateGoalPlatform(340, 140, 50, 15),
			CreateGoalPlatform(410, 140, 50, 15),
		},
	}
}
