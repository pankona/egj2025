#!/bin/bash

# ステージファイルの形式チェックスクリプト
# 規定サイズ: 40文字×31行

set -e

EXPECTED_WIDTH=40
EXPECTED_HEIGHT=31
EXIT_CODE=0

echo "ステージファイルの形式チェックを開始します..."
echo "規定サイズ: ${EXPECTED_WIDTH}文字×${EXPECTED_HEIGHT}行"
echo ""

# stage00.txt から stage10.txt をチェック
for i in 00 {01..10}; do
    STAGE_FILE="stage${i}.txt"
    
    if [ ! -f "$STAGE_FILE" ]; then
        echo "❌ $STAGE_FILE: ファイルが見つかりません"
        EXIT_CODE=1
        continue
    fi
    
    # 行数チェック
    LINE_COUNT=$(wc -l < "$STAGE_FILE")
    if [ "$LINE_COUNT" -ne "$EXPECTED_HEIGHT" ]; then
        echo "❌ $STAGE_FILE: 行数が正しくありません (実際: ${LINE_COUNT}行, 期待: ${EXPECTED_HEIGHT}行)"
        EXIT_CODE=1
    fi
    
    # 各行の文字数チェック
    LINE_NUM=1
    WIDTH_ERROR=0
    while IFS= read -r line; do
        WIDTH=$(echo -n "$line" | wc -c)
        if [ "$WIDTH" -ne "$EXPECTED_WIDTH" ]; then
            if [ "$WIDTH_ERROR" -eq 0 ]; then
                echo "❌ $STAGE_FILE: 文字数が正しくない行があります"
                WIDTH_ERROR=1
                EXIT_CODE=1
            fi
            echo "   行 $LINE_NUM: ${WIDTH}文字 (期待: ${EXPECTED_WIDTH}文字)"
        fi
        LINE_NUM=$((LINE_NUM + 1))
    done < "$STAGE_FILE"
    
    if [ "$LINE_COUNT" -eq "$EXPECTED_HEIGHT" ] && [ "$WIDTH_ERROR" -eq 0 ]; then
        echo "✅ $STAGE_FILE: 形式が正しいです"
    fi
done

echo ""
if [ "$EXIT_CODE" -eq 0 ]; then
    echo "✅ すべてのステージファイルが正しい形式です"
else
    echo "❌ 一部のステージファイルに問題があります"
fi

exit $EXIT_CODE