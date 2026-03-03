import { Section } from "@/components/ui/section";
import { ScrollReveal } from "@/components/ui/scroll-reveal";

const stats = [
  {
    number: "19%",
    label: "slower with AI",
    detail:
      "A randomized controlled trial found developers are 19% slower with AI tools, yet still believe they sped up by 20%.",
    source: "METR, 2025",
    icon: "/logos/metr.png",
    url: "https://metr.org/blog/2025-07-10-early-2025-ai-experienced-os-dev-study/",
  },
  {
    number: "67%",
    label: "drop in junior roles",
    detail:
      "Junior software engineering job postings have collapsed, with the learning curve itself being automated away.",
    source: "Stanford Digital Economy Lab",
    icon: "/logos/stanford.ico",
    url: "https://www.sundeepteki.org/advice/impact-of-ai-on-the-2025-software-engineering-job-market",
  },
  {
    number: "66%",
    label: "fixing 'almost-right' code",
    detail:
      "Two-thirds of developers say they spend more time fixing AI-generated code that looks correct but isn't.",
    source: "Stack Overflow, 2025",
    icon: "/logos/stackoverflow.ico",
    url: "https://stackoverflow.blog/2025/12/29/developers-remain-willing-but-reluctant-to-use-ai-the-2025-developer-survey-results-are-here",
  },
  {
    number: "4x",
    label: "more code cloning",
    detail:
      "AI-assisted coding has led to 4x more code duplication. Developers are pasting more than they're understanding.",
    source: "GitClear analysis",
    icon: "/logos/gitclear.ico",
    url: "https://www.gitclear.com/coding_on_copilot_data_shows_ais_downward_pressure_on_code_quality",
  },
];

export function IndustryData() {
  return (
    <Section id="data">
      <ScrollReveal>
        <p className="mb-4 text-center text-sm font-semibold uppercase tracking-wider text-pear">
          The data is clear
        </p>
        <h2 className="mb-16 text-center font-(family-name:--font-jetbrains) text-3xl font-bold md:text-4xl">
          Engineering craft is in decline.
        </h2>
      </ScrollReveal>

      <div className="mx-auto grid max-w-4xl grid-cols-1 gap-8 sm:grid-cols-2">
        {stats.map((stat, i) => (
          <ScrollReveal key={i} delay={i * 80} className="flex">
            <div className="flex flex-1 flex-col rounded-xl border border-border bg-background p-6">
              <p className="font-(family-name:--font-jetbrains) text-4xl font-bold text-pear">
                {stat.number}
              </p>
              <p className="mt-1 font-(family-name:--font-jetbrains) text-sm font-semibold text-(--fg)">
                {stat.label}
              </p>
              <p className="mt-3 flex-1 text-sm leading-relaxed text-(--fg)/70">
                {stat.detail}
              </p>
              <a
                href={stat.url}
                target="_blank"
                rel="noopener noreferrer"
                className="mt-4 flex items-center gap-2 text-xs text-(--fg)/40 transition-colors hover:text-(--fg)/60"
              >
                {/* eslint-disable-next-line @next/next/no-img-element */}
                <img
                  src={stat.icon}
                  alt=""
                  width={14}
                  height={14}
                  className="rounded-sm"
                />
                {stat.source}
              </a>
            </div>
          </ScrollReveal>
        ))}
      </div>

      <ScrollReveal delay={400}>
        <p className="mt-16 text-center font-(family-name:--font-jetbrains) text-xl font-semibold text-(--fg)/80 italic sm:text-2xl">
          What if your tools didn&rsquo;t just write code &mdash; but made sure you understood it?
        </p>
      </ScrollReveal>
    </Section>
  );
}
