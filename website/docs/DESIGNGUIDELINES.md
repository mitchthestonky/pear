# Pear — Design Guidelines

> Visual design system for pearcode.dev. Every design decision traces back to this document.
> Aesthetic references: Raycast.com (developer-native, dark-first, accent pops), Linear.app (clean, precise, subtle animation).

---

## Design Principles

1. **Developer-native.** This site is for people who live in terminals. The design should feel like home — monospace type, dark mode default, visible code, no stock photos.
2. **Minimal but not empty.** Every element earns its place. White space is deliberate, not lazy. Animations enhance comprehension, not decoration.
3. **Confident.** The design matches the copy tone — direct, precise, no hedging. Bold type. Clear hierarchy. Strong contrast.
4. **Fast.** No heavy images, no video autoplay above the fold, no layout shift. The site should feel as fast as a terminal.

---

## Color System

### Light Mode (default for first visit — system preference respected)

| Token | Value | Usage |
|---|---|---|
| `--bg` | `#FFFFFF` | Page background |
| `--bg-subtle` | `#FAFAFA` | Card backgrounds, alternating sections |
| `--fg` | `#0A0A0A` | Primary text |
| `--fg-muted` | `#6B7280` | Secondary text, captions |
| `--border` | `#E5E7EB` | Card borders, dividers |
| `--accent` | `#22C55E` | CTAs, highlights, brand (green-500) |
| `--accent-hover` | `#16A34A` | CTA hover state (green-600) |
| `--accent-subtle` | `#DCFCE7` | Accent backgrounds, badges (green-100) |
| `--dot` | `#0A0A0A` opacity 0.06 | Background floating dots |

### Dark Mode

| Token | Value | Usage |
|---|---|---|
| `--bg` | `#0A0A0A` | Page background |
| `--bg-subtle` | `#111111` | Card backgrounds, alternating sections |
| `--fg` | `#FAFAFA` | Primary text |
| `--fg-muted` | `#9CA3AF` | Secondary text, captions |
| `--border` | `#1F2937` | Card borders, dividers |
| `--accent` | `#22C55E` | CTAs, highlights, brand (same green) |
| `--accent-hover` | `#4ADE80` | CTA hover state (green-400, lighter in dark) |
| `--accent-subtle` | `#052E16` | Accent backgrounds (green-950) |
| `--dot` | `#FFFFFF` opacity 0.06 | Background floating dots |

### Accent Usage Rules

- Green is used sparingly: CTAs, the Pear logo mark, active states, and key highlights.
- Body text is never green. Green draws the eye — reserve it for action and emphasis.
- The green should feel *natural and confident*, not neon. `#22C55E` (Tailwind green-500) hits this balance.

---

## Typography

### Font Stack

| Role | Font | Fallback | Weight |
|---|---|---|---|
| **Headings** | JetBrains Mono | `monospace` | 700 (bold) |
| **Body** | Inter | `system-ui, sans-serif` | 400 (regular), 500 (medium) |
| **Code / Terminal** | JetBrains Mono | `monospace` | 400 |
| **CTA buttons** | Inter | `system-ui, sans-serif` | 600 (semibold) |

### Type Scale

| Element | Size | Line Height | Weight | Font |
|---|---|---|---|---|
| Hero tagline | `text-5xl` (48px) / `text-6xl` (60px) on desktop | 1.1 | 700 | JetBrains Mono |
| Hero subheadline | `text-xl` (20px) | 1.5 | 400 | Inter |
| Section headers | `text-3xl` (30px) / `text-4xl` (36px) on desktop | 1.2 | 700 | JetBrains Mono |
| Section body | `text-lg` (18px) | 1.7 | 400 | Inter |
| Card titles | `text-xl` (20px) | 1.3 | 600 | Inter |
| Card body | `text-base` (16px) | 1.6 | 400 | Inter |
| Code blocks | `text-sm` (14px) | 1.6 | 400 | JetBrains Mono |
| CTA button text | `text-base` (16px) | 1 | 600 | Inter |
| Micro-copy | `text-sm` (14px) | 1.5 | 400 | Inter, muted color |

### Monospace Headings

All section headers use JetBrains Mono. This is the single strongest visual signal that this product is for developers. Body text uses Inter for readability — monospace body text is fatiguing at length.

---

## Background: Floating Dots

### Concept

The background features a uniform grid of small circular dots that float with subtle, slow movement. This creates depth without distraction — a living texture that makes the page feel dynamic without competing with content.

### Specification

| Property | Value |
|---|---|
| Dot size | 2px diameter (circles) |
| Grid spacing | 40px uniform grid |
| Dot color (light mode) | `#0A0A0A` at 6% opacity |
| Dot color (dark mode) | `#FFFFFF` at 6% opacity |
| Movement | Slow sinusoidal drift: ±3px horizontal, ±3px vertical |
| Animation speed | 8-12 second cycle (randomized per dot for organic feel) |
| Layer | Behind all content (`z-index: 0`) |
| Performance | Canvas-based or CSS-only for dots visible in viewport. No heavy JS libraries. |

### Parallax Behavior

The dot background scrolls at **50% of the foreground scroll speed** (`transform: translateY(scrollY * 0.5)`). This creates a subtle depth effect — the content moves over the dots, giving the page a layered, spatial feel.

Implementation: CSS `background-attachment: fixed` or JS-driven `translateY` on the canvas container. Prefer CSS for performance; fall back to rAF-driven JS if CSS doesn't achieve the desired smoothness.

---

## Animations

### Typewriter (Hero Tagline)

The tagline types out character by character on page load.

| Property | Value |
|---|---|
| Characters per step | 1 |
| Speed | 50ms per character |
| Cursor | Blinking block cursor (`▊`), green (`--accent`), blinks at 530ms interval |
| Cursor behavior | Blinks during typing, stays solid for 500ms after completion, then blinks indefinitely |
| Trigger | On page load (no scroll trigger) |

After the tagline finishes typing (~2.5s), the subheadline fades in over 400ms.

### Section Reveals

Sections fade in and slide up slightly as they enter the viewport.

| Property | Value |
|---|---|
| Trigger | Intersection Observer, 15% visible |
| Animation | `opacity: 0 → 1`, `translateY(20px) → 0` |
| Duration | 500ms |
| Easing | `cubic-bezier(0.16, 1, 0.3, 1)` (ease-out-expo) |
| Stagger | If section has multiple cards, stagger each by 100ms |

### Terminal Mockup (How It Works)

The "How it works" section contains a terminal mockup that animates step by step:

1. Prompt appears: `$ pear`
2. Status line types: `🍐 Pear active. Hold [Space] to talk.`
3. Voice indicator: `🎤 Recording...` (pulses for 1.5s)
4. Transcript appears: `"Walk me through what changed in this file."`
5. Context injection: `📎 Context: git diff (3 files), src/auth/rbac.go, file tree`
6. Response streams in, line by line (simulated streaming at 30ms/char)

This runs once when the section scrolls into view. Users can replay via a small "replay" button.

---

## Components

### Waitlist Input

```
┌─────────────────────────────────────────────────────────┐
│  [email@example.com          ] [ Join the waitlist → ]  │
└─────────────────────────────────────────────────────────┘
```

| Property | Spec |
|---|---|
| Layout | Inline: input + button on one row (stacks on mobile) |
| Input | `border: 1px solid var(--border)`, `bg: var(--bg-subtle)`, `rounded-lg`, padding 12px 16px |
| Button | `bg: var(--accent)`, `color: white`, `font-weight: 600`, `rounded-lg`, padding 12px 24px |
| Hover | Button: `bg: var(--accent-hover)`. Input: `border-color: var(--accent)` |
| Success state | Input + button replaced with: "✓ You're on the list." (green text, fade transition) |
| Error state | Red border on input, "Please enter a valid email" below in red text |
| Mobile | Input full width, button full width below |

### Pricing Cards

Two cards side by side (stacked on mobile). Pro card has a subtle green border and "Early Access" badge.

| Property | Free Card | Pro Card |
|---|---|---|
| Border | `var(--border)` | `var(--accent)` 1px solid |
| Background | `var(--bg-subtle)` | `var(--bg-subtle)` |
| Badge | None | "Early Access" — `bg: var(--accent-subtle)`, `color: var(--accent)`, `text-xs`, `rounded-full` |
| CTA | "Get started free" (outline button) | "Lock in $20/mo" (filled green button) |
| Price | "$0" large + "forever" small | "~~$30~~ $20/month" with strikethrough on $30 |

### Teaching Comparison (Side-by-Side)

Two columns. Left: muted, plain. Right: highlighted, detailed.

| Property | Left (Assistant) | Right (Pear) |
|---|---|---|
| Header | "A coding assistant says:" | "Pear says:" |
| Header color | `var(--fg-muted)` | `var(--accent)` |
| Border | `var(--border)` | `var(--accent)` |
| Content style | Plain monospace text, short | Monospace text with **bold** keywords, longer, structured |
| Background | `var(--bg-subtle)` | `var(--bg-subtle)` with subtle green left-border (3px) |

### Problem Cards

Three cards in a row (stacked on mobile). Each has:
- A short bold title (JetBrains Mono, `text-lg`, bold)
- Body text (Inter, `text-base`, muted)
- Subtle top-border in `var(--border)` (or accent for emphasis)

---

## Layout

### Grid

- Max content width: `1200px` centered
- Section vertical padding: `96px` (desktop), `64px` (mobile)
- Card gap: `24px`
- Responsive breakpoints: Tailwind defaults (`sm: 640px`, `md: 768px`, `lg: 1024px`, `xl: 1280px`)

### Section Rhythm

Each section alternates between `var(--bg)` and `var(--bg-subtle)` backgrounds. This creates visual separation without heavy dividers.

```
Hero          → --bg
Problem       → --bg-subtle
Solution      → --bg
How it works  → --bg-subtle
Comparison    → --bg
Pricing       → --bg-subtle
Final CTA     → --bg
Footer        → --bg-subtle
```

---

## Dark/Light Mode

### Implementation

- Respect `prefers-color-scheme` on first visit.
- Toggle switch in top-right of navbar (sun/moon icon).
- Store preference in `localStorage`.
- Use CSS custom properties (tokens above) switched via `data-theme="dark"` on `<html>`.
- Transition: all color properties transition over 200ms for smooth switching.

### Key Differences

| Element | Light | Dark |
|---|---|---|
| Background | White (`#FFFFFF`) | Near-black (`#0A0A0A`) |
| Text | Near-black | Near-white |
| Dots | Black at 6% opacity | White at 6% opacity |
| Cards | `#FAFAFA` bg | `#111111` bg |
| Accent | Green-500 (`#22C55E`) | Same green (works on both) |
| Accent hover | Green-600 (darker) | Green-400 (lighter) |
| Code blocks | Light gray bg | Dark gray bg |

---

## Imagery & Media

- **No stock photos.** Ever.
- **No illustrations** unless custom and developer-themed.
- **Terminal mockups** are the primary visual content. Styled as realistic terminal windows (title bar with dots, monospace content, dark bg even in light mode).
- **Demo video placeholder:** A bordered rectangle with a play button and "Watch the demo — 60s" text. Video embeds after launch when the real demo is recorded.
- **Pear logo:** Simple, monochrome. A pear silhouette or the word "pear" in JetBrains Mono bold. Green accent optional. Keep it simple enough to work as a favicon.

---

## Navbar

| Element | Spec |
|---|---|
| Position | Fixed top, blur backdrop (`backdrop-filter: blur(12px)`), `bg: var(--bg)` at 80% opacity |
| Left | Pear logo (text or icon) |
| Center/Right | Links: Pricing · Docs · Blog (grayed out / "coming soon" for launch) |
| Far right | Dark/light toggle (sun/moon) · "Join waitlist" button (small, green, outline) |
| Mobile | Hamburger menu with same links |

---

## Footer

Minimal. One row.

```
© 2026 Pear · Built by Mitch · LinkedIn · Substack · X
```

---

## Tech Stack

| Tool | Purpose |
|---|---|
| Next.js 15 (App Router) | Framework, SSR for SEO |
| TypeScript | Type safety |
| Tailwind CSS v4 | Utility-first styling |
| JetBrains Mono (Google Fonts) | Monospace headings + code |
| Inter (Google Fonts or next/font) | Body text |
| Framer Motion | Typewriter, scroll reveals, section transitions |
| Canvas API or CSS | Floating dot background |
| Resend or Loops | Waitlist email collection (or simple POST to API) |
| Vercel | Hosting, preview deploys |

---

## File Structure (Next.js)

```
website/
├── app/
│   ├── layout.tsx              # root layout, fonts, theme provider
│   ├── page.tsx                # one-pager landing page
│   ├── api/
│   │   └── waitlist/route.ts   # POST endpoint for email capture
│   ├── globals.css             # CSS custom properties, base styles
│   └── fonts/                  # local font files if not using Google Fonts CDN
├── components/
│   ├── Hero.tsx
│   ├── Problem.tsx
│   ├── Solution.tsx
│   ├── HowItWorks.tsx
│   ├── TeachingDifference.tsx
│   ├── Pricing.tsx
│   ├── FinalCTA.tsx
│   ├── Navbar.tsx
│   ├── Footer.tsx
│   ├── WaitlistForm.tsx
│   ├── TypewriterText.tsx
│   ├── TerminalMockup.tsx
│   ├── FloatingDots.tsx        # background dot canvas/CSS
│   ├── ThemeToggle.tsx
│   └── ScrollReveal.tsx        # intersection observer wrapper
├── lib/
│   └── theme.ts                # theme context + localStorage
├── docs/
│   ├── DESIGNGUIDELINES.md     # this file
│   └── POSITIONING.md          # copy & messaging source of truth
├── public/
│   └── favicon.ico
├── tailwind.config.ts
├── next.config.ts
├── package.json
└── tsconfig.json
```
