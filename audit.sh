#!/bin/bash

echo "Building ascii-art..."
go build -o ascii-art .

if [ $? -ne 0 ]; then
    echo "Build failed. Please check your Go code."
    exit 1
fi

echo "Build successful! Running audit tests..."
echo "=========================================="

run_test() {
    echo -e "\n---> Testing: $1"
    ./ascii-art "$2" "$3"
    echo "------------------------------------------"
}

# 1. Basic tests
run_test 'Passing "hello" standard' 'hello' 'standard'
run_test 'Passing "hello world" shadow' "hello world" 'shadow'
run_test 'Passing "nice 2 meet you" thinkertoy' "nice 2 meet you" 'thinkertoy'
run_test 'Passing "you & me" standard' "you & me" 'standard'
run_test 'Passing "123" shadow' "123" 'shadow'
run_test 'Passing "/(\")" thinkertoy' "/(\")" 'thinkertoy'
run_test 'Passing "ABCDEFGHIJKLMNOPQRSTUVWXYZ" shadow' "ABCDEFGHIJKLMNOPQRSTUVWXYZ" 'shadow'
run_test 'Passing "\"#$%&/()*+,-./" thinkertoy' "\"#$%&/()*+,-./" 'thinkertoy'
run_test "Passing \"It's Working\" thinkertoy" "It's Working" "thinkertoy"

# 2. Extra tests
run_test 'Passing "WeEeEe" shadow' 'WeEeEe' 'shadow'
run_test 'Passing "w33333" thinkertoy' 'w33333' 'thinkertoy'
run_test 'Passing "ayo!!!" standard' 'ayo!!!' 'standard'
run_test 'Passing "reb00t01 something?     S P A C E !" shadow' "reb00t01 something?     S P A C E !" 'shadow'

echo -e "\nAudit tests completed."

# Clean up the compiled binary
rm ascii-art