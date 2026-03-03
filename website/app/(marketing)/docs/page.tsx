import type { Metadata } from "next";

export const dynamic = "force-static";

export const metadata: Metadata = {
  title: "Docs | pear",
  description:
    "Get started with pear — install, configure your API key, and run your first session.",
  alternates: { canonical: "https://pearcode.dev/docs" },
  openGraph: {
    title: "Docs | pear",
    description:
      "Get started with pear — install, configure your API key, and run your first session.",
    url: "https://pearcode.dev/docs",
    siteName: "pear",
    type: "website",
  },
};

const jsonLd = {
  "@context": "https://schema.org",
  "@type": "WebPage",
  name: "Docs | pear",
  description:
    "Get started with pear — install, configure your API key, and run your first session.",
  url: "https://pearcode.dev/docs",
};

export default function DocsPage() {
  return (
    <>
      <script
        type="application/ld+json"
        dangerouslySetInnerHTML={{ __html: JSON.stringify(jsonLd) }}
      />

      <div className="prose-pear">
        <h1 className="font-(family-name:--font-jetbrains) text-3xl font-bold text-(--fg)">
          Getting Started
        </h1>
        <p className="mt-4 text-lg text-(--fg)/70">
          pear is a CLI teaching tool that watches you code and proactively
          teaches during natural pauses. It doesn&apos;t write code for you — it
          helps you understand what you&apos;re writing and why.
        </p>

        <h2 className="mt-12 font-(family-name:--font-jetbrains) text-xl font-semibold text-(--fg)">
          Install
        </h2>

        <h3 className="mt-6 font-(family-name:--font-jetbrains) text-base font-semibold text-(--fg)">
          Go install
        </h3>
        <pre className="mt-2 overflow-x-auto rounded-lg bg-[var(--bg-subtle)] p-4 text-sm text-(--fg)">
          <code>go install github.com/MitchTheStonky/pear/cli@latest</code>
        </pre>

        <h3 className="mt-6 font-(family-name:--font-jetbrains) text-base font-semibold text-(--fg)">
          Homebrew
        </h3>
        <pre className="mt-2 overflow-x-auto rounded-lg bg-[var(--bg-subtle)] p-4 text-sm text-(--fg)">
          <code>brew install MitchTheStonky/pear/pear</code>
        </pre>

        <h3 className="mt-6 font-(family-name:--font-jetbrains) text-base font-semibold text-(--fg)">
          curl
        </h3>
        <pre className="mt-2 overflow-x-auto rounded-lg bg-[var(--bg-subtle)] p-4 text-sm text-(--fg)">
          <code>curl -fsSL https://raw.githubusercontent.com/MitchTheStonky/pear/main/install.sh | sh</code>
        </pre>

        <h2 className="mt-12 font-(family-name:--font-jetbrains) text-xl font-semibold text-(--fg)">
          Setup
        </h2>
        <p className="mt-4 text-(--fg)/70">
          Run the init wizard to configure your name, preferred languages,
          experience level, and API key:
        </p>
        <pre className="mt-2 overflow-x-auto rounded-lg bg-[var(--bg-subtle)] p-4 text-sm text-(--fg)">
          <code>pear init</code>
        </pre>
        <p className="mt-4 text-(--fg)/70">
          This creates <code className="rounded bg-[var(--bg-subtle)] px-1.5 py-0.5 text-sm">~/.pear/config.toml</code> with
          your settings. You can re-run <code className="rounded bg-[var(--bg-subtle)] px-1.5 py-0.5 text-sm">pear init</code> anytime
          to reconfigure.
        </p>

        <h2 className="mt-12 font-(family-name:--font-jetbrains) text-xl font-semibold text-(--fg)">
          Your first session
        </h2>
        <p className="mt-4 text-(--fg)/70">
          Navigate to a project directory and start watching:
        </p>
        <pre className="mt-2 overflow-x-auto rounded-lg bg-[var(--bg-subtle)] p-4 text-sm text-(--fg)">
          <code>{`cd /path/to/your/project
pear watch`}</code>
        </pre>
        <p className="mt-4 text-(--fg)/70">
          Pear opens an interactive TUI. Start coding — when you pause, Pear
          reviews your recent changes and teaches you something relevant.
          Type a question directly or use slash commands like{" "}
          <code className="rounded bg-[var(--bg-subtle)] px-1.5 py-0.5 text-sm">/review</code> and{" "}
          <code className="rounded bg-[var(--bg-subtle)] px-1.5 py-0.5 text-sm">/help</code>.
        </p>

        <h2 className="mt-12 font-(family-name:--font-jetbrains) text-xl font-semibold text-(--fg)">
          Verify your setup
        </h2>
        <p className="mt-4 text-(--fg)/70">
          Run the doctor command to check your config, API key, and provider
          connectivity:
        </p>
        <pre className="mt-2 overflow-x-auto rounded-lg bg-[var(--bg-subtle)] p-4 text-sm text-(--fg)">
          <code>pear doctor</code>
        </pre>
      </div>
    </>
  );
}
