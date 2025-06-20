package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"text/template"
)

// StageData represents the parsed stage data from ASCII art
type StageData struct {
	StageNumber     int
	Platforms       []PlatformData
	GoalPlatforms   []PlatformData
	BlueStartX      int
	BlueStartY      int
	RedStartX       int
	RedStartY       int
	BlueStartPixelX int
	BlueStartPixelY int
	RedStartPixelX  int
	RedStartPixelY  int
}

// PlatformData represents a platform in grid coordinates
type PlatformData struct {
	X      int
	Y      int
	Width  int
	Height int
}

// parseASCIIArt parses ASCII art file and returns stage data
func parseASCIIArt(filename string) (*StageData, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("ファイルを開けませんでした: %v", err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("ファイル読み込みエラー: %v", err)
	}

	if len(lines) == 0 {
		return nil, fmt.Errorf("空のファイルです")
	}

	// Extract stage number from filename
	baseName := filepath.Base(filename)
	// Extract number from stageN.txt format
	stageNum := 1
	if strings.HasPrefix(baseName, "stage") && strings.HasSuffix(baseName, ".txt") {
		numStr := strings.TrimPrefix(strings.TrimSuffix(baseName, ".txt"), "stage")
		if num, err := strconv.Atoi(numStr); err == nil {
			stageNum = num
		} else if numStr != "" {
			log.Printf("警告: ファイル名 '%s' からステージ番号を抽出できませんでした。デフォルトの1を使用します。", baseName)
		}
	}

	stageData := &StageData{StageNumber: stageNum}

	// Create a 2D grid to track processed cells
	height := len(lines)
	width := 0
	for _, line := range lines {
		if len(line) > width {
			width = len(line)
		}
	}

	processed := make([][]bool, height)
	for i := range processed {
		processed[i] = make([]bool, width)
	}

	// Parse the grid
	for y, line := range lines {
		for x, char := range line {
			if processed[y][x] {
				continue
			}

			switch char {
			case 'O':
				// Find rectangular platform starting from this position
				platform := findRectangularPlatform(lines, processed, x, y, 'O')
				if platform != nil {
					stageData.Platforms = append(stageData.Platforms, *platform)
				}
			case 'G':
				// Find rectangular goal platform starting from this position
				platform := findRectangularPlatform(lines, processed, x, y, 'G')
				if platform != nil {
					stageData.GoalPlatforms = append(stageData.GoalPlatforms, *platform)
				}
			case 'L':
				stageData.BlueStartX = x
				stageData.BlueStartY = y
				stageData.BlueStartPixelX = x * 20
				stageData.BlueStartPixelY = y * 20
				processed[y][x] = true
			case 'R':
				stageData.RedStartX = x
				stageData.RedStartY = y
				stageData.RedStartPixelX = x * 20
				stageData.RedStartPixelY = y * 20
				processed[y][x] = true
			case '.':
				processed[y][x] = true
			default:
				return nil, fmt.Errorf("不明な文字 '%c' が座標 (%d, %d) で見つかりました", char, x, y)
			}
		}
	}

	return stageData, nil
}

// findRectangularPlatform finds a rectangular platform starting from (startX, startY)
func findRectangularPlatform(lines []string, processed [][]bool, startX, startY int, targetChar rune) *PlatformData {
	if startY >= len(lines) || startX >= len(lines[startY]) {
		return nil
	}

	if rune(lines[startY][startX]) != targetChar {
		return nil
	}

	// Find the width by scanning right
	width := 0
	for x := startX; x < len(lines[startY]) && rune(lines[startY][x]) == targetChar; x++ {
		width++
	}

	// Find the height by scanning down, ensuring all rows have the same width
	height := 0
	for y := startY; y < len(lines); y++ {
		// Check if this row has the required width of target characters
		if len(lines[y]) < startX+width {
			break
		}
		hasFullWidth := true
		for x := startX; x < startX+width; x++ {
			if rune(lines[y][x]) != targetChar {
				hasFullWidth = false
				break
			}
		}
		if !hasFullWidth {
			break
		}
		height++
	}

	// Mark all cells in this rectangle as processed
	for y := startY; y < startY+height; y++ {
		for x := startX; x < startX+width; x++ {
			if y < len(processed) && x < len(processed[y]) {
				processed[y][x] = true
			}
		}
	}

	return &PlatformData{
		X:      startX,
		Y:      startY,
		Width:  width,
		Height: height,
	}
}

// generateStageFile generates a Go stage file from stage data
func generateStageFile(stageData *StageData, outputPath string) error {
	tmpl := `package main

// LoadStage{{.StageNumber}} creates stage {{.StageNumber}} - Generated from ASCII art
// Grid layout: 40x30 cells (800x600 pixels with 20px cells)
func LoadStage{{.StageNumber}}() *Stage {
	return &Stage{
		Platforms: []Platform{
			// Ground platform (full width, 3 cells high at bottom)
			CreateGridGroundPlatform(),
{{range .Platforms}}
			// Regular platform at ({{.X}}, {{.Y}}) size {{.Width}}x{{.Height}}
			CreateGridPlatform({{.X}}, {{.Y}}, {{.Width}}, {{.Height}}),
{{end}}{{range .GoalPlatforms}}
			// Goal platform at ({{.X}}, {{.Y}}) size {{.Width}}x{{.Height}}
			CreateGridGoalPlatform({{.X}}, {{.Y}}, {{.Width}}, {{.Height}}),
{{end}}
		},
	}
}

// GetStage{{.StageNumber}}StartPositions returns the starting positions for stage {{.StageNumber}}
func GetStage{{.StageNumber}}StartPositions() (blueX, blueY, redX, redY float64) {
	// Convert grid coordinates to pixel coordinates
	return {{.BlueStartPixelX}}, {{.BlueStartPixelY}}, {{.RedStartPixelX}}, {{.RedStartPixelY}}
}
`

	t, err := template.New("stage").Parse(tmpl)
	if err != nil {
		return fmt.Errorf("テンプレート解析エラー: %v", err)
	}

	file, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("出力ファイル作成エラー: %v", err)
	}
	defer file.Close()

	if err := t.Execute(file, stageData); err != nil {
		return fmt.Errorf("テンプレート実行エラー: %v", err)
	}

	return nil
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "使用方法: %s <input.txt>\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "例: %s stage1.txt\n", os.Args[0])
		os.Exit(1)
	}

	inputFile := os.Args[1]

	// Parse ASCII art
	stageData, err := parseASCIIArt(inputFile)
	if err != nil {
		log.Fatalf("ASCII art解析エラー: %v", err)
	}

	// Generate output filename
	baseName := filepath.Base(inputFile)
	outputName := strings.TrimSuffix(baseName, ".txt") + ".go"

	// Generate stage file
	if err := generateStageFile(stageData, outputName); err != nil {
		log.Fatalf("ステージファイル生成エラー: %v", err)
	}

	fmt.Printf("ステージファイルを生成しました: %s\n", outputName)
	fmt.Printf("ステージ番号: %d\n", stageData.StageNumber)
	fmt.Printf("プラットフォーム数: %d\n", len(stageData.Platforms))
	fmt.Printf("ゴールプラットフォーム数: %d\n", len(stageData.GoalPlatforms))
	fmt.Printf("青キャラ開始位置: (%d, %d)\n", stageData.BlueStartX, stageData.BlueStartY)
	fmt.Printf("赤キャラ開始位置: (%d, %d)\n", stageData.RedStartX, stageData.RedStartY)
}
