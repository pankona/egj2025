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
- **グリッドサイズ**: 横40×縦30（固定仕様）
- **ピクセルサイズ**: 800×600ピクセル（各セル20×20ピクセル）
- **座標系**: 左上が(0,0)、右下が(39,29)
- **ASCII表記**: 
  - `O` = プラットフォーム（足場）
  - `.` = 空白スペース
  - `L` = 青キャラクタースタート位置
  - `R` = 赤キャラクタースタート位置
  - `G` = ゴールプラットフォーム
  - `^` = スパイク（トゲ）
  - `u` = スピードアップ足場（上に乗ると30%速くなる）
  - `d` = スピードダウン足場（上に乗ると30%遅くなる）
- **ステージファイル**: stage01.txt〜stage10.txt（40文字×30行のASCII）
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

## メモリ管理ガイドライン
- レビュー指摘や開発作業で得られた重要な知見はCLAUDE.mdに記録する
- 追記する際は事前にユーザーに確認を求める
- プロジェクト固有のルールやベストプラクティスを蓄積する
