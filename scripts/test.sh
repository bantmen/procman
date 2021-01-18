set -e

assert_equal() {
    expected=$1
    observed=$2
    if [ "$expected" != "$observed" ]; then
        echo "Test failed. Expected:"
        echo "$expected"
        echo "Observed:"
        echo "$observed"
        exit 1 
    fi
}

cd bins

# No flags
expected="Running process:  [echo testing]
testing
Process stopped without error
testing
Process stopped without error
testing"
assert_equal "$expected" "$(./procman_linux_amd64 echo testing | head -n6)"

# mem flag
expected="Running process:  [echo testing]
With memory threshold: 50.00 percent
testing
Process stopped without error
testing
Process stopped without error
testing"
assert_equal "$expected" "$(./procman_linux_amd64 -mem 50 echo testing | head -n7)"

# logfile flag with empty error file
temp_file=$(mktemp)
expected="Running process:  [echo testing]
Logging command stdout to '$temp_file' and stderr to '$temp_file.error'
Process stopped without error"
assert_equal "$expected" "$(./procman_linux_amd64 -logfile $temp_file echo testing | head -n3)"
logfile_expected="testing
testing"
assert_equal "$logfile_expected" "$(cat $temp_file)"
logfile_err_out=$(cat $temp_file.error)
assert_equal "" "$logfile_err_out"

cd - > /dev/null
echo "Test success!"
