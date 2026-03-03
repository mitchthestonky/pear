import { ScrollReveal } from "@/components/ui/scroll-reveal";

export function Transition() {
  return (
    <section className="flex items-center justify-center px-6 py-20 md:py-32">
      <ScrollReveal>
        <h2 className="text-center font-(family-name:--font-jetbrains) text-2xl font-bold leading-tight sm:text-4xl md:text-6xl">
          Introducing <span className="text-pear">pear</span>.
          <br />
          <span className="text-lg md:text-2xl text-(--fg)/60 font-normal">The learning system that makes sure you understand what AI writes.</span>
        </h2>
      </ScrollReveal>
    </section>
  );
}
