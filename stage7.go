package main

// LoadStage7 creates stage 7 - Asymmetric: Master challenge
// This stage is the most complex asymmetric stage before character swapping begins
func LoadStage7() *Stage {
	return &Stage{
		Platforms: []Platform{
			// Ground platform
			CreateGroundPlatform(),
			
			// Left side (Blue) - maze-like path
			CreatePlatform(50, 480, 70, 15),
			CreatePlatform(160, 450, 40, 15),
			CreatePlatform(80, 410, 60, 15),
			CreatePlatform(180, 380, 50, 15),
			CreatePlatform(120, 340, 45, 15),
			CreatePlatform(200, 300, 55, 15),
			CreatePlatform(140, 260, 50, 15),
			CreatePlatform(220, 220, 45, 15),
			CreatePlatform(160, 180, 60, 15),
			
			// Right side (Red) - vertical challenge path
			CreatePlatform(650, 480, 100, 15),
			CreatePlatform(700, 440, 50, 15),
			CreatePlatform(620, 400, 60, 15),
			CreatePlatform(680, 360, 45, 15),
			CreatePlatform(600, 320, 70, 15),
			CreatePlatform(660, 280, 50, 15),
			CreatePlatform(580, 240, 60, 15),
			CreatePlatform(640, 200, 55, 15),
			CreatePlatform(570, 160, 80, 15),
			
			// Central obstacle course
			CreatePlatform(280, 400, 30, 15),
			CreatePlatform(340, 380, 25, 15),
			CreatePlatform(300, 340, 35, 15),
			CreatePlatform(360, 320, 30, 15),
			CreatePlatform(320, 280, 40, 15),
			CreatePlatform(380, 260, 35, 15),
			CreatePlatform(340, 220, 30, 15),
			CreatePlatform(400, 200, 40, 15),
			
			// Final platforms before goal
			CreatePlatform(280, 140, 50, 15),
			CreatePlatform(370, 120, 60, 15),
			CreatePlatform(470, 140, 50, 15),
			
			// Single central goal requiring perfect coordination
			CreateGoalPlatform(385, 80, 50, 20),
		},
	}
}