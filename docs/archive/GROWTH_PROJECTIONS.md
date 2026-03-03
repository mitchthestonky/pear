# Pear — Growth Projections & Trajectory Analysis

> February 2026. Based on market research of Cursor, Bolt.new, Lovable, Windsurf, Claude Code, and industry benchmarks.

---

## Market Context

### AI Developer Tool Growth Benchmarks (2024-2026)

| Company | $0 → $100M ARR | Team Size at $100M | Total Funding |
|---------|----------------|-------------------|---------------|
| Cursor | ~12 months | 15-50 | $400M+ |
| Lovable | ~8 months | 20-40 | $20M+ |
| Bolt.new | ~12 months (projected) | 30+ | $105M |
| Windsurf | ~8 months ($12M→$100M) | 100+ | $260M |

### Conversion Rate Benchmarks

| Model | Median | Top Quartile |
|-------|--------|-------------|
| Freemium (self-serve) | 3-5% | 6-8% |
| Free trial (opt-in) | 18.2% | 35-45% |
| Lovable actual | 7.8% | — |

### AI-Native vs Traditional SaaS Growth

| ARR Band | AI-Native Median Growth | Traditional SaaS |
|----------|------------------------|-------------------|
| <$1M | 100% YoY | 75% YoY |
| $1M-$5M | 110% YoY | 40% YoY |
| $5M-$20M | 90% YoY | 30% YoY |

---

## Pear Growth Scenarios

### Scenario 1: Conservative (Solo Founder, Organic Growth)

Steady content-driven growth, no viral moments. Product is good, word-of-mouth is slow.

| Timeline | Total Users | Paid Users | MRR | ARR |
|----------|------------|------------|-----|-----|
| Launch day (PH/HN) | 200-500 | 5-15 | $100-300 | — |
| Month 1 | 800 | 25 | $500 | $6K |
| Month 3 | 2,500 | 80 | $1,600 | $19K |
| Month 6 | 6,000 | 250 | $5,000 | $60K |
| Month 12 | 15,000 | 700 | $14,000 | $168K |
| Month 18 | 30,000 | 1,500 | $30,000 | $360K |
| Month 24 | 50,000 | 3,000 | $60,000 | $720K |

**Assumptions:** 5% trial-to-paid, $30 avg revenue/user, ~15% monthly user growth, zero paid marketing.

**Outcome:** Profitable lifestyle SaaS. Proves the concept. ~$170K year 1 revenue, ~$500K year 2.

---

### Scenario 2: Viral (Content Hit + Strong Product-Market Fit)

Demo video hits 500K+ views. HN front page. MCP integration drives adoption in Cursor/Claude Code ecosystems. Word-of-mouth compounds.

| Timeline | Total Users | Paid Users | MRR | ARR |
|----------|------------|------------|-----|-----|
| Launch week | 5,000-10,000 | 50-100 | $1,000-2,000 | — |
| Month 1 | 15,000 | 500 | $10,000 | $120K |
| Month 3 | 40,000 | 2,000 | $40,000 | $480K |
| Month 6 | 100,000 | 6,000 | $120,000 | $1.4M |
| Month 9 | 180,000 | 12,000 | $240,000 | $2.9M |
| Month 12 | 250,000 | 18,000 | $360,000 | $4.3M |

**Assumptions:** 6-7% conversion, 30%+ monthly user growth sustained by virality, one 50-seat enterprise deal ($10K ARR) by month 6.

**Outcome:** Fundable. Pre-seed/seed at month 3-6. Hire 2-3 engineers. $4M+ year 1 ARR.

---

### Scenario 3: Rocketship (Funded + Ecosystem Distribution)

Raise pre-seed ($500K-$1M) at month 3-6. MCP becomes default in Claude Code/Cursor. Enterprise pipeline opens. Team of 4-5 by month 6.

| Timeline | Total Users | Paid Users | MRR | ARR |
|----------|------------|------------|-----|-----|
| Month 3 | 50,000 | 3,000 | $60K | $720K |
| Month 6 | 200,000 | 15,000 | $300K | $3.6M |
| Month 9 | 500,000 | 40,000 | $800K | $9.6M |
| Month 12 | 1,000,000 | 80,000 | $1.6M | $19M |

**Assumptions:** 8%+ conversion, 50%+ monthly growth from ecosystem distribution, multiple enterprise deals, funded team executing on marketing + product simultaneously.

**Outcome:** Series A territory. ~$19M ARR year 1. This follows ~50% of Lovable's trajectory (they had 50K GitHub stars + team).

---

## What Determines Which Scenario

| Factor | Conservative | Viral | Rocketship |
|--------|-------------|-------|------------|
| Demo video views | <10K | 500K+ | 1M+ |
| HN/PH placement | Top 10 | #1 front page | #1 + sustained discussion |
| Second-session rate | 20-30% | 40-50% | 50%+ |
| Free-to-paid conversion | 3-5% | 6-8% | 8%+ |
| MCP adoption | Nice feature | 1K+ installs | Default in ecosystems |
| Enterprise interest | 0 deals | 2-5 teams | 10+ teams |
| Content consistency | 3x/week | Daily + viral hits | Daily + media coverage |
| Team size | Solo | Solo + 1 contractor | 4-5 full-time |

---

## Revenue Composition at Scale

### At $1M ARR (Viral scenario, ~month 8)
| Source | % Revenue | Notes |
|--------|-----------|-------|
| Individual Pro ($30/mo) | 90% | Core revenue |
| Team/Enterprise | 10% | First 2-3 team deals |

### At $5M ARR (Rocketship scenario, ~month 10)
| Source | % Revenue | Notes |
|--------|-----------|-------|
| Individual Pro | 50% | Still majority but declining share |
| Team/Enterprise | 40% | Enterprise contracts drive growth |
| Additional hosted usage | 10% | Power users buying extra credits |

---

## Unit Economics

### Per-User Cost (BYOK mode — majority of users)
| Item | Monthly Cost |
|------|-------------|
| Infrastructure (Fly.io, Vercel) | ~$0.10 |
| Upstash Redis | ~$0.01 |
| Analytics/logging | ~$0.05 |
| **Total** | **~$0.16** |

### Per-User Cost (Hosted mode — 50 requests/mo)
| Item | Monthly Cost |
|------|-------------|
| LLM API (~$0.05/request × 50) | ~$2.50 |
| Infrastructure | ~$0.10 |
| **Total** | **~$2.60** |

### Margin at $30/mo
- BYOK user: **89% margin** ($26.70 profit)
- Hosted user: **78% margin** ($23.40 profit)
- Blended (80% BYOK / 20% hosted): **87% margin**

---

## Key Metrics to Track

| Metric | Target (Month 1) | Target (Month 6) | Why |
|--------|-----------------|-------------------|-----|
| Second-session rate | 40%+ | 50%+ | Product-market fit signal |
| Trial-to-paid conversion | 5%+ | 7%+ | Revenue efficiency |
| Monthly user growth | 20%+ | 15%+ | Sustainable acquisition |
| Churn (monthly) | <8% | <5% | Retention health |
| MCP installs | 100+ | 1,000+ | Ecosystem distribution |
| NPS | 40+ | 50+ | Word-of-mouth predictor |

---

## The Honest Summary

- **Conservative ($260K ARR):** Achievable solo. Proves the model. Bootstrap territory.
- **Viral ($2.03M ARR):** Requires one viral moment + sustained quality. Fundable.
- **Rocketship ($26.2M ARR):** Requires funding + team + ecosystem luck. Series A territory.

The gap between conservative and viral is almost entirely **distribution, not product.** The same product with 500K video views vs 5K produces a 25x revenue difference.

**The single highest-leverage asset:** The 60-second demo video.

**The single most important metric:** Second-session rate.

**The single biggest risk:** Speed to market. Every week of delay is a week competitors could ship a teach mode.
