# Website Architecture — pearcode.dev

> Implementation reference for building the site. Pair with PRD.md (what to build), POSITIONING.md (copy), and DESIGNGUIDELINES.md (visuals).

---

## Tech Stack

| Tool | Version | Purpose |
|---|---|---|
| Next.js | 15 (App Router) | Framework. SSR for SEO, RSC for performance. |
| TypeScript | 5.x | Type safety throughout. |
| Tailwind CSS | v4 | Utility-first styling. Use CSS custom properties for theme tokens. |
| Framer Motion | latest | Typewriter animation, scroll reveals, section transitions. |
| next/font | built-in | Load JetBrains Mono + Inter. No Google Fonts CDN — self-hosted for performance. |
| Vercel | — | Hosting. Zero-config deploy. Free tier. |

---

## File Structure

```
website/
├── app/
│   ├── layout.tsx                 # Root layout: fonts, theme provider, metadata, navbar, footer
│   ├── page.tsx                   # One-pager: composes all sections in order
│   ├── api/
│   │   └── waitlist/route.ts      # POST handler for email capture
│   ├── globals.css                # CSS custom properties (theme tokens), Tailwind base
│   └── favicon.ico
├── components/
│   ├── navbar.tsx                 # Fixed nav: logo, links, theme toggle, CTA
│   ├── hero.tsx                   # Typewriter tagline + subheadline + waitlist form
│   ├── problem.tsx                # 3 problem cards
│   ├── solution.tsx               # Solution intro + 3 capability highlights
│   ├── how-it-works.tsx           # Terminal mockup with step animation
│   ├── teaching-difference.tsx    # Side-by-side comparison (assistant vs Pear)
│   ├── pricing.tsx                # Free + Pro cards + Teams coming soon
│   ├── final-cta.tsx              # Repeated tagline + waitlist form
│   ├── footer.tsx                 # Minimal footer
│   ├── ui/
│   │   ├── waitlist-form.tsx      # Reusable email input + submit (used in hero + final CTA)
│   │   ├── typewriter-text.tsx    # Character-by-character typing animation
│   │   ├── terminal-mockup.tsx    # Dark terminal window with animated content
│   │   ├── floating-dots.tsx      # Canvas-based dot background with parallax
│   │   ├── theme-toggle.tsx       # Sun/moon toggle, localStorage persistence
│   │   ├── scroll-reveal.tsx      # Intersection Observer wrapper for fade-in
│   │   └── section.tsx            # Reusable section wrapper (max-width, padding, alternating bg)
│   └── icons/                     # Small SVG icon components (sun, moon, menu, arrow, check)
├── lib/
│   ├── theme-provider.tsx         # React context for theme state
│   └── utils.ts                   # cn() classname merger, any shared helpers
├── docs/                          # Design docs (not deployed)
│   ├── PRD.md
│   ├── ARCHITECTURE.md
│   ├── DESIGNGUIDELINES.md
│   └── POSITIONING.md
├── public/
│   ├── og-image.png               # Open Graph image (1200x630)
│   └── favicon.ico
├── tailwind.config.ts
├── next.config.ts
├── package.json
└── tsconfig.json
```

---

## Theme System

### CSS Custom Properties (`globals.css`)

```css
:root {
  --bg: #FFFFFF;
  --bg-subtle: #FAFAFA;
  --fg: #0A0A0A;
  --fg-muted: #6B7280;
  --border: #E5E7EB;
  --accent: #22C55E;
  --accent-hover: #16A34A;
  --accent-subtle: #DCFCE7;
  --dot-color: rgba(10, 10, 10, 0.06);
}

[data-theme="dark"] {
  --bg: #0A0A0A;
  --bg-subtle: #111111;
  --fg: #FAFAFA;
  --fg-muted: #9CA3AF;
  --border: #1F2937;
  --accent: #22C55E;
  --accent-hover: #4ADE80;
  --accent-subtle: #052E16;
  --dot-color: rgba(255, 255, 255, 0.06);
}

* {
  transition: background-color 200ms, color 200ms, border-color 200ms;
}
```

### Theme Provider (`lib/theme-provider.tsx`)

- React context wrapping `<html>` element
- Reads `localStorage.getItem('theme')` on mount
- Falls back to `prefers-color-scheme` media query
- Sets `data-theme` attribute on `<html>`
- Provides `{ theme, toggleTheme }` to consumers

---

## Component Implementation Notes

### `floating-dots.tsx` — Background

- **Canvas-based** for performance. Single `<canvas>` element covering the viewport, `position: fixed`, `z-index: 0`, `pointer-events: none`.
- Generate a grid of dots: 2px circles, 40px spacing, offset to fill viewport.
- Each dot drifts ±3px using `Math.sin(time * speed + offset)` for x and y. Randomize `speed` (0.5-0.8) and `offset` (0-2π) per dot for organic feel.
- Redraw on `requestAnimationFrame`. Only draw dots visible in current viewport bounds.
- **Parallax:** Offset all dot y-positions by `window.scrollY * 0.5`. Listen to scroll via `passive: true` event listener.
- Read `--dot-color` from CSS custom property for theme reactivity. Re-read on theme change.
- On resize: recalculate grid dimensions.

### `typewriter-text.tsx` — Hero Tagline

Props: `text: string`, `speed?: number` (default 50ms), `onComplete?: () => void`

- Renders characters one at a time using `useState` index + `setInterval`.
- Appends a blinking cursor (`▊`) in `--accent` color.
- Cursor blinks at 530ms via CSS `animation: blink 1.06s step-end infinite`.
- Calls `onComplete` when all characters are rendered (triggers subheadline fade-in).

### `terminal-mockup.tsx` — How It Works

- Dark background (`#0A0A0A`) with title bar (three colored dots: red/yellow/green circles).
- Content is an array of lines with `type` and `delay`:
  ```ts
  const lines = [
    { type: 'command', text: '$ pear', delay: 0 },
    { type: 'output', text: '🍐 Pear active. Hold [Space] to talk.', delay: 800 },
    { type: 'voice', text: '🎤 "Walk me through what changed in this file."', delay: 1500 },
    { type: 'context', text: '📎 Context: git diff (3 files), src/auth/rbac.go, file tree', delay: 2500 },
    { type: 'response', text: 'Looking at your diff, Claude made 3 changes to rbac.go...\n\n1. Line 23: Switched from sync.Mutex to sync.RWMutex — This is\n   a good change. RWMutex lets multiple goroutines read\n   concurrently...\n\n2. Line 41: Added context.WithTimeout — Without this, a slow\n   DB call could hang your request indefinitely...', delay: 3500 },
  ];
  ```
- Lines appear sequentially. `command` and `voice` lines use typewriter effect. `output`, `context`, and `response` lines fade in. `response` streams character-by-character at 15ms/char.
- Triggered by Intersection Observer (play once on scroll into view).
- Small "↻ Replay" button in top-right corner of terminal, muted color.

### `scroll-reveal.tsx` — Section Animations

Generic wrapper component using Intersection Observer.

Props: `children`, `delay?: number` (stagger offset in ms), `className?: string`

- Initial state: `opacity: 0`, `translateY: 20px`
- On intersection (15% threshold): animate to `opacity: 1`, `translateY: 0`
- Duration: 500ms, easing: `[0.16, 1, 0.3, 1]` (Framer Motion ease-out-expo)
- Animate once — don't re-trigger on scroll back up.

### `waitlist-form.tsx` — Email Capture

- Uncontrolled form with `<input type="email">` + submit button
- Client-side validation: basic email regex before submit
- On submit: `POST /api/waitlist` with `{ email }` JSON body
- States: `idle` → `submitting` (button shows spinner) → `success` (replace form with "✓ You're on the list.") or `error` (red border + message)
- Reusable: used in `hero.tsx` and `final-cta.tsx`

### `section.tsx` — Reusable Section Wrapper

Props: `id`, `children`, `alternate?: boolean` (toggles `--bg` vs `--bg-subtle`), `className?: string`

- `<section>` with `max-w-[1200px] mx-auto`, vertical padding `py-24 md:py-32`
- Alternating background via `alternate` prop
- All sections get an `id` for anchor linking

---

## API Route

### `POST /api/waitlist`

```ts
// app/api/waitlist/route.ts
import { NextRequest, NextResponse } from 'next/server';

export async function POST(req: NextRequest) {
  const { email } = await req.json();

  // Validate
  if (!email || !/^[^\s@]+@[^\s@]+\.[^\s@]+$/.test(email)) {
    return NextResponse.json({ error: 'Invalid email' }, { status: 400 });
  }

  // Store — pick ONE:
  // Option A: Vercel KV (simplest, free tier)
  // Option B: POST to Resend audience API
  // Option C: Google Sheets API
  // Option D: Append to a Supabase table

  // For MVP, Vercel KV with a Set (deduplicates automatically):
  // await kv.sadd('waitlist', email);

  return NextResponse.json({ success: true });
}
```

Keep the storage simple. Vercel KV with a Set is the lowest-friction option for a solo founder — free tier, no external service, auto-deduplication.

---

## Metadata (`layout.tsx`)

```ts
export const metadata: Metadata = {
  title: 'Pear — AI makes you fast. Pear makes you good.',
  description: 'Pear is the voice-first AI tutor that teaches software engineering while you code. Talk to your codebase, learn the patterns, understand the why. macOS CLI.',
  openGraph: {
    title: 'Pear — AI makes you fast. Pear makes you good.',
    description: 'The voice-first AI tutor that teaches software engineering in your terminal.',
    url: 'https://pearcode.dev',
    siteName: 'Pear',
    images: [{ url: '/og-image.png', width: 1200, height: 630 }],
    type: 'website',
  },
  twitter: {
    card: 'summary_large_image',
    title: 'Pear — AI makes you fast. Pear makes you good.',
    description: 'The voice-first AI tutor that teaches software engineering in your terminal.',
    images: ['/og-image.png'],
  },
  icons: { icon: '/favicon.ico' },
};
```

---

## Performance Targets

| Metric | Target |
|---|---|
| Lighthouse Performance | 95+ |
| Lighthouse Accessibility | 100 |
| First Contentful Paint | <1.0s |
| Largest Contentful Paint | <2.0s |
| Cumulative Layout Shift | <0.05 |
| Total JS bundle | <150KB gzipped |
| Floating dots framerate | 60fps on M1 MacBook |

---

## Dependencies (package.json)

```json
{
  "dependencies": {
    "next": "^15",
    "react": "^19",
    "react-dom": "^19",
    "framer-motion": "^11",
    "@vercel/kv": "^2"
  },
  "devDependencies": {
    "typescript": "^5",
    "@types/react": "^19",
    "@types/node": "^22",
    "tailwindcss": "^4",
    "@tailwindcss/postcss": "^4"
  }
}
```

Minimal. No component libraries (no shadcn for the landing page — it's a single page with custom components). No animation libraries beyond Framer Motion.

---

## Deploy

- Push to GitHub → Vercel auto-deploys from `website/` directory
- Set Vercel root directory to `website/`
- Domain: `pearcode.dev` (configure DNS in Vercel)
- Preview deploys on every PR
- Environment variables: `KV_REST_API_URL`, `KV_REST_API_TOKEN` (if using Vercel KV)
