#!/bin/bash

for dir in contracts/programs/*; do
  name=$(basename "$dir")
  keypair_path="contracts/target/deploy/${name}-keypair.json"
  #program_path="${dir}/src/lib.rs"

  echo $dir
  echo $name

  # Generate keypair
  # solana-keygen new -o "$keypair_path" --no-passphrase --silent

  program_id=$(solana-keygen pubkey $keypair_path)

  # Extract address
  # new_id=$(solana address -k "$keypair_path")

  # Update declare_id!
  # sed -i '' -E "s@declare_id!\(\".*\"\)@declare_id!(\"$new_id\")@" "$program_path"

  solana program close $program_id --bypass-warning
done
