import type { Metadata } from "next";
import { Pricing } from "@/components/pricing";

export const dynamic = "force-static";

export const metadata: Metadata = {
  title: "Pricing | pear | Pair Programmer for AI-Enabled Engineers",
  description:
    "Pear Free (OSS) is free forever with your own API key. Pro at $20/month or $130/year. Teams at $30/seat/month. BYOK means you control your LLM costs.",
  alternates: { canonical: "https://pearcode.dev/pricing" },
  openGraph: {
    title: "Pricing | pear",
    description:
      "Free forever (OSS). Pro at $20/month. Teams at $30/seat/month. BYOK — you control your LLM costs.",
    url: "https://pearcode.dev/pricing",
    siteName: "pear",
    type: "website",
  },
};

const jsonLd = {
  "@context": "https://schema.org",
  "@type": "WebPage",
  "@id": "https://pearcode.dev/pricing",
  name: "pear Pricing",
  description:
    "Pear Free (OSS) is free forever. Pro at $20/month or $130/year. Teams at $30/seat/month. BYOK means you control your LLM costs.",
  url: "https://pearcode.dev/pricing",
  isPartOf: { "@id": "https://pearcode.dev/#website" },
};

export default function PricingPage() {
  return (
    <main className="relative z-10 pt-16">
      <script
        type="application/ld+json"
        dangerouslySetInnerHTML={{ __html: JSON.stringify(jsonLd) }}
      />
      <Pricing />
    </main>
  );
}
