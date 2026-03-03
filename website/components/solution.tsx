import { Section } from "@/components/ui/section";
import { ScrollReveal } from "@/components/ui/scroll-reveal";

const pillars = [
  {
    num: "01",
    title: "Learning state memory",
    tagline: "Pear remembers what you understand.",
    body: "Every session builds on the last. It never explains the same thing twice, and never skips what you need.",
  },
  {
    num: "02",
    title: "Concept graph",
    tagline: "Pear maps how concepts connect.",
    body: "A structured graph of engineering concepts and their prerequisites. When you hit something new, pear knows what you already have.",
  },
  {
    num: "03",
    title: "Adaptive pedagogy",
    tagline: "Pear changes how it teaches you.",
    body: "Explanations adapt based on your behaviour - how you respond, what you skip, what you revisit.",
  },
  {
    num: "04",
    title: "Intervention timing",
    tagline: "Pear teaches at the right moment.",
    body: "Insights surface in the gaps - while your agent thinks, before you accept a plan, after a diff lands.",
  },
  {
    num: "05",
    title: "Skill progression",
    tagline: "Pear shows you how far you\u2019ve come.",
    body: "Track concepts mastered, gaps closed, and patterns learned across languages and frameworks.",
  },
];

export function Solution() {
  return (
    <Section id="solution">
      <ScrollReveal>
        <h2 className="mb-4 text-center text-3xl font-bold text-(--fg) md:text-4xl font-(family-name:--font-jetbrains)">
          Five things no AI tool does.
        </h2>
        <p className="mx-auto mb-16 max-w-xl text-center text-lg text-(--fg)/70">
          Any AI tool can explain code. Only pear builds a learning system around you.
        </p>
      </ScrollReveal>

      <div className="mx-auto max-w-2xl space-y-12">
        {pillars.map((pillar, i) => (
          <ScrollReveal key={i} delay={i * 100}>
            <div className="flex gap-4 sm:gap-6">
              <span className="shrink-0 font-(family-name:--font-jetbrains) text-2xl font-bold text-pear/70 sm:text-3xl">
                {pillar.num}
              </span>
              <div>
                <h3 className="font-(family-name:--font-jetbrains) text-lg font-semibold text-(--fg)">
                  {pillar.title}
                </h3>
                <p className="mt-1 text-sm font-medium text-pear">
                  {pillar.tagline}
                </p>
                <p className="mt-2 text-lg leading-relaxed text-(--fg)/70">
                  {pillar.body}
                </p>
              </div>
            </div>
          </ScrollReveal>
        ))}
      </div>
    </Section>
  );
}
