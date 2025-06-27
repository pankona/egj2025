# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## 言語設定
- Claude Codeは日本語で応対する

## プロジェクト概要
- プロジェクト名: egj2025 (Ebitengine Game Jam 2025)
- ゲーム名: UNION JUMPERS
- 説明: Ebitengineを使用した2Dプラットフォームゲーム
- ディレクトリ: /home/pankona/go/src/github.com/pankona/egj2025
- Go バージョン: 1.24.4
- メインライブラリ: github.com/hajimehoshi/ebiten/v2

## 開発コマンド
- `make lint` - コード品質チェック（`go vet ./...`）
- `make test` - テスト実行（`GOOS=windows go test -v ./...`、WSL環境対応）
- `make build-wasm` - WebAssembly向けビルド（distフォルダに出力）
- `make serve-wasm` - ローカル開発サーバー起動（ポート8080）
- `make clean` - distフォルダのクリーンアップ

## ゲームアーキテクチャ
- **単一ファイル構成**: main.goにすべてのゲームロジックを集約
- **Ebitengine Game interface**: Update(), Draw(), Layout()メソッドの実装
- **ゲーム状態管理**: StatePlaying, StateGameOver, StateCleared
- **物理演算**: 重力、衝突判定、ジャンプ機能を自作実装
- **入力処理**: キーボード（F/Jキー）とタッチ（左右画面半分）の両対応

## 主要コンポーネント
- `Game` - メインゲーム構造体
- `Unit` - プレイヤーキャラクター（Blue/Red）
- `Platform` - ゲーム内のプラットフォーム（通常/ゴール/スピード変更）
- `Stage` - ステージ情報とプラットフォーム配置
- `Spike` - トゲ障害物（接触でゲームオーバー）

## スピード足場機能
- **スピードアップ足場（u）**: 緑色、上に乗っている間は移動速度が1.3倍（30%増加）
- **スピードダウン足場（d）**: オレンジ色、上に乗っている間は移動速度が0.7倍（30%減少）
- **効果範囲**: 足場の上に完全に乗っている間のみ効果が適用される
- **衝突判定**: 通常足場と同様に横からの衝突でキャラクターが反転する

## ステージ仕様（重要）
- **グリッドサイズ**: 横40×縦31（固定仕様）
- **ピクセルサイズ**: 800×620ピクセル（各セル20×20ピクセル）
- **座標系**: 左上が(0,0)、右下が(39,30)
- **ASCII表記**: 
  - `O` = プラットフォーム（足場）
  - `.` = 空白スペース
  - `L` = 青キャラクタースタート位置
  - `R` = 赤キャラクタースタート位置
  - `G` = ゴールプラットフォーム
  - `^` = スパイク（トゲ）
  - `u` = スピードアップ足場（上に乗ると30%速くなる）
  - `d` = スピードダウン足場（上に乗ると30%遅くなる）
- **ステージファイル**: stage01.txt〜stage10.txt（40文字×31行のASCII）
- **生成ファイル**: stage1.go〜stage10.go（自動生成）

## WASMビルドとデプロイメント
- `make build-wasm` でWebAssembly向けにビルドし、distフォルダに出力
- `make serve-wasm` でローカル開発サーバーを起動（ポート8080）  
- GitHub Actionsで自動的にGitHub Pagesにデプロイ
- mainブランチにプッシュすると自動でhttps://pankona.github.io/egj2025/でゲームが公開される
- web/index.htmlにゲーム説明とコントロール説明が含まれる

## 開発ワークフロー
- 実装完了時には必ず `make lint` を実行してコード品質をチェックする
- 実装完了時には必ず `make test` を実行してテストをチェックする
- 実装完了時には必ず `make fmt` を実行してフォーマットを整える
- **コードをコミットする前に必ず `make fmt` を実行すること**
- WASMビルドのテストは `make serve-wasm` でローカル確認

## デバッグモード
- **WASM環境**: URLパラメータ `?debug=true` または `?debug=1` を追加してデバッグモードを有効化
  - 例: `http://localhost:8080/?debug=true`
- **ネイティブ環境**: 環境変数 `DEBUG=true` または `DEBUG=1` を設定
  - 例: `DEBUG=true go run .`
- デバッグモード有効時は開発用のstage0（stage00.txt）から開始する
- 通常モードではstage1から開始する

## 開発時の重要な知見

### スピード足場の衝突判定
- **問題**: スピード足場に横からあたった時にキャラクターがすり抜ける
- **原因**: `!u.OnGround`条件により、地面にいる時に横方向の衝突判定がスキップされる
- **修正方法**: `isOnTopOfPlatform`チェックを使用し、その特定の足場の上に立っている場合のみ横方向衝突をスキップ
- **実装**: main.go:712の`updatePhysics`関数内で適切な衝突判定ロジックを実装

### Makefileの自動化機能
- **ステージ生成の自動フォーマット**: `make generate-stages`実行後に自動的に`make fmt`が実行される
- **実装場所**: Makefileの`generate-stages`ターゲット末尾で`$(MAKE) fmt`を呼び出し
- **効果**: ステージファイル生成後の手動フォーマット作業が不要になる

### ステージグリッド拡張時の注意点
- **対象ファイル**: main.go（ScreenHeight）、CLAUDE.md（仕様）、cmd/stagegen/main.go（テンプレート）、全stage*.txt、全stage*.go
- **手順**: 
  1. main.goでScreenHeightを更新（620px = 31行 × 20px）
  2. 全stage*.txtファイルに底面プラットフォーム行を追加
  3. stagegen toolのテンプレートコメントを更新
  4. `make generate-stages`で全stage*.goファイルを再生成
  5. CLAUDE.mdの仕様を更新
- **重要**: ステージファイルの行数は必ず指定されたグリッドサイズと一致させる

### ステージ構成ポリシー
- **ステージ1-3**: ギミックをひとつずつ理解するための簡単なステージ（左右対称）
  - Stage1: スパイクの理解
  - Stage2: スピード足場の理解  
  - Stage3: 基本的な協調動作の理解
- **ステージ4-6**: ギミックを攻略するステージ（左右対称、チュートリアルより難しい）
  - 複数のギミックを組み合わせた応用問題
  - 左右対称を維持して操作の混乱を避ける
- **ステージ7-10**: ギミックを攻略しつつ左右非対称になる高難易度ステージ
  - LのキャラとRのキャラの位置が入れ替わり、空間認識の混乱を誘発
  - 全ギミックの複合的な攻略が必要

## メモリ管理ガイドライン
- レビュー指摘や開発作業で得られた重要な知見はCLAUDE.mdに記録する
- 追記する際は事前にユーザーに確認を求める
- プロジェクト固有のルールやベストプラクティスを蓄積する
