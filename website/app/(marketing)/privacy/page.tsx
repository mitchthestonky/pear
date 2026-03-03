import type { Metadata } from "next";
import { PolicyLayout } from "@/components/policy-layout";

export const dynamic = "force-static";

export const metadata: Metadata = {
  title: "Privacy Policy | pear",
  description:
    "How pear handles your data, code context, and API keys. BYOK-first privacy.",
  alternates: { canonical: "https://pearcode.dev/privacy" },
  openGraph: {
    title: "Privacy Policy | pear",
    description:
      "How pear handles your data, code context, and API keys. BYOK-first privacy.",
    url: "https://pearcode.dev/privacy",
    siteName: "pear",
    type: "website",
  },
};

export default function PrivacyPage() {
  return (
    <PolicyLayout title="Privacy Policy" lastUpdated="February 11, 2026">
      <p className="text-base! italic border-l-2 border-pear/30 pl-4">
        This document outlines the best-practice structure for pear&apos;s
        privacy policy. It is pending legal review before becoming binding.
      </p>

      <section>
        <h2 className="font-(family-name:--font-jetbrains) text-xl font-bold text-(--fg)">1. Information We Collect</h2>
        <p className="mt-4">We collect minimal information to provide the Service:</p>
        <ul className="mt-3 list-disc space-y-2 pl-6">
          <li>
            <strong className="text-(--fg)">Account information:</strong> email address when you sign up
            for the waitlist or create an account.
          </li>
          <li>
            <strong className="text-(--fg)">Usage analytics:</strong> anonymous, aggregated data about
            feature usage to improve the product (via Vercel Analytics).
          </li>
        </ul>
      </section>

      <section>
        <h2 className="font-(family-name:--font-jetbrains) text-xl font-bold text-(--fg)">2. How We Use Your Information</h2>
        <p className="mt-4">
          We use collected information to operate and improve the Service,
          communicate product updates, and provide customer support.
        </p>
      </section>

      <section>
        <h2 className="font-(family-name:--font-jetbrains) text-xl font-bold text-(--fg)">3. BYOK and Code Context</h2>
        <p className="mt-4">
          pear operates on a BYOK (Bring Your Own Key) model. When you use pear,
          code context (git diffs, file structure, recent changes) is sent{" "}
          <strong className="text-(--fg)">directly from your machine to your chosen LLM provider</strong>
          . This data does not pass through pear&apos;s servers. Your API keys
          are stored locally on your machine and are never transmitted to us.
        </p>
      </section>

      <section>
        <h2 className="font-(family-name:--font-jetbrains) text-xl font-bold text-(--fg)">4. Data Storage</h2>
        <p className="mt-4">
          Learning profiles and session history are stored locally on your
          machine. We do not store your code, prompts, or LLM responses on our
          servers.
        </p>
      </section>

      <section>
        <h2 className="font-(family-name:--font-jetbrains) text-xl font-bold text-(--fg)">5. Third Parties</h2>
        <p className="mt-4">We use the following third-party services:</p>
        <ul className="mt-3 list-disc space-y-2 pl-6">
          <li>
            <strong className="text-(--fg)">Vercel:</strong> hosting and analytics
          </li>
          <li>
            <strong className="text-(--fg)">Stripe:</strong> payment processing (when billing is active)
          </li>
        </ul>
        <p className="mt-4">
          Your use of third-party LLM providers (Anthropic, OpenAI, Google) is
          governed by their respective privacy policies.
        </p>
      </section>

      <section>
        <h2 className="font-(family-name:--font-jetbrains) text-xl font-bold text-(--fg)">6. Cookies</h2>
        <p className="mt-4">
          We use essential cookies only (theme preference). We do not use
          tracking cookies or sell data to advertisers.
        </p>
      </section>

      <section>
        <h2 className="font-(family-name:--font-jetbrains) text-xl font-bold text-(--fg)">7. Your Rights</h2>
        <p className="mt-4">
          You may request access to, correction of, or deletion of your personal
          data at any time by contacting us. Since most data is stored locally,
          you can delete it directly from your machine.
        </p>
      </section>

      <section>
        <h2 className="font-(family-name:--font-jetbrains) text-xl font-bold text-(--fg)">8. Children&apos;s Privacy</h2>
        <p className="mt-4">
          pear is not directed at children under 13. We do not knowingly collect
          personal information from children.
        </p>
      </section>

      <section>
        <h2 className="font-(family-name:--font-jetbrains) text-xl font-bold text-(--fg)">9. Changes to This Policy</h2>
        <p className="mt-4">
          We may update this Privacy Policy from time to time. Changes will be
          posted on this page with an updated date.
        </p>
      </section>

      <section>
        <h2 className="font-(family-name:--font-jetbrains) text-xl font-bold text-(--fg)">10. Contact</h2>
        <p className="mt-4">
          Questions about privacy? Email us at{" "}
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
