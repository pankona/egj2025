# ステージ生成ツール (Stage Generator)

ASCII artからEbitengineゲーム用のステージファイルを生成するツールです。

## 使用方法

```bash
go run cmd/stagegen/main.go <input.txt>
```

例:
```bash
go run cmd/stagegen/main.go stage1.txt
```

## ASCII art記号

- `.` = 空間（穴）
- `O` = 足場、壁
- `G` = ゴール
- `L` = 青キャラの初期位置（右向きに歩く）
- `R` = 赤キャラの初期位置（左向きに歩く）

## 入力例

```
OOOOOOOOOO
O........O
O........O
O........O
OL..GG..RO
OOOOOOOOOO
```

## 出力

- `stageN.go` ファイルが生成されます
- `LoadStageN()` 関数とキャラクター開始位置を返す `GetStageNStartPositions()` 関数が含まれます
- グリッド座標系（40x30セル、20px/セル）を使用します
- 生成されるコードには常に `CreateGridGroundPlatform()` が含まれ、ASCIIアートで定義したプラットフォームが追加されます

## 生成されるコードの例

```go
package main

// LoadStage1 creates stage 1 - Generated from ASCII art
func LoadStage1() *Stage {
    return &Stage{
        Platforms: []Platform{
            CreateGridGroundPlatform(),
            CreateGridPlatform(0, 0, 10, 1),
            // ... その他のプラットフォーム
            CreateGridGoalPlatform(4, 4, 2, 1),
        },
    }
}

// GetStage1StartPositions returns the starting positions for stage 1
func GetStage1StartPositions() (blueX, blueY, redX, redY float64) {
    return 20, 80, 160, 80
}
```