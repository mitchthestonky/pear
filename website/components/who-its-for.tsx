import { Section } from "@/components/ui/section";
import { ScrollReveal } from "@/components/ui/scroll-reveal";

const personas = [
  {
    title: "Developers using AI agents",
    body: "Claude Code, Cursor, OpenCode \u2014 they write your code. Pear makes sure you actually understand it. Stay sharp while your agent ships.",
  },
  {
    title: "Career switchers & self-taught devs",
    body: "You learned by building, not by studying. Pear fills in the fundamentals you skipped \u2014 patterns, trade-offs, and the deeper concepts that bootcamps and tutorials don\u2019t have time for.",
  },
  {
    title: "Founders shipping with AI",
    body: "You\u2019re shipping fast and breaking things. Pear helps you understand the architecture decisions you\u2019re making so you don\u2019t build a house of cards.",
  },
  {
    title: "Senior engineers switching stacks",
    body: "You know how to build systems, just not in this language yet. Pear accelerates your ramp-up on Go, Rust, or whatever you\u2019re picking up next.",
  },
];

export function WhoItsFor() {
  return (
    <Section id="who-its-for">
      <ScrollReveal>
        <h2 className="mb-4 text-center font-(family-name:--font-jetbrains) text-3xl font-bold md:text-4xl">
          Who it&rsquo;s for
        </h2>
        <p className="mx-auto mb-16 max-w-xl text-center text-lg text-(--fg-muted)">
          If you ship with AI but want to understand what you ship, pear is for you.
        </p>
      </ScrollReveal>

      <div className="mx-auto grid max-w-4xl grid-cols-1 gap-6 sm:grid-cols-2">
        {personas.map((persona, i) => (
          <ScrollReveal key={i} delay={i * 60}>
            <div className="flex h-full flex-col rounded-xl border border-border bg-background p-6">
              <h3 className="font-(family-name:--font-jetbrains) text-sm font-bold text-pear">
                {persona.title}
              </h3>
              <p className="mt-3 flex-1 text-sm leading-relaxed text-(--fg)/70">
                {persona.body}
              </p>
            </div>
          </ScrollReveal>
        ))}
      </div>
    </Section>
  );
}
