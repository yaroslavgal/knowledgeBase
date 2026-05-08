#!/bin/bash
set -euo pipefail

ROOT="$(cd "$(dirname "$0")/../.." && pwd)"
VAULT="${OBSIDIAN_VAULT_PATH:-${1:-}}"

SRC_RUNTIME="$ROOT/runtime"
SRC_MEMORY="$ROOT/memory"

fail() {
  echo "FAIL: $*" >&2
  exit 1
}

require_dir() {
  [ -d "$1" ] || fail "missing directory: $1"
}

require_cmd() {
  command -v "$1" >/dev/null 2>&1 || fail "$1 is required"
}

sync_dir() {
  local src="$1"
  local dst="$2"
  rm -rf "$dst"
  mkdir -p "$dst"
  cp -a "$src/." "$dst/"
  find "$dst" -name '.DS_Store' -delete 2>/dev/null || true
}

[ -n "$VAULT" ] || fail "set OBSIDIAN_VAULT_PATH or pass the vault path as the first argument"
require_dir "$VAULT"
require_dir "$SRC_RUNTIME"
require_dir "$SRC_MEMORY"
require_cmd qmd

rm -f "$VAULT/AGENTS.md"
cp "$SRC_RUNTIME/AGENTS.md" "$VAULT/AGENTS.md"
rm -f "$VAULT/CLAUDE.md" "$VAULT/.mcp.json"
cp "$SRC_RUNTIME/CLAUDE.md" "$VAULT/CLAUDE.md"
cp "$SRC_RUNTIME/.mcp.json" "$VAULT/.mcp.json"

mkdir -p "$VAULT/.vault-agent"
for memory_file in RULES.md STYLE.md PROFILE.md STRUCTURE.md PATTERNS.md LESSONS.md; do
  if [ ! -f "$VAULT/.vault-agent/$memory_file" ]; then
    cp "$SRC_MEMORY/$memory_file" "$VAULT/.vault-agent/$memory_file"
  fi
done
mkdir -p "$VAULT/.vault-agent/examples"
if [ ! -f "$VAULT/.vault-agent/examples/README.md" ]; then
  cp "$SRC_RUNTIME/core/examples/README.md" "$VAULT/.vault-agent/examples/README.md"
fi

sync_dir "$SRC_RUNTIME/.opencode" "$VAULT/.opencode"
sync_dir "$SRC_RUNTIME/core/skills" "$VAULT/.opencode/skills"

sync_dir "$SRC_RUNTIME/.claude" "$VAULT/.claude"
for skill_dir in "$SRC_RUNTIME/core/skills"/*; do
  [ -d "$skill_dir" ] || continue
  sync_dir "$skill_dir" "$VAULT/.claude/skills/$(basename "$skill_dir")"
done

find "$VAULT/AGENTS.md" "$VAULT/CLAUDE.md" "$VAULT/.mcp.json" "$VAULT/.vault-agent" "$VAULT/.opencode" "$VAULT/.claude" -name '.DS_Store' -delete

echo "Installed Obsidian Agent into $VAULT"
