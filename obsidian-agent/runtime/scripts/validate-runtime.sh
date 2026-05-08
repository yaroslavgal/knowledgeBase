#!/bin/bash
set -euo pipefail

ROOT="$(cd "$(dirname "$0")/../.." && pwd)"
RUNTIME="$ROOT/runtime"
VAULT="${OBSIDIAN_VAULT_PATH:-${1:-}}"

fail() { echo "FAIL: $*" >&2; exit 1; }
ok() { echo "ok: $*"; }
require_file() { [ -f "$1" ] || fail "missing file: $1"; }
require_dir() { [ -d "$1" ] || fail "missing directory: $1"; }
require_cmd() { command -v "$1" >/dev/null 2>&1 || fail "$1 is required"; }

require_cmd python3
require_cmd qmd

require_file "$RUNTIME/AGENTS.md"
require_file "$RUNTIME/CLAUDE.md"
cmp -s "$RUNTIME/AGENTS.md" "$RUNTIME/CLAUDE.md" || fail "AGENTS.md and CLAUDE.md must be identical"
require_file "$RUNTIME/.mcp.json"
require_file "$RUNTIME/.opencode/opencode.json"
require_file "$RUNTIME/.claude/settings.json"
require_file "$RUNTIME/core/examples/README.md"
require_dir "$RUNTIME/core/skills"
require_dir "$RUNTIME/.opencode/agents"
require_dir "$RUNTIME/.claude/agents"
ok "required runtime files exist"

for name in onboard-vault create-note improve-note organize-vault find-links visual-diagram setup-git check-vault save-lesson update-rules; do
  skill="$RUNTIME/core/skills/$name/SKILL.md"
  require_file "$skill"
  grep -q "^name: $name$" "$skill" || fail "skill name mismatch: $skill"
  claude_skill="$RUNTIME/.claude/skills/$name/SKILL.md"
  require_file "$claude_skill"
  grep -q "^name: $name$" "$claude_skill" || fail "Claude skill name mismatch: $claude_skill"
  cmp -s "$skill" "$claude_skill" || fail "Claude skill must match shared skill: $name"
done
for removed in fix-links repo-note note-create note-improve note-refactor repo-analyst next; do
  [ ! -e "$RUNTIME/core/skills/$removed" ] || fail "removed obsidian-agent skill exists: $removed"
done
ok "obsidian-agent skills validate"

for name in connection-scout structure-scout reviewer vault-scanner; do
  require_file "$RUNTIME/.opencode/agents/$name.md"
  require_file "$RUNTIME/.claude/agents/$name.md"
  grep -q "bash: deny" "$RUNTIME/.opencode/agents/$name.md" || fail "OpenCode read-only agent must deny bash: $name"
  ! grep -q '"\\*": allow' "$RUNTIME/.opencode/agents/$name.md" || fail "OpenCode read-only agent must not allow all skills: $name"
  grep -q "mcp__qmd__query" "$RUNTIME/.claude/agents/$name.md" || fail "Claude read-only agent lacks QMD tool: $name"
done
[ ! -e "$RUNTIME/.opencode/agents/repo-analyst.md" ] || fail "repo-analyst must not be in obsidian-agent"
[ ! -e "$RUNTIME/.claude/agents/repo-analyst.md" ] || fail "repo-analyst must not be in obsidian-agent"
ok "obsidian-agent agents validate"

for command in onboard check learn reflect; do
  require_file "$RUNTIME/.opencode/commands/$command.md"
done
ok "OpenCode commands validate"

python3 - <<PY
import json, pathlib
runtime = pathlib.Path("$RUNTIME")
for path in [runtime / ".opencode/opencode.json", runtime / ".claude/settings.json", runtime / ".mcp.json"]:
    json.loads(path.read_text())
config = json.loads((runtime / ".opencode/opencode.json").read_text())
qmd = config.get("mcp", {}).get("qmd")
if not qmd or qmd.get("command") != ["qmd", "mcp"]:
    raise SystemExit("OpenCode qmd MCP must use ['qmd', 'mcp']")
mcp = json.loads((runtime / ".mcp.json").read_text())
qmd_project = mcp.get("mcpServers", {}).get("qmd")
if not qmd_project or qmd_project.get("command") != "qmd" or qmd_project.get("args") != ["mcp"]:
    raise SystemExit("Claude qmd MCP must use command qmd args ['mcp']")
for name in ["connection-scout", "structure-scout", "reviewer", "vault-scanner"]:
    permission = config.get("agent", {}).get(name, {}).get("permission", {})
    if "skill" in permission:
        raise SystemExit(f"OpenCode read-only agent must not allow skills: {name}")
PY
ok "json and QMD config validate"

for memory_file in RULES.md STYLE.md PROFILE.md STRUCTURE.md PATTERNS.md LESSONS.md; do
  require_file "$ROOT/memory/$memory_file"
done
[ ! -e "$ROOT/memory/TASKS.md" ] || fail "obsidian-agent must not contain TASKS.md"
ok "obsidian-agent memory validates"

if find "$ROOT" -name '.DS_Store' | grep -q .; then
  find "$ROOT" -name '.DS_Store' >&2
  fail "obsidian-agent must not contain .DS_Store files"
fi
ok "no macOS metadata files"

if grep -rEn 'fix-links|repo-analyst|TASKS\.md|connection-discovery|independent-reviewer|GUIDE\.md|QUEUE\.md|INBOX\.md|playwright|ruff_cache' \
  "$RUNTIME/AGENTS.md" "$RUNTIME/CLAUDE.md" "$RUNTIME/.opencode" "$RUNTIME/.claude" "$RUNTIME/core" "$ROOT/memory" "$ROOT/README.md" >/tmp/obsidian-agent-stale.txt; then
  cat /tmp/obsidian-agent-stale.txt >&2
  fail "stale personal-runtime reference found"
fi
ok "no stale personal-runtime references"

if [ -n "$VAULT" ]; then
  "$RUNTIME/scripts/validate-vault-runtime.sh" "$VAULT"
fi

echo "Obsidian Agent runtime validation passed"
