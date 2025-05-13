#!/usr/bin/env bash
set -euo pipefail

# Directory to scan (default to current dir)
dir="${1:-./contracts/target/deploy}"

# lamports per SOL constant
LAMPORTS_PER_SOL=1000000000

total_lamports=0

# Header
printf "%-32s %10s %12s\n" "File" "Size" "Rent (SOL)"
printf "%-32s %10s %12s\n" "----" "----" "---------"

# Loop over each .so
for so in "$dir"/*.so; do
  [[ -f $so ]] || continue

  # 1) raw file size in bytes
  size=$(wc -c < "$so")
  size=$((size))
  # 2) human-readable size
  size_hr=$(numfmt --to=iec --suffix=B "$size")

  # 5) ask solana-cli for rent‐exempt lamports
  lam=$(solana rent "$size" --lamports \
    | awk '/Rent-exempt minimum/ {print $3}')

  # accumulate
  total_lamports=$((total_lamports + lam))

  # convert to SOL
  rent_sol=$(awk -v L="$lam" -v P="$LAMPORTS_PER_SOL" \
    'BEGIN { printf("%.3f", L/P) }')

  printf "%-32s %10s %12s\n" \
    "$(basename "$so")" "$size_hr" "$rent_sol"
done

# grand total
total_sol=$(awk -v TL="$total_lamports" -v P="$LAMPORTS_PER_SOL" \
  'BEGIN { printf("%.3f", TL/P) }')

printf "\nTotal rent-exempt deposit: %s SOL\n" "$total_sol"