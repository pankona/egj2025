package main

// LoadStage3 creates stage 3 - Tutorial: Advanced symmetric layout
// This stage teaches timing and coordination with symmetric design
func LoadStage3() *Stage {
	return &Stage{
		Platforms: []Platform{
			// Ground platform
			CreateGroundPlatform(),

			// Starting platforms (symmetric)
			CreatePlatform(50, 450, 100, 20),
			CreatePlatform(650, 450, 100, 20),

			// Lower symmetric platforms with gaps
			CreatePlatform(180, 420, 60, 15),
			CreatePlatform(560, 420, 60, 15),

			// Mid platforms (staggered but symmetric)
			CreatePlatform(120, 360, 70, 15),
			CreatePlatform(610, 360, 70, 15),

			CreatePlatform(220, 300, 70, 15),
			CreatePlatform(510, 300, 70, 15),

			// Higher platforms
			CreatePlatform(160, 240, 80, 15),
			CreatePlatform(560, 240, 80, 15),

			// Central bridge area
			CreatePlatform(270, 200, 60, 15),
			CreatePlatform(470, 200, 60, 15),

			// Final approach platforms
			CreatePlatform(340, 160, 50, 15),
			CreatePlatform(410, 160, 50, 15),

			// Goal platform - single central goal
			CreateGoalPlatform(370, 120, 60, 20),
		},
	}
}
