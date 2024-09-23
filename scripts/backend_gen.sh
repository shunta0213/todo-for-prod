#!/bin/bash

# 置換対象のフォルダパスを指定
folder_path="./backend"

# 置換対象の文字列と新しい文字列
old_string="github.com/GIT_USER_ID/GIT_REPO_ID/"
new_string="todo/"

# フォルダ内の全てのファイルを走査して、指定の文字列を置換
find "$folder_path" -type f | while read -r file; do
    # ファイル内の文字列を置換して一時ファイルに書き出し
    if grep -q "$old_string" "$file"; then
        sed "s|$old_string|$new_string|g" "$file" > temp_file && mv temp_file "$file"
        echo "Updated: $file"
    fi
done
