#!/bin/bash

# Run the server
go run main.go &
echo -e "🚀 Server is running..."

# Welcome message
echo -e "🚀 Click on http://localhost:8080"

# message for file change
# echo -e "📄 File changed: $file"

# Get a list of all .go, .html, and .css files
files=$(find . -type f \( -name "*.go" -o -name "*.html" -o -name "*.css" \))

# Create an associative array to store the modification times of each file
declare -A mod_times

for file in $files; do
    mod_times["$file"]=$(stat -c "%y" "$file")
done

while true; do
    sleep 1 # Poll every second

    # Check if any .go, .html, or .css files have changed
    for file in $files; do
        new_mod_time=$(stat -c "%y" "$file")
        old_mod_time=${mod_times["$file"]}

        if [ "$old_mod_time" != "$new_mod_time" ]; then
            mod_times["$file"]="$new_mod_time"
            kill %1
            go run main.go &
            echo -e "📄 File changed: $file"

            break
        fi

    done
done
