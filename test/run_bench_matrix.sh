#!/usr/bin/env bash
set -euo pipefail

# Runs foreach parallel benchmarks and produces a CSV summary.
# Usage: test/run_bench_matrix.sh [OUT_CSV]

OUT_CSV=${1:-bench_summary.csv}

echo "Running go benchmarks..." >&2
go test -bench 'ForeachParallel(Matrix|ExecTrue)' -run ^$ -benchmem | tee bench_raw.out

echo "Parsing output to CSV: ${OUT_CSV}" >&2
awk -f "$(dirname "$0")/bench_parse.awk" bench_raw.out >"${OUT_CSV}"

echo "Done. Summary written to ${OUT_CSV}" >&2

