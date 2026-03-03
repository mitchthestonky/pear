import type { Metadata } from "next";
import { PolicyLayout } from "@/components/policy-layout";

export const dynamic = "force-static";

export const metadata: Metadata = {
  title: "Terms & Conditions | pear",
  description: "Terms and conditions for using pear, the AI coding tutor.",
  alternates: { canonical: "https://pearcode.dev/terms" },
  openGraph: {
    title: "Terms & Conditions | pear",
    description: "Terms and conditions for using pear, the AI coding tutor.",
    url: "https://pearcode.dev/terms",
    siteName: "pear",
    type: "website",
  },
};

export default function TermsPage() {
  return (
    <PolicyLayout title="Terms & Conditions" lastUpdated="February 11, 2026">
      <p className="text-base! italic border-l-2 border-pear/30 pl-4">
        This document outlines the best-practice structure for pear&apos;s terms
        of service. It is pending legal review before becoming binding.
      </p>

      <section>
        <h2 className="font-(family-name:--font-jetbrains) text-xl font-bold text-(--fg)">1. Acceptance of Terms</h2>
        <p className="mt-4">
          By accessing or using pear (&quot;the Service&quot;), you agree to be
          bound by these Terms &amp; Conditions. If you do not agree, do not use
          the Service.
        </p>
      </section>

      <section>
        <h2 className="font-(family-name:--font-jetbrains) text-xl font-bold text-(--fg)">2. Description of Service</h2>
        <p className="mt-4">
          pear is a command-line AI coding tutor that analyses your coding
          context and provides educational explanations using third-party large
          language model providers. The Service operates on a BYOK (Bring Your
          Own Key) model.
        </p>
      </section>

      <section>
        <h2 className="font-(family-name:--font-jetbrains) text-xl font-bold text-(--fg)">3. BYOK and API Keys</h2>
        <p className="mt-4">
          You are responsible for obtaining and managing your own API keys from
          supported LLM providers. pear does not store, transmit, or have access
          to your API keys beyond your local machine. Usage of third-party APIs
          is subject to their respective terms of service.
        </p>
      </section>

      <section>
        <h2 className="font-(family-name:--font-jetbrains) text-xl font-bold text-(--fg)">4. User Responsibilities</h2>
        <p className="mt-4">You agree to:</p>
        <ul className="mt-3 list-disc space-y-2 pl-6">
          <li>Use the Service only for lawful purposes</li>
          <li>Not attempt to reverse-engineer, decompile, or disassemble the software</li>
          <li>Not redistribute, sublicense, or resell access to the Service</li>
          <li>Maintain the security of your account credentials and API keys</li>
        </ul>
      </section>

      <section>
        <h2 className="font-(family-name:--font-jetbrains) text-xl font-bold text-(--fg)">5. Intellectual Property</h2>
        <p className="mt-4">
          pear and its original content, features, and functionality are owned
          by pear and are protected by applicable intellectual property laws.
          Your code remains yours. pear claims no ownership of code analysed
          during sessions.
        </p>
      </section>

      <section>
        <h2 className="font-(family-name:--font-jetbrains) text-xl font-bold text-(--fg)">6. Limitation of Liability</h2>
        <p className="mt-4">
          pear is provided &quot;as is&quot; without warranty of any kind. We are
          not liable for any damages arising from the use of the Service,
          including but not limited to errors in AI-generated explanations, data
          loss, or interruptions of service.
        </p>
      </section>

      <section>
        <h2 className="font-(family-name:--font-jetbrains) text-xl font-bold text-(--fg)">7. Subscription and Billing</h2>
        <p className="mt-4">
          Pro subscriptions are billed monthly ($20/month) or annually
          ($200/year). You may cancel at any time. Refunds are handled on a
          case-by-case basis.
        </p>
      </section>

      <section>
        <h2 className="font-(family-name:--font-jetbrains) text-xl font-bold text-(--fg)">8. Termination</h2>
        <p className="mt-4">
          We may suspend or terminate your access to the Service at our
          discretion, with or without notice, for conduct that we believe
          violates these Terms or is harmful to other users or the Service.
        </p>
      </section>

      <section>
        <h2 className="font-(family-name:--font-jetbrains) text-xl font-bold text-(--fg)">9. Changes to Terms</h2>
        <p className="mt-4">
          We reserve the right to modify these Terms at any time. Continued use
          of the Service after changes constitutes acceptance of the new Terms.
        </p>
      </section>

      <section>
        <h2 className="font-(family-name:--font-jetbrains) text-xl font-bold text-(--fg)">10. Contact</h2>
        <p className="mt-4">
          Questions about these Terms? Email us at{" "}
          <a
            href="mailto:mitch@pearcode.dev"
            className="text-pear underline underline-offset-4 hover:text-pear-hover"
          >
            mitch@pearcode.dev
          </a>
          .
        </p>
      </section>
    </PolicyLayout>
  );
}
