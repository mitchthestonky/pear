import type { Metadata } from "next";
import Link from "next/link";
import { Section } from "@/components/ui/section";
import { TeachingDifference } from "@/components/teaching-difference";

export const dynamic = "force-static";

export const metadata: Metadata = {
  title: "Pear vs Claude Code vs Cursor vs OpenCode | pear",
  description:
    "See how pear compares to Claude Code, Cursor, and OpenCode. AI coding tools generate code. Pear makes sure you understand it.",
  alternates: { canonical: "https://pearcode.dev/compare" },
  openGraph: {
    title: "Pear vs Claude Code vs Cursor vs OpenCode",
    description:
      "AI coding tools generate code. Pear makes sure you understand it.",
    url: "https://pearcode.dev/compare",
    siteName: "pear",
    type: "website",
  },
};

const jsonLd = {
  "@context": "https://schema.org",
  "@type": "WebPage",
  name: "Pear vs Claude Code vs Cursor vs OpenCode",
  description:
    "See how pear compares to Claude Code, Cursor, and OpenCode. AI coding tools generate code. Pear makes sure you understand it.",
  url: "https://pearcode.dev/compare",
  publisher: { "@id": "https://pearcode.dev/#organization" },
};

export default function ComparePage() {
  return (
    <main className="relative z-10 pt-16">
      <script
        type="application/ld+json"
        dangerouslySetInnerHTML={{ __html: JSON.stringify(jsonLd) }}
      />

      <Section>
        <div className="mx-auto max-w-3xl">
          <h1 className="text-center font-(family-name:--font-jetbrains) text-3xl font-bold text-(--fg) md:text-4xl">
            Pear vs Claude Code vs Cursor vs OpenCode
          </h1>
          <p className="mt-6 text-center text-lg leading-relaxed text-(--fg-muted)">
            AI coding tools generate code. Pear makes sure you understand it.
            They&rsquo;re not competitors &mdash; they&rsquo;re complementary.
          </p>
        </div>
      </Section>

      <TeachingDifference />

      <Section>
        <div className="mx-auto max-w-2xl text-lg leading-relaxed text-(--fg)/70">
          <h2 className="mt-4 font-(family-name:--font-jetbrains) text-2xl font-bold text-(--fg)">
            Claude Code
          </h2>
          <p className="mt-6">
            Anthropic&apos;s agentic coding tool lives in your terminal and can
            autonomously write, refactor, and debug code across your entire
            project. It excels at turning natural-language instructions into
            working implementations. What it doesn&apos;t do is teach you why those
            implementations work.
          </p>

          <h2 className="mt-16 font-(family-name:--font-jetbrains) text-2xl font-bold text-(--fg)">
            Cursor
          </h2>
          <p className="mt-6">
            Cursor is an AI-first code editor built on VS Code. It offers inline
            completions, chat, and multi-file edits. It&apos;s fast, polished,
            and great at generating code in context. But it optimises for
            shipping speed, not for your understanding of the code it produces.
          </p>

          <h2 className="mt-16 font-(family-name:--font-jetbrains) text-2xl font-bold text-(--fg)">
            OpenCode
          </h2>
          <p className="mt-6">
            OpenCode is an open-source terminal-based AI coding assistant. Like
            Claude Code, it can read your project and generate changes. It&apos;s
            a solid option for developers who want AI assistance without vendor
            lock-in. And like the others, it focuses on code generation, not
            comprehension.
          </p>

          <h2 className="mt-16 font-(family-name:--font-jetbrains) text-2xl font-bold text-(--fg)">
            Where pear fits
          </h2>
          <p className="mt-6">
            Pear is not a code generator. It doesn&apos;t autocomplete, write
            functions, or manage your git history. Instead, it watches the code
            your AI tools produce and teaches you what it means. It tags
            concepts, adapts explanations to your level, tracks your knowledge
            gaps, and shows your growth over time.
          </p>
          <p className="mt-6">
            Use Claude Code, Cursor, or OpenCode to write your code. Use pear to
            make sure you actually understand it.
          </p>

          <div className="mt-16 text-center">
            <Link
              href="/pricing"
              className="inline-block whitespace-nowrap rounded-lg bg-pear px-6 py-3 font-(family-name:--font-inter) font-semibold text-white cursor-pointer transition-colors hover:bg-pear-hover"
            >
              See pricing
            </Link>
          </div>
        </div>
      </Section>
    </main>
  );
}
