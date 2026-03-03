# Pear Distribution Guide: From Zero to Self-Sustaining Growth

> A practical playbook for distributing pear, based on the product decisions, positioning, and growth strategies developed across multiple sessions. February 2026.

---

## Part 1: Pear's Distribution Advantages (What You're Working With)

### The Product Has Natural Distribution Hooks

Most developer tools have to manufacture shareability. Pear has several built-in:

**The response format is inherently screenshottable.** A pear teaching response (diagnosis → concept → production context → depth offer, with concept tags) looks visually different from a ChatGPT response. When a developer screenshots a pear response and posts it, viewers immediately see something unfamiliar. "What tool is that?" is the reaction you're engineering for.

**The learning profile is identity-reinforcing.** When `pear progress` shows "47 concepts across 12 sessions, 8-day streak," that's something developers want to share. It signals "I'm actively investing in getting better." It's the same psychology that makes people share Duolingo streaks, Strava runs, and GitHub contribution graphs. The progress display should be designed with shareability in mind from day one. Clean, compact, terminal-aesthetic, looks good in a screenshot.

**MCP puts pear inside existing viral channels.** Cursor and Claude Code already have active communities sharing tips, configs, and workflows. "Here's my MCP setup" posts are common. Pear appearing inside someone's Cursor workflow is organic, contextual exposure to exactly the right audience.

**The thesis itself is shareable content.** "AI is making developers fast but not good" is a take that generates discussion without being rage bait. It's the kind of thing a senior developer shares with the comment "this is what I've been saying." The positioning does distribution work independent of the product.

### Your Specific Assets

You're not starting from zero. Map your existing assets to distribution:

- **LinkedIn (~4.3K followers):** dev-heavy professional audience. This is your primary channel. 4.3K engaged followers in the right niche outperforms 50K generic followers. Your network skews toward startup founders, product leaders, and engineers from your recruitment and product leadership career. This is the exact audience for pear.
- **Substack:** long-form credibility builder. Essays drive email subscribers, email subscribers convert to users.
- **X/Twitter:** secondary channel for developer community engagement. Lower existing presence but higher potential for breakout content.
- **YouTube (planned):** highest leverage long-term asset. Videos compound for years.
- **The personal narrative:** this is your unfair advantage. Science degree → health IT operations → tech recruitment (where you assessed hundreds of engineering teams) → self-taught engineer → shipping production SaaS platforms → Head of Product leading engineering teams → now building pear. This arc is genuinely unusual. You understand the developer skills gap from three angles: you recruited engineers, you became one via AI-assisted self-teaching, and now you lead them. Most people posting about AI-assisted development have one perspective. You have all three.
- **Real production receipts:** IwiComply ($1M+ pipeline, 5 months as sole engineer), Uptio ($360K+ pipeline, leading engineering), RocketTags (full platform rebuild), plus freelance projects across AI chatbots, data pipelines, and community platforms. These aren't side projects. They're real systems with real stakes. Every one of them is a story.
- **Technical range:** TypeScript, Go, Python, LangChain, Supabase, AWS. You build across the stack. This credibility matters when your product is about teaching developers to understand what they're building.

---

## Part 2: The Content Engine (Practical Automation)

### The Content Supply Chain (Post-Launch Framework)

> **During the 4-6 week launch sprint, only Tier 2 matters.** Tier 1 (YouTube, Substack essays) and Tier 3 (automated derivatives) are post-launch plays. See Part 4 for the launch-specific content plan.

The goal is to produce maximum content with minimum original creation time. Everything flows from a single source and gets transformed into multiple formats.

**The source content hierarchy:**

```
Tier 1: Weekly long-form (1 per week, 60-90 min to create)  [POST-LAUNCH]
  → YouTube video OR Substack essay OR livestream recording

Tier 2: Daily short-form (1 per day, 10-15 min to create)   [NOW]
  → LinkedIn post OR X thread

Tier 3: Automated derivatives (0 min to create)              [POST-LAUNCH]
  → Clips, reposts, scheduled variations
```

The key insight: you create Tier 1 content once per week. Everything else is either Tier 2 (fast original content) or Tier 3 (automated derivatives of Tier 1 and 2).

### LinkedIn Automation System

**What to post (content pillars, tailored to your story):**

> **Note:** The percentages below are the steady-state ratios for post-launch content. During the 4-6 week launch sprint, pillar ratios shift by phase — see the Phase-Specific Content Pillar Ratios table in Part 4. The pillar descriptions and example angles below apply regardless of phase.

1. **The Self-Taught Builder (35% steady-state):** Your path from science to recruitment to self-taught engineer to product leader is genuinely unusual. This is your most differentiated angle. Most people posting about AI-assisted development are either career engineers looking down at it, or non-technical people cheerleading it. You're the rare person who learned to code in the AI era, shipped production systems, and now leads engineering teams. That gives you credibility on both sides.

   Example angles:
   - "I taught myself to code in 2022. Here's what AI tools got right and what they got dangerously wrong."
   - "From recruiting engineers to being one. What I learned about what actually makes developers good."
   - "I shipped a production SaaS platform in 5 months as a solo self-taught engineer. AI helped. But here's the part AI couldn't do."
   - Your IwiComply, Uptio, and RocketTags stories are real-stakes narratives, not hypotheticals.

2. **The AI Learning Crisis (25% steady-state):** The pear thesis, sharpened through your specific lens. You're not just theorising. You've recruited developers (pattern recognition on what makes engineers succeed), managed developers (seeing the skills gap from the leadership seat), and been the developer learning with AI (living the problem yourself). All within three years.

   Example angles:
   - "I used to assess technical capability for a living. Then I taught myself to code with AI. Here's the gap nobody's talking about."
   - Data-driven takes on vibe coding, AI reliance, skill atrophy.
   - "At Uptio I need engineers who understand what they're building, not just ship it. Here's why that's getting harder."
   - Stats, opinions, observations that position you as the authority on this problem.

3. **Building Pear in Public (25% steady-state):** Product progress, decisions, learnings. Terminal screenshots. Architecture decisions. The meta-narrative is compelling: you're learning Go with AI while building a tool that teaches people to learn better with AI. That's recursive and it resonates.

   Example angles:
   - "Today I shipped concept tracking. Here's what it looks like when pear remembers what you learned last week."
   - "Why I chose Go for a CLI tool as someone who primarily knows TypeScript."
   - Product decisions: pricing, MCP strategy, CLI design, open source vs. proprietary.
   - Honest metrics: waitlist numbers, user feedback, what broke.

4. **Product Engineering & 0-to-1 (15% steady-state):** You've done 0-to-1 at least four times. Most content creators talk about building products theoretically. You have real receipts: $360K pipeline at Uptio, $1M pipeline at IwiComply, production systems with real users. This pillar attracts the indie hacker and solo founder audience.

   Example angles:
   - "How I scope an MVP in a weekend" (design thinking + engineering background).
   - Player-coach leadership: when to code vs. when to delegate.
   - "How I built an interactive demo environment in one day that replaced weeks of engineering overhead."
   - Technical architecture decisions for early-stage products.

**The automation stack for LinkedIn:**

**Batch creation:** block 90 minutes every Sunday. The Agentic Content Engine (see Part 5) generates 7 to 10 draft posts across all four pillars, weighted to the correct ratios. You review, tweak voice and specifics, approve the best 5 to 7, and the system handles scheduling and cross-platform reformatting. This replaces writing from scratch. Your job is curation and authenticity injection, not blank-page creation.

**Scheduling:** use Taplio, Shield, or Buffer to schedule posts. **Post at midday AEST (~12:00 PM AEST / 8:00 PM ET previous day).** This deliberately prioritises engagement over peak-audience timing. You're awake, at your desk, and can execute the 15-minute comment engagement rule immediately — which matters more for the algorithm than catching the US morning window. LinkedIn's algorithm rewards early comment velocity above all else; a midday AEST post with 10 comments in the first hour outperforms a 7:30 AM ET post with zero engagement until you wake up. Tuesday through Thursday outperform Monday and Friday. Your APAC, UK-morning, and EU-afternoon audience gets served natively; US audience sees the post in their evening feed and again next morning (LinkedIn resurfaces posts with high engagement).

**Engagement automation:** the 15-minute rule. After every post goes live at midday AEST, spend 15 minutes responding to every comment. You're at your desk, the post is fresh, and your replies count as comments. LinkedIn's algorithm massively boosts posts that get comments in the first 60 minutes. With midday posting, you're live for the entire critical first hour — this is the single biggest advantage of posting on your timezone instead of scheduling for US mornings. This is the one thing you cannot automate and it's the highest leverage activity on the platform.

**Repurposing automation:** every LinkedIn post with above-average engagement gets automatically reformatted for X/Twitter. Use Typefully or Hypefury to maintain a queue of X posts derived from your LinkedIn content. Adjust the tone (LinkedIn is professional narrative, X is punchy and direct) but the ideas are the same.

**Analytics to track:** impressions per post, engagement rate (likes + comments / impressions), profile views per week, follower growth rate, and most importantly which content pillar drives the most profile views (profile views correlate with waitlist signups).

### X/Twitter Automation System (Post-Launch)

> **During the launch sprint, X is for launch-day amplification only.** Don't spend time building X presence in weeks 1-5. Reallocate those 10 minutes daily to LinkedIn comment engagement instead. The strategies below apply from month 2 onward when you have bandwidth to build a second channel.

**What works on X that doesn't work on LinkedIn:**

X rewards frequency, speed, and engagement with other accounts more than LinkedIn does. The strategy is different:

**Reply-first strategy.** Before you have a large following on X, your best distribution is showing up in other people's threads. Find tweets about AI coding tools, developer education, or learning to code. Reply with genuine insight, not a pitch. "This is exactly why I'm building pear" is a pitch. "The data actually shows AI-assisted developers take 19% longer on real tasks because they skip understanding" is insight. The insight gets clicked through to your profile. Your profile links to the waitlist.

**Thread format for teaching content.** Take your LinkedIn teaching posts and expand them into 5 to 8 tweet threads. Threads get 3 to 5x the engagement of single tweets because each tweet in the thread is a separate impression in the algorithm.

**Quote tweet pear responses.** When pear is live, screenshot a teaching response, post it with a one-line caption. "Asked pear why my Go server was leaking connections. It didn't just fix it, it taught me about connection pooling in 30 seconds." This is the atomic unit of pear marketing. It should happen daily once the product exists.

**Timing for X:** Schedule all X posts and threads for 8 to 10 AM US Eastern. This catches the US developer audience at their morning peak. Use Typefully or Hypefury's scheduling to queue posts the day before. Reddit posts follow the same US-timezone logic.

**Automation tools for X:**
- Typefully for scheduling and thread creation
- Hypefury for auto-retweeting your best performing tweets at off-peak hours (recycling good content)
- TweetHunter for finding relevant conversations to reply to (saves 30 min/day of manual searching)
- Set up saved searches for keywords: "AI coding," "cursor vs," "learning to code," "don't understand the code," "vibe coding." These surface conversations where your reply is genuinely relevant.

### The Cross-Platform Automation Flow

Here's the concrete daily and weekly workflow:

**Sunday (90 minutes):**
- Open the Agentic Content Engine review queue. The system has pre-generated 7 to 10 drafts across all four pillars with LinkedIn and X variants.
- Review and tweak each draft. Add personal details, sharpen takes, inject voice. Approve 5 to 7.
- The engine handles scheduling across LinkedIn (Taplio/Buffer) and X (Typefully) at optimal times.
- If you filmed a YouTube video that week, use Opus Clip or similar to auto-generate 5 to 10 short clips.
- Schedule clips to X and LinkedIn throughout the week.

**Daily (20 to 30 minutes):**
- 15 minutes at midday AEST: respond to all comments on today's LinkedIn post (posted at midday, you're live for the first hour)
- 10 minutes: engage with 3 to 5 relevant threads on X or Reddit (reply strategy, US-timezone content scheduled separately)
- 5 minutes: check analytics, note which topics resonated

**Monthly (1 hour):**
- Review which content pillar performed best
- Double down on what's working, cut what isn't
- Update your content bank with new ideas from user feedback and product progress

---

## Part 3: YouTube and Streaming Strategy (Post-Launch — Month 2+)

> **Skip this section during the 4-6 week launch sprint.** YouTube and streaming are deferred to post-launch. They're high-leverage long-term assets but zero-leverage for launch-week virality. Come back to this section once pear is live and the content engine is running.

### Why YouTube is Your Highest Leverage Long-Term Channel

YouTube has two properties that no other platform has:

1. **Videos compound indefinitely.** A LinkedIn post dies after 48 hours. A tweet dies after 6 hours. A YouTube video from 2024 still drives traffic in 2026. Every video you make is a permanent asset that continues to acquire users while you sleep.

2. **YouTube is a search engine.** Developers search YouTube for explanations the same way they search Stack Overflow. "Go concurrency explained," "why use goroutines," "AI coding tools compared." If pear has videos ranking for these searches, you acquire users with zero ongoing effort.

### The Self-Taught Engineer Channel Concept

Your channel identity: a self-taught engineer who builds products and startups using a systematic, planned process. This is a compelling niche because it combines three audiences:

1. **Aspiring developers** who want to see a self-taught path that works
2. **Indie hackers and solo founders** who want to see a real startup being built
3. **AI-curious developers** who want to see how AI tools fit into real product development

The pear product naturally appears in this content without being the focus. You're building a startup. You're using AI tools. You're learning Go. Pear is what you're building, so it shows up in every video organically.

### Content Format Recommendations

**Format 1: The Build Session (weekly, 15 to 30 minutes)**

Screen recording of you actually building pear. Real code, real decisions, real problems. Edited to remove dead time but not overly polished. Think of it as a documentary of the build, not a tutorial.

Structure:
- 0:00 to 0:30 — what you're building today and why
- 0:30 to 20:00 — the actual build session, narrating decisions as you make them
- 20:00 to 25:00 — what you learned, what went wrong, what's next

This format works because viewers feel like they're sitting next to you while you work. It's authentic, educational, and serialized (people come back for the next episode).

**Format 2: The Concept Explainer (2 per month, 8 to 12 minutes)**

Take a concept you learned while building pear and teach it well. "Why Go's error handling is actually genius." "What happens when your CLI tool reads a 10,000-file repo." "How I built a context collector that actually works."

These are your search-engine plays. They rank for developer queries and they demonstrate the exact teaching quality that pear delivers. The video IS a pear demo without being a product video.

**Format 3: The Startup Process (monthly, 15 to 20 minutes)**

Zoom out from the code. Talk about the product decisions, the pricing research, the GTM strategy. "How I priced my developer tool at $20/mo and why." "My launch plan for a solo-founder startup." "What I learned analyzing how Cursor got to $100M ARR."

These attract the indie hacker and solo founder audience. They position you as thoughtful and strategic, not just a coder. They also naturally reference pear's positioning, pricing, and distribution without being sales content.

### Streaming Strategy

Live streaming is high-effort but high-reward if you enjoy it. Here's how to make it work:

**Platform: YouTube Live or Twitch**

YouTube Live is better if your primary goal is building a YouTube channel, because the VOD (recorded version) stays on your channel and gets recommended alongside your regular videos. Twitch is better if you want to build a real-time community with chat interaction, but the content doesn't persist the same way.

Recommendation: YouTube Live. The recordings become Tier 1 source content that gets clipped and repurposed across all other channels.

**Stream format: "Building pear LIVE"**

2 to 3 hour sessions, 1 to 2 times per week. You're actually building the product on stream. Chat can ask questions, suggest approaches, vote on decisions. This creates a parasocial investment in pear's success. Viewers feel like co-builders.

The automation angle: every stream recording gets fed through Opus Clip or similar AI clipping tool. A 2-hour stream produces 10 to 15 potential short clips. You review them (10 minutes), pick the best 5, and schedule them across LinkedIn, X, YouTube Shorts, and TikTok for the following week.

**This means one 2-hour stream generates 5 days of cross-platform content with 10 minutes of curation effort.**

### YouTube Automation and Production

**Thumbnail strategy:** for developer content, thumbnails with terminal screenshots, code snippets, or "before/after" comparisons perform best. Use Canva templates so each thumbnail takes under 5 minutes. Consistency in thumbnail style builds channel recognition.

**Title formula:** [Emotion/Hook] + [Specific Topic] + [Implicit Promise]. Examples:
- "I Rebuilt My Entire CLI Architecture in One Day (here's why)"
- "The AI Learning Crisis Nobody's Talking About"
- "How I'm Building a $1M Solo Startup with AI Tools"

**SEO:** include target search terms in title, description, and tags. Use TubeBuddy or vidIQ to find what developers are actually searching for. Optimize for search-driven discovery, not just subscriber-driven discovery.

**Chaptering:** add timestamps to every video. This improves watch time (viewers jump to relevant sections instead of leaving) and YouTube surfaces individual chapters in search results.

---

## Part 4: The Pear-Specific Distribution Playbook

### The Compressed 4-6 Week Launch Timeline

The original phasing assumed 7+ weeks. With a 4-6 week window and full-time employment at Uptio, every week needs to carry more weight. YouTube, streaming, and X audience-building are deferred to post-launch. LinkedIn is the only channel that matters pre-launch. X is launch-day amplification only.

### Phase-Specific Content Pillar Ratios

The pillar weights should shift by phase. The thesis creates problem awareness. The reveal converts awareness to waitlist. The launch activates the waitlist. Post-launch settles into the balanced split.

| Phase | Self-Taught Builder | AI Learning Crisis | Build in Public | Product Eng & 0-to-1 |
|-------|--------------------|--------------------|-----------------|----------------------|
| Weeks 1-2 (thesis) | 20% | 50% | 10% | 20% |
| Weeks 3-4 (reveal) | 25% | 25% | 40% | 10% |
| Week 5+ (launch) | 20% | 20% | 40% | 20% |
| Post-launch (steady) | 35% | 25% | 25% | 15% |

**Critical framing note:** your post analysis shows personal/identity posts outperform professional insight posts by 2-3x on likes and 6-22x on comments. This means ALL pillar content — including the thesis — must be delivered through personal narrative, not through insight-first framing. "I watched a junior dev at Uptio ship a feature in 20 minutes with Cursor. When I asked how it worked: silence." Not: "The data shows AI-assisted developers take 19% longer on real tasks."

### Phase 0: Thesis Establishment (Weeks 1-2)

**Objective:** establish the AI learning crisis narrative. Build problem awareness. No product mention yet. pearcode.dev waitlist link in bio and post footers.

**Daily actions:**
- 1 LinkedIn post at midday AEST (from Sunday batch, 50% thesis-weighted)
- 15 minutes responding to every comment in the first hour (non-negotiable, highest-leverage activity — you're at your desk)
- Optional: 3 to 5 replies on relevant X threads or Reddit threads when you have spare minutes (X/Reddit content scheduled for 8-10 AM ET separately)

**Content themes (all delivered through personal narrative):**
- The AI skills gap from your recruitment perspective
- What you've observed leading engineers at Uptio
- Your own experience learning to code with AI — what it taught you and what it didn't
- Data points wrapped in personal stories (METR study, vibe coding discourse)

**Breakout post strategy (see detailed section below):** invest disproportionate effort in 2 breakout-candidate posts in these two weeks. These are the posts designed to escape your existing audience.

**Preparation running in parallel:**
- Product Hunt listing: create account, prepare assets, recruit 20-30 people to upvote/comment on launch day
- Show HN: draft the post, prepare a technical write-up for the HN audience
- Agentic Content Engine: building async via Claude Code (see Part 5)
- pearcode.dev: waitlist live, positioning clear, demo video placeholder, Plausible/Fathom analytics installed
- UTM attribution: set up UTM parameters for all links, configure waitlist source tracking in Supabase/Upstash
- GitHub: public repo set up with optimised README, install scripts, open-source shell committed
- Blog: first 1-2 thesis posts live on pearcode.dev/blog, cross-posted to Dev.to and Hashnode with canonical URLs
- Reddit: identify target subreddits, establish presence by commenting on relevant threads (don't post about pear yet)

**Goal by end of week 2:** 100+ waitlist signups, 2-3 posts above 50 likes, at least 1 breakout post above 100 likes, 200+ new followers.

### Phase 1: The Reveal (Weeks 3-4)

**Objective:** shift from "there's a problem" to "I'm building the solution." First screenshots. First terminal demos. Waitlist urgency increases.

**Content themes:**
- "I've been building something." The first screenshot of pear in action (the atomic demo moment — see section below).
- Technical decisions as personal narrative: "Why I chose Go as someone who primarily knows TypeScript."
- Product philosophy: "I don't want pear to just fix your code. I want it to make you understand why it was broken."
- Continued thesis content, now with "and here's what I'm doing about it" as the second half.

**The atomic demo moment:** by week 3, you need ONE perfect terminal screenshot of pear teaching something. This screenshot appears in: the reveal post, the PH listing, the Show HN post, the waitlist landing page, and the launch announcement. Engineer this specific response. It IS your marketing.

**Beta preparation:**
- Identify 15-20 developers in your LinkedIn network who engaged with your thesis posts
- DM them personally: "Hey, I've been building something related to those posts about AI and learning. Would you be up for early access in a couple weeks?"
- Build the list before you need it. Don't cold-DM on launch day.

**Goal by end of week 4:** 300+ waitlist signups, 15-20 confirmed beta testers, the atomic demo screenshot ready, PH listing prepared, Show HN draft written, GitHub repo public with optimised README, 3-5 blog posts live on pearcode.dev.

### Phase 2: Compressed Beta (Week 5)

**Objective:** get 15-20 hand-picked users into pear. Collect screenshots, testimonials, and feedback. Iterate on prompts.

**This is a 1-week sprint, not a 2-week phase.** You're not looking for product-market fit data yet. You're looking for: (a) does the product work reliably, (b) are there any response quality issues to fix before public launch, and (c) can you get 3-5 shareable screenshots from real users.

**The beta feedback loop:**
1. User installs pear, completes wizard, asks first question
2. You personally follow up within 24 hours: "How was your first session? Screenshot your favorite response."
3. If the response is good, ask: "Mind if I share this?" Post it on LinkedIn with their permission.
4. If the response is bad, ask: "What went wrong?" Fix the prompts.

**Every beta user screenshot that gets posted is a micro-ad for pear.** This is your most powerful content during beta. Real users, real code, real teaching responses. It's credible in a way that marketing copy never will be.

**Content this week:** 2-3 LinkedIn posts featuring beta user screenshots and reactions. "One of our early testers asked pear about [concept]. Here's what happened." These posts bridge the reveal phase into the launch and generate FOMO for the public launch.

### Phase 3: Public Launch (Week 6)

**Preparation checklist (must be done before launch day):**
- [ ] Product Hunt listing finalised: logo, tagline, description, screenshots, maker comment drafted, 60-second demo video
- [ ] 20-30 people confirmed to upvote and comment on PH within first 2 hours
- [ ] Show HN post drafted: technical, honest, links to open-source GitHub repo, leads with the thesis
- [ ] Waitlist email nurture sequence scheduled (see section below)
- [ ] LinkedIn announcement post written and reviewed (this is your highest-effort post ever)
- [ ] X thread drafted (8-10 tweets, scheduled for 8-10 AM ET)
- [ ] Reddit posts drafted for r/golang, r/programming, r/cursor (scheduled for 8-10 AM ET)
- [ ] Substack essay ready: "AI makes you fast. Pear makes you good." Cross-posted to Dev.to and Hashnode.
- [ ] 3-5 beta user screenshots/testimonials collected and permission granted to share
- [ ] pearcode.dev updated with demo, screenshots, install command, social proof section, analytics tracking live
- [ ] GitHub repo public with optimised README: demo GIF, install command, quickstart, star request
- [ ] UTM parameters configured for all launch day links (PH, LinkedIn, X, Reddit, HN, email)
- [ ] Post-install onboarding email sequence configured (Day 0, 2, 5, 14)
- [ ] "How did you hear about pear?" question added to setup wizard

**The staggered multi-channel launch:**

PH and Show HN both require constant founder engagement in comments to succeed. Running them simultaneously splits your attention and weakens both. Stagger them: PH on Day 0, Show HN on Day +1 or +2. This also gives you Day 1 install numbers and user screenshots to reference in the HN post, adding credibility.

**Day 0 (Launch Day) — Product Hunt + LinkedIn + X + Email:**

**PH launch (7:00 PM AEST / midnight Pacific):**
- Product Hunt listing goes live when PH resets at midnight Pacific (7:00 PM AEST)
- Spend the evening in PH comments (7 PM - 10 PM AEST)

**Next morning (midday AEST):**
- LinkedIn announcement post goes live at midday AEST — your highest-effort post ever. You're live for the full first-hour engagement window.
- Email to full waitlist: final email in the nurture sequence with install command and demo video
- Substack essay published: the long-form version of the thesis + product + story (shared via LinkedIn and X, not as its own distribution channel unless you have active subscribers — see note below)

**US morning (8-10 AM ET / ~11 PM - 1 AM AEST):**
- X thread goes live: 8 to 10 tweets, punchy, demo GIF/screenshot in tweet 2 (scheduled for US developer audience)
- Reddit posts go live in relevant subreddits (see Developer Community Channels below)

**All Day 0:**
- Respond to every PH comment (this is how you stay on the front page)
- Respond to every LinkedIn comment (15-minute rule, all day)
- Respond to every X reply
- Be available for new users hitting issues

**Evening Day 0:**
- Post a "Day 1" update with real numbers. "247 installs. 43 questions asked. Here's a screenshot of the best teaching response so far." Transparency is content. This post often outperforms the launch announcement itself.

**Day +1 or +2 — Show HN:**

- Post Show HN in the morning (8-10 AM ET / 12-2 AM AEST)
- Reference Day 0 numbers: "[X] developers installed pear yesterday. Here's what I learned." HN values real data.
- Include the best user screenshot or teaching response from Day 0
- Be available in HN comments for 6+ hours. Respond to every question and criticism thoughtfully.
- This day is HN-focused. LinkedIn and PH can coast on momentum from Day 0.

### The Breakout Post Strategy

Virality is non-linear. You don't need 30 good posts — you need 2-3 posts that escape your existing audience and reach 50K+ impressions. These need disproportionate preparation.

**Breakout candidate 1: "The recruiter who became an engineer" (week 1)**

Frame: "I used to assess technical capability for a living. I interviewed hundreds of engineers. Then I taught myself to code with AI. Here's the gap nobody's talking about."

Why this works: it combines your unique arc with the controversial thesis. The hook is identity-based (your top-performing format) but the payload is the AI learning crisis thesis. People who recruited or were recruited will engage. Engineers will debate. It creates the "who is this person?" reaction.

**Breakout candidate 2: "The data post" (week 2)**

Frame: Take a specific data point (METR study, or your own observation managing engineers at Uptio) and wrap it in personal narrative. "At Uptio, I need engineers who understand what they're building. Last month, something happened that made me realise how hard that's getting."

Why this works: data + personal stake + debatable conclusion. The professional insight posts underperformed because they lacked tension. This one has tension: there's a real problem in your real team.

**Breakout candidate 3: "The reveal" (week 3)**

Frame: the first screenshot of pear. "I've been posting about AI and developer learning for weeks. Here's what I've been building." Terminal screenshot. One-line explanation. Link to waitlist.

Why this works: if breakout posts 1 and 2 landed, your audience has context for why this matters. The reveal post converts accumulated problem-awareness into waitlist signups.

**For each breakout post:**
- Write 3-5 hook variants. Test the opening line with 2-3 trusted people before posting.
- Post on Tuesday or Wednesday at midday AEST. You're at your desk for the entire first-hour engagement window — this is when your replies drive the algorithm boost.
- Spend the full first 30 minutes replying to every comment in real time.
- Have 3-5 people from your network ready to comment substantively (not just "great post") in the first hour. If they're in US timezones (8 PM ET their time), they'll be online.
- For X cross-posts of breakout content: schedule for 8-10 AM ET the next morning to catch the US developer audience fresh.

**Fallback plan if breakout posts don't break out:**

Not every engineered viral post actually goes viral. If none of the three candidates hit 100+ likes, don't panic — adjust:

1. **Double down on the personal narrative format that already works.** Your top performers (95-97 likes) were personal/identity posts, not thesis posts. If the thesis framing isn't landing, lean harder into your story and weave the thesis in as a secondary element rather than the lead. "Here's what I've been building and why" instead of "Here's the problem with AI coding."
2. **Analyse the comments, not just the likes.** A post with 40 likes and 15 comments is outperforming a post with 80 likes and 2 comments. If the thesis posts are generating debate (even at lower like counts), they're working — the algorithm is boosting them to the right people.
3. **Try shorter formats.** If long-form thesis posts underperform, try the 3-4 sentence hook format: one provocative observation, one personal stake sentence, one question. Your freelancing post was essentially this format and it hit 95 likes.
4. **The reveal post (breakout candidate 3) has the highest floor.** Even if thesis posts underperform, the first screenshot of pear in action has inherent novelty. If breakout candidates 1 and 2 don't land, invest extra preparation into making the reveal post carry the load.

### LinkedIn Headline and Profile Optimisation

Your current headline: "Leading product @ Uptio | Vibing and Coding." When someone clicks through from a viral thesis post, the headline is the first thing they see. "Vibing and Coding" undercuts the serious thesis about AI making developers not good. This is a 30-second change with outsized impact on profile-visit-to-waitlist conversion.

**Phase-specific headline changes:**

- **Weeks 1-2 (thesis phase):** "Self-taught engineer | Head of Product @ Uptio | Writing about what AI gets wrong about developer learning"
- **Weeks 3-6 (reveal through launch):** "Building pear — the CLI that teaches while you code | Head of Product @ Uptio"
- **Post-launch (steady state):** "Founder @ pear | Head of Product @ Uptio | Teaching developers to understand, not just ship"

Also update your LinkedIn bio/About section during weeks 3-4 to include: what pear is (one sentence), link to pearcode.dev, and your personal narrative (the science → recruitment → self-taught engineer arc). Every profile visit during the content sprint should have a clear path to the waitlist.

### Strategic Amplification Playbook

The fastest path to virality from 4.3K followers isn't just your content — it's getting people with larger audiences to engage with it.

**Build your amplification list (do this in week 1):**
- 5-10 LinkedIn connections with 15K+ followers who would authentically engage with the thesis (startup founders you placed at Talent Army, the Uptio CEO, senior engineers you've worked with, dev influencers you're connected to)
- People who've commented on your posts before (they're already warm)
- People posting about AI coding tools, vibe coding, developer education (aligned interests)

**How to activate them (do NOT ask people to share your posts):**
- DM before your breakout posts: "Hey, I've been thinking about this AI learning gap I keep seeing. Writing something about it — would love your take when it goes up."
- This is asking for their opinion, not their amplification. But if they comment, the algorithm amplifies.
- Tag sparingly: one well-placed tag per post maximum, only people who will genuinely engage.
- After your reveal post: DM your amplification list with early access to pear. "I built the thing I've been posting about. Want to try it before launch?"

**The network effect:** every comment from a 15K+ follower account exposes your post to a slice of their audience. 5 substantive comments from large accounts in the first hour can 5-10x your reach.

### The Atomic Demo Moment

Every dev tool that launched with virality had ONE clear visual moment. Cursor had autocomplete. v0 had UI generation. Devin had the SWE-bench video.

**For pear, you need to engineer this specific moment before week 3.**

Criteria for the perfect demo screenshot:
1. **Visually distinctive.** Must look immediately different from a ChatGPT or Claude response. The concept tags, the teaching structure (diagnosis → concept → production context), the terminal aesthetic — these should be visible at a glance.
2. **Immediately comprehensible.** A developer scrolling LinkedIn should understand what's happening in 3 seconds. The question should be relatable ("why is my Go server leaking connections?") and the response should visibly teach, not just fix.
3. **Creates "what tool is that?" curiosity.** The screenshot should make people stop scrolling and check the comments. That reaction drives engagement.

**How to engineer it:**
- Generate 20-30 pear teaching responses across different scenarios and languages
- Screenshot all of them in a clean terminal with good contrast
- Pick the top 3 that are most visually striking
- Test them: send to 5 developer friends and ask "what's your first reaction?"
- The winner becomes your atomic demo asset. It appears in: your reveal post, PH listing, Show HN, waitlist landing page, launch announcement, X thread header image.

### Product Hunt Preparation (Start Week 1)

PH is not a launch-day task. It's a 4-week preparation project.

**Week 1:** Create your PH maker account. Study top-performing developer tool launches. Draft your tagline and description. Begin recruiting upvoters.

**Week 2-3:** Prepare all assets: logo (512x512), gallery images (5-6 screenshots including the atomic demo moment), 60-second demo video (screen recording with voiceover, or animated GIF walkthrough), maker comment draft (personal, honest, tells the story of why you built pear).

**Week 3-4:** Recruit 20-30 people to upvote and leave genuine comments in the first 2 hours. These should be real users or supporters, not engagement pods. Prepare 5-10 responses to likely questions ("how is this different from X?" "does this work with Y language?" "what's the pricing?").

**Launch day timing:** PH resets at midnight Pacific. Launch at 12:01 AM Pacific (6:01 PM AEST). Have your maker comment ready to post immediately. Be in the PH comments all day.

### Show HN Strategy

Show HN can be the single highest-traffic event for a developer CLI tool. But it's high-variance and the audience is unforgiving of anything that feels like marketing.

**The post format that works for Show HN:**

Title: "Show HN: Pear – a CLI that teaches you while you code (instead of just fixing your code)"

Body: 2-3 paragraphs. Technical, honest, no hype. Lead with the problem (AI tools fix code but developers aren't learning). Describe what pear does technically (context collection, concept tracking, teaching responses). Link to the repo/install command. End with "I'd love feedback."

**What HN values:** technical substance, honesty about limitations, founder engagement in comments, open source credibility, novel approach to a real problem.

**What HN punishes:** marketing language, inflated claims, not responding to comments, "AI-powered" buzzword soup, anything that looks like an ad.

**Timing:** post between 8-10 AM Eastern on a weekday (Tuesday-Thursday best). Be available in the comments for at least 6 hours. Respond to every question and criticism thoughtfully.

### Pre-Launch Email Nurture Sequence

A cold email to a stale waitlist on launch day converts poorly. Warm them up:

**Day -7 (one week before launch):**
Subject: "Pear launches next week"
Body: "You signed up for the pear waitlist [X] weeks ago because you care about actually understanding the code AI helps you write. Next week, pear goes live. Here's a 30-second preview of what it looks like." [Atomic demo screenshot]. Short, visual, re-engagement.

**Day -3:**
Subject: "Early access in 3 days"
Body: "Pear launches publicly on [date]. But waitlist members get access 24 hours early. Here's exactly what to expect on install day." [Brief walkthrough: install command, first question, what the response looks like]. Builds anticipation, reduces friction.

**Day -1:**
Subject: "Tomorrow morning"
Body: "Pear goes live tomorrow at [time]. Here's the install command: `[command]`. Save this email. I'll send you a follow-up the moment it's live." One-line email. Urgency.

**Day 0 (launch day):**
Subject: "Pear is live. Install now."
Body: "It's here. [Install command]. [One-line description]. [Link to demo video]. [Link to PH page if you want to support us]." Direct, no fluff.

**Day +1:**
Subject: "[X] developers installed pear yesterday"
Body: "Day 1 numbers: [installs] installs, [questions] questions asked, [concepts] concepts tracked. Here's the best teaching response from yesterday: [screenshot]. If you haven't installed yet: [command]." Social proof. Transparency. Second-chance activation.

**Scaling the sequence:** if your waitlist is under 100 people by Day -7, simplify to Day -1 and Day 0 emails only. The full 5-email sequence is highest-leverage when you have 200+ subscribers to activate.

### Substack: Role and Preparation

Substack is **not a distribution channel** during the launch sprint unless you already have active subscribers. It's a **landing page for the long-form version of your story.** The Substack essay on launch day ("AI makes you fast. Pear makes you good.") gets shared via LinkedIn and X — it doesn't have its own audience to notify.

**If you want Substack to have distribution power by launch day:**
- Publish 2-3 thesis essays during weeks 1-4 and share them on LinkedIn as "full essay in comments" posts
- Each essay drives a handful of Substack subscribers organically
- By launch day, you may have 50-100 subscribers who get notified directly when the launch essay drops
- This is a nice-to-have, not a priority. Don't let Substack essay writing compete with LinkedIn post creation.

**Post-launch:** Substack becomes the weekly/bi-weekly long-form channel. Deeper analysis, user data insights, product philosophy. This is where the Tier 1 content from the Content Supply Chain lives.

### Contingency: If the Product Isn't Ready at Week 5

Software timelines slip. If pear development takes 1-2 weeks longer than expected, the content strategy is resilient:

**Weeks 1-4 content is product-independent.** The thesis establishment (weeks 1-2) and the reveal (weeks 3-4) work even with work-in-progress screenshots. Architecture decisions, design philosophy, "here's what I'm building and why" posts don't require a finished product. This content can stretch to week 5-6 without feeling stale.

**If launch slips by 1-2 weeks:**
- Continue reveal-phase content (40% build-in-public, architecture and design decisions)
- Convert the beta phase to an "alpha preview" — give 5-10 close contacts access to whatever exists, even if it's rough, to start collecting real reactions
- Shift PH listing and Show HN prep forward accordingly
- The email nurture sequence timing is relative (Day -7, -3, -1, 0) so it moves with the launch date
- Your audience momentum from weeks 1-4 is durable — a 1-2 week delay won't erode it if you keep posting

**What NOT to do:** don't launch the product before it's ready just to hit the timeline. A buggy Day 1 experience is worse than a delayed launch. The content sprint builds audience that persists. The product experience determines whether they stay.

### Phase 3: Growth (months 2 to 6)

**The content machine is now self-fueling.** You have users generating screenshots. You have usage data generating insights. You have concept tracking data telling you what developers struggle with most (this is content gold).

**Weekly content:**
- 1 YouTube video (build session, explainer, or startup update)
- Clips from that video spread across LinkedIn, X, Shorts
- 1 Substack essay (deeper analysis of something you learned from user data)
- 5 LinkedIn posts (batch-created Sunday)
- Daily X engagement

**The shareable output feature.** Build `/share` into pear. User types `/share` after a teaching response and gets a formatted snippet optimized for posting on LinkedIn or X. It includes the question, the concept tags, and the first 3 to 4 sentences of the teaching response with a "Learn more with pear" link. Every share is an organic ad created by the user, for the user's audience, featuring pear's actual teaching quality.

**The affiliate play (month 3+).** Recruit 10 to 20 dev content creators with 10K to 100K followers. 30% lifetime commission on conversions through their referral link. Provide them with pear Pro access, talking points, and demo scripts. Based on what worked for Submagic ($1M ARR in 90 days with 50 to 70 affiliates).

### Phase 4: MCP Distribution (month 2+)

**MCP is pear's Trojan horse.** It puts the product inside existing workflows without requiring anyone to change their behavior.

**The MCP distribution sequence:**
1. Ship `pear mcp` server
2. List on mcp.so and PulseMCP
3. Post "how to add pear to your Cursor setup" tutorial on YouTube, LinkedIn, X
4. Engage in Cursor and Claude Code community discussions about MCP configs
5. Every "share your MCP setup" thread becomes distribution for pear

**The MCP interjection angle is your viral mechanic.** When pear inside Cursor surfaces a teaching moment that genuinely helps a developer understand something, that developer screenshots it and posts it. "My Cursor setup just taught me about connection pooling without me asking." That's a share that creates curiosity, demonstrates value, and links back to pear.

### Phase 5: Open Source & GitHub Strategy (Launch Day Onward)

For a CLI tool, GitHub is as important a distribution channel as LinkedIn. HN heavily rewards open source. GitHub trending drives discovery to thousands of developers actively looking for tools. But pear's competitive advantage is the teaching engine — prompts, context collection strategy, concept tracking — and that stays proprietary. The strategy is **selective open source:** enough to earn GitHub credibility and community goodwill, without exposing the moat.

> **The exact scope of what goes open source is a pending decision** (see PRODUCT_DECISIONS.md). The principles below are settled; the specific components and repo structure are not. This section focuses on the GitHub distribution strategy regardless of how much source is public.

**What definitely stays proprietary (the IP):**

| Component | Why It's Protected |
|-----------|-------------------|
| Teaching prompts and mode system | The pedagogical format (diagnosis → concept → production context → depth offer) is the core differentiator. The actual prompt engineering is the moat. |
| Context collection implementation | Token budgeting (~25K tokens), parallel gathering, middle-truncation, priority weighting. This is what makes pear responses contextually better than "paste your code into ChatGPT." |
| Prompt assembly pipeline | How context, mode frames, user profile, learning history, and the teaching prompt compose the final LLM request. The orchestration logic is the product. |
| Concept tracking algorithms | How concepts are extracted, scored, and tracked. This becomes the data flywheel and the enterprise proficiency mapping feature. |
| Role frame system | The pre-built lenses (senior-staff, security, performance) that shape teaching responses. |

**Candidates for open source (decision pending):**

These are components that could be open-sourced without exposing the teaching engine. The distribution benefit of each is noted — the decision will weigh these benefits against the effort of maintaining a public codebase:

- **CLI framework and TUI shell** — generic Cobra + Bubble Tea scaffolding. Distribution: developers see clean Go code, star the repo.
- **MCP server transport** — standard JSON-RPC stdio. Distribution: drives MCP ecosystem adoption.
- **LLM adapter interfaces** — provider abstraction for Claude/OpenAI/Gemini. Distribution: community contributes new adapters.
- **Config system** — TOML parsing. Distribution: user trust ("I can see what data it reads").
- **Doctor command** — system health check. Distribution: useful standalone utility.
- **Install scripts and build tooling** — Goreleaser, Homebrew tap. Distribution: `brew install` requires a public tap repo.
- **Extension framework** — if pear supports custom role frames. Distribution: community investment in the ecosystem.

**Key constraint:** the Homebrew tap formula must live in a public repo for `brew install pearcode/tap/pear` to work. At minimum, the tap repo and install scripts will be public. Beyond that, the scope is open.

**GitHub README as a conversion page:**

The README is the landing page for every developer who discovers pear through GitHub. Structure it for conversion:

1. **One-line description:** "The CLI that teaches you while you code."
2. **Atomic demo GIF/screenshot** — the same terminal screenshot from your PH listing and launch posts. Immediately shows what pear does.
3. **Install command** — `brew install pearcode/tap/pear` or `go install github.com/pearcode/pear@latest`. Copy-pasteable, one line.
4. **30-second quickstart** — "Run `pear`, complete the 3-question setup, ask your first question. Pear reads your codebase, teaches you while it helps."
5. **What makes pear different** — 3 bullet points: context-aware, teaches not just fixes, tracks your learning.
6. **Link to pearcode.dev** — for pricing, full docs, and the waitlist.
7. **Badges** — star count, release version, Go version, license.

**GitHub Stars strategy (launch week):**

Stars are social proof. Getting to 100 stars in the first week puts pear on GitHub "trending" lists for Go and CLI tools, which is organic discovery to thousands of high-intent developers.

- Ask all 15-20 beta testers to star the repo
- Include "Star us on GitHub" in the launch day LinkedIn post and X thread
- Add to the PH listing gallery: a screenshot showing the GitHub repo with the star button visible
- Post the repo link in relevant Discord servers and Reddit threads on launch day
- Track stars as a distribution metric alongside installs

**Open source credibility for HN:**

The Show HN post should link to the GitHub repo (whatever is public by launch). HN commenters will inspect the code — having clean, well-documented Go code demonstrates engineering quality. Prepare for the "why isn't it fully open source?" question in your HN comment: "The teaching engine (prompts, context collection, concept tracking) is proprietary — that's the product. [Whatever is open-sourced] is public." Be direct and honest about it. HN respects clarity on business model more than pretending to be fully open source.

### Phase 6: Developer Community Channels (Launch Day + Ongoing)

The guide was over-indexed on LinkedIn + X. Developers discover and discuss tools in communities the content strategy wasn't reaching.

**Reddit (launch day + weekly):**

Reddit is allergic to self-promotion but rewards genuine utility. The strategy is value-first, product-second.

- **r/programming** (5M+ members): launch day post framed as "Show r/programming: I built a CLI that teaches instead of just fixing your code." Technical, show the architecture, invite feedback. One post only — don't spam.
- **r/golang** (200K+): "Show r/golang: Built a teaching CLI in Go — here's the architecture." The Go community loves seeing real Go projects. Talk about your experience learning Go with AI (meta-narrative is compelling here). Link to the open-source repo.
- **r/learnprogramming** (4M+): post about the AI learning crisis thesis, not the product. "I taught myself to code with AI. Here's the skill gap nobody warned me about." If it resonates, people will check your profile and find pear organically.
- **r/ExperiencedDevs**: thesis content about what you're seeing managing engineers at Uptio. Senior developers engage with this topic.
- **r/cursor** and **r/ClaudeAI**: "I built an MCP tool that adds teaching to your Cursor workflow." Directly relevant audience.

**Timing:** all Reddit posts at 8-10 AM US Eastern. Reddit's algorithm heavily favors US-morning posts.

**Rules:** never post the same content to multiple subreddits simultaneously (cross-posting penalty). Space Reddit posts across different days. Engage authentically in comments — Reddit detects and punishes "post and ghost."

**Discord servers (ongoing):**

- **Cursor Discord**: participate in MCP discussions, share pear's MCP integration, help people with setup issues. Be a community member first.
- **Claude Discord**: same approach. Share how pear uses Claude's API for teaching responses.
- **Dev community servers** (various): look for "what tools do you use?" threads and share pear where genuinely relevant.

Discord is not a scalable channel — it's a credibility and early-adopter channel. 10 engaged Discord users who love pear become 10 organic advocates.

**Slack communities:**

- **Indie Hackers Slack**: your build-in-public content resonates here.
- **Rands Leadership Slack**: your player-coach and engineering leadership perspective fits.
- **Language-specific Slacks** (Gophers Slack, etc.): share the Go learning story.

**Dev.to and Hashnode (weeks 2+, then ongoing):**

These platforms have built-in SEO and developer discovery. Posts on Dev.to and Hashnode rank in Google searches, driving traffic for months after publication. The content is the same as your Substack essays and blog posts — publish across all three.

- **Week 2:** publish your thesis essay ("AI is making developers fast but not good") on Dev.to and Hashnode. Use canonical URL pointing to pearcode.dev/blog if you have the blog live, or to Substack.
- **Week 3-4:** publish "Why I'm building pear" — the product reveal essay. Cross-post from Substack.
- **Launch day:** publish "How pear teaches developers while they code" — the technical deep-dive.
- **Ongoing:** every Substack essay gets cross-posted to Dev.to and Hashnode with canonical URLs. Zero additional writing effort, multiplied distribution.

**Why canonical URLs matter:** set the canonical URL to your own domain (pearcode.dev/blog) so Google credits your site, not Dev.to, for the content. Dev.to and Hashnode both support canonical URLs.

**Package registries (launch day):**

Being discoverable through package managers is zero-effort organic distribution:
- **Homebrew tap:** `brew install pearcode/tap/pear` — requires the tap repo to be public on GitHub.
- **`go install`:** `go install github.com/pearcode/pear@latest` — works automatically if the public repo is set up correctly.
- **AUR package** (Arch Linux): community contribution opportunity. Arch users are power users who adopt CLI tools.
- **Scoop** (Windows): if Windows is supported.

### Phase 7: SEO Strategy for pearcode.dev (Weeks 2+, Compounds Over Time)

Social media traffic is ephemeral — a LinkedIn post dies after 48 hours, a tweet after 6 hours. SEO-optimised content on your own domain drives traffic for years. This is compound growth that social media can't provide.

**The blog as a distribution channel:**

pearcode.dev/blog should host the same content you write for Substack, Dev.to, and Hashnode — but on your own domain. Your domain builds authority. Every blog post is a permanent landing page that ranks in Google and funnels traffic to the waitlist/install page.

**SEO content strategy (mapped to content pillars):**

| Content Pillar | Target Keywords | Example Blog Posts |
|---------------|----------------|-------------------|
| AI Learning Crisis | "AI coding problems," "vibe coding risks," "AI developer skills gap" | "Why AI Is Making Developers Fast but Not Good" |
| Self-Taught Builder | "self-taught engineer," "learn to code with AI 2026," "career change to software engineering" | "What I Learned Recruiting 200 Engineers Then Becoming One" |
| Building Pear | "AI coding assistant that teaches," "learn to code CLI tool," "alternative to ChatGPT for learning code" | "How Pear Teaches Developers While They Code" |
| Product Engineering | "build MVP in a weekend," "solo founder technical decisions," "0 to 1 product engineering" | "How I Scope an MVP in a Weekend" |

**Technical SEO for pearcode.dev:**
- Each blog post should have: meta title (60 chars), meta description (155 chars), Open Graph image, structured data (Article schema).
- Internal linking: every blog post links to the install page and 2-3 other relevant posts.
- Sitemap: auto-generated and submitted to Google Search Console.
- Page speed: Next.js with static generation means fast pages by default.

**The content workflow:** write once, publish everywhere. Each piece of content starts as a LinkedIn post or Substack essay, then gets expanded into a blog post on pearcode.dev with SEO optimisation, then cross-posted to Dev.to and Hashnode with canonical URL pointing to pearcode.dev. The LinkedIn post drives immediate traffic; the blog post drives traffic for years.

**SEO targets:**
- Month 1: 5 blog posts live, Google Search Console connected, sitemap submitted.
- Month 3: 15+ posts, first posts appearing in Google search results.
- Month 6: 1,000+ monthly organic visitors from search. Key posts ranking page 1 for long-tail developer queries.

### Funnel Model and Attribution Infrastructure

The content and distribution strategy above describes *activities*. This section defines the *measurement* that tells you which activities are actually working.

**The pear conversion funnel:**

```
LinkedIn impression → Profile visit → pearcode.dev visit → Waitlist signup
→ Email open → Install → Wizard completion → First question asked
→ First teaching response (aha moment) → Second session → Day 7 retention
→ Free-to-paid conversion
```

**Conversion rate targets at each stage:**

| Stage | Target Conversion | Volume (from 50K impressions) |
|-------|------------------|-------------------------------|
| Impression → Profile visit | 3-5% | 1,500-2,500 |
| Profile visit → Site visit | 15-25% | 225-625 |
| Site visit → Waitlist signup | 25-35% | 56-220 |
| Waitlist → Install (launch day) | 30-40% | 17-88 |
| Install → First question | 70-80% | 12-70 |
| First question → Second session | 40%+ (target) | 5-28 |
| Second session → Paid (month 1) | 5-15% | 0.3-4 |

**What this model tells you:** even with a breakout post (50K impressions), you're looking at 50-200 waitlist signups and 15-70 installs. You need multiple high-performing posts or sustained content output to hit 300+ waitlist by week 4. The guide's targets are achievable but tight. The funnel model makes this visible.

**Attribution infrastructure (set up in week 1, 1-2 hours):**

1. **UTM parameters on every link.** Every LinkedIn post, X tweet, Reddit post, PH listing, HN post, Dev.to article, and email should use distinct UTM tags:
   - `pearcode.dev?utm_source=linkedin&utm_medium=social&utm_campaign=thesis_week1&utm_content=recruiter_post`
   - `pearcode.dev?utm_source=producthunt&utm_medium=referral&utm_campaign=launch_day`
   - `pearcode.dev?utm_source=reddit&utm_medium=social&utm_campaign=launch_day&utm_content=r_golang`

2. **Waitlist signup source tracking.** The Supabase/Upstash waitlist table should record: referral source (from UTM), timestamp, and which page they signed up from (homepage vs. blog post vs. direct link).

3. **Post-install attribution.** When a user runs `pear` for the first time, the setup wizard should include "How did you hear about pear?" with options (LinkedIn, Twitter/X, Reddit, Product Hunt, Hacker News, GitHub, friend/colleague, search, other). Low-friction, one question, high-signal.

4. **Connect content analytics to funnel.** The content engine's "Learn" phase should cross-reference post performance against waitlist signups that occurred within 48 hours, using UTM data. "This post got 100 likes and drove 12 signups at a 2.4% conversion rate" is 10x more useful than "this post got 100 likes."

**Without attribution:** you optimise for likes and impressions. **With attribution:** you optimise for signups and installs. These are different posts.

**Landing page conversion optimisation for pearcode.dev:**

pearcode.dev is the single point through which ALL content traffic must convert. If it converts at 15% instead of 30%, you need 2x the content effort to hit the same waitlist numbers. Treat the landing page as a conversion funnel, not a marketing site.

- **Above the fold:** positioning line, atomic demo screenshot/GIF, single CTA (join waitlist → changes to "install now" at launch). Nothing else. No navigation, no features list, no pricing — just the hook and the action.
- **Social proof section (added incrementally):** beta user screenshots (added week 5), install count badge (added launch day), testimonials with real names and roles.
- **Conversion rate target:** 25-35% of visitors should sign up for the waitlist. If it's below 20%, fix the page before investing more in content.
- **Tracking:** Plausible or Fathom analytics (privacy-respecting, developers prefer these over Google Analytics), plus heatmap tool (Hotjar free tier) during launch week to see where visitors drop off.
- **A/B testing:** test the positioning line. "The only learning tool that teaches while you work" vs. "AI makes you fast. Pear makes you good." vs. "Stop copying code you don't understand." Simple A/B tests on the headline can move conversion 20-50%.

**K-factor (virality coefficient) framework:**

The distribution mechanics listed in this guide each have a viral potential. Estimating the k-factor tells you whether growth depends entirely on your content output (k < 0.3) or is near-self-sustaining (k > 0.7).

| Mechanic | Estimated k-factor | Assumptions |
|----------|-------------------|-------------|
| Screenshottable responses | ~0.5 | 5% of users screenshot, reaches 500 people, 2% of viewers install |
| `/share` command | ~0.3 | 10% of sessions use `/share`, reaches 300 people, 1% convert |
| MCP exposure (pair programming, screen shares) | ~0.1-0.2 | Hard to estimate; organic exposure to colleagues |
| `pear progress` sharing | ~0.1 | Duolingo-style streak sharing, lower than screenshots |

**Combined estimated k: 0.7-1.1.** If this holds, growth is near-self-sustaining and content becomes a booster rather than the engine. But these are estimates — instrument k-factor measurement from month 1 by tracking referral source on every install.

### Onboarding, Activation, and Retention

The guide was heavily acquisition-focused. This section covers what happens after a developer installs pear — the part that determines whether they come back.

**Time-to-first-value target: under 90 seconds.**

From `pear` being run → setup wizard (30 seconds) → first question asked → first teaching response received. If this is more than 2 minutes, it's too long. The wizard must be fast and the first response must be the aha moment.

**The aha moment:** the first time pear teaches something the developer didn't know, grounded in their actual code. This must happen in session 1, ideally within the first 3 interactions. If the first response is generic or unhelpful, the user won't come back. This is why context injection is the launch feature, not voice.

**Post-install onboarding email sequence:**

After a user installs pear (tracked via opt-in telemetry or manual trigger), send:

- **Day 0 — "Welcome + first tip":** "You just installed pear. Here's something most people don't try in their first session: ask pear about code you wrote last week that you're not sure about. The context injection makes the explanation specific to your codebase." Short, actionable, shows a feature they might miss.

- **Day 2 — "Something cool you might not have tried":** "Did you know pear tracks concepts you've learned? Run `pear progress` to see your learning profile. Here's what it looks like after a few sessions: [screenshot]." Introduces the retention mechanic early.

- **Day 5 — "Your first week with pear":** "You've been using pear for 5 days. Here's your learning profile so far: [personalized if possible, or generic example]. The developers who get the most from pear ask about code they've just written or just reviewed — it's the difference between learning in theory and learning in context." Reinforces the habit, shows progress.

- **Day 14 — "Upgrade prompt" (end of trial):** "Your 14-day trial is ending. You've asked [X] questions and learned [Y] concepts. Pro unlocks unlimited questions, concept tracking, all modes, and voice at $30/mo. [Upgrade link]." Timed to when the trial is ending.

**Retention mechanics (built into the product):**

These turn one-time installers into daily users:

- **Streak notifications:** when a user opens their terminal, pear can display a subtle one-liner: "Day 5 streak. 12 concepts this week." This is a `pear` startup message, not a push notification. Developers see it every time they use pear without it being intrusive.

- **Weekly learning summaries:** "This week you learned 3 new concepts: goroutines, connection pooling, error wrapping. Your streak: 5 days." Emailed weekly (opt-in). Each summary is shareable — "my pear learning summary this week" is organic content for the user's network.

- **Progress milestones:** "You've learned 25 concepts. You're in the top 10% of pear users this month." Gamification that's tasteful for developers. Milestone notifications appear in the CLI, not as emails.

- **Re-engagement triggers:** if a user hasn't opened pear in 7 days, trigger an email: "Your last session covered goroutines. Ready to go deeper? Here's a question to try: 'pear, explain how Go channels differ from goroutines in my project.'" Gives them a specific action, not a generic "come back."

- **Concept decay nudges (v2):** if a concept was taught 30+ days ago and never re-encountered, surface it subtly: "It's been a while since you worked with connection pooling. Want a refresher?" This is spaced-repetition for developers, built into the workflow.

---

## Part 5: The Agentic Content Engine

### Why Build This Instead of Using Off-the-Shelf Tools

Tools like Taplio, Buffer, and Typefully handle scheduling. But they don't understand your content pillars, your voice, or which topics are resonating. You end up staring at a blank text box every Sunday. The agentic content engine replaces blank-page creation with an opinionated draft queue that you curate and refine.

The additional meta-benefit: building this system is itself a content pillar. "I built an AI agent that drafts my LinkedIn posts so I can focus on authenticity" is exactly the kind of post that resonates in developer and indie hacker communities.

### How It Works

The system operates on a weekly cycle with four stages: Generate, Review, Publish, Learn.

```
┌──────────────────────────────────────────────────────────┐
│                    1. GENERATE (automated)                │
│                                                           │
│  Content Pillars + Voice Guide + Performance Data         │
│           ↓                                               │
│  LLM generates 7-10 draft posts for the week              │
│  - Weighted to pillar ratios (35/25/25/15)                │
│  - Each draft has LinkedIn version + X version            │
│  - Pillar tag, suggested schedule, hook variants          │
│                                                           │
│  Optional inputs:                                         │
│  - Topic seeds you drop in (Slack msg, quick note)        │
│  - Trending topics in AI/dev space (web scraping agent)   │
│  - Pear usage data once product is live                   │
└──────────────────────────┬───────────────────────────────┘
                           │
                           ▼
┌──────────────────────────────────────────────────────────┐
│               2. REVIEW (human-in-the-loop)               │
│                                                           │
│  Sunday session: you open the review queue                │
│  - Each draft shows pillar, platform, suggested time      │
│  - Approve as-is, edit and approve, or reject             │
│  - Add personal details, sharpen takes, inject voice      │
│  - Reorder or swap scheduling slots                       │
│                                                           │
│  Interface: Slack-based OR lightweight Next.js dashboard   │
│  Target: 90 minutes to review and approve 5-7 posts       │
└──────────────────────────┬───────────────────────────────┘
                           │
                           ▼
┌──────────────────────────────────────────────────────────┐
│               3. PUBLISH (agentic)                        │
│                                                           │
│  Approved posts enter the scheduling queue                │
│  - LinkedIn: posted via API at midday AEST (Tue-Thu)     │
│    You're live for first-hour engagement                  │
│  - X/Twitter: platform-adapted version posted             │
│    via API at 8-10 AM ET for US dev audience              │
│  - Cross-platform: high-performing LinkedIn posts         │
│    auto-queued as X threads (after threshold met)         │
│                                                           │
│  APIs: LinkedIn API (via Arcade), X API (via Arcade),     │
│  or routed through Taplio/Typefully APIs                  │
└──────────────────────────┬───────────────────────────────┘
                           │
                           ▼
┌──────────────────────────────────────────────────────────┐
│               4. LEARN (feedback loop)                    │
│                                                           │
│  After 48 hours, the analytics agent collects:            │
│  - Impressions, engagement rate, comments, profile views  │
│  - Performance by pillar, by topic, by hook style         │
│  - Which posts you edited heavily vs. approved as-is      │
│  - Which posts you rejected entirely                      │
│                                                           │
│  This data feeds back into the generation step:           │
│  - Pillar ratios adjust based on what resonates           │
│  - Hook styles that outperform get reinforced             │
│  - Topics that underperform get deprioritised             │
│  - Your editing patterns teach the system your voice      │
│                                                           │
│  Weekly summary: Slack notification with top performer,   │
│  pillar breakdown, and suggested ratio adjustments        │
└──────────────────────────────────────────────────────────┘
```

### The Pillar Configuration

The generation agent needs a structured definition of each pillar. This is the core input:

```yaml
pillars:
  - name: "The Self-Taught Builder"
    weight: 0.35
    description: >
      Your path from science to recruitment to self-taught engineer to
      product leader. Authentic stories about learning to code in the AI
      era, shipping production systems, and bridging the gap between
      non-traditional background and technical execution.
    tone: "honest, reflective, specific. Never preachy."
    example_hooks:
      - "I taught myself to code in 2022. Here's what nobody tells you."
      - "From recruiting engineers to being one."
      - "I shipped a production platform in 5 months as a solo self-taught engineer."
    avoid: "humble-bragging, generic inspiration, anything that sounds like a LinkedIn influencer"

  - name: "The AI Learning Crisis"
    weight: 0.25
    description: >
      The pear thesis. AI is making developers fast but not good.
      Grounded in your triple perspective: recruiting developers,
      being a developer learning with AI, and leading developers.
    tone: "data-informed, opinionated but not rage bait. Provocative enough to generate discussion."
    example_hooks:
      - "I used to assess technical capability for a living. Here's the gap nobody's talking about."
      - "The data shows AI-assisted developers take 19% longer on real tasks."
      - "Vibe coding is a skill accelerator with a hidden cost."
    avoid: "doomerism, anti-AI takes, anything that sounds like a boomer yelling at clouds"

  - name: "Building Pear in Public"
    weight: 0.25
    description: >
      Product progress, technical decisions, honest metrics. The
      meta-narrative: learning Go with AI while building a tool that
      teaches people to learn better with AI.
    tone: "transparent, technical but accessible. Show the work."
    example_hooks:
      - "Today I shipped concept tracking. Here's what it does."
      - "Why I chose Go for a CLI tool as someone who primarily knows TypeScript."
      - "Week 3 metrics: here's what's working and what's broken."
    avoid: "hype, vaporware energy, premature celebration"

  - name: "Product Engineering & 0-to-1"
    weight: 0.15
    description: >
      Lessons from building 0-to-1 four times. Real stories from
      IwiComply, Uptio, RocketTags, freelance projects. Practical
      product engineering for early-stage teams.
    tone: "practical, experienced, specific examples over general advice."
    example_hooks:
      - "How I scope an MVP in a weekend."
      - "Player-coach leadership: when to code vs. when to delegate."
      - "I built an interactive demo in one day that replaced weeks of engineering."
    avoid: "generic startup advice, anything without a concrete example"
```

### The Voice Guide (Based on Top-Performing Post Analysis)

Analysis of your highest-performing LinkedIn posts reveals clear patterns the generation agent must encode. Your top posts (95-97 likes, 12-22 comments) are personal/identity posts. Your professional insight posts get 37-41 likes with 1-2 comments. The audience engages with Mitch the person first, Mitch the thought leader second.

**What works (from your actual top performers):**

- **Personal narrative first, insight embedded.** Your New Year post (97 likes, 22 comments) and freelancing post (95 likes, 12 comments) dramatically outperform your professional insight posts. The thesis content must be delivered through story, not through statement. Not "AI is making developers fast but not good" as an opening — instead "Last week I watched a junior dev ship a feature in 20 minutes with Cursor. When I asked how it worked, silence."
- **Warm, direct, emoji-natural.** You use emojis organically, not performatively. Short sentences. First person throughout. This is your actual voice — don't let the generation agent polish it into generic LinkedIn thought-leader tone.
- **Values-driven framing.** Your top posts lead with who you are (empathy, curiosity, intuition) not what you know. The freelancing post got 95 likes because it was about your character, not your skills.
- **Casual confidence, not authority posturing.** "I'm only 26, I don't have decades of domain experience" — this honesty is magnetic. Never let the agent produce posts that sound like someone with 20 years of experience pontificating.
- **Community-oriented endings.** Posts that invite responses ("don't hesitate to reach out," "look forward to many coffee conversations") outperform posts that conclude with statements.

**What underperforms (also from your posts):**

- **Consensus professional takes.** "The first job of a product leader is to listen" (37 likes, 2 comments). Nobody disagrees, so nobody comments. Every thesis post needs a debatable edge.
- **Statement-ending posts.** "Let's cook" closes a conversation. "What's your experience with this?" opens one. Comments are the algorithm multiplier — your professional posts lack the invitation.
- **Observation without personal stake.** "AI has completely shifted the paradigm" is an observation. "I can now deliver production code that would have taken me weeks two years ago" is personal stake. The second gets engagement.

**Voice characteristics to encode:**

- First person, direct, specific. Opens with a concrete moment or honest admission.
- Short paragraphs. One idea per paragraph. LinkedIn formatting (line breaks between paragraphs).
- Emojis used naturally at structural points, not as decoration.
- Technical enough to be credible, warm enough to be approachable.
- Ends with an open question, an invitation, or a forward-looking statement — never a conclusion that closes discussion.
- The overall register is "smart friend telling you about their week over coffee" not "thought leader dispensing wisdom from a stage."

**Few-shot examples for the agent (from your actual high performers):**

```
EXAMPLE 1 (Personal/aspirational — 97 likes, 22 comments):
"Unsure if sharing the dogs on LinkedIn is kosher so cropped to be safe🦶
After a restful end to 2025...
[personal context, casual tone, priorities list, forward-looking energy]
Hope you all had a brilliant break and look forward to many coffee
conversations and new friends to come ☕️"

EXAMPLE 2 (Values-driven entrepreneurial — 95 likes, 12 comments):
"Lots of interest in my 0 to 1 software product freelancing so far and
I'm beyond excited about some of the ideas I've heard so far💡
I'm only 26, I don't have decades of domain experience, but what I do
pride myself on is empathy, curiosity and intuition.
[values explained with emojis, personal and honest, CTA at end]"
```

The generation agent should produce content that reads like these examples — personal, warm, specific — but applied to the four content pillars. The thesis content should feel like you're telling a friend about something you noticed, not writing a position paper.

### Technical Architecture

**Option A: LangGraph-based (recommended if you want full control)**

Fork or reference the LangChain social-media-agent repo (2.2K stars, TypeScript, LangGraph). Strip out the URL-ingestion pipeline and replace with pillar-driven generation. Keep the human-in-the-loop infrastructure (Agent Inbox or Slack-based review) and the Arcade-based social media auth for posting.

Stack:
- LangGraph for agent orchestration and state management
- Anthropic Claude for content generation
- Supabase for post storage, analytics, and performance tracking
- Arcade for LinkedIn and X API authentication and posting
- Slack for review notifications and quick approvals
- LangSmith for tracing and debugging the generation quality

**Option B: n8n-based (recommended if you want speed to deploy)**

Build the pipeline in n8n, which you already use. Use n8n's AI agent nodes with Claude or GPT-4 for generation. Store drafts in Supabase or Airtable. Use Slack for review workflow. Post via LinkedIn and X API nodes or via Taplio/Typefully APIs.

Stack:
- n8n for workflow orchestration
- Claude API or OpenAI for generation
- Supabase or Airtable for draft storage and performance data
- Slack for human-in-the-loop review
- LinkedIn API + X API (or Taplio + Typefully) for publishing
- Cron triggers for weekly generation and post-publish analytics collection

**Option C: Hybrid (best of both)**

Use n8n for the orchestration layer (scheduling, API calls, notifications) and LangGraph for the generation agent (pillar weighting, voice matching, performance-informed drafting). This gives you n8n's visual workflow builder for the plumbing and LangGraph's sophisticated agent capabilities for the intelligent parts.

### Build Approach: Async with Claude Code

The content engine is being built asynchronously using Claude Code and agentic agents running in the background — it does not compete with pear development time or content creation time. This means the engine can be developed in parallel with the 4-6 week launch sprint without any trade-off.

**Build phases (running in background):**

**Phase 1: Manual-with-structure (Week 1, 2 hours of setup)**

Before the engine is ready, validate the pillar framework manually. Create a Supabase table or Notion database with columns for: pillar, platform, draft text, status (draft/approved/posted), post date, impressions, engagement rate. Manually write posts using the phase-specific pillar weights and track everything. This gives the agent training data and validates the framework.

**Phase 2: Generation agent (Weeks 1-3, built async via Claude Code)**

Build the core generation pipeline using Claude Code agents running in background sessions. Input: pillar config + voice examples (from your top-performing posts) + performance data. Output: 7 to 10 draft posts with LinkedIn and X variants. Review via Slack. Target: engine producing usable drafts by week 3, in time for the reveal and launch phases. Even if only the generation piece is ready, having an AI draft queue that understands your pillars saves 60% of your Sunday batch time.

**Phase 3: Automated publishing (Week 4+)**

Wire up the posting APIs. Approved posts get scheduled and published automatically. This is the plumbing. Use Arcade for social media auth or go through Taplio/Typefully APIs if you want their scheduling intelligence. Can be completed during or after launch.

**Phase 4: Feedback loop (Month 2)**

Build the analytics collection agent. Pull engagement data 48 hours after each post. Score by pillar, topic, and hook style. Feed scores back into the generation prompt. Track your editing patterns (what you change, what you reject) and use that to refine the voice model. This phase is most valuable once you have 4-6 weeks of post data to learn from.

### Cost Breakdown

| Component | Tool | Cost |
|-----------|------|------|
| Generation | Claude API (Anthropic) | ~$5 to 10/mo at 40 posts/mo |
| Orchestration | n8n (self-hosted) or LangGraph | Free (self-hosted) |
| Storage | Supabase | Free tier sufficient |
| Social auth | Arcade | Free tier for personal use |
| Publishing | Direct API or Taplio + Typefully | $0 (direct) or $50 to 80/mo (tools) |
| Video clipping | Opus Clip or Vizard (when YouTube starts) | $15 to 30/mo |
| Tracing | LangSmith | Free tier |

**Total: $5 to $120/month** depending on whether you use direct APIs or scheduling tools. The agent itself costs almost nothing to run.

### Legacy Automation Workflows (Still Useful)

These workflows from the original plan still apply and complement the agentic engine:

**Workflow 1: Stream to Multi-Platform Content**

```
Stream on YouTube Live (2 hours)
    ↓
Recording auto-saved to YouTube
    ↓
Upload to Opus Clip → auto-generates 10-15 short clips
    ↓
Review clips (10 min) → select best 5
    ↓
Schedule to YouTube Shorts, X, LinkedIn
    ↓
5 days of content from 1 session
```

**Workflow 2: High-Performing Post Amplification**

```
LinkedIn post exceeds engagement threshold (tracked by analytics agent)
    ↓
System auto-generates X thread variant (punchier, shorter)
    ↓
Queued for your next review session
    ↓
Approved version scheduled via X API
```

**Workflow 3: User Feedback to Content (post-launch)**

```
Beta user sends screenshot of pear response
    ↓
Ask permission to share
    ↓
Post to LinkedIn: "One of our beta users asked pear about [concept]. Here's what happened."
    ↓
Same screenshot to X with shorter caption
    ↓
Add to "wall of responses" on pearcode.dev
```

**Workflow 4: Usage Data to Content (post-launch)**

```
Pear concept tracking shows "goroutines" asked 47 times this week
    ↓
Analytics agent flags trending concept
    ↓
Generation agent creates pillar-appropriate draft
    ↓
Queued for review → approved → published
    ↓
Expand into YouTube explainer if engagement warrants
```

The content flywheel: the product generates the data that generates the content that drives users to the product. The agentic engine makes this flywheel spin with minimal manual effort.

---

## Part 6: Metrics and Milestones

### Pre-Launch Sprint (Weeks 1-5)

| Milestone | Target | Why |
|-----------|--------|-----|
| Week 2: waitlist signups | 100+ | Thesis content is driving interest |
| Week 2: LinkedIn followers gained | +200 | Content is reaching new audiences |
| Week 2: breakout post | At least 1 post above 100 likes | Escaped existing audience |
| Week 4: waitlist signups | 300+ | Reveal is converting |
| Week 4: beta list | 15-20 confirmed testers | Personal outreach is working |
| Week 5: beta screenshots | 3-5 shareable user screenshots | Social proof for launch day |

### Launch Day and Week 1

| Metric | Target | Why |
|--------|--------|-----|
| Day 1 installs | 50 to 100 | PH + LinkedIn + email + Reddit drove action |
| Week 1 total installs | 200 to 500 | Multi-channel launch generated sustained interest |
| PH finish position | Top 10 of the day | Listing was well-prepared |
| Show HN upvotes | 50+ | HN audience found it relevant |
| GitHub stars | 100+ in week 1 | Hits trending lists, drives organic discovery |
| Second-session rate | 40%+ | The only metric that matters for product-market fit |
| Waitlist to install conversion | 30%+ | Nurture sequence and launch messaging worked |
| Landing page conversion rate | 25-35% of visitors sign up | Page is optimised for conversion |

### Month 1 Post-Launch

| Metric | Target | Why |
|--------|--------|-----|
| Total installs | 300 to 700 | Sustained growth beyond launch spike |
| LinkedIn followers | +500 from start | Content is compounding |
| Paying users | 10 to 25 | Early signal on willingness to pay |
| Content engagement rate | 5%+ on LinkedIn | Audience is real and engaged |
| Profile views per week | 500+ | Content is driving curiosity |
| GitHub stars | 200+ | Sustained organic discovery |
| Blog posts live on pearcode.dev | 5+ | SEO compounding has started |
| Onboarding email open rate | 50%+ | Users are engaged post-install |

### Month 3

| Metric | Target | Why |
|--------|--------|-----|
| Total users | 1,000 to 2,500 | Sustainable growth |
| MRR | $500 to $1,600 | Revenue exists |
| MCP installs | 100+ | Distribution channel is working |
| YouTube subscribers | 500+ | Long-term channel growing (started month 2) |
| Content engine approval rate | 60%+ of drafts approved | Engine is learning your voice |
| GitHub stars | 500+ | Organic discovery is compounding |
| Monthly organic search visitors | 500+ | SEO is starting to work |
| Day 7 retention rate | 30%+ | Onboarding and retention mechanics are working |
| K-factor (measured) | 0.3+ | Product-led growth loops are active |

### Month 6

| Metric | Target | Why |
|--------|--------|-----|
| Total users | 5,000 to 10,000 | Approaching inflection |
| MRR | $3,000 to $10,000 | Approaching sustainability |
| Organic acquisition | 50%+ of new users | Product-led growth is working |
| Share command usage | 10%+ of sessions | Viral loop is active |
| YouTube views/month | 10,000+ | Compounding has started |
| Monthly organic search visitors | 3,000+ | SEO is a real channel |
| GitHub stars | 1,000+ | Social proof for enterprise conversations |
| K-factor (measured) | 0.5+ | Growth is near-self-sustaining |

---

## Part 7: The One-Page Summary

**Positioning:** "The only learning tool that teaches while you work."

**Primary channel:** LinkedIn (~4.3K followers, dev-heavy professional audience, posts at midday AEST)

**Secondary channels:** X/Twitter (US timezone, 8-10 AM ET), Reddit (US timezone, launch day + ongoing), Dev.to/Hashnode (SEO cross-posts), YouTube (deferred to post-launch)

**Content pillars (tailored to your story, ratios shift by phase):**
1. The Self-Taught Builder (20-35%) — your unusual path, real production stories
2. The AI Learning Crisis (25-50%) — the pear thesis, front-loaded pre-launch
3. Building Pear in Public (10-40%) — ramps up at reveal, peaks at launch
4. Product Engineering & 0-to-1 (10-20%) — lessons from building four products from scratch

**Voice (from top-performing post analysis):** personal narrative first, insight embedded. Warm, direct, emoji-natural. Values-driven. Ends with open questions. "Smart friend over coffee" not "thought leader from a stage." Personal/identity posts get 2-3x the engagement of professional insight posts.

**Timezone strategy:** LinkedIn at midday AEST (you're live for the full first-hour engagement window). X, Reddit, and Dev.to scheduled for 8-10 AM US Eastern (peak US developer activity). PH launches at 7 PM AEST (midnight Pacific reset). Show HN at 8-10 AM ET.

**The 4-6 week launch sprint:**
- Weeks 1-2: thesis establishment (50% AI Learning Crisis). Engineer 2 breakout posts. PH preparation begins. LinkedIn headline updated. UTM attribution set up. Blog posts start on pearcode.dev.
- Weeks 3-4: the reveal (40% Build in Public). Atomic demo moment. Beta list recruited via DM. LinkedIn headline updated again. Dev.to/Hashnode cross-posts.
- Week 5: compressed beta (15-20 users, 1 week, collect screenshots and testimonials). Landing page updated with social proof.
- Week 6: staggered launch. Day 0: PH (7 PM AEST) + LinkedIn (midday AEST) + email + Reddit + X (8 AM ET). Day +1/+2: Show HN with Day 1 data.

**Critical launch assets:**
- The atomic demo moment (one perfect terminal screenshot)
- Strategic amplification list (5-10 accounts with 15K+ followers)
- Product Hunt listing (prepared weeks in advance, 20-30 upvoters recruited)
- Show HN post (technical, honest, links to open-source GitHub repo, posted Day +1/+2 with real data)
- Email nurture sequence (Day -7, -3, -1, 0, +1)
- LinkedIn headline optimised for each phase
- Substack essay as long-form landing page (shared via LinkedIn/X, not standalone distribution)
- GitHub repo with optimised README (demo GIF, install command, star request)
- UTM parameters on every link, source tracking on waitlist signups
- Post-install onboarding emails (Day 0, 2, 5, 14)

**Open source strategy (decision pending):** teaching engine (prompts, context collection, concept tracking, prompt assembly) is definitively proprietary. Scope of what goes open source is TBD — candidates include CLI shell, MCP transport, config, LLM adapters, build tooling. Homebrew tap requires at minimum a public tap repo. GitHub stars drive discovery. HN rewards open source. Whatever goes public, fork risk is contained: the teaching engine IS the product.

**Automation stack:** Agentic Content Engine built async via Claude Code (LangGraph or n8n + Claude API + Supabase). Generation agent target ready by week 3. Publishing automation wired by week 4+. Feedback loop month 2. $5 to $120/month.

**Distribution mechanics (8 channels):**
1. Screenshottable responses (output virality)
2. Shareable progress (`pear progress` streaks and concept counts)
3. `/share` command (engineered sharing)
4. MCP inside Cursor/Claude Code (embedded distribution)
5. Content flywheel (usage data → educational content → new users)
6. Agentic Content Engine (pillar-driven generation → human review → automated publishing → performance learning loop)
7. GitHub (open-source repo, stars, trending, README as conversion page)
8. SEO (pearcode.dev/blog + Dev.to + Hashnode, compound organic traffic)

**Growth engineering infrastructure:**
- Funnel model with conversion targets at each stage (impression → profile → site → waitlist → install → aha → retention → paid)
- UTM attribution connecting every post to downstream signups/installs
- Landing page conversion optimisation (target 25-35% visitor-to-signup)
- K-factor measurement to quantify viral growth loops
- Post-install onboarding sequence (Day 0, 2, 5, 14) pushing second-session rate above 40%
- Retention mechanics: streak display, weekly learning summaries, progress milestones, re-engagement triggers

**Developer communities:**
- Reddit: r/programming, r/golang, r/learnprogramming, r/ExperiencedDevs, r/cursor (launch day, then value-first ongoing)
- Discord: Cursor, Claude, dev community servers (credibility + early-adopter channel)
- Slack: Indie Hackers, Gophers, Rands Leadership (niche communities, ongoing)
- Dev.to + Hashnode: SEO-optimised cross-posts with canonical URLs to pearcode.dev

**The non-negotiable daily habit:** 15 minutes responding to every comment within 60 minutes of each LinkedIn post going live at midday AEST. You're at your desk for the entire critical first hour. This is the single highest-leverage activity on LinkedIn and it cannot be automated.

**The one rule:** every piece of content should either teach something valuable or show pear teaching something valuable. If it does neither, don't post it.
