import type { Metadata } from "next";
import Link from "next/link";
import { Section } from "@/components/ui/section";

export const dynamic = "force-static";

export const metadata: Metadata = {
  title: "The 5 Engineering Skills AI Cannot Teach You | pear blog",
  description:
    "AI can generate code, write tests, and explain syntax. But five core engineering skills require human experience to develop. Here is what they are and how to build them.",
  alternates: {
    canonical:
      "https://pearcode.dev/blog/five-engineering-skills-ai-cant-teach",
  },
  openGraph: {
    title: "The 5 Engineering Skills AI Cannot Teach You",
    description:
      "AI can generate code, write tests, and explain syntax. But five core engineering skills require human experience to develop.",
    url: "https://pearcode.dev/blog/five-engineering-skills-ai-cant-teach",
    siteName: "pear",
    type: "article",
    publishedTime: "2026-02-17T00:00:00Z",
    authors: ["Mitch Hazelhurst"],
  },
};

const jsonLd = {
  "@context": "https://schema.org",
  "@type": "TechArticle",
  headline: "The 5 Engineering Skills AI Cannot Teach You",
  description:
    "AI can generate code, write tests, and explain syntax. But five core engineering skills require human experience to develop.",
  author: {
    "@type": "Person",
    name: "Mitch Hazelhurst",
    url: "https://pearcode.dev/about",
  },
  publisher: {
    "@id": "https://pearcode.dev/#organization",
  },
  datePublished: "2026-02-17",
  dateModified: "2026-02-17",
  mainEntityOfPage: {
    "@type": "WebPage",
    "@id": "https://pearcode.dev/blog/five-engineering-skills-ai-cant-teach",
  },
  wordCount: 1000,
  url: "https://pearcode.dev/blog/five-engineering-skills-ai-cant-teach",
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
              The 5 Engineering Skills AI Cannot Teach You
            </h1>
            <p className="mt-4 text-sm text-(--fg-muted)">
              By Mitch Hazelhurst &middot;{" "}
              <time dateTime="2026-02-17">February 17, 2026</time>
            </p>
          </header>

          <div className="mt-14 text-lg leading-relaxed text-(--fg)/70">
            <p>
              AI can generate code, write tests, explain syntax, refactor
              functions, and even architect systems. It is an extraordinary tool
              for productivity. But there are five engineering skills that AI
              cannot teach you on its own, no matter how good the model gets.
            </p>
            <p className="mt-8">
              These are the skills that determine whether you are an engineer or
              just someone who operates an AI tool.
            </p>

            <h2 className="mt-20 font-(family-name:--font-jetbrains) text-2xl font-bold text-(--fg)">
              1. Systems thinking
            </h2>
            <p className="mt-8">
              AI generates code one prompt at a time. It optimises locally. But
              real engineering is about understanding how components interact
              across a system. How does this database change affect the cache
              layer? What happens to the queue when this service scales to 10x
              traffic? What are the second-order effects of adding this
              dependency?
            </p>
            <p className="mt-8">
              Systems thinking comes from building things, watching them break,
              and understanding why. No amount of generated code teaches you to
              see the invisible connections between components.
            </p>

            <h2 className="mt-20 font-(family-name:--font-jetbrains) text-2xl font-bold text-(--fg)">
              2. Debugging under pressure
            </h2>
            <p className="mt-8">
              When production is down and the CEO is in the Slack channel, you
              need more than prompting skills. You need the ability to form
              hypotheses, isolate variables, read logs with intuition, and
              navigate unfamiliar code under time pressure. This is a skill
              built through hundreds of debugging sessions, not through watching
              AI explain error messages.
            </p>
            <p className="mt-8">
              The developers who are dangerous in an incident room are the ones
              who have personally struggled through enough failures to develop
              instinct. AI can help you debug. It cannot teach you how to debug.
            </p>

            <h2 className="mt-20 font-(family-name:--font-jetbrains) text-2xl font-bold text-(--fg)">
              3. Trade-off reasoning
            </h2>
            <p className="mt-8">
              Every engineering decision is a trade-off. Speed vs safety.
              Simplicity vs flexibility. Consistency vs availability. AI will
              happily generate code for any approach you ask for, but it will
              not teach you when to choose one over the other.
            </p>
            <p className="mt-8">
              Trade-off reasoning requires context that AI does not have: your
              team size, your deployment cadence, your risk tolerance, your
              customers. It also requires judgement that only develops through
              making decisions and living with their consequences.
            </p>

            <h2 className="mt-20 font-(family-name:--font-jetbrains) text-2xl font-bold text-(--fg)">
              4. Codebase navigation
            </h2>
            <p className="mt-8">
              Senior engineers can open an unfamiliar codebase and within an
              hour, understand how it is structured, where the critical paths
              are, and where the complexity hides. This is not about reading
              code line by line. It is about pattern recognition: spotting the
              architecture, identifying the conventions, understanding the
              history behind decisions.
            </p>
            <p className="mt-8">
              AI can explain individual files. It cannot teach you to read a
              codebase like a map. That skill comes from navigating many
              codebases, many times, with genuine curiosity about how things
              fit together.
            </p>

            <h2 className="mt-20 font-(family-name:--font-jetbrains) text-2xl font-bold text-(--fg)">
              5. Technical communication
            </h2>
            <p className="mt-8">
              The ability to explain a technical decision to a non-technical
              stakeholder, write a design doc that survives scrutiny, or mentor
              a junior through a concept they are struggling with. These are
              deeply human skills that require empathy, clarity, and the ability
              to model someone else&apos;s understanding.
            </p>
            <p className="mt-8">
              AI can generate documentation. It cannot teach you to communicate
              with the precision and empathy that builds trust in a team.
            </p>

            <h2 className="mt-20 font-(family-name:--font-jetbrains) text-2xl font-bold text-(--fg)">
              What this means for you
            </h2>
            <p className="mt-8">
              None of this means AI is not valuable. It is incredibly valuable.
              But if your entire development process is prompt, accept, ship,
              you are optimising for output while starving the skills that
              actually make you employable long-term.
            </p>
            <p className="mt-8">
              The engineers who will thrive are the ones who use AI for speed
              and invest in understanding. They use the time AI saves them to
              go deeper on the concepts that matter, not just to ship more
              features.
            </p>
            <p className="mt-8">
              That is the philosophy behind pear. We built a learning engine
              that sits in your terminal and teaches you these skills in the
              context of your actual work. It does not replace the struggle. It
              makes sure the struggle teaches you something.
            </p>
          </div>

          <hr className="mt-20 border-border" />
          <div className="mt-10">
            <h3 className="font-(family-name:--font-jetbrains) text-sm font-semibold uppercase tracking-wider text-(--fg-muted)">
              Related reading
            </h3>
            <ul className="mt-4 space-y-3 text-base">
              <li>
                <Link href="/blog/learning-at-point-of-execution" className="text-pear underline underline-offset-4 hover:text-pear-hover">
                  The Case for Learning at the Point of Execution
                </Link>
              </li>
              <li>
                <Link href="/blog/vibe-coding-junior-engineers" className="text-pear underline underline-offset-4 hover:text-pear-hover">
                  What Vibe Coding Is Doing to Junior Engineers
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
