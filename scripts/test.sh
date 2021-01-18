set -e

assert_count_occurrence() {
    out=$1
    pattern=$2
    count=$3
    found=$(echo "$out" | grep -c "$pattern")
    if [ "$count" -ne "$found" ]; then
        echo "Expected $count counts but found $found for: $pattern"
        exit 1 
    fi
}

cd bins

out=$(./procman_linux_amd64 echo testing | head -n6)

assert_count_occurrence "$out" "^Running process:  \[echo testing\]$" 1
assert_count_occurrence "$out" '^testing$' 3
assert_count_occurrence "$out" '^Process stopped without error' 2

cd - > /dev/null

echo "Test success!"
