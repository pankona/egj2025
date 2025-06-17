package main

// LoadStage6 creates stage 6 - Asymmetric: Timing challenge
// This stage requires precise timing between the two characters
func LoadStage6() *Stage {
	return &Stage{
		Platforms: []Platform{
			// Ground platform
			CreateGroundPlatform(),

			// Left side (Blue) - long jumps required
			CreatePlatform(60, 460, 80, 20),
			CreatePlatform(220, 400, 60, 15),
			CreatePlatform(140, 340, 50, 15),
			CreatePlatform(280, 280, 70, 15),
			CreatePlatform(180, 220, 60, 15),

			// Right side (Red) - precise small platforms
			CreatePlatform(660, 470, 80, 15),
			CreatePlatform(600, 430, 40, 15),
			CreatePlatform(640, 390, 35, 15),
			CreatePlatform(580, 350, 45, 15),
			CreatePlatform(620, 310, 40, 15),
			CreatePlatform(560, 270, 50, 15),
			CreatePlatform(600, 230, 45, 15),

			// Central timing platforms
			CreatePlatform(320, 360, 80, 15),
			CreatePlatform(440, 320, 60, 15),
			CreatePlatform(380, 260, 50, 15),

			// Convergence area with timing challenge
			CreatePlatform(300, 180, 45, 15),
			CreatePlatform(360, 160, 40, 15),
			CreatePlatform(420, 180, 45, 15),

			// Dual goal platforms requiring simultaneous arrival
			CreateGoalPlatform(330, 120, 40, 15),
			CreateGoalPlatform(430, 120, 40, 15),
		},
	}
}
