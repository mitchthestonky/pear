import type { Metadata } from "next";
import Image from "next/image";
import Link from "next/link";
import { Section } from "@/components/ui/section";

export const dynamic = "force-static";

export const metadata: Metadata = {
  title: "About pear | Meet the founder",
  description:
    "pear was built by Mitch Hazelhurst, an engineer who wished he had a learning engine that made him smarter while he shipped. Learn the story behind pear.",
  alternates: { canonical: "https://pearcode.dev/about" },
  openGraph: {
    title: "About pear | Meet the founder",
    description:
      "Built by an engineer who wished he had a learning engine that made him smarter while he shipped. Learn the story behind pear.",
    url: "https://pearcode.dev/about",
    siteName: "pear",
    type: "website",
  },
};

const jsonLd = {
  "@context": "https://schema.org",
  "@graph": [
    {
      "@type": "WebPage",
      "@id": "https://pearcode.dev/about",
      name: "About pear | Meet the founder",
      description:
        "pear was built by Mitch Hazelhurst, an engineer who wished he had a learning engine that made him smarter while he shipped.",
      url: "https://pearcode.dev/about",
      isPartOf: { "@id": "https://pearcode.dev/#website" },
      about: { "@id": "https://pearcode.dev/about#founder" },
    },
    {
      "@type": "Person",
      "@id": "https://pearcode.dev/about#founder",
      name: "Mitch Hazelhurst",
      jobTitle: "Founder",
      url: "https://pearcode.dev/about",
      image: "https://pearcode.dev/mitch.webp",
      sameAs: [
        "https://github.com/MitchTheStonky",
        "https://www.linkedin.com/in/mitchhazel/",
        "https://x.com/mitchthedev",
      ],
      worksFor: { "@id": "https://pearcode.dev/#organization" },
    },
    {
      "@type": "Organization",
      "@id": "https://pearcode.dev/#organization",
      name: "pear",
      url: "https://pearcode.dev",
      founder: { "@id": "https://pearcode.dev/about#founder" },
      description:
        "pear is the learning engine for AI-enabled engineers. Tracks what you understand, detects gaps, teaches at the moment of execution.",
    },
  ],
};

export default function AboutPage() {
  return (
    <main className="relative z-10 pt-16">
      <script
        type="application/ld+json"
        dangerouslySetInnerHTML={{ __html: JSON.stringify(jsonLd) }}
      />

      <Section>
        <div className="mx-auto max-w-3xl">
          <h1 className="text-center font-(family-name:--font-jetbrains) text-3xl font-bold text-(--fg) md:text-4xl">
            AI made me faster. It didn&rsquo;t make me better.
          </h1>

          {/* Row 1: Photo + intro text — side by side on md+, stacked on mobile */}
          <div className="mt-12 flex flex-col gap-10 md:flex-row md:items-start">
            <div className="flex shrink-0 justify-center">
              <Image
                src="/mitch.webp"
                alt="Mitch Hazelhurst, founder of pear"
                width={400}
                height={400}
                className="w-[220px] h-auto rounded-2xl"
                priority
              />
            </div>
            <div className="text-lg leading-relaxed text-(--fg)/70">
              <p>
                I&apos;m Mitch. I didn&apos;t start in tech. I started in
                science. I studied biomedical science, worked in health IT, then
                moved into tech recruitment where I spent years talking to
                engineers about what they actually do every day.
              </p>
              <p className="mt-6">
                That&apos;s when I decided I wanted to build, not just recruit
                builders. I taught myself to code, eventually became Head of
                Product at a startup, and shipped software across the stack.
              </p>
            </div>
          </div>

          {/* Row 2: Remaining story — full width */}
          <div className="mt-8 text-lg leading-relaxed text-(--fg)/70">
            <p>
              Along the way I used every AI coding tool I could find. They all
              made me faster. None of them made me smarter. I could ship features
              in hours instead of days, but when something broke I still
              didn&apos;t understand why. I was assembling code, not engineering
              it.
            </p>
            <p className="mt-6">
              That&apos;s the gap pear fills. The thesis is simple:{" "}
              <strong className="text-(--fg)">
                AI makes shipping faster, but it doesn&apos;t make engineers better.
              </strong>{" "}
              Speed without understanding is technical debt you carry in your
              head. Pear is a learning engine that remembers what you know,
              detects what you don&apos;t, and teaches you at the moment of
              execution.
            </p>
            <p className="mt-6">
              It doesn&apos;t write code for you. It watches what you build,
              tracks your learning state, and adapts its teaching to your level.
              Think of it as the senior engineer sitting next to you who
              remembers every conversation you&apos;ve ever had, knows exactly
              where your gaps are, and asks the right questions at the right
              time.
            </p>
          </div>

          <div className="mt-12 text-center">
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
