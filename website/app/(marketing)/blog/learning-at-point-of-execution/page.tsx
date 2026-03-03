import type { Metadata } from "next";
import Link from "next/link";
import { Section } from "@/components/ui/section";

export const dynamic = "force-static";

export const metadata: Metadata = {
  title:
    "The Case for Learning at the Point of Execution | pear blog",
  description:
    "Why context-switched learning fails for working developers. Courses, docs, and tutorials break your flow. Learning should happen where you code.",
  alternates: {
    canonical:
      "https://pearcode.dev/blog/learning-at-point-of-execution",
  },
  openGraph: {
    title: "The Case for Learning at the Point of Execution",
    description:
      "Why context-switched learning fails for working developers.",
    url: "https://pearcode.dev/blog/learning-at-point-of-execution",
    siteName: "pear",
    type: "article",
    publishedTime: "2026-02-21T00:00:00Z",
    authors: ["Mitch Hazelhurst"],
  },
};

const jsonLd = {
  "@context": "https://schema.org",
  "@type": "BlogPosting",
  headline: "The Case for Learning at the Point of Execution",
  description:
    "Why context-switched learning fails for working developers. Learning should happen where you code.",
  author: {
    "@type": "Person",
    name: "Mitch Hazelhurst",
    url: "https://pearcode.dev/about",
  },
  publisher: { "@id": "https://pearcode.dev/#organization" },
  datePublished: "2026-02-21",
  dateModified: "2026-02-21",
  mainEntityOfPage: {
    "@type": "WebPage",
    "@id": "https://pearcode.dev/blog/learning-at-point-of-execution",
  },
  wordCount: 1050,
  url: "https://pearcode.dev/blog/learning-at-point-of-execution",
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
              The Case for Learning at the Point of Execution
            </h1>
            <p className="mt-4 text-sm text-(--fg-muted)">
              By Mitch Hazelhurst &middot;{" "}
              <time dateTime="2026-02-21">February 21, 2026</time>
            </p>
          </header>

          <div className="mt-14 text-lg leading-relaxed text-(--fg)/70">
            <p>
              Every developer has done this: you hit a concept you don&apos;t understand, so
              you open a new tab. You Google it, find a blog post or a Stack Overflow answer,
              read for five minutes, then switch back to your code. By the time you&apos;re
              back, you&apos;ve lost your train of thought. The context you built up is gone.
            </p>
            <p className="mt-8">
              This is context-switched learning, and it&apos;s how almost all developer
              education works today. Courses, documentation, tutorials, bootcamps &mdash;
              they all require you to leave the place where you work to go to a place where
              you learn.
            </p>
            <p className="mt-8">
              It doesn&apos;t have to be this way.
            </p>

            <h2 className="mt-20 font-(family-name:--font-jetbrains) text-2xl font-bold text-(--fg)">
              Why context switching kills learning
            </h2>
            <p className="mt-8">
              Cognitive science has a clear finding: learning is most effective when it
              happens in context. When the material is directly relevant to a problem
              you&apos;re solving right now, retention goes up dramatically. When it&apos;s
              abstract and disconnected from your work, most of it evaporates within hours.
            </p>
            <p className="mt-8">
              Traditional developer education ignores this. You take a course on system
              design, then go back to building a CRUD app. You read about concurrency
              patterns, then spend the day writing React components. The knowledge exists in
              a vacuum, disconnected from the moments where it would actually be useful.
            </p>
            <p className="mt-8">
              The result is a learning industry that generates completion certificates but
              not competence. Developers finish courses and can&apos;t apply what they
              learned because they learned it in the wrong context.
            </p>

            <h2 className="mt-20 font-(family-name:--font-jetbrains) text-2xl font-bold text-(--fg)">
              The case for in-flow learning
            </h2>
            <p className="mt-8">
              What if learning happened at the exact moment you needed it? Not in a separate
              tab, not in a 40-minute video, not in a course you&apos;ll forget by next week.
              Right here, in your terminal, while you&apos;re looking at the code that
              prompted the question.
            </p>
            <p className="mt-8">
              This is what &ldquo;point of execution&rdquo; learning means. The teaching
              surfaces when you&apos;re actively working with the concept. Your AI agent just
              implemented a retry mechanism with exponential backoff? That&apos;s the moment
              to learn why exponential backoff matters, what the alternatives are, and when
              you&apos;d choose a different strategy.
            </p>
            <p className="mt-8">
              The context is already loaded in your head. The code is in front of you. The
              problem is real, not hypothetical. This is when learning sticks.
            </p>

            <h2 className="mt-20 font-(family-name:--font-jetbrains) text-2xl font-bold text-(--fg)">
              What in-flow learning requires
            </h2>
            <p className="mt-8">
              You can&apos;t just bolt a chatbot onto a terminal and call it in-flow learning.
              Real contextual education requires several things that don&apos;t exist in
              traditional tools:
            </p>
            <p className="mt-8">
              <strong className="text-(--fg)">Awareness of what you&apos;re working on.</strong>{" "}
              The tool needs to understand your code, your project structure, and what just
              changed. Generic explanations aren&apos;t contextual.
            </p>
            <p className="mt-8">
              <strong className="text-(--fg)">Memory of what you already know.</strong>{" "}
              Teaching the same concept twice wastes time. Skipping concepts you haven&apos;t
              seen wastes opportunity. The tool needs a model of your understanding.
            </p>
            <p className="mt-8">
              <strong className="text-(--fg)">Timing that doesn&apos;t break flow.</strong>{" "}
              Interrupting a developer mid-thought is worse than not teaching at all. Insights
              should surface in natural gaps: while your agent thinks, before you accept a
              plan, after a diff lands.
            </p>
            <p className="mt-8">
              <strong className="text-(--fg)">Adaptation to your level.</strong>{" "}
              A senior engineer switching to Rust needs different explanations than a bootcamp
              grad encountering async for the first time. One-size-fits-all teaching is
              no-size-fits-anyone teaching.
            </p>

            <h2 className="mt-20 font-(family-name:--font-jetbrains) text-2xl font-bold text-(--fg)">
              Pear&apos;s thesis
            </h2>
            <p className="mt-8">
              This is the thesis behind{" "}
              <Link
                href="/compare"
                className="text-pear underline underline-offset-4 hover:text-pear-hover"
              >
                pear
              </Link>
              . Developer education shouldn&apos;t happen in a classroom, a course platform,
              or a documentation site. It should happen in the terminal, at the moment of
              execution, adapted to what you&apos;re building and what you already understand.
            </p>
            <p className="mt-8">
              Pear watches your code, tags concepts, tracks your knowledge, and teaches you
              in the gaps between actions. It&apos;s not a course. It&apos;s not a chatbot.
              It&apos;s a learning system that lives where you work.
            </p>
            <p className="mt-8">
              The best time to learn something is when you need it. The best place to learn
              it is where you&apos;re using it. Everything else is a compromise.
            </p>
          </div>

          <div className="mt-20 text-center">
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
