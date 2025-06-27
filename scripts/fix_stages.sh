#!/bin/bash

# ステージファイルを40文字×31行に自動修正するスクリプト

set -e

EXPECTED_WIDTH=40
EXPECTED_HEIGHT=31

echo "ステージファイルを${EXPECTED_WIDTH}文字×${EXPECTED_HEIGHT}行に修正します..."

for i in {01..10}; do
    STAGE_FILE="stage${i}.txt"
    
    if [ ! -f "$STAGE_FILE" ]; then
        echo "❌ $STAGE_FILE: ファイルが見つかりません"
        continue
    fi
    
    echo "修正中: $STAGE_FILE"
    
    # 一時ファイルを作成
    TEMP_FILE="${STAGE_FILE}.tmp"
    
    # 各行を40文字に調整し、31行にする
    LINE_COUNT=0
    while IFS= read -r line && [ "$LINE_COUNT" -lt "$EXPECTED_HEIGHT" ]; do
        # 40文字に調整（長い場合は切り詰め、短い場合は.で埋める）
        if [ ${#line} -gt $EXPECTED_WIDTH ]; then
            # 40文字に切り詰め
            echo "${line:0:$EXPECTED_WIDTH}" >> "$TEMP_FILE"
        elif [ ${#line} -lt $EXPECTED_WIDTH ]; then
            # 40文字になるまで.で埋める
            padding_needed=$((EXPECTED_WIDTH - ${#line}))
            padding=$(printf '.%.0s' $(seq 1 $padding_needed))
            echo "${line}${padding}" >> "$TEMP_FILE"
        else
            # 既に40文字の場合
            echo "$line" >> "$TEMP_FILE"
        fi
        LINE_COUNT=$((LINE_COUNT + 1))
    done < "$STAGE_FILE"
    
    # 31行に満たない場合は底面プラットフォーム行を追加
    while [ "$LINE_COUNT" -lt "$EXPECTED_HEIGHT" ]; do
        echo "OOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOO" >> "$TEMP_FILE"
        LINE_COUNT=$((LINE_COUNT + 1))
    done
    
    # 元ファイルを置き換え
    mv "$TEMP_FILE" "$STAGE_FILE"
    echo "✅ $STAGE_FILE: 修正完了"
done

echo ""
echo "✅ 全てのステージファイルの修正が完了しました"