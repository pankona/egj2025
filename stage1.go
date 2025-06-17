package main

// LoadStage1 creates stage 1 - Tutorial: Simple symmetric layout
// This stage teaches basic movement and jumping
func LoadStage1() *Stage {
	return &Stage{
		Platforms: []Platform{
			// Ground platform
			CreateGroundPlatform(),

			// Simple symmetric platforms for tutorial
			// Left side platform
			CreatePlatform(150, 450, 100, 20),
			// Right side platform (mirror)
			CreatePlatform(550, 450, 100, 20),

			// Higher symmetric platforms
			CreatePlatform(200, 350, 100, 20),
			CreatePlatform(500, 350, 100, 20),

			// Goal platforms - symmetric at bottom center
			CreateGoalPlatform(320, float64(ScreenHeight-70), 60, 20),
			CreateGoalPlatform(420, float64(ScreenHeight-70), 60, 20),
		},
	}
}
