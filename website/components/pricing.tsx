import { Check } from "lucide-react";
import { Section } from "@/components/ui/section";
import { ScrollReveal } from "@/components/ui/scroll-reveal";
import { Button } from "@/components/ui/button";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Badge } from "@/components/ui/badge";

const freeFeatures = [
  "Watch mode (auto-reviews changes)",
  "Interactive Q&A with codebase context",
  "BYO any LLM provider",
  "Subtle teaching moments & concept tags",
];

const proFeatures = [
  "Learning state memory",
  "Adaptive learning (pear teach)",
  "Concept tracking & progress",
  "Personal knowledge gap visibility",
  "Cross-machine sync",
];

export function Pricing() {
  return (
    <Section id="pricing">
      <ScrollReveal>
        <h2 className="mb-4 text-center font-(family-name:--font-jetbrains) text-3xl font-bold md:text-4xl">
          Invest in yourself, not another subscription.
        </h2>
        <p className="mx-auto mb-16 max-w-2xl text-center text-lg text-muted-foreground">
          BYOK means you control your LLM costs. Pear handles the learning system.
        </p>
      </ScrollReveal>

      <div className="mx-auto grid max-w-3xl grid-cols-1 gap-8 md:grid-cols-2">
        {/* Free */}
        <ScrollReveal>
          <Card className="flex h-full flex-col">
            <CardHeader>
              <CardTitle className="text-2xl">Free (OSS)</CardTitle>
              <div className="mt-2">
                <span className="text-3xl font-bold sm:text-4xl">$0</span>
                <span className="ml-1 text-muted-foreground">/forever</span>
              </div>
              <p className="mt-1 text-sm text-muted-foreground">
                Free forever with your own API key.
              </p>
              <p className="mt-1 text-xs text-muted-foreground/70">
                Works with Anthropic, OpenAI, and OpenRouter.
              </p>
            </CardHeader>
            <CardContent className="flex flex-1 flex-col">
              <ul className="flex-1 space-y-3">
                {freeFeatures.map((f) => (
                  <li key={f} className="flex items-start gap-2 text-muted-foreground">
                    <Check className="mt-0.5 h-4 w-4 shrink-0 text-muted-foreground" />
                    {f}
                  </li>
                ))}
              </ul>
              <Button asChild variant="outline" className="mt-8 w-full" size="lg">
                <a href="#hero">Join the waitlist</a>
              </Button>
            </CardContent>
          </Card>
        </ScrollReveal>

        {/* Pro */}
        <ScrollReveal delay={100}>
          <Card className="relative flex h-full flex-col border-2 border-pear">
            <Badge className="absolute right-4 top-4 bg-pear-subtle text-pear hover:bg-pear-subtle">
              Most Popular
            </Badge>
            <CardHeader>
              <CardTitle className="text-2xl">Pro</CardTitle>
              <div className="mt-2">
                <span className="text-3xl font-bold sm:text-4xl">$20</span>
                <span className="ml-1 text-muted-foreground">/month</span>
              </div>
              <p className="mt-1 text-sm text-muted-foreground">
                $130/year (4 months free)
              </p>
            </CardHeader>
            <CardContent className="flex flex-1 flex-col">
              <ul className="flex-1 space-y-3">
                {proFeatures.map((f) => (
                  <li key={f} className="flex items-start gap-2">
                    <Check className="mt-0.5 h-4 w-4 shrink-0 text-pear" />
                    {f}
                  </li>
                ))}
              </ul>
              <Button asChild className="mt-8 w-full bg-pear text-white hover:bg-pear-hover" size="lg">
                <a href="#hero">Join the waitlist</a>
              </Button>
            </CardContent>
          </Card>
        </ScrollReveal>

      </div>

      <p className="mt-10 text-center text-sm text-muted-foreground">
        Need team-wide learning metrics?{" "}
        <a href="mailto:mitch@pearcode.dev" className="text-pear underline underline-offset-4 hover:text-pear-hover">
          Get in touch
        </a>{" "}
        about Teams pricing.
      </p>
    </Section>
  );
}
