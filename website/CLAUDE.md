# Pear Website — Claude Code Instructions

## What This Is
Marketing website for Pear (pearcode.dev). Next.js 15 App Router on Vercel.

## Tech Stack
- Next.js 15 (App Router, Turbopack)
- React 19
- Tailwind CSS 4
- shadcn/ui + Radix UI
- Framer Motion (animations)
- Three.js (3D hero element)
- Resend (email)
- Upstash Redis (KV, rate limiting)
- Vercel Analytics + Speed Insights

## Structure
```
app/
├── (marketing)/          # Public pages (home, blog, pricing, docs, about, FAQ, compare, terms, privacy)
├── (dashboard)/          # Future authenticated dashboard
└── api/waitlist/         # Waitlist API endpoint
components/               # React components
emails/                   # Resend email templates
lib/                      # Utilities, theme provider
public/                   # Static assets, logos
docs/                     # Website-specific docs (positioning, design guidelines)
```

## Reference Docs
- `docs/POSITIONING.md` — source of truth for all website copy and messaging
- `docs/DESIGNGUIDELINES.md` — design system and visual language
- `docs/PRD.md` — website-specific product requirements
- `docs/ARCHITECTURE.md` — website architecture details

## Rules
- Follow positioning and tone from `docs/POSITIONING.md`
- Use shadcn/ui components where possible
- Keep bundle size small — no unnecessary dependencies
- All pages must have proper meta tags, OG images, and sitemap entries
