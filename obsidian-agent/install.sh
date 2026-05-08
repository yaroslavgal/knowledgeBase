#!/bin/bash
set -euo pipefail

ROOT="$(cd "$(dirname "$0")" && pwd)"
VAULT="${OBSIDIAN_VAULT_PATH:-${1:-}}"

if [ -z "$VAULT" ]; then
  echo "Usage: OBSIDIAN_VAULT_PATH=/path/to/ObsidianVault ./install.sh" >&2
  echo "   or: ./install.sh /path/to/ObsidianVault" >&2
  exit 1
fi

"$ROOT/runtime/scripts/materialize-vault-runtime.sh" "$VAULT"
"$ROOT/runtime/scripts/validate-vault-runtime.sh" "$VAULT"
