import type { Metadata } from "next";
import Link from "next/link";
import { Section } from "@/components/ui/section";

export const dynamic = "force-static";

export const metadata: Metadata = {
  title: "AI Is Making Developers Fast, But Not Good | pear blog",
  description:
    "AI coding tools increase output but degrade understanding. The data from METR, Stanford, and GitClear shows why speed without comprehension is dangerous, and what needs to change.",
  alternates: {
    canonical:
      "https://pearcode.dev/blog/ai-making-developers-fast-not-good",
  },
  openGraph: {
    title: "AI Is Making Developers Fast, But Not Good",
    description:
      "AI coding tools increase output but degrade understanding. Here's why that matters.",
    url: "https://pearcode.dev/blog/ai-making-developers-fast-not-good",
    siteName: "pear",
    type: "article",
    publishedTime: "2026-02-11T00:00:00Z",
    authors: ["Mitch Hazelhurst"],
  },
};

const jsonLd = {
  "@context": "https://schema.org",
  "@type": "TechArticle",
  headline: "AI Is Making Developers Fast, But Not Good",
  description:
    "AI coding tools increase output but degrade understanding. The data from METR, Stanford, and GitClear shows why speed without comprehension is dangerous.",
  author: {
    "@type": "Person",
    name: "Mitch Hazelhurst",
    url: "https://pearcode.dev/about",
  },
  publisher: {
    "@id": "https://pearcode.dev/#organization",
  },
  datePublished: "2026-02-11",
  dateModified: "2026-02-11",
  mainEntityOfPage: {
    "@type": "WebPage",
    "@id": "https://pearcode.dev/blog/ai-making-developers-fast-not-good",
  },
  wordCount: 750,
  url: "https://pearcode.dev/blog/ai-making-developers-fast-not-good",
};

export default function BlogPost() {
  return (
    <main className="relative z-10 pt-16">
      <script
        type="application/ld+json"
        dangerouslySetInnerHTML={{ __html: JSON.stringify(jsonLd) }}
      />

      <Section>
        <article className="mx-auto max-w-2xl">
          <header>
            <Link
              href="/blog"
              className="text-sm text-(--fg-muted) transition-colors hover:text-(--fg)"
            >
              &larr; Back to blog
            </Link>
            <h1 className="mt-10 font-(family-name:--font-jetbrains) text-3xl font-bold text-(--fg) md:text-4xl">
              AI Is Making Developers Fast, But Not Good
            </h1>
            <p className="mt-4 text-sm text-(--fg-muted)">
              By Mitch Hazelhurst &middot;{" "}
              <time dateTime="2026-02-11">February 11, 2026</time>
            </p>
          </header>

          <div className="mt-14 text-lg leading-relaxed text-(--fg)/70">
            <p>
              Every developer I talk to says the same thing: AI tools made them
              faster. Ship features in hours instead of days. Generate entire
              modules from a prompt. Autocomplete that actually works.
            </p>
            <p className="mt-8">
              But ask them to explain what their code does. Why it&apos;s
              structured that way, what trade-offs were made, what would break
              if you changed the architecture. The confidence disappears.
              They assembled the code. They didn&apos;t engineer it.
            </p>

            <h2 className="mt-20 font-(family-name:--font-jetbrains) text-2xl font-bold text-(--fg)">
              The data tells the same story
            </h2>
            <p className="mt-8">
              A 2025 METR study found that AI tools actually made experienced
              open-source developers <strong className="text-(--fg)">19% slower</strong> on real-world
              tasks in repos they knew well, despite those developers
              predicting a 24% speedup. The tools helped generate code, but the
              overhead of reviewing, debugging, and correcting AI output ate
              the gains.
            </p>
            <p className="mt-8">
              Stanford research on AI-assisted coding found that developers who
              used AI assistants produced code with{" "}
              <strong className="text-(--fg)">more security vulnerabilities</strong> than those who
              didn&apos;t, and were simultaneously more confident their code
              was secure. Speed bred false confidence.
            </p>
            <p className="mt-8">
              GitClear&apos;s analysis of 150 million lines of code showed that
              AI-assisted codebases have significantly{" "}
              <strong className="text-(--fg)">higher churn rates</strong>. Code that gets written then
              quickly rewritten or deleted. More output, but less of it
              survives. Stack Overflow traffic dropped over 50% as developers
              stopped researching and started prompting, trading deep
              understanding for fast answers.
            </p>

            <h2 className="mt-20 font-(family-name:--font-jetbrains) text-2xl font-bold text-(--fg)">
              Speed without understanding is debt
            </h2>
            <p className="mt-8">
              When you generate code you don&apos;t understand, you&apos;re not
              saving time. You&apos;re borrowing it. Every function you
              can&apos;t explain is a future debugging session. Every
              architectural decision you didn&apos;t consciously make is a
              refactor waiting to happen.
            </p>
            <p className="mt-8">
              This isn&apos;t an argument against AI tools. They&apos;re
              genuinely useful and they&apos;re not going away. The problem is
              that the entire ecosystem optimises for one metric:{" "}
              <strong className="text-(--fg)">speed of code production</strong>, while ignoring the
              metric that actually determines engineering quality:{" "}
              <strong className="text-(--fg)">depth of understanding</strong>.
            </p>

            <h2 className="mt-20 font-(family-name:--font-jetbrains) text-2xl font-bold text-(--fg)">
              What needs to change
            </h2>
            <p className="mt-8">
              Developers need tools that close the gap between what they ship
              and what they understand. Not tools that slow them down. Tools
              that teach them as they go. The best learning happens in context:
              when you&apos;re looking at your own code, solving your own
              problems, and the explanation is relevant to what you&apos;re
              building right now.
            </p>
            <p className="mt-8">
              That&apos;s why I&apos;m building pear. It&apos;s a CLI tutor
              that sits in your terminal and teaches you software engineering
              concepts in the context of your actual work. It reads your diffs,
              understands your project structure, and explains what you&apos;re
              building and why it matters.
            </p>
            <p className="mt-8">
              It doesn&apos;t write code for you. It doesn&apos;t autocomplete.
              It teaches. Because the goal isn&apos;t to ship faster. It&apos;s
              to ship faster <em>and</em> actually understand what you shipped.
            </p>

            <h2 className="mt-20 font-(family-name:--font-jetbrains) text-2xl font-bold text-(--fg)">
              The developers who will thrive
            </h2>
            <p className="mt-8">
              The next decade belongs to developers who use AI as a lever, not
              a crutch. The ones who can ship at AI speed but debug at human
              depth. Who can read a codebase, not just generate one. Who treat
              understanding as the competitive advantage it is.
            </p>
            <p className="mt-8">
              AI makes developers fast. The question is whether you&apos;re also
              getting good.
            </p>
          </div>

          <hr className="mt-20 border-border" />
          <div className="mt-10">
            <h3 className="font-(family-name:--font-jetbrains) text-sm font-semibold uppercase tracking-wider text-(--fg-muted)">
              Related reading
            </h3>
            <ul className="mt-4 space-y-3 text-base">
              <li>
                <Link href="/blog/understanding-ai-generated-code" className="text-pear underline underline-offset-4 hover:text-pear-hover">
                  Why Understanding Your AI-Generated Code Actually Matters
                </Link>
              </li>
              <li>
                <Link href="/blog/claude-code-vs-cursor-vs-opencode" className="text-pear underline underline-offset-4 hover:text-pear-hover">
                  Claude Code vs Cursor vs OpenCode: What They Do (and What They Don&apos;t)
                </Link>
              </li>
              <li>
                <Link href="/compare" className="text-pear underline underline-offset-4 hover:text-pear-hover">
                  See the full comparison: Pear vs AI coding tools
                </Link>
              </li>
            </ul>
          </div>

          <div className="mt-16 text-center">
            <Link
              href="/pricing"
              className="inline-block whitespace-nowrap rounded-lg bg-pear px-6 py-3 font-(family-name:--font-inter) font-semibold text-white cursor-pointer transition-colors hover:bg-pear-hover"
            >
              Try pear free
            </Link>
          </div>
        </article>
      </Section>
    </main>
  );
}
