#!/bin/sh

run_timeout() {
	echo "Executing: ./cmd/timeout/timeout $@"
	./cmd/timeout/timeout "$@"
}

ok_case() {
	echo "--- OK Case (should succeed) ---"
	run_timeout --timeout 3s sleep 1
	echo "Exit code: $?"
	echo
}

fail_case() {
	echo "--- Fail Case (command exits with non-zero) ---"
	run_timeout --timeout 3s false
	echo "Exit code: $?"
	echo
}

timeout_case() {
	echo "--- Timeout Case (should be killed) ---"
	run_timeout --timeout 3s sleep 10
	echo "Exit code: $?"
	echo
}

ok_case
fail_case
timeout_case
