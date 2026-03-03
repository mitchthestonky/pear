# Pear — Financial Projections: Year One

> Version: 1.0
> Last updated: February 2026
> Author: Mitch

---

## Assumptions (All Scenarios)

### Pricing

| Tier | Monthly | Annual | Effective Monthly (Annual) |
|---|---|---|---|
| 14-day free trial | $0 | — | $0 |
| Pro | $30/mo | $300/yr | $25/mo |
| Team | TBD (v2) | TBD | TBD |

### Revenue Model

All Pro users pay $30/mo from day 1. ~30% of Pro users choose annual billing at $300/yr (effective $25/mo).

### Cost Structure

| Cost | Monthly | Notes |
|---|---|---|
| Fly.io (Go backend + Postgres) | $15-50 | Scales with usage. Free tier covers early months. |
| Vercel (Next.js) | $0-20 | Free tier for launch, Pro ($20/mo) if traffic exceeds limits. |
| Domain (pearcode.dev) | ~$2/mo | Annual payment amortized. |
| Stripe fees | 2.9% + $0.30 per transaction | ~$1 per monthly sub, ~$6 per annual sub. |
| Hosted LLM costs | ~$0.05/request × hosted requests | 50 requests/user/month = ~$2.50/user/month for Pro users using hosted. Estimated ~40% of Pro users use any hosted requests. |
| PostHog | $0 | Free tier sufficient for first year. |
| sox/Whisper/TTS | $0 | Users pay their own LLM providers (BYOK). Whisper + TTS costs are on user's OpenAI key. |

**Effective cost per Pro subscriber:** ~$3-4/mo (Stripe fees + proportional infra + partial hosted usage)

**Gross margin per Pro subscriber:** ~87-90% at $30/mo

### Conversion Assumptions

| Metric | Conservative | Likely | Ambitious |
|---|---|---|---|
| Free → Pro conversion rate | 3% | 5% | 8% |
| Monthly Pro churn rate | 8% | 6% | 4% |
| Annual billing adoption | 20% | 30% | 40% |
| Organic growth rate (month-over-month installs) | +10% | +20% | +35% |

---

## Scenario 1: Conservative

**Narrative:** Decent launch, moderate reception. PH does okay but doesn't hit top 3. HN gets some traction but doesn't front-page. LinkedIn content drives steady but not explosive growth. Product-market fit takes time to find. Teaching quality is good but not yet viral-worthy.

### Monthly Progression

| Month | New Installs | Cumulative Trial Users | New Pro Subs | Pro Churn | Active Pro | MRR | Cumulative Revenue |
|---|---|---|---|---|---|---|---|
| **1** (launch) | 300 | 200 | 6 | 0 | 6 | $180 | $180 |
| **2** | 180 | 320 | 5 | 0 | 11 | $330 | $510 |
| **3** | 200 | 440 | 6 | 1 | 16 | $480 | $990 |
| **4** | 220 | 560 | 6 | 1 | 21 | $630 | $1,620 |
| **5** | 240 | 680 | 7 | 2 | 26 | $780 | $2,400 |
| **6** | 260 | 800 | 7 | 2 | 31 | $930 | $3,330 |
| **7** | 290 | 930 | 8 | 2 | 37 | $1,110 | $4,440 |
| **8** | 320 | 1,060 | 9 | 3 | 43 | $1,290 | $5,730 |
| **9** | 350 | 1,200 | 10 | 3 | 50 | $1,500 | $7,230 |
| **10** | 380 | 1,340 | 10 | 4 | 56 | $1,680 | $8,910 |
| **11** | 420 | 1,480 | 12 | 4 | 64 | $1,920 | $10,830 |
| **12** | 460 | 1,630 | 13 | 5 | 72 | $2,160 | $12,990 |

### Conservative Year 1 Summary

| Metric | Value |
|---|---|
| **Total installs** | ~3,620 |
| **Active trial users (month 12)** | ~1,560 |
| **Active Pro subscribers (month 12)** | 72 |
| **MRR at month 12** | ~$2,160 |
| **ARR run-rate at month 12** | ~$25,920 |
| **Total Year 1 revenue** | ~$12,990 |
| **Total Year 1 costs** | ~$3,500 |
| **Year 1 profit** | ~$9,490 |
| **LinkedIn followers (month 12)** | ~12,000 |
| **Substack subscribers (month 12)** | ~2,000 |

**Milestone check:** At 72 Pro subscribers and ~$1.5k MRR, this validates demand but doesn't yet justify quitting a day job or raising funding. The product works, people pay for it, but growth is linear, not exponential.

---

## Scenario 2: Likely

**Narrative:** Strong launch. PH hits top 5. HN makes front page for a few hours. Several LinkedIn posts go semi-viral (500+ likes). The demo video resonates — people share it. Word of mouth kicks in by month 3. Teaching quality gets genuine "this is better than asking Claude directly" feedback. Build-in-public content creates a small but loyal following.

### Monthly Progression

| Month | New Installs | Cumulative Trial Users | New Pro Subs | Pro Churn | Active Pro | MRR | Cumulative Revenue |
|---|---|---|---|---|---|---|---|
| **1** (launch) | 600 | 400 | 20 | 0 | 20 | $600 | $600 |
| **2** | 350 | 600 | 15 | 1 | 34 | $1,020 | $1,620 |
| **3** | 420 | 840 | 21 | 2 | 53 | $1,590 | $3,210 |
| **4** | 500 | 1,100 | 25 | 3 | 75 | $2,250 | $5,460 |
| **5** | 600 | 1,400 | 30 | 4 | 101 | $3,030 | $8,490 |
| **6** | 720 | 1,750 | 36 | 6 | 131 | $3,930 | $12,420 |
| **7** | 860 | 2,200 | 43 | 8 | 166 | $4,980 | $17,400 |
| **8** | 1,030 | 2,700 | 51 | 10 | 207 | $6,210 | $23,610 |
| **9** | 1,240 | 3,300 | 62 | 12 | 257 | $7,710 | $31,320 |
| **10** | 1,490 | 4,000 | 74 | 15 | 316 | $9,480 | $40,800 |
| **11** | 1,780 | 4,800 | 89 | 19 | 386 | $11,580 | $52,380 |
| **12** | 2,140 | 5,700 | 107 | 23 | 470 | $14,100 | $66,480 |

### Likely Year 1 Summary

| Metric | Value |
|---|---|
| **Total installs** | ~11,730 |
| **Active trial users (month 12)** | ~5,230 |
| **Active Pro subscribers (month 12)** | 470 |
| **MRR at month 12** | ~$14,100 |
| **ARR run-rate at month 12** | ~$169,200 |
| **Total Year 1 revenue** | ~$66,480 |
| **Total Year 1 costs** | ~$12,000 |
| **Year 1 profit** | ~$54,480 |
| **LinkedIn followers (month 12)** | ~20,000 |
| **Substack subscribers (month 12)** | ~5,000 |

**Milestone check:** At 470 Pro subscribers and ~$10k MRR, this is a real business. ARR run-rate of ~$124k is approaching the "credible solo founder SaaS" threshold. Fundable on these metrics if growth trajectory holds. Team tier waitlist should have 20+ companies by now.

**Funding readiness:** At this trajectory, you could approach pre-seed investors at month 9-10 with: $6-8k MRR, 250+ Pro subscribers, 20% MoM growth, clear product-market fit signal, and a team tier waitlist to show B2B potential. A $500k-$1M pre-seed at this stage is realistic.

---

## Scenario 3: Ambitious

**Narrative:** The product goes viral. The demo video gets shared by a developer influencer (ThePrimeagen, Theo, Fireship, or similar). PH hits #1 Product of the Day. HN front-pages for 12+ hours with 300+ points. The "AI makes you fast, Pear makes you good" framing resonates with the broader conversation about AI replacing developers. Multiple tech publications pick up the story. Organic word of mouth drives exponential growth from month 2.

### Monthly Progression

| Month | New Installs | Cumulative Trial Users | New Pro Subs | Pro Churn | Active Pro | MRR | Cumulative Revenue |
|---|---|---|---|---|---|---|---|
| **1** (launch) | 2,000 | 1,400 | 112 | 0 | 112 | $3,360 | $3,360 |
| **2** | 1,500 | 2,200 | 80 | 4 | 188 | $5,640 | $9,000 |
| **3** | 2,000 | 3,400 | 120 | 8 | 300 | $9,000 | $18,000 |
| **4** | 2,700 | 4,800 | 162 | 12 | 450 | $13,500 | $31,500 |
| **5** | 3,600 | 6,500 | 216 | 18 | 648 | $19,440 | $50,940 |
| **6** | 4,900 | 8,800 | 294 | 26 | 916 | $27,480 | $78,420 |
| **7** | 6,600 | 11,800 | 396 | 37 | 1,275 | $38,250 | $116,670 |
| **8** | 8,900 | 15,500 | 534 | 51 | 1,758 | $52,740 | $169,410 |
| **9** | 12,000 | 19,800 | 720 | 70 | 2,408 | $72,240 | $241,650 |
| **10** | 16,200 | 25,000 | 972 | 96 | 3,284 | $98,520 | $340,170 |
| **11** | 21,900 | 31,500 | 1,314 | 131 | 4,467 | $134,010 | $474,180 |
| **12** | 29,500 | 39,500 | 1,770 | 179 | 6,058 | $181,740 | $655,920 |

### Ambitious Year 1 Summary

| Metric | Value |
|---|---|
| **Total installs** | ~111,800 |
| **Active trial users (month 12)** | ~33,400 |
| **Active Pro subscribers (month 12)** | 6,058 |
| **MRR at month 12** | ~$181,740 |
| **ARR run-rate at month 12** | ~$2.18M |
| **Total Year 1 revenue** | ~$655,920 |
| **Total Year 1 costs** | ~$80,000 |
| **Year 1 profit** | ~$575,920 |
| **LinkedIn followers (month 12)** | ~50,000+ |
| **Substack subscribers (month 12)** | ~15,000+ |

**Milestone check:** This is the "category-defining" scenario. $133k MRR and $1.6M ARR run-rate at month 12. This is fundable at seed stage ($2-5M) with strong terms. The team tier becomes urgent — enterprise demand will be knocking.

**Funding readiness:** Raise by month 6-8 when MRR is $12-25k and the growth curve is undeniable. At 35% MoM growth with 8% conversion rates, you'd attract top-tier pre-seed/seed investors focused on developer tools. A $2-3M seed at $15-20M valuation is realistic.

---

## Revenue Comparison: All Scenarios

### MRR Trajectory

```
Month     Conservative    Likely       Ambitious
  1          $180          $600         $3,360
  2          $330        $1,020         $5,640
  3          $480        $1,590         $9,000
  4          $630        $2,250        $13,500
  5          $780        $3,030        $19,440
  6          $930        $3,930        $27,480
  7        $1,110        $4,980        $38,250
  8        $1,290        $6,210        $52,740
  9        $1,500        $7,710        $72,240
 10        $1,680        $9,480        $98,520
 11        $1,920       $11,580       $134,010
 12        $2,160       $14,100       $181,740
```

### Year 1 Totals

| Metric | Conservative | Likely | Ambitious |
|---|---|---|---|
| Total revenue | $12,990 | $66,480 | $655,920 |
| Total costs | $3,500 | $12,000 | $80,000 |
| Profit | $9,490 | $54,480 | $575,920 |
| Month 12 MRR | $2,160 | $14,100 | $181,740 |
| Month 12 ARR run-rate | $25,920 | $169,200 | $2,180,880 |
| Pro subscribers (month 12) | 72 | 470 | 6,058 |
| Total installs | 3,620 | 11,730 | 111,800 |

---

## Funding Scenarios

### If Conservative: Bootstrap

- $9k revenue is healthy side-project income
- Continue building while employed
- Focus on improving teaching quality and conversion rate
- Raise only if growth accelerates in year 2

### If Likely: Pre-Seed ($500k-$1M)

**When to raise:** Month 9-10, when MRR is $5-7k with clear upward trajectory.

**The pitch:**
- "AI developer education" — new category, zero competition
- $5-7k MRR growing 20% MoM
- 250+ paying subscribers
- BYOK-first model = ~85% gross margins
- Team tier waitlist validates B2B potential
- Solo founder, profitable from month 1

**Use of funds:**
- First hire: developer advocate / content marketer ($80-100k)
- Second hire: senior Go engineer to accelerate CLI development ($120-150k)
- 12 months of runway to reach $50k MRR

### If Ambitious: Seed ($2-5M)

**When to raise:** Month 6-8, when MRR is $12-25k and growing 30%+ MoM.

**The pitch:**
- Category creator in AI developer education
- $15-25k MRR, 30%+ MoM growth
- 500+ paying subscribers, 8% conversion rate
- BYOK-first = capital-efficient (no upstream LLM costs)
- Team tier demand from 20+ companies
- The building-in-public story + 20k+ LinkedIn following = built-in distribution

**Use of funds:**
- Engineering team (3 hires): Go backend, ML/teaching intelligence, platform
- Go-to-market team (2 hires): developer advocate, growth marketer
- Team tier development (v2)
- 18 months of runway to reach $100k MRR

**Investor targets:**
- Developer tools focused: Heavybit, Boldstart, Redpoint
- Education focused: Reach Capital, Owl Ventures, Learn Capital
- Solo-founder friendly: Y Combinator (apply at month 4-5 for the next batch)

---

## Key Assumptions & Sensitivity

### What Swings Revenue Most

1. **Free → Pro conversion rate** is the biggest lever. The difference between 3% and 8% is the difference between $9k and $475k in year 1. The teaching quality and first-run experience are what drive this number.

2. **Churn rate** compounds fast. 8% monthly churn means you're replacing 60%+ of your base annually. Getting churn below 5% (through teaching quality, multi-turn value, feature expansion) is existential.

3. **Organic growth rate** determines whether you're linear or exponential. Linear (10% MoM) means a solid side project. Exponential (35% MoM) means a fundable startup. The content marketing flywheel is what drives this — the product alone won't grow fast enough.

### What Could Break the Model

| Risk | Impact | Mitigation |
|---|---|---|
| Teaching quality isn't visibly better than raw Claude | Conversion collapses | Invest in prompt engineering before launch. The prompt is the product. |
| Voice latency >3s kills UX | Retention drops | Parallel context collection, Whisper turbo, streaming responses. Measure obsessively. |
| Big player (Anthropic, Cursor) adds teaching mode | Growth ceiling | Move fast, build brand, accumulate data. Category ownership beats features. |
| Content marketing doesn't convert to installs | Growth stays linear | Test different content formats. Video > text for developer tools. Invest in the demo. |
| BYOK friction too high for juniors | Free → Pro drops | Improve `pear init` wizard. Consider making hosted mode more generous. |

---

## Team-Building Roadmap (If Funded)

### First 3 Hires (Pre-Seed / Early Seed)

| Role | Why | When |
|---|---|---|
| **Developer Advocate / Content Marketer** | Scale the content flywheel beyond founder bandwidth. Creates tutorials, docs, community content. Maintains LinkedIn/Substack/X. | $500k+ raised or $5k MRR |
| **Senior Go Engineer** | Accelerate CLI development. Build team tier infrastructure. Someone who can own the backend while founder focuses on product + GTM. | $500k+ raised or $8k MRR |
| **Teaching Quality / Prompt Engineer** | Full-time focus on prompt engineering, teaching methodology, response quality. The most important role for the product. | $1M+ raised or $15k MRR |

### Team of 6-8 (Seed Stage)

| Role | Purpose |
|---|---|
| Founder (Mitch) | Product, strategy, content, vision |
| Developer Advocate | Community, content, tutorials, events |
| Senior Go Engineer | CLI, backend, infrastructure |
| Prompt/Teaching Engineer | Teaching quality, prompt iteration, knowledge graph |
| Growth Marketer | SEO, paid experiments, conversion optimization |
| Full-Stack Engineer | Web dashboard, team tier, billing |
| (Optional) ML Engineer | Intent detection, concept tracking, teaching memory |
| (Optional) Designer | Brand, marketing site, dashboard UX |

### Hiring Philosophy

Hire people who are **passionate about developer education and L&D**. This isn't a company that happens to teach developers — teaching IS the company. Every hire should believe the thesis: AI is making developers faster but not better, and that's a problem worth solving.

Look for:
- Former teachers / educators who became engineers
- Senior engineers frustrated by mentoring overhead
- Developer advocates who care about education, not just marketing
- People who've built courses, written technical books, or mentored extensively

---

## The Viral Path

For Pear to go viral, one of these must happen:

1. **The demo video gets shared by a developer influencer.** ThePrimeagen, Theo, Fireship, or similar. One share to 500k+ followers changes the trajectory entirely. This is why the demo video is the most important marketing asset.

2. **The "AI learning crisis" framing resonates with mainstream tech media.** If The Verge, TechCrunch, or Ars Technica covers "developers are losing skills to AI," Pear is the only product that addresses it. Pitch this angle to tech journalists proactively.

3. **A viral LinkedIn post.** "I built an AI tutor because Copilot is making developers worse" — if one post breaks 1,000+ likes, the algorithm amplifies it to 100k+ impressions. Build-in-public content has the highest viral potential on LinkedIn.

4. **HN stays on the front page for 8+ hours.** This drives 5-10k visits in a day. The blog post angle matters more than the Show HN — "I built X because Y" stories outperform product launches on HN.

5. **Engineering managers share it internally.** If Pear solves the "my juniors don't understand their own code" problem visibly, managers will share it on Slack. This is organic B2B distribution that costs nothing.

The product doesn't need all five. It needs one. Build for all five, optimize for the demo video and the LinkedIn flywheel — those are within your control.
