# Description: A simple script to run a Go server and watch for file changes

# PID of the server
pid=$(lsof -ti :8080)

# Run the server
go run main.go &
echo -e "🚀 Server is running..."
echo -e "🚀 Click on http://localhost:8080"

# Get a list of all .go, .html, and .css files
files=$(find . -type f \( -name "*.go" -o -name "*.html" -o -name "*.css" \))

# Create an associative array to store the modification times of each file
declare -A mod_times

for file in $files; do
    mod_times["$file"]=$(stat -c "%y" "$file")
done

while true; do
    sleep 1

    for file in $files; do
        old_mod_time=${mod_times["$file"]}
        new_mod_time=$(stat -c "%y" "$file")

        if [ "$old_mod_time" != "$new_mod_time" ]; then
            mod_times["$file"]="$new_mod_time"
            kill -9 $pid 2> /dev/null
            go run main.go &
            echo -e "📄 File changed: $file"
            break
        fi

    done
done

# Kill the server when the script exits
trap "kill -9 $pid 2> /dev/null" EXIT
