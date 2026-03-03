# SEO & Agentic Search Strategy

> February 2026. How to make Pear discoverable by both traditional search engines and AI agents (ChatGPT, Perplexity, Claude, Gemini).

---

## The Landscape: GEO + SEO

Traditional SEO gets you indexed. **Generative Engine Optimization (GEO)** gets you cited in AI-generated answers. You need both.

AI agents from ChatGPT, Perplexity, Claude, and Gemini are actively crawling websites. If your content isn't machine-readable, AI agents skip you and cite competitors instead.

**Key insight:** SEO makes you discoverable. GEO makes you quotable.

---

## 1. Technical SEO — Foundation

### Sitemap Structure

Use silo architecture — group related pages in clear hierarchies:

```
/
├── /docs/
│   ├── /docs/getting-started/
│   ├── /docs/installation/
│   ├── /docs/configuration/
│   ├── /docs/commands/
│   │   ├── /docs/commands/init
│   │   ├── /docs/commands/teach
│   │   ├── /docs/commands/record
│   │   └── /docs/commands/progress
│   └── /docs/troubleshooting/
├── /guides/
│   ├── /guides/byok-setup/
│   ├── /guides/workflows/
│   └── /guides/examples/
├── /pricing/
├── /blog/
│   ├── /blog/tutorials/
│   ├── /blog/announcements/
│   └── /blog/deep-dives/
├── /faq/
└── /changelog/
```

### JSON-LD Structured Data

Implement on every page. Priority schemas:

| Schema Type | Where | Why |
|---|---|---|
| `SoftwareApplication` | Homepage, docs | Tells engines "this is a CLI tool" |
| `Product` | Pricing page | Surfaces pricing in AI answers |
| `Organization` | All pages | Brand entity recognition |
| `FAQPage` | FAQ, docs | Direct answer extraction |
| `HowTo` | Tutorials | Step-by-step citation |
| `TechArticle` | Blog deep dives | Technical content classification |

Example for homepage:
```json
{
  "@context": "https://schema.org",
  "@type": "SoftwareApplication",
  "name": "Pear",
  "applicationCategory": "DeveloperApplication",
  "operatingSystem": "macOS",
  "description": "AI-powered CLI that teaches you about your code as you work",
  "offers": {
    "@type": "Offer",
    "price": "20.00",
    "priceCurrency": "USD",
    "billingIncrement": "month"
  }
}
```

### Semantic HTML

Every page must use: `<header>`, `<nav>`, `<main>`, `<article>`, `<section>`, `<footer>`. Proper heading hierarchy (one `<h1>`, structured `<h2>`/`<h3>`). Alt text on all images.

---

## 2. Agentic Search Optimization

### robots.txt — Allow AI Crawlers

```
# Traditional crawlers
User-agent: *
Allow: /
Sitemap: https://pear.dev/sitemap.xml

# Explicitly allow AI search crawlers
User-agent: GPTBot
Allow: /

User-agent: ChatGPT-User
Allow: /

User-agent: ClaudeBot
Allow: /

User-agent: Claude-Web
Allow: /

User-agent: Google-Extended
Allow: /

User-agent: PerplexityBot
Allow: /

User-agent: Anthropic-AI
Allow: /
```

**Note:** Allowing AI search crawlers does NOT affect Google/Bing rankings. You can block training-only crawlers (GPTBot) while allowing search crawlers (OAI-SearchBot) if desired.

### llms.txt — AI Documentation Index

Create `/llms.txt` at site root. This is a Markdown file that tells AI agents what your site offers. Cursor explicitly checks for it. 844k+ websites have adopted it.

```markdown
# Pear — Learn while you code

> An AI-powered CLI that teaches you about your code as you work. BYOK (bring your own key) model — works with Claude, GPT, Gemini, and more.

## Documentation
- [Getting Started](/docs/getting-started): Install and configure Pear
- [Commands Reference](/docs/commands): All CLI commands
- [Configuration](/docs/configuration): API keys, preferences, config file format
- [Teach Mode](/docs/teach-mode): How Pear's teaching system works

## Key Information
- [Pricing](/pricing): $30/mo or $300/yr
- [FAQ](/faq): Common questions
- [Changelog](/changelog): Release history

## Guides
- [BYOK Setup](/guides/byok-setup): Configure your own API keys
- [Workflows](/guides/workflows): Common usage patterns
```

### Content Format for AI Citation

AI agents prefer:
- **Concise paragraphs** (2–3 sentences)
- **Bullet points and numbered lists**
- **Answer first, then context** (inverted pyramid)
- **Specific numbers and data points** (makes content verifiable/citable)
- **Unique, quotable statements** (your positioning line is perfect for this)

---

## 3. Content Strategy for SEO + GEO

### High-Priority Pages

| Page | Target Query | Type |
|---|---|---|
| Homepage | "pear cli", "learn code with AI" | Brand + category |
| Docs: Getting Started | "pear setup", "pear install" | Navigation |
| Docs: Each Command | "pear teach", "pear record" | Feature |
| Pricing | "pear pricing", "ai coding tutor price" | Commercial |
| FAQ | Long-tail questions | Informational |
| Blog: "AI makes devs fast not good" | "AI developer skills gap" | Thought leadership |
| Blog: Pear vs X comparisons | "aider vs pear", "cursor learning" | Competitive |
| Troubleshooting | Error messages (exact match) | Problem-solution |

### Content That Ranks for Developers

1. **One page per CLI command** — complete docs with examples, copy-paste ready
2. **Troubleshooting guides** — target exact error messages
3. **"How to" tutorials** — "How to understand your codebase with Pear"
4. **Architecture deep dives** — "How Pear's context engine works"
5. **Comparison posts** — "Pear vs reading docs vs asking ChatGPT"
6. **Code examples** — syntax-highlighted, with copy buttons

### Developer-Specific UX

- Code copy buttons on all examples
- Dark mode (essential)
- Search (Algolia DocSearch — free for qualifying projects)
- Syntax highlighting (Shiki or Prism.js)
- CLI demo animations (Asciinema)

---

## 4. Implementation Tasks

### Week 1 — Foundation
- [ ] Set up Google Search Console
- [ ] Create comprehensive `sitemap.xml`
- [ ] Implement `robots.txt` with AI crawler permissions
- [ ] Create `/llms.txt`
- [ ] Add JSON-LD `SoftwareApplication` schema to homepage
- [ ] Add JSON-LD `Product` schema to pricing page
- [ ] Add JSON-LD `Organization` schema site-wide
- [ ] Run Lighthouse audit, fix critical issues
- [ ] Ensure all pages use semantic HTML

### Week 2 — Content Structure
- [ ] Create one-page-per-command documentation structure
- [ ] Add JSON-LD `FAQPage` schema to FAQ
- [ ] Add JSON-LD `HowTo` schema to tutorial pages
- [ ] Implement internal linking strategy (docs ↔ guides ↔ blog)
- [ ] Add code copy buttons to all documentation
- [ ] Set up Algolia DocSearch or equivalent

### Week 3 — Content Creation
- [ ] Write 5 tutorial blog posts (target long-tail keywords)
- [ ] Write "How Pear Works Under the Hood" deep dive
- [ ] Create troubleshooting page targeting common error messages
- [ ] Write FAQ page (minimum 15 questions)
- [ ] Draft 3 comparison posts (Pear vs alternatives)

### Week 4 — Distribution & Measurement
- [ ] Syndicate blog content to Dev.to, Hashnode
- [ ] Share on LinkedIn, X, Reddit (r/programming, r/devtools)
- [ ] Set up weekly blog publishing cadence
- [ ] Record first CLI demo video for YouTube
- [ ] Create OpenAPI spec if public API exists

### Ongoing — Measurement
- [ ] Track Google Search Console: impressions, clicks, CTR
- [ ] Monitor AI citations (manual checks on ChatGPT, Perplexity, Claude)
- [ ] Track docs page views
- [ ] Monthly content audit: what ranks, what gets cited, what to create next

---

## 5. Quick Wins vs Long-Term Plays

### Quick Wins (< 1 day each)
- `robots.txt` with AI crawler permissions
- `/llms.txt` file
- JSON-LD on homepage and pricing
- Google Search Console setup

### Medium Effort (1–3 days each)
- Full sitemap with silo architecture
- Semantic HTML audit and fix
- FAQ page with schema
- Lighthouse performance fixes

### Long-Term Compounds
- Weekly blog content (SEO compounds over 6–12 months)
- YouTube tutorials (videos compound for years)
- Community examples repo (ongoing SEO from GitHub)
- Comparison pages (capture competitive search traffic)
