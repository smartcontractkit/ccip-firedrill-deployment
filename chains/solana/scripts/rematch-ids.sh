#!/bin/bash

for dir in contracts/programs/*; do
  name=$(basename "$dir")
  toml_path="contracts/Anchor.toml"
  keypair_path="contracts/target/deploy/${name}-keypair.json"
  program_path="${dir}/src/lib.rs"

  echo $dir
  echo $name

  # Generate keypair
  # solana-keygen new -o "$keypair_path" --no-passphrase --silent

  # Extract address
  new_id=$(solana address -k "$keypair_path")

  # Update declare_id!
  sed -i '' -E "s@declare_id!\(\".*\"\)@declare_id!(\"$new_id\")@" "$program_path"

  # Update Anchor.toml
  sed -i '' -E "s@$name = \".*\"@$name = \"$new_id\"@" "$toml_path"

  echo "Updated $name to $new_id"
done
