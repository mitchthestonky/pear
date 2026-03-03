import type { Metadata } from "next";
import Link from "next/link";
import { Section } from "@/components/ui/section";

export const dynamic = "force-static";

export const metadata: Metadata = {
  title: "Blog | pear",
  description:
    "Thoughts on AI-assisted development, software engineering education, and building tools that teach.",
  alternates: { canonical: "https://pearcode.dev/blog" },
  openGraph: {
    title: "Blog | pear",
    description:
      "Thoughts on AI-assisted development, software engineering education, and building tools that teach.",
    url: "https://pearcode.dev/blog",
    siteName: "pear",
    type: "website",
  },
};

const posts = [
  {
    slug: "claude-code-vs-cursor-vs-opencode",
    title: "Claude Code vs Cursor vs OpenCode: What They Do (and What They Don't)",
    description:
      "An honest comparison of the top AI coding tools in 2026. What Claude Code, Cursor, and OpenCode do well, what they don't, and where pear fits in.",
    date: "2026-03-03",
  },
  {
    slug: "understanding-ai-generated-code",
    title: "Why Understanding Your AI-Generated Code Actually Matters",
    description:
      "The hidden cost of shipping code you can't explain. 66% of developers spend more time fixing almost-right AI code. Here's why understanding matters more than speed.",
    date: "2026-02-27",
  },
  {
    slug: "learning-at-point-of-execution",
    title: "The Case for Learning at the Point of Execution",
    description:
      "Why context-switched learning fails for working developers. Courses, docs, and tutorials break your flow. Learning should happen where you code.",
    date: "2026-02-21",
  },
  {
    slug: "five-engineering-skills-ai-cant-teach",
    title: "The 5 Engineering Skills AI Cannot Teach You",
    description:
      "AI can generate code, write tests, and explain syntax. But five core engineering skills require human experience to develop. Here is what they are and how to build them.",
    date: "2026-02-17",
  },
  {
    slug: "vibe-coding-junior-engineers",
    title: "What Vibe Coding Is Doing to Junior Engineers",
    description:
      "Vibe coding feels productive. But when juniors skip understanding and ship prompts instead of code, the long-term cost is steep.",
    date: "2026-02-13",
  },
  {
    slug: "ai-making-developers-fast-not-good",
    title: "AI Is Making Developers Fast, But Not Good",
    description:
      "The data is clear: AI tools increase output while degrading understanding. Here is why that matters and what we can do about it.",
    date: "2026-02-11",
  },
];

const jsonLd = {
  "@context": "https://schema.org",
  "@type": "CollectionPage",
  name: "pear blog",
  description:
    "Thoughts on AI-assisted development, software engineering education, and building tools that teach.",
  url: "https://pearcode.dev/blog",
  publisher: { "@id": "https://pearcode.dev/#organization" },
  mainEntity: {
    "@type": "ItemList",
    itemListElement: posts.map((post, i) => ({
      "@type": "ListItem",
      position: i + 1,
      url: `https://pearcode.dev/blog/${post.slug}`,
      name: post.title,
    })),
  },
};

export default function BlogPage() {
  return (
    <main className="relative z-10 pt-16">
      <script
        type="application/ld+json"
        dangerouslySetInnerHTML={{ __html: JSON.stringify(jsonLd) }}
      />
      <Section>
        <div className="mx-auto max-w-3xl">
          <h1 className="text-center font-(family-name:--font-jetbrains) text-3xl font-bold text-(--fg) md:text-4xl">
            Blog
          </h1>
          <p className="mt-6 text-center text-lg leading-relaxed text-(--fg-muted)">
            Thoughts on AI-assisted development, software engineering education,
            and building tools that teach.
          </p>

          <div className="mt-20">
            {posts.map((post, i) => (
              <article key={post.slug} className={i > 0 ? "mt-10" : ""}>
                <Link
                  href={`/blog/${post.slug}`}
                  className="group block rounded-2xl border border-border bg-background/70 px-6 py-8 backdrop-blur-lg transition-all hover:border-pear/40 md:px-12 md:py-10"
                >
                  <time className="text-sm text-(--fg-muted)">
                    {new Date(post.date).toLocaleDateString("en-US", {
                      year: "numeric",
                      month: "long",
                      day: "numeric",
                    })}
                  </time>
                  <h2 className="mt-4 font-(family-name:--font-jetbrains) text-2xl font-bold text-(--fg) transition-colors group-hover:text-pear">
                    {post.title}
                  </h2>
                  <p className="mt-4 text-base leading-relaxed text-(--fg-muted)">
                    {post.description}
                  </p>
                </Link>
              </article>
            ))}
          </div>
        </div>
      </Section>
    </main>
  );
}
