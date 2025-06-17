package main

// LoadStage5 creates stage 5 - Asymmetric: Complex different paths
// This stage features more complex asymmetric design with challenging coordination
func LoadStage5() *Stage {
	return &Stage{
		Platforms: []Platform{
			// Ground platform
			CreateGroundPlatform(),

			// Left side (Blue) - complex zigzag path
			CreatePlatform(80, 480, 60, 15),
			CreatePlatform(200, 440, 50, 15),
			CreatePlatform(120, 380, 70, 15),
			CreatePlatform(250, 340, 60, 15),
			CreatePlatform(160, 280, 80, 15),
			CreatePlatform(280, 220, 50, 15),

			// Right side (Red) - stepped descent path
			CreatePlatform(620, 470, 100, 15),
			CreatePlatform(580, 420, 80, 15),
			CreatePlatform(640, 370, 60, 15),
			CreatePlatform(560, 320, 90, 15),
			CreatePlatform(600, 270, 70, 15),
			CreatePlatform(540, 200, 80, 15),

			// Central challenge area
			CreatePlatform(350, 300, 50, 15),
			CreatePlatform(420, 260, 40, 15),
			CreatePlatform(380, 200, 60, 15),

			// Final convergence
			CreatePlatform(340, 160, 40, 15),
			CreatePlatform(420, 160, 40, 15),

			// Goal platform - single target requiring coordination
			CreateGoalPlatform(370, 120, 60, 20),
		},
	}
}
