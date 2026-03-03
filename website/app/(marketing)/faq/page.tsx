import type { Metadata } from "next";
import Link from "next/link";
import { Section } from "@/components/ui/section";

export const dynamic = "force-static";

export const metadata: Metadata = {
  title: "FAQ | pear | AI Pair Programmer",
  description:
    "Frequently asked questions about pear: pricing, BYOK, learning state memory, supported LLMs, privacy, and how pear compares to Claude Code, Cursor, and OpenCode.",
  alternates: { canonical: "https://pearcode.dev/faq" },
  openGraph: {
    title: "FAQ | pear | AI Pair Programmer",
    description:
      "Everything you need to know about pear. Pricing, learning system, BYOK, privacy, and more.",
    url: "https://pearcode.dev/faq",
    siteName: "pear",
    type: "website",
  },
};

const faqs = [
  {
    q: "What is pear?",
    a: "pear is a pair programmer for AI-enabled engineers. It runs in your terminal alongside your coding agent, watches your diffs and context, and tells you what matters. The free open-source tier auto-reviews changes and surfaces insights. Pro remembers what you understand, detects knowledge gaps, and teaches you at the moment of execution.",
  },
  {
    q: "Why can't I just ask ChatGPT or Claude to explain my code?",
    a: "You can — and you'll get a decent answer. But that's an explainer, not a learning system. ChatGPT forgets everything between sessions. It doesn't know what you already understand. It can't detect gaps you don't know you have. It won't proactively teach you when a concept is relevant. pear Pro remembers your learning state, adapts its teaching to your level, and surfaces insights at the right moment — without you having to ask.",
  },
  {
    q: "How is Pear different from Claude Code / Cursor / OpenCode?",
    a: "Claude Code, Cursor, and OpenCode are AI coding tools — they generate and edit code for you. Pear doesn't write code. It's the layer that makes sure you understand what those tools write. They make you faster; Pear makes you smarter. Use them together.",
  },
  {
    q: "How does pear work?",
    a: "You run pear in your terminal alongside your editor. It reads context from your coding session (diffs, file tree, recent changes), checks your learning state to understand what you already know, and teaches you only what's new or misunderstood. Over time, it builds a concept graph of your engineering knowledge and adapts explanations accordingly.",
  },
  {
    q: "What does BYOK mean?",
    a: "BYOK stands for Bring Your Own Key. Instead of routing your requests through our servers, pear sends prompts directly to your LLM provider using your own API key. This means you control your costs, your data stays between you and your provider, and you can use whichever model you prefer.",
  },
  {
    q: "Do I need an API key?",
    a: "Yes. Pear is BYOK (Bring Your Own Key) — you provide your own API key from Anthropic, OpenAI, or OpenRouter. This keeps your data private and your costs transparent. Managed mode (no API key needed) is planned for the future.",
  },
  {
    q: "Which LLMs does pear support?",
    a: "At launch, pear supports Claude (Anthropic), GPT (OpenAI), and any model available through OpenRouter. Since pear is BYOK, you'll use whichever provider you already have an API key for.",
  },
  {
    q: "How much does pear cost?",
    a: "Pear Free (open source) is free forever with your own API key — it includes watch mode, auto-reviews, and concept tagging. Pear Pro is $20/month or $130/year (4 months free) and adds learning state memory, adaptive teaching, knowledge gap tracking, and cross-machine sync. Need team-wide learning metrics? Get in touch about Teams pricing.",
  },
  {
    q: "Is there a free tier?",
    a: "Yes. Pear Free is open source and free forever. Bring your own API key and you get watch mode (auto-reviews your changes), interactive Q&A with codebase context, subtle teaching moments, and concept tagging. No credit card, no time limit.",
  },
  {
    q: "What about privacy? Where does my code go?",
    a: "With BYOK, your code context is sent directly from your machine to your chosen LLM provider. pear doesn't store your code on our servers. Your API key is stored locally in your config file and never transmitted to us.",
  },
  {
    q: "Is pear available on Linux or Windows?",
    a: "pear launches on macOS only. Linux support is planned for v1.6. Windows support is on the roadmap but not yet scheduled — WSL may work but isn't officially supported at launch.",
  },
  {
    q: "Do I need to change my editor to use pear?",
    a: "No. pear runs in your terminal independently of your editor. Use VS Code, Neovim, Zed, or whatever you prefer — pear works alongside any editor by reading your project's git state and file structure.",
  },
  {
    q: "How do I install pear?",
    a: "pear installs via go install, Homebrew (brew install MitchTheStonky/pear/pear), or a one-line curl script. Run the installer, add your API key with pear init, and you're ready to go. The whole setup takes under a minute.",
  },
];

const jsonLd = {
  "@context": "https://schema.org",
  "@type": "FAQPage",
  mainEntity: faqs.map((faq) => ({
    "@type": "Question",
    name: faq.q,
    acceptedAnswer: {
      "@type": "Answer",
      text: faq.a,
    },
  })),
};

export default function FAQPage() {
  return (
    <main className="relative z-10 pt-16">
      <script
        type="application/ld+json"
        dangerouslySetInnerHTML={{ __html: JSON.stringify(jsonLd) }}
      />

      <Section>
        <div className="mx-auto max-w-2xl">
          <h1 className="text-center font-(family-name:--font-jetbrains) text-3xl font-bold text-(--fg) md:text-4xl">
            Frequently asked questions
          </h1>
          <p className="mt-4 text-center text-lg text-(--fg)/70">
            Everything you need to know about pear. Can&apos;t find what
            you&apos;re looking for?{" "}
            <a
              href="mailto:mitch@pearcode.dev"
              className="text-pear underline underline-offset-4 hover:text-pear-hover"
            >
              Get in touch
            </a>
            .
          </p>

          <div className="mt-16 divide-y divide-border">
            {faqs.map((faq) => (
              <details key={faq.q} className="group py-6">
                <summary className="cursor-pointer flex list-none items-center justify-between font-(family-name:--font-jetbrains) text-base font-semibold text-(--fg) [&::-webkit-details-marker]:hidden">
                  {faq.q}
                  <span className="ml-6 shrink-0 text-lg text-(--fg-muted) transition-transform duration-200 group-open:rotate-45">
                    +
                  </span>
                </summary>
                <p className="mt-4 text-base leading-relaxed text-(--fg)/70">
                  {faq.a}
                </p>
              </details>
            ))}
          </div>

          <div className="mt-16 text-center">
            <Link
              href="/#hero"
              className="inline-block whitespace-nowrap rounded-lg bg-pear px-6 py-3 font-(family-name:--font-inter) font-semibold text-white cursor-pointer transition-colors hover:bg-pear-hover"
            >
              Join the waitlist
            </Link>
          </div>
        </div>
      </Section>
    </main>
  );
}
