# Website PRD — pearcode.dev

> Single-page marketing site + waitlist capture for Pear.
> Read POSITIONING.md for exact copy. Read DESIGNGUIDELINES.md for visual specs.
> This doc is the build spec — what to implement and why.

---

## What This Is

A one-page marketing site for Pear — a voice-first AI tutor CLI for software engineers. The site's job: educate the visitor on the problem, position Pear as the solution, capture waitlist emails, and establish SEO foundations. No product is live yet — this is pre-launch.

**URL:** `pearcode.dev`
**Deploy:** Vercel
**Timeline:** Weekend build

---

## Page Sections (top to bottom)

The page follows a progressive education → conversion flow. Each section builds on the last.

### 1. Navbar (fixed)

- Fixed top, blurred backdrop (`backdrop-filter: blur(12px)`), semi-transparent background
- Left: "pear" wordmark in JetBrains Mono bold, with a small 🍐 emoji or green dot
- Right: `Pricing` anchor link · dark/light toggle (sun/moon icon) · "Join waitlist" small green outline button (scrolls to hero CTA)
- Mobile: hamburger menu

### 2. Hero (full viewport height)

- **Tagline:** "AI makes you fast. Pear makes you good." — typewriter animation on load, JetBrains Mono, `text-5xl`/`text-6xl`
- **Subheadline:** "The voice-first AI tutor that teaches software engineering in your terminal." — fades in 400ms after typewriter completes, Inter, `text-xl`
- **Waitlist CTA:** inline email input + "Join the waitlist →" green button
- **Micro-copy:** "Early access. macOS. Free to try with your own API key."
- Green blinking block cursor on typewriter (`▊`, 530ms blink, `--accent` color)

### 3. Problem (3 cards)

- **Section header:** "The AI-assisted learning crisis is real." (JetBrains Mono)
- Three cards, side-by-side on desktop, stacked on mobile:
  1. **"Developers are shipping code they don't understand."** — 40% of junior devs deploy AI code they can't explain. AI optimizes for output, not understanding.
  2. **"Senior engineers can't scale mentoring."** — 10:1 junior-to-senior ratio. AI makes juniors faster without making them better.
  3. **"Learning is disconnected from real work."** — Courses exist in a vacuum. Learning should happen at the point of work.
- Cards fade in with 100ms stagger on scroll

### 4. Solution

- **Section header:** "Pear teaches you while you code." (JetBrains Mono)
- Brief paragraph: Pear is a CLI companion alongside existing AI tools. Voice in, teaching out, grounded in your actual codebase.
- Three capability highlights (icon + title + one-liner):
  - **Voice-first** — Hold space, ask, hear the answer
  - **Context-aware** — Reads your git diff, file tree, error logs automatically
  - **Teaching-first** — Three modes: teach, mentor, pair

### 5. How It Works (terminal mockup)

- **Section header:** "Three steps. Under three seconds." (JetBrains Mono)
- Animated terminal mockup showing a simulated Pear session:
  1. `$ pear` → status line types in
  2. Voice recording indicator (pulses)
  3. Transcript appears
  4. Context injection line shows enrichment
  5. Teaching response streams in line-by-line
- Animation triggers once on scroll into view. Replay button (small, subtle) in corner.
- Terminal styled as dark window even in light mode (realistic terminal aesthetic)

### 6. Teaching Difference (side-by-side comparison)

- **Section header:** "Not just the fix. The lesson." (JetBrains Mono)
- Two columns:
  - **Left:** "A coding assistant says:" — short, plain answer about the `==` vs `slices.Contains()` bug
  - **Right:** "Pear says:" — full teaching response explaining the bug, the Go gotcha, why it matters for security, offering to go deeper
- Right column has green left border (3px) and green header text
- See POSITIONING.md Section 5 for exact copy

### 7. Pricing

- **Section header:** "Simple pricing. Bring your own AI." (JetBrains Mono)
- **Intro line:** "Pear is a teaching tool, not an LLM reseller. Bring your own API key and pay for the intelligence layer."
- Two pricing cards side-by-side:
  - **Free:** $0/forever, BYOK, 10 voice-min/day, text only, teach mode only. Outline CTA: "Get started free"
  - **Pro (Early Access):** ~~$30~~ $20/mo, everything in Free plus unlimited voice, TTS, all modes, 50 hosted requests. Green border, "Early Access" badge. Filled green CTA: "Lock in $20/mo →". Annual: $200/yr (2 months free) noted below price.
- **Teams — Coming Soon** card or row below: "Pear for your engineering team. Reduce senior mentoring overhead." with "Join team waitlist →" link
- See POSITIONING.md Section 6 for full copy

### 8. Final CTA

- **Tagline repeated:** "AI makes you fast. Pear makes you good."
- **Body:** "Join the waitlist for early access. macOS. Free tier available day one."
- Same waitlist email input + button as hero
- **Footer below:** `© 2026 Pear · Built by Mitch · LinkedIn · Substack · X`

---

## Functional Requirements

### Waitlist Email Capture

- POST to `/api/waitlist` route
- Store emails somewhere persistent (options: Vercel KV, Resend audience, Google Sheet via API, or Supabase — keep it simple)
- Validate email client-side and server-side
- Success state: input + button replaced with "✓ You're on the list." (green, fade transition)
- Error state: red border, "Please enter a valid email" message
- Duplicate emails: accept silently (don't reveal if email already exists)
- The same `WaitlistForm` component is used in both hero and final CTA sections

### Dark/Light Mode

- Respect `prefers-color-scheme` on first visit
- Toggle in navbar (sun/moon icon)
- Persist preference in `localStorage`
- CSS custom properties on `<html data-theme="dark|light">`
- 200ms transition on color changes

### SEO

- Meta title: "Pear — AI makes you fast. Pear makes you good."
- Meta description: "Pear is the voice-first AI tutor that teaches software engineering while you code. Talk to your codebase, learn the patterns, understand the why. macOS CLI."
- Open Graph image: simple branded card (can be static PNG for now)
- Canonical URL: `https://pearcode.dev`
- Semantic HTML: proper heading hierarchy (single `h1` in hero, `h2` per section)
- All sections have `id` attributes for anchor linking

### Performance

- No heavy JS libraries beyond Framer Motion
- Fonts: preload JetBrains Mono and Inter via `next/font`
- No autoplay video — demo video is a placeholder with play button
- Lighthouse target: 95+ on Performance, 100 Accessibility
- Floating dots: canvas-based or pure CSS. Must not drop below 60fps on M1 MacBook.

---

## Non-Goals

- No blog, docs, or multi-page content in v1. Just the one-pager.
- No authentication or dashboard. Those come with the product launch.
- No Stripe integration. Pricing section is informational + waitlist.
- No actual product download. Install commands shown are for positioning.
- No analytics beyond basic Vercel Analytics (free tier). PostHog deferred.

---

## Copy Source

All copy lives in `POSITIONING.md`. Do not hardcode copy that deviates from that doc. If a section needs copy, reference POSITIONING.md for the exact words.
