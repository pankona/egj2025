package main

import (
	_ "embed"
)

// Embed the jump sound effect
//
//go:embed assets/jump.mp3
var jumpSoundBytes []byte

// Embed the background music
//
//go:embed assets/bgm.mp3
var bgmSoundBytes []byte

//go:embed assets/bakuhatsu.mp3
var deadSoundBytes []byte

//go:embed assets/clear.mp3
var clearSoundBytes []byte

//go:embed assets/shot.mp3
var shotSoundBytes []byte
