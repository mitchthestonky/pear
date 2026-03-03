# Pear — Viral GTM Playbook

> Actionable growth strategies based on what actually worked for Cursor, Bolt.new, Lovable, Cluely, v0, Replit, and Kinso.ai. February 2026.

---

## The Kinso.ai Blueprint (Most Relevant to Pear)

Kinso.ai (Sydney, Australia) is the closest comp to pear's situation:
- **Founders:** Frank & Jacques Greeff, previously sold Realbase to Domain Group for **$180M**
- **Strategy:** Build-in-public YouTube show ("Founders Table"), documenting everything — strategy sessions, pitch rejections, product decisions
- **Results:** **12 million monthly impressions** in 4 months. Zero paid advertising. 15,000+ waitlist before product launch.
- **Content machine:** Long episodes → chopped into TikToks, LinkedIn threads, Instagram clips, newsletter snippets
- **Self-funded** after VCs said no — documented the rejections on camera (which became some of their best content)

### What pear can steal from Kinso:
1. Document the build publicly — film/write about product decisions, Go learning curve, AI-assisted development
2. The rejections and struggles ARE the content — "I'm junior in Go building a Go product" is a compelling narrative
3. Repurpose everything — one LinkedIn post becomes a tweet thread, a short video, a newsletter section
4. Build the waitlist through content consistency, not one-off launches

---

## Proven Viral Tactics by Company

### 1. Cursor — Product IS Marketing ($0 spend → $1B ARR)

| Tactic | How it worked | Pear adaptation |
|--------|--------------|-----------------|
| **Fork existing UX** | Built on VS Code (73% of devs already use it). Zero learning curve. | Build on terminal conventions devs already know. `pear` should feel like a native CLI tool from keystroke one. |
| **Founder Twitter presence** | Aman Sanger posted consistently for years before launch. One viral tweet in 2022 seeded everything. | Start posting daily on LinkedIn NOW. Document the build, share insights about the AI learning crisis. Build the audience before the product ships. |
| **Podcast appearances** | Lex Fridman interview gave Cursor a face and philosophy. | Target dev podcasts (Changelog, Syntax, devtools.fm). The "AI is making devs fast but not good" thesis is a compelling interview topic. |
| **Third-party validation** | OpenAI reportedly tried to acquire Cursor. Best possible social proof. | The "pear" name strategy + positioning against the AI learning crisis is contrarian enough to generate discussion. |
| **36% freemium conversion** | Product was so good 1 in 3 free users paid. | The 14-day free trial must be undeniably better than asking Claude directly. If conversion is <5%, the product needs work, not more marketing. |

### 2. Bolt.new — Single Tweet to $40M ARR

| Tactic | How it worked | Pear adaptation |
|--------|--------------|-----------------|
| **The demo was the tweet** | One video showing an app building itself in real-time. Inherently shareable. | The 60-second pear demo: ask about real code → watch context injection happen → get teaching response. Must be visually compelling enough to screenshot/share. |
| **Zero friction trial** | Browser-based, no signup for first interaction. | `curl install | sh && pear ask "why is this failing?"` — first value in under 60 seconds. |
| **Every output is marketing** | Users shared their Bolt creations, each one an implicit ad. | Every pear teaching response could be shareable. Add a `/share` command that creates a formatted snippet for Twitter/LinkedIn. |
| **Founder shared real revenue** | Eric Simons posted "$20M ARR in 2 months!" on LinkedIn. These posts went viral. | Share real metrics publicly. "50 signups on day 1." "First paying user." Transparency is content. |
| **Near-death backstory** | 7 years of StackBlitz struggle → overnight success. Compelling narrative. | Solo founder, junior in Go, building a learning tool while learning Go. The meta-narrative writes itself. |

### 3. Lovable — Open Source → $200M ARR

| Tactic | How it worked | Pear adaptation |
|--------|--------------|-----------------|
| **40K GitHub stars** | GPT Engineer went viral. 40K stars in 2 months. Massive top-of-funnel. | Consider open-sourcing a non-core component (context collector, prompt templates) for GitHub distribution. Or: stay closed but build an MCP server that gets listed in MCP directories. |
| **27K waitlist before launch** | Open-source reputation converted to commercial demand. | Waitlist is live. Content consistency is the conversion mechanism. Every LinkedIn post should mention the waitlist. |
| **Rebrand to expand TAM** | GPT Engineer → Lovable. From "devs" to "anyone." | Future consideration: if pear catches on with vibe-coders, the audience is broader than "junior devs." |
| **12+ growth channels simultaneously** | Social, podcasts, SEO, PH, partnerships, etc. | Don't go wide too early. LinkedIn + Substack + X + PH + HN is enough for launch. Add channels as you prove which ones convert. |

### 4. Cluely — Chaos Marketing to $7M ARR

| Tactic | How it worked | Pear adaptation |
|--------|--------------|-----------------|
| **"100 to 1" rule** | Make 100 videos. 1 goes viral. Repost it 100 times across accounts. | Film 10+ short clips per week of pear in action. Different angles, different code scenarios. Post the winners multiple times across platforms. |
| **700 clippers** | Army of people reposting content across platforms. $20-40/video. $1K bonus for 1M views. | At micro-scale: recruit 5-10 dev content creators as affiliates. 30% lifetime commission on conversions through their links. |
| **Controversy as fuel** | "Cheat on everything" was polarizing but unforgettable. | "AI is making you dumber" is pear's version. Not rage-bait, but contrarian enough to generate discussion. Lead with the problem, not the product. |
| **Founder as main character** | Roy Lee became the face/personality of Cluely. | You're the solo founder learning Go while building a Go learning tool. Lean into the personal story. |

**Warning:** Cluely's growth may have plateaued. Roy Lee admitted "maybe we launched too early." Controversy drives awareness but not retention. Product quality is what sustains growth.

### 5. v0 by Vercel — Shareable Output as Viral Loop

| Tactic | How it worked | Pear adaptation |
|--------|--------------|-----------------|
| **Every output has a URL** | Generated UIs could be shared as links. Each share = marketing. | Add `/share` to pear — generates a formatted code snippet + teaching explanation that looks great on Twitter/LinkedIn. Each share shows pear's teaching quality to the poster's audience. |
| **Screenshot-to-recreation** | Users shared side-by-side comparisons. | Side-by-side: same question to ChatGPT vs pear. The difference in teaching quality should be visually obvious. |
| **Existing ecosystem** | v0 launched into Vercel's 6M developer base. | MCP integration launches pear into Cursor and Claude Code ecosystems. Every MCP user is a potential pear user. |

---

## Pear's Viral GTM Playbook (Specific Actions)

### Phase 0: Pre-Launch (Now → Launch Day)

**Content cadence:** Daily LinkedIn posts, 3x/week Substack, daily X posts.

| Week | Content Theme | Format |
|------|--------------|--------|
| 1-2 | "The AI learning crisis" — stats, data, hot takes | LinkedIn text posts, Substack essay |
| 3-4 | "Building pear" — screenshots, terminal recordings, decisions | LinkedIn + X with images/video |
| 5-6 | "Sneak peek" — short clips of pear teaching in action | Video clips (15-30s), Twitter threads |
| 7 (launch week) | "pear is live" — full demo + install link | 60s video everywhere, PH, HN |

**Specific content pieces to create:**
1. "AI makes developers 19% slower — the data is clear" (essay with METR data)
2. "I'm a junior Go developer building a Go learning tool" (build-in-public story)
3. "What Cursor can't teach you" (positioning piece)
4. "I asked Claude and pear the same question. Here's what happened." (side-by-side demo)
5. "The $101.8B reason AI tutoring is inevitable" (L&D market thesis)

### Phase 1: Launch Week

**Day 1:** Product Hunt launch + HN "Show HN" post + LinkedIn announcement + X thread + Substack issue + email to waitlist

**The PH listing:**
- Title: "pear — AI tutor that teaches while you code"
- Tagline: "Ask a question about your code. Learn the concept, not just the fix."
- First comment: the 60-second demo video
- Maker story: solo founder, junior in Go, building a learning tool

**The HN post:**
- "Show HN: pear — CLI tutor that teaches why, not just what"
- Lead with the problem (AI learning crisis), show the product, include install command
- Be in the comments responding for 8+ hours

**The demo video (most important asset):**
- 60 seconds max
- Show: real codebase → ask question → visible context injection → structured teaching response with concept tags
- End with install command on screen
- No intro, no logo animation. Start with the action.

### Phase 2: Post-Launch (Weeks 2-8)

| Channel | Action | Frequency |
|---------|--------|-----------|
| LinkedIn | Share user stories, metrics, learnings | Daily |
| X/Twitter | Short demo clips, reply to coding discussions with "pear would explain this as..." | Daily |
| Substack | Deep-dive essays on the teaching methodology | Weekly |
| YouTube | Full pear sessions (10-20 min) showing real teaching | Weekly |
| Dev podcasts | Pitch to Changelog, Syntax, devtools.fm, CoRecursive | 2-3 appearances |
| MCP directories | List pear MCP server on mcp.so, PulseMCP | One-time |
| Dev communities | r/programming, r/ChatGPTCoding, HN comments | 3x/week |

### Phase 3: Scale (Months 2-6)

**Micro-influencer affiliate program:**
- Recruit 10-20 dev content creators (10K-100K followers)
- 30% lifetime commission on paid conversions
- Provide them with pear Pro access + talking points + demo scripts
- Based on Submagic's model ($1M ARR in 90 days with 50-70 affiliates)

**Content repurposing machine:**
- Record one 15-min pear teaching session per week
- Extract 5-10 short clips (15-30s each) using AI clipping tools
- Post shorts to: X, LinkedIn, TikTok, YouTube Shorts, Instagram Reels
- Based on Cluely's "100 to 1" principle at micro-scale

**Shareable output feature:**
- `/share` command generates a formatted teaching snippet
- Includes: question asked, context used, pear's teaching response, concept tags
- Optimized for Twitter/LinkedIn cards
- Every share is an organic ad for pear's teaching quality

---

## The Metrics That Matter

| Metric | Signal | Target |
|--------|--------|--------|
| **Second-session rate** | Product-market fit | 40%+ |
| **Free-to-paid conversion** | Value delivery | 5-8% |
| **Demo video view-to-install rate** | Marketing effectiveness | 3%+ |
| **Content engagement rate** | Audience resonance | 5%+ on LinkedIn |
| **MCP installs** | Ecosystem distribution | 1,000+ month 1 |
| **Waitlist → active user** | Launch conversion | 30%+ |
| **Share command usage** | Organic viral loop | 10%+ of sessions |

---

## Budget

| Item | Monthly Cost | Notes |
|------|-------------|-------|
| Infrastructure | $0-20 | Vercel free, Fly.io free tier, Upstash free |
| Content tools | $0-50 | Opus Clip or similar for video clipping |
| Affiliate commissions | Variable | 30% of conversions, starts at ~$0 |
| Podcast guest spots | $0 | Free distribution |
| **Total pre-revenue** | **$0-70/mo** | |

Zero paid marketing. The product and content do the work.

---

## The One Thing That Matters Most

**The 60-second demo video.**

Every successful AI tool launch had one moment that compressed the value proposition into something shareable. For Cursor it was "type English, get code." For Bolt it was "describe an app, watch it build." For pear it needs to be: **"ask about your code, learn the concept."**

If that video doesn't make a developer stop scrolling and think "I want that," no amount of content strategy will compensate.

Film it last (when the product is polished). Post it everywhere. Link every piece of content back to it. It's the centerpiece of the entire GTM.
