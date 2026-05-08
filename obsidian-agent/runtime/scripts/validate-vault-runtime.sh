#!/bin/bash
set -euo pipefail

VAULT="${OBSIDIAN_VAULT_PATH:-${1:-}}"

fail() { echo "FAIL: $*" >&2; exit 1; }
warn() { echo "WARN: $*" >&2; }
require_file() { [ -f "$1" ] || fail "missing file: $1"; }
require_dir() { [ -d "$1" ] || fail "missing directory: $1"; }
require_cmd() { command -v "$1" >/dev/null 2>&1 || fail "$1 is required"; }

[ -n "$VAULT" ] || fail "set OBSIDIAN_VAULT_PATH or pass the vault path"
require_cmd python3
require_cmd qmd

require_file "$VAULT/AGENTS.md"
require_file "$VAULT/CLAUDE.md"
cmp -s "$VAULT/AGENTS.md" "$VAULT/CLAUDE.md" || fail "installed AGENTS.md and CLAUDE.md must be identical"
require_file "$VAULT/.mcp.json"
require_file "$VAULT/.opencode/opencode.json"
require_file "$VAULT/.claude/settings.json"
require_dir "$VAULT/.vault-agent"
require_dir "$VAULT/.vault-agent/examples"

for memory_file in RULES.md STYLE.md PROFILE.md STRUCTURE.md PATTERNS.md LESSONS.md; do
  require_file "$VAULT/.vault-agent/$memory_file"
done
[ ! -e "$VAULT/.vault-agent/TASKS.md" ] || warn "stale .vault-agent/TASKS.md exists; obsidian-agent does not install or use it"

for skill in onboard-vault create-note improve-note organize-vault find-links visual-diagram setup-git check-vault save-lesson update-rules; do
  require_file "$VAULT/.opencode/skills/$skill/SKILL.md"
  require_file "$VAULT/.claude/skills/$skill/SKILL.md"
  cmp -s "$VAULT/.opencode/skills/$skill/SKILL.md" "$VAULT/.claude/skills/$skill/SKILL.md" || fail "OpenCode and Claude skill differ: $skill"
done
for removed in fix-links repo-note note-create note-improve note-refactor; do
  [ ! -e "$VAULT/.opencode/skills/$removed" ] || fail "removed skill installed: $removed"
  [ ! -e "$VAULT/.claude/skills/$removed" ] || fail "removed skill installed: $removed"
done

for agent in connection-scout structure-scout reviewer vault-scanner; do
  require_file "$VAULT/.opencode/agents/$agent.md"
  require_file "$VAULT/.claude/agents/$agent.md"
  grep -q "bash: deny" "$VAULT/.opencode/agents/$agent.md" || fail "OpenCode read-only agent must deny bash: $agent"
  ! grep -q '"\\*": allow' "$VAULT/.opencode/agents/$agent.md" || fail "OpenCode read-only agent must not allow all skills: $agent"
  grep -q "mcp__qmd__query" "$VAULT/.claude/agents/$agent.md" || fail "Claude read-only agent lacks QMD tool: $agent"
done
[ ! -e "$VAULT/.opencode/agents/repo-analyst.md" ] || fail "repo-analyst installed"
[ ! -e "$VAULT/.claude/agents/repo-analyst.md" ] || fail "repo-analyst installed"

for command in onboard check learn reflect; do
  require_file "$VAULT/.opencode/commands/$command.md"
done

if find "$VAULT/.vault-agent" "$VAULT/.opencode" "$VAULT/.claude" -name '.DS_Store' | grep -q .; then
  find "$VAULT/.vault-agent" "$VAULT/.opencode" "$VAULT/.claude" -name '.DS_Store' >&2
  fail "installed runtime must not contain .DS_Store files"
fi

python3 - <<PY
import json, pathlib
vault = pathlib.Path("$VAULT")
config = json.loads((vault / ".opencode/opencode.json").read_text())
json.loads((vault / ".claude/settings.json").read_text())
mcp = json.loads((vault / ".mcp.json").read_text())
qmd = config.get("mcp", {}).get("qmd")
if not qmd or qmd.get("command") != ["qmd", "mcp"]:
    raise SystemExit("OpenCode qmd MCP must use ['qmd', 'mcp']")
qmd_project = mcp.get("mcpServers", {}).get("qmd")
if not qmd_project or qmd_project.get("command") != "qmd" or qmd_project.get("args") != ["mcp"]:
    raise SystemExit("Claude qmd MCP must use command qmd args ['mcp']")
for name in ["connection-scout", "structure-scout", "reviewer", "vault-scanner"]:
    permission = config.get("agent", {}).get(name, {}).get("permission", {})
    if "skill" in permission:
        raise SystemExit(f"OpenCode read-only agent must not allow skills: {name}")
PY

if grep -rEn 'fix-links|repo-analyst|TASKS\.md|connection-discovery|independent-reviewer|GUIDE\.md|QUEUE\.md|INBOX\.md|playwright|ruff_cache' \
  "$VAULT/AGENTS.md" "$VAULT/CLAUDE.md" "$VAULT/.vault-agent" "$VAULT/.opencode" "$VAULT/.claude" >/tmp/obsidian-agent-vault-stale.txt; then
  cat /tmp/obsidian-agent-vault-stale.txt >&2
  fail "installed runtime has stale personal-runtime reference"
fi

echo "Obsidian Agent vault validation passed"
