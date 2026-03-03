import { Section } from "@/components/ui/section";
import { ScrollReveal } from "@/components/ui/scroll-reveal";

const problems = [
  {
    num: "01",
    title: "You\u2019re shipping code you don\u2019t understand.",
    body: "AI tools optimize for output, not understanding. You\u2019re building faster than ever, but the gap between \u2018code that works\u2019 and \u2018code you could explain\u2019 is widening every day.",
  },
  {
    num: "02",
    title: "Nobody has time to teach you anymore.",
    body: "Senior engineers are drowning in AI-generated PRs. The mentorship pipeline is broken. You\u2019re on your own with code you didn\u2019t write.",
  },
  {
    num: "03",
    title: "Learning happens somewhere else. It shouldn\u2019t.",
    body: "Courses, docs, and tutorials exist in a vacuum. By the time you context-switch to learn, the moment is gone. Learning should happen at the point of execution.",
  },
];

export function Problem() {
  return (
    <Section id="problem">
      <h2 className="mb-4 text-center text-3xl font-bold text-(--fg) md:text-4xl font-(family-name:--font-jetbrains)">
        The AI-assisted learning crisis is real.
      </h2>
      <p className="mx-auto mb-16 max-w-xl text-center text-lg text-(--fg)/70">
        AI is making engineers faster. Not better.
      </p>

      <div className="mx-auto max-w-2xl space-y-12">
        {problems.map((problem, i) => (
          <ScrollReveal key={i} delay={i * 100}>
            <div className="flex gap-4 sm:gap-6">
              <span className="shrink-0 font-(family-name:--font-jetbrains) text-2xl font-bold text-pear/70 sm:text-3xl">
                {problem.num}
              </span>
              <div>
                <h3 className="mb-2 font-(family-name:--font-jetbrains) text-lg font-semibold text-(--fg)">
                  {problem.title}
                </h3>
                <p className="text-lg leading-relaxed text-(--fg)/70">
                  {problem.body}
                </p>
              </div>
            </div>
          </ScrollReveal>
        ))}
      </div>
    </Section>
  );
}
