package main

// LoadStage2 creates stage 2 - Tutorial: Symmetric with more platforms
// This stage teaches precision jumping with symmetric design
func LoadStage2() *Stage {
	return &Stage{
		Platforms: []Platform{
			// Ground platform
			CreateGroundPlatform(),

			// Bottom symmetric platforms
			CreatePlatform(100, 480, 80, 15),
			CreatePlatform(620, 480, 80, 15),

			// Mid-level symmetric platforms
			CreatePlatform(180, 400, 80, 15),
			CreatePlatform(540, 400, 80, 15),

			// Higher symmetric platforms
			CreatePlatform(120, 320, 80, 15),
			CreatePlatform(600, 320, 80, 15),

			// Top symmetric platforms
			CreatePlatform(200, 240, 80, 15),
			CreatePlatform(520, 240, 80, 15),

			// Central stepping stones
			CreatePlatform(300, 380, 60, 15),
			CreatePlatform(440, 380, 60, 15),

			// Goal platforms - symmetric at center
			CreateGoalPlatform(350, 180, 50, 15),
			CreateGoalPlatform(400, 180, 50, 15),
		},
	}
}
