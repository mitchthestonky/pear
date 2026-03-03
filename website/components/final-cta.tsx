import { Section } from "@/components/ui/section";
import { ScrollReveal } from "@/components/ui/scroll-reveal";
import { WaitlistForm } from "@/components/ui/waitlist-form";

export function FinalCTA() {
  return (
    <Section id="cta">
      <ScrollReveal>
        <div className="mx-auto max-w-2xl text-center">
          <h2 className="mb-6 font-(family-name:--font-jetbrains) text-3xl font-bold md:text-4xl">

            Don&rsquo;t let AI make you a worse engineer.{" "}
            Start learning with <span className="text-pear">pear</span>.
          </h2>
          <p className="mb-8 text-lg text-muted-foreground">
            Free and open source. Pro from $20/mo.
          </p>
          <div className="mx-auto max-w-md">
            <WaitlistForm />
          </div>
          <p className="mt-10 text-sm text-muted-foreground italic">
            &ldquo;I taught myself to code, shipped with AI, and realized I was getting faster without getting smarter. So I built pear.&rdquo;
            <span className="mt-2 block not-italic text-muted-foreground/70">&mdash; Mitch, founder</span>
          </p>
        </div>
      </ScrollReveal>
    </Section>
  );
}
