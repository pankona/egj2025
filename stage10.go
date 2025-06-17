package main

// LoadStage10 creates stage 10 - Final Challenge: Ultimate coordination test
// This is the final stage with maximum complexity, character swapping, and asymmetric design
func LoadStage10() *Stage {
	return &Stage{
		Platforms: []Platform{
			// Ground platform
			CreateGroundPlatform(),
			
			// Starting challenge - immediate decision required
			CreatePlatform(100, 480, 50, 15),
			CreatePlatform(375, 480, 50, 15),
			CreatePlatform(650, 480, 50, 15),
			
			// Triple path beginning
			CreatePlatform(50, 440, 40, 15),   // Far left
			CreatePlatform(380, 440, 40, 15),  // Center
			CreatePlatform(700, 440, 40, 15),  // Far right
			
			// Complex weaving section
			CreatePlatform(120, 400, 30, 15),
			CreatePlatform(200, 380, 25, 15),
			CreatePlatform(280, 360, 35, 15),
			CreatePlatform(360, 340, 30, 15),
			CreatePlatform(440, 360, 35, 15),
			CreatePlatform(520, 380, 25, 15),
			CreatePlatform(600, 400, 30, 15),
			CreatePlatform(680, 420, 35, 15),
			
			// Mid-section with multiple crossing points
			CreatePlatform(160, 320, 40, 15),
			CreatePlatform(240, 300, 30, 15),
			CreatePlatform(320, 280, 35, 15),
			CreatePlatform(400, 260, 40, 15),
			CreatePlatform(480, 280, 35, 15),
			CreatePlatform(560, 300, 30, 15),
			CreatePlatform(640, 320, 40, 15),
			
			// Timing-critical section - small platforms
			CreatePlatform(100, 240, 25, 15),
			CreatePlatform(180, 220, 20, 15),
			CreatePlatform(260, 200, 30, 15),
			CreatePlatform(340, 180, 25, 15),
			CreatePlatform(420, 160, 35, 15),
			CreatePlatform(500, 180, 25, 15),
			CreatePlatform(580, 200, 30, 15),
			CreatePlatform(660, 220, 20, 15),
			CreatePlatform(720, 240, 25, 15),
			
			// Final convergence - requires perfect coordination
			CreatePlatform(200, 140, 50, 15),
			CreatePlatform(300, 120, 40, 15),
			CreatePlatform(380, 100, 40, 15),
			CreatePlatform(460, 120, 40, 15),
			CreatePlatform(550, 140, 50, 15),
			
			// Ultimate challenge - single goal requiring perfect timing
			CreateGoalPlatform(375, 60, 50, 20),
		},
	}
}