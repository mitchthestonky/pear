import nextDynamic from "next/dynamic";
import { Problem } from "@/components/problem";
import { IndustryData } from "@/components/industry-data";
import { Transition } from "@/components/transition";
import { Solution } from "@/components/solution";
import { WhoItsFor } from "@/components/who-its-for";
import { TeachingDifference } from "@/components/teaching-difference";
import { Pricing } from "@/components/pricing";
import { FinalCTA } from "@/components/final-cta";

export const dynamic = "force-static";

const FloatingSphere = nextDynamic(
  () =>
    import("@/components/ui/floating-sphere").then((m) => ({
      default: m.FloatingSphere,
    }))
);

const Hero = nextDynamic(
  () => import("@/components/hero").then((m) => ({ default: m.Hero })),
  {
    loading: () => (
      <section className="min-h-screen" aria-hidden="true" />
    ),
  }
);

const HowItWorks = nextDynamic(
  () =>
    import("@/components/how-it-works").then((m) => ({
      default: m.HowItWorks,
    }))
);

const jsonLd = {
  "@context": "https://schema.org",
  "@graph": [
    {
      "@type": "Organization",
      "@id": "https://pearcode.dev/#organization",
      name: "pear",
      url: "https://pearcode.dev",
      description:
        "pear is a pair programmer for AI-enabled engineers. Watches your code, surfaces insights, and helps you understand what your AI tools write.",
    },
    {
      "@type": "WebSite",
      "@id": "https://pearcode.dev/#website",
      url: "https://pearcode.dev",
      name: "pear",
      publisher: { "@id": "https://pearcode.dev/#organization" },
    },
    {
      "@type": "SoftwareApplication",
      "@id": "https://pearcode.dev/#app",
      name: "pear",
      applicationCategory: "DeveloperApplication",
      operatingSystem: "macOS",
      description:
        "A pair programmer for AI-enabled engineers. Watches your code, surfaces insights, and helps you understand what your AI tools write.",
      offers: {
        "@type": "Offer",
        price: "20.00",
        priceCurrency: "USD",
        billingIncrement: "month",
      },
      provider: { "@id": "https://pearcode.dev/#organization" },
    },
  ],
};

export default function Home() {
  return (
    <>
      <script
        type="application/ld+json"
        dangerouslySetInnerHTML={{ __html: JSON.stringify(jsonLd) }}
      />
      <FloatingSphere />
      <main className="relative z-10">
        <Hero />
        <Problem />
        <IndustryData />
        <Transition />
        <HowItWorks />
        <Solution />
        <WhoItsFor />
        <TeachingDifference />
        <Pricing />
        <FinalCTA />
      </main>
    </>
  );
}
