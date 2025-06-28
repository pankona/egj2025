package main

import (
	"bytes"
	"io"
	"log"
	"time"

	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/mp3"
)

// BGM Loop settings - specify in bytes for precise control
// Note: Audio samples are typically 16-bit (2 bytes) per channel
// For stereo audio: 1 sample = 4 bytes (2 bytes × 2 channels)
// For mono audio: 1 sample = 2 bytes (2 bytes × 1 channel)
const (
	BGM_LOOP_START_BYTES = int64(160_000 * 0.65 * 4)  // Loop start position in bytes (0 = beginning of file)
	BGM_LOOP_END_BYTES   = int64(160_000 * 40.85 * 4) // Loop end position in bytes (0 = auto-calculate to 90% of file)
	BGM_BYTES_PER_SAMPLE = 4                          // Bytes per sample (4 = stereo 16-bit, 2 = mono 16-bit)
	BGM_AUTO_END_PERCENT = 90                         // Used when BGM_LOOP_END_BYTES is 0
)

type SoundManager struct {
	audioContext   *audio.Context
	jumpSoundBytes []byte          // Store decoded jump sound data for creating new players
	jumpPlayerPool []*audio.Player // Pool of audio players for concurrent playback
	deadSoundBytes []byte          // Store decoded dead sound data for creating new players
	deadPlayerPool []*audio.Player // Pool of audio players for dead sound
	maxConcurrent  int             // Maximum number of concurrent sounds
	bgmPlayer      *audio.Player   // BGM player for background music (with infinite loop)
}

// alignBytesToSampleBoundary ensures the byte position aligns to a sample boundary
// This prevents audio artifacts by ensuring we don't cut in the middle of a sample
func alignBytesToSampleBoundary(bytes int64, bytesPerSample int) int64 {
	return (bytes / int64(bytesPerSample)) * int64(bytesPerSample)
}

// convertBytesToSamples converts byte position to sample position
func convertBytesToSamples(bytes int64, bytesPerSample int) int64 {
	return bytes / int64(bytesPerSample)
}

func NewSoundManager() *SoundManager {
	audioContext := audio.NewContext(SampleRate)
	maxConcurrent := 5 // Allow up to 5 concurrent jump sounds

	// Decode jump MP3 file once and store the decoded data
	jumpSoundReader := bytes.NewReader(jumpSoundBytes)
	decodedJumpSound, err := mp3.DecodeWithSampleRate(SampleRate, jumpSoundReader)
	if err != nil {
		log.Printf("Failed to decode jump sound: %v", err)
		return &SoundManager{
			audioContext:   audioContext,
			jumpSoundBytes: nil,
			jumpPlayerPool: nil,
			deadSoundBytes: nil,
			deadPlayerPool: nil,
			maxConcurrent:  maxConcurrent,
		}
	}

	// Read all decoded jump data into bytes for reuse
	var jumpBuf bytes.Buffer
	_, err = jumpBuf.ReadFrom(decodedJumpSound)
	if err != nil {
		log.Printf("Failed to read decoded jump sound: %v", err)
		return &SoundManager{
			audioContext:   audioContext,
			jumpSoundBytes: nil,
			jumpPlayerPool: nil,
			deadSoundBytes: nil,
			deadPlayerPool: nil,
			maxConcurrent:  maxConcurrent,
		}
	}
	jumpSoundData := jumpBuf.Bytes()

	// Decode dead MP3 file once and store the decoded data
	deadSoundReader := bytes.NewReader(deadSoundBytes)
	decodedDeadSound, err := mp3.DecodeWithSampleRate(SampleRate, deadSoundReader)
	if err != nil {
		log.Printf("Failed to decode dead sound: %v", err)
		return &SoundManager{
			audioContext:   audioContext,
			jumpSoundBytes: jumpSoundData,
			jumpPlayerPool: nil,
			deadSoundBytes: nil,
			deadPlayerPool: nil,
			maxConcurrent:  maxConcurrent,
		}
	}

	// Read all decoded dead data into bytes for reuse
	var deadBuf bytes.Buffer
	_, err = deadBuf.ReadFrom(decodedDeadSound)
	if err != nil {
		log.Printf("Failed to read decoded dead sound: %v", err)
		return &SoundManager{
			audioContext:   audioContext,
			jumpSoundBytes: jumpSoundData,
			jumpPlayerPool: nil,
			deadSoundBytes: nil,
			deadPlayerPool: nil,
			maxConcurrent:  maxConcurrent,
		}
	}
	deadSoundData := deadBuf.Bytes()

	// Pre-create a pool of jump players
	jumpPlayerPool := make([]*audio.Player, 0, maxConcurrent)
	for i := 0; i < maxConcurrent; i++ {
		player, err := audioContext.NewPlayer(bytes.NewReader(jumpSoundData))
		if err != nil {
			log.Printf("Failed to create jump sound player %d: %v", i, err)
			continue
		}
		jumpPlayerPool = append(jumpPlayerPool, player)
	}

	// Pre-create a pool of dead players
	deadPlayerPool := make([]*audio.Player, 0, maxConcurrent)
	for i := 0; i < maxConcurrent; i++ {
		player, err := audioContext.NewPlayer(bytes.NewReader(deadSoundData))
		if err != nil {
			log.Printf("Failed to create dead sound player %d: %v", i, err)
			continue
		}
		deadPlayerPool = append(deadPlayerPool, player)
	}

	// Initialize BGM player
	var bgmPlayer *audio.Player

	if bgmSoundBytes != nil {
		bgmReader := bytes.NewReader(bgmSoundBytes)
		decodedBGM, err := mp3.DecodeWithSampleRate(SampleRate, bgmReader)
		if err != nil {
			log.Printf("Failed to decode BGM: %v", err)
		} else {
			// Get the total length of the BGM
			bgmLength := decodedBGM.Length()
			totalDuration := time.Duration(bgmLength) * time.Second / time.Duration(SampleRate)

			// Calculate total audio data size in bytes
			totalBytes := bgmLength * int64(BGM_BYTES_PER_SAMPLE)

			// Calculate loop start position from bytes, aligned to sample boundary
			loopStartBytesAligned := alignBytesToSampleBoundary(int64(BGM_LOOP_START_BYTES), BGM_BYTES_PER_SAMPLE)
			loopStartSamples := convertBytesToSamples(loopStartBytesAligned, BGM_BYTES_PER_SAMPLE)

			// Calculate loop end position from bytes
			var loopEndBytesAligned int64
			var loopEndSamples int64

			if BGM_LOOP_END_BYTES > 0 {
				// Use specified byte position, aligned to sample boundary
				loopEndBytesAligned = alignBytesToSampleBoundary(int64(BGM_LOOP_END_BYTES), BGM_BYTES_PER_SAMPLE)
				loopEndSamples = convertBytesToSamples(loopEndBytesAligned, BGM_BYTES_PER_SAMPLE)
			} else {
				// Auto-calculate based on percentage
				autoEndBytes := totalBytes * int64(BGM_AUTO_END_PERCENT) / 100
				loopEndBytesAligned = alignBytesToSampleBoundary(autoEndBytes, BGM_BYTES_PER_SAMPLE)
				loopEndSamples = convertBytesToSamples(loopEndBytesAligned, BGM_BYTES_PER_SAMPLE)
			}

			// Ensure loop end doesn't exceed file length
			if loopEndSamples > bgmLength {
				loopEndSamples = bgmLength
				loopEndBytesAligned = bgmLength * int64(BGM_BYTES_PER_SAMPLE)
			}

			// Calculate loop length (end - start)
			loopLengthSamples := loopEndSamples - loopStartSamples
			loopLengthBytes := loopEndBytesAligned - loopStartBytesAligned

			log.Printf("BGM total: %d samples (%.2f seconds, %d bytes)",
				bgmLength, totalDuration.Seconds(), totalBytes)
			log.Printf("Loop start: %d bytes → %d samples (%.2f seconds)",
				loopStartBytesAligned, loopStartSamples, float64(loopStartSamples)/float64(SampleRate))
			log.Printf("Loop end: %d bytes → %d samples (%.2f seconds)",
				loopEndBytesAligned, loopEndSamples, float64(loopEndSamples)/float64(SampleRate))
			log.Printf("Loop length: %d bytes → %d samples (%.2f seconds)",
				loopLengthBytes, loopLengthSamples, float64(loopLengthSamples)/float64(SampleRate))
			log.Printf("Bytes per sample: %d (configured), Sample alignment: OK", BGM_BYTES_PER_SAMPLE)

			// Create an infinite loop from the BGM
			// If loop start is not 0, we need to use NewInfiniteLoopWithIntro
			var loopedBGM io.ReadSeeker
			if loopStartSamples > 0 {
				// Use NewInfiniteLoopWithIntro for custom start position
				loopedBGM = audio.NewInfiniteLoopWithIntro(decodedBGM, loopStartSamples, loopLengthSamples)
			} else {
				// Use simple NewInfiniteLoop for start from beginning
				loopedBGM = audio.NewInfiniteLoop(decodedBGM, loopLengthSamples)
			}

			bgmPlayer, err = audioContext.NewPlayer(loopedBGM)
			if err != nil {
				log.Printf("Failed to create BGM player: %v", err)
			}
		}
	}

	return &SoundManager{
		audioContext:   audioContext,
		jumpSoundBytes: jumpSoundData,
		jumpPlayerPool: jumpPlayerPool,
		deadSoundBytes: deadSoundData,
		deadPlayerPool: deadPlayerPool,
		maxConcurrent:  maxConcurrent,
		bgmPlayer:      bgmPlayer,
	}
}

func (sm *SoundManager) PlayJumpSound() {
	if sm.jumpSoundBytes == nil {
		return
	}

	// Find an available player from the pool
	for _, player := range sm.jumpPlayerPool {
		if player != nil && !player.IsPlaying() {
			player.Rewind()
			player.Play()
			return
		}
	}

	// If all players are busy, create a new temporary player
	// This ensures we can always play a sound even if the pool is exhausted
	tempPlayer, err := sm.audioContext.NewPlayer(bytes.NewReader(sm.jumpSoundBytes))
	if err != nil {
		log.Printf("Failed to create temporary jump sound player: %v", err)
		return
	}
	tempPlayer.Play()
}

func (sm *SoundManager) PlayDeadSound() {
	if sm.deadSoundBytes == nil {
		return
	}

	// Find an available player from the pool
	for _, player := range sm.deadPlayerPool {
		if player != nil && !player.IsPlaying() {
			player.Rewind()
			player.Play()
			return
		}
	}

	// If all players are busy, create a new temporary player
	// This ensures we can always play a sound even if the pool is exhausted
	tempPlayer, err := sm.audioContext.NewPlayer(bytes.NewReader(sm.deadSoundBytes))
	if err != nil {
		log.Printf("Failed to create temporary dead sound player: %v", err)
		return
	}
	tempPlayer.Play()
}

func (sm *SoundManager) StartBGM() {
	if sm.bgmPlayer != nil && !sm.bgmPlayer.IsPlaying() {
		sm.bgmPlayer.Rewind()
		sm.bgmPlayer.Play()
	}
}

func (sm *SoundManager) StopBGM() {
	if sm.bgmPlayer != nil && sm.bgmPlayer.IsPlaying() {
		sm.bgmPlayer.Pause()
	}
}
