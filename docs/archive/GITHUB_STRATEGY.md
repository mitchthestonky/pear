# GitHub Repository Strategy — Open Source vs Proprietary

> February 2026. How to structure Pear's GitHub presence while keeping core IP closed source.

---

## Industry Analysis: What Competitors Open Source

### Fully Open Source
- **Zed** (75k+ stars) — entire editor open source, paid collaboration features
- **Ghostty** — entire terminal open source (Apache 2)

### Closed Core, Open Extensions
- **Warp** (25k+ stars) — closed terminal, open themes/workflows/commands
- **Raycast** (7k+ stars extensions) — closed app, open extension marketplace

### Fully Closed
- **Cursor** (32k+ stars community repo) — closed IDE, GitHub repo used for issue tracking only
- **Fig** (acquired by Amazon) — previously partially open

**Pattern:** The most commercially successful dev tools (Cursor, Warp, Raycast) keep their core closed and open-source the ecosystem layer. This is the right model for Pear.

---

## Pear's GitHub Structure

### Phase 1: Launch (Weeks 1–5)

Private monorepo only. No public repos. Focus on shipping.

```
pear/ (private)
├── cli/
├── api/
├── web/
└── docs/
```

### Phase 2: Post-Launch (Week 6+)

Create strategic public repositories:

#### 1. `pear-cli/community` (Priority 1)
Issue tracking, feature requests, discussions — without exposing code. This is the Cursor model.

```
community/
├── README.md
├── ROADMAP.md
├── .github/
│   ├── ISSUE_TEMPLATE/
│   └── DISCUSSIONS/
```

#### 2. `pear-cli/examples` (Priority 2)
Show Pear in action across languages/frameworks. SEO benefit, shareable for content marketing.

```
examples/
├── nodejs-api/
├── python-flask/
├── react-app/
├── go-cli/
└── README.md
```

#### 3. `pear-cli/context-providers` (Priority 3, v1.7+)
Community-extensible context collection. Opens ecosystem without opening core.

```
context-providers/
├── git/
├── docker/
├── terraform/
└── README.md          # How to build providers
```

#### 4. `pear-cli/docs` (Priority 4)
Public docs site. Community can submit improvements. Better SEO than docs on main domain alone.

---

## What Stays Proprietary

- Teaching algorithm and prompt engineering
- Context collection strategy and budget logic
- LLM routing and intelligence layer
- API backend, auth, billing
- CLI core logic

## What Can Open Source Later (v2.0+)

- Basic context providers (Git, Docker)
- Config schemas and types
- CLI utilities/helpers
- Testing utilities
- Official integrations

---

## Extension Protocol Design (v1.7+)

**Naming convention:** `pear-context-{name}` (e.g., `pear-context-kubernetes`)

**Installation:**
```bash
pear install context kubernetes
# Fetches from pear-cli/context-kubernetes → ~/.pear/contexts/kubernetes
```

**Discovery:**
- GitHub topic: `pear-context-provider`
- Official registry (JSON in main org)
- Auto-discover via GitHub API + topic search

---

## GitHub Topics (All Public Repos)

`pear`, `developer-tools`, `cli-tool`, `ai-coding-assistant`, `code-education`, `developer-productivity`

---

## Why This Model Works

- **Protects core IP** while building community
- **Generates SEO** through examples and docs repos
- **Enables "open ecosystem" messaging** even with closed core
- **GitHub stars across multiple repos** serve as social proof
- **Community contributions** (docs, examples, extensions) without exposing secret sauce
