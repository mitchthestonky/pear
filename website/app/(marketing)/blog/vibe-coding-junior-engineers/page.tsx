import type { Metadata } from "next";
import Link from "next/link";
import { Section } from "@/components/ui/section";

export const dynamic = "force-static";

export const metadata: Metadata = {
  title: "What Vibe Coding Is Doing to Junior Engineers | pear blog",
  description:
    "Vibe coding feels productive. But when juniors skip understanding and ship prompts instead of code, the long-term cost is steep. Here is what is actually happening.",
  alternates: {
    canonical:
      "https://pearcode.dev/blog/vibe-coding-junior-engineers",
  },
  openGraph: {
    title: "What Vibe Coding Is Doing to Junior Engineers",
    description:
      "Vibe coding feels productive. But when juniors skip understanding and ship prompts instead of code, the long-term cost is steep.",
    url: "https://pearcode.dev/blog/vibe-coding-junior-engineers",
    siteName: "pear",
    type: "article",
    publishedTime: "2026-02-13T00:00:00Z",
    authors: ["Mitch Hazelhurst"],
  },
};

const jsonLd = {
  "@context": "https://schema.org",
  "@type": "TechArticle",
  headline: "What Vibe Coding Is Doing to Junior Engineers",
  description:
    "Vibe coding feels productive. But when juniors skip understanding and ship prompts instead of code, the long-term cost is steep.",
  author: {
    "@type": "Person",
    name: "Mitch Hazelhurst",
    url: "https://pearcode.dev/about",
  },
  publisher: {
    "@id": "https://pearcode.dev/#organization",
  },
  datePublished: "2026-02-13",
  dateModified: "2026-02-13",
  mainEntityOfPage: {
    "@type": "WebPage",
    "@id": "https://pearcode.dev/blog/vibe-coding-junior-engineers",
  },
  wordCount: 900,
  url: "https://pearcode.dev/blog/vibe-coding-junior-engineers",
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
              What Vibe Coding Is Doing to Junior Engineers
            </h1>
            <p className="mt-4 text-sm text-(--fg-muted)">
              By Mitch Hazelhurst &middot;{" "}
              <time dateTime="2026-02-13">February 13, 2026</time>
            </p>
          </header>

          <div className="mt-14 text-lg leading-relaxed text-(--fg)/70">
            <p>
              There is a new workflow spreading through junior engineering teams.
              Open Cursor. Describe what you want. Accept the code. Ship it.
              Repeat.
            </p>
            <p className="mt-8">
              People are calling it &quot;vibe coding&quot; and it feels like a
              superpower. You can build features you would have struggled with
              six months ago. You can ship pull requests that look like a senior
              wrote them. Your output metrics are through the roof.
            </p>
            <p className="mt-8">
              But here is the thing nobody talks about: you are not learning
              anything.
            </p>

            <h2 className="mt-20 font-(family-name:--font-jetbrains) text-2xl font-bold text-(--fg)">
              The illusion of competence
            </h2>
            <p className="mt-8">
              Vibe coding creates a specific kind of confidence. You can
              build things, so you feel like you understand things. But building
              and understanding are different skills. A junior who vibe-codes
              an authentication system can ship it in an afternoon. Ask them
              why they chose JWTs over sessions, what happens when the signing
              key rotates, or how token refresh works under the hood, and you
              get silence.
            </p>
            <p className="mt-8">
              This is not a failure of intelligence. It is a failure of process.
              The AI did the thinking for them, and nothing in the workflow
              prompted them to think for themselves.
            </p>

            <h2 className="mt-20 font-(family-name:--font-jetbrains) text-2xl font-bold text-(--fg)">
              What juniors are actually skipping
            </h2>
            <p className="mt-8">
              When you vibe-code, you skip the part of software engineering that
              builds real skill: the struggle. Reading error messages. Tracing
              execution. Understanding why one approach works and another does
              not. That friction is not a bug in the learning process. It is
              the learning process.
            </p>
            <p className="mt-8">
              Specifically, vibe coding lets you bypass:
            </p>
            <ul className="mt-6 list-inside list-disc space-y-3">
              <li>
                <strong className="text-(--fg)">Pattern recognition</strong> -
                seeing recurring structures across different problems
              </li>
              <li>
                <strong className="text-(--fg)">Trade-off analysis</strong> -
                understanding why you would choose one approach over another
              </li>
              <li>
                <strong className="text-(--fg)">Debugging intuition</strong> -
                developing a sense for where things break and why
              </li>
              <li>
                <strong className="text-(--fg)">Architectural thinking</strong> -
                understanding how individual decisions compound into systems
              </li>
            </ul>
            <p className="mt-8">
              These are the skills that separate a junior from a mid-level
              engineer. They cannot be prompt-engineered. They have to be lived.
            </p>

            <h2 className="mt-20 font-(family-name:--font-jetbrains) text-2xl font-bold text-(--fg)">
              The career ceiling nobody sees coming
            </h2>
            <p className="mt-8">
              Companies are starting to notice. A junior who ships fast but
              cannot debug their own code is a liability, not an asset. When
              production breaks at 2am, prompt skills do not help. When a
              principal engineer asks you to explain your design decisions in a
              review, &quot;the AI suggested it&quot; is not an answer.
            </p>
            <p className="mt-8">
              The developers who will hit a ceiling are the ones who used AI to
              avoid learning instead of using AI to accelerate learning. The
              distinction is subtle but the outcomes are wildly different.
            </p>

            <h2 className="mt-20 font-(family-name:--font-jetbrains) text-2xl font-bold text-(--fg)">
              A different approach
            </h2>
            <p className="mt-8">
              The answer is not to stop using AI tools. That ship has sailed,
              and honestly, they are too useful to ignore. The answer is to
              pair AI-assisted development with deliberate learning. Every time
              AI generates code for you, that is a teaching moment: what
              patterns did it use? What trade-offs did it make? What would you
              have done differently if you understood the problem deeply?
            </p>
            <p className="mt-8">
              This is exactly why we built pear. It watches your coding sessions,
              detects the concepts you are working with, and teaches you what
              you need to know in real time. Not after the fact. Not in a
              separate learning app. Right there in your terminal, while you
              are building.
            </p>
            <p className="mt-8">
              Vibe coding is not the enemy. Vibe coding without learning is.
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
                <Link href="/blog/five-engineering-skills-ai-cant-teach" className="text-pear underline underline-offset-4 hover:text-pear-hover">
                  The 5 Engineering Skills AI Cannot Teach You
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
