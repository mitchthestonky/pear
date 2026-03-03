"use client";

import { useRef, useState, useEffect } from "react";
import { Section } from "@/components/ui/section";
import { ScrollReveal } from "@/components/ui/scroll-reveal";
import { TerminalMockup } from "@/components/ui/terminal-mockup";

export function HowItWorks() {
  const ref = useRef<HTMLDivElement>(null);
  const [visible, setVisible] = useState(false);

  useEffect(() => {
    if (!ref.current) return;
    const observer = new IntersectionObserver(
      ([entry]) => {
        if (entry.isIntersecting) setVisible(true);
      },
      { threshold: 0.2 }
    );
    observer.observe(ref.current);
    return () => observer.disconnect();
  }, []);

  return (
    <Section id="how-it-works">
      <ScrollReveal>
        <h2 className="mb-10 text-center font-(family-name:--font-jetbrains) text-3xl font-bold md:text-4xl">
          You code. Pear watches. You understand.
        </h2>
      </ScrollReveal>

      <div ref={ref}>
        <ScrollReveal>
          <TerminalMockup isVisible={visible} />
        </ScrollReveal>
      </div>

      <div className="mx-auto mt-16 grid max-w-3xl grid-cols-1 gap-8 sm:grid-cols-3">
        <ScrollReveal delay={0}>
          <div className="text-center">
            <span className="inline-flex h-10 w-10 items-center justify-center rounded-full bg-pear/10 font-(family-name:--font-jetbrains) text-sm font-bold text-pear">
              1
            </span>
            <h3 className="mt-4 font-(family-name:--font-jetbrains) text-base font-semibold text-(--fg)">
              Run pear alongside your agent
            </h3>
            <p className="mt-2 text-base leading-relaxed text-(--fg)/60">
              Start <code className="text-pear/80">pear watch</code> in your terminal. It reads your diffs, file tree, and recent changes in real time.
            </p>
          </div>
        </ScrollReveal>
        <ScrollReveal delay={100}>
          <div className="text-center">
            <span className="inline-flex h-10 w-10 items-center justify-center rounded-full bg-pear/10 font-(family-name:--font-jetbrains) text-sm font-bold text-pear">
              2
            </span>
            <h3 className="mt-4 font-(family-name:--font-jetbrains) text-base font-semibold text-(--fg)">
              Pear reviews and teaches
            </h3>
            <p className="mt-2 text-base leading-relaxed text-(--fg)/60">
              It auto-reviews changes, tags concepts, and surfaces explanations in the gaps between actions.
            </p>
          </div>
        </ScrollReveal>
        <ScrollReveal delay={200}>
          <div className="text-center">
            <span className="inline-flex h-10 w-10 items-center justify-center rounded-full bg-pear/10 font-(family-name:--font-jetbrains) text-sm font-bold text-pear">
              3
            </span>
            <h3 className="mt-4 font-(family-name:--font-jetbrains) text-base font-semibold text-(--fg)">
              Pro builds your learning profile
            </h3>
            <p className="mt-2 text-base leading-relaxed text-(--fg)/60">
              Over time, pear remembers what you understand, tracks gaps, and adapts its teaching to your level.
            </p>
          </div>
        </ScrollReveal>
      </div>
    </Section>
  );
}
