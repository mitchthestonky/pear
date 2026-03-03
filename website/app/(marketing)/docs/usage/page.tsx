import type { Metadata } from "next";

export const dynamic = "force-static";

export const metadata: Metadata = {
  title: "Workflows | pear docs",
  description:
    "Learn how to use pear's core workflows: watch, ask, review, and teach.",
  alternates: { canonical: "https://pearcode.dev/docs/usage" },
  openGraph: {
    title: "Workflows | pear docs",
    description:
      "Learn how to use pear's core workflows: watch, ask, review, and teach.",
    url: "https://pearcode.dev/docs/usage",
    siteName: "pear",
    type: "website",
  },
};

const jsonLd = {
  "@context": "https://schema.org",
  "@type": "WebPage",
  name: "Workflows | pear docs",
  url: "https://pearcode.dev/docs/usage",
};

export default function UsagePage() {
  return (
    <>
      <script
        type="application/ld+json"
        dangerouslySetInnerHTML={{ __html: JSON.stringify(jsonLd) }}
      />

      <div className="prose-pear">
        <h1 className="font-(family-name:--font-jetbrains) text-3xl font-bold text-(--fg)">
          Workflows
        </h1>
        <p className="mt-4 text-lg text-(--fg)/70">
          Pear has four core workflows. Each one reads your codebase context
          (git diff, file tree, recent changes) and responds with teaching
          intent.
        </p>

        <h2 className="mt-12 font-(family-name:--font-jetbrains) text-xl font-semibold text-(--fg)">
          Watch
        </h2>
        <pre className="mt-2 overflow-x-auto rounded-lg bg-[var(--bg-subtle)] p-4 text-sm text-(--fg)">
          <code>pear watch</code>
        </pre>
        <p className="mt-4 text-(--fg)/70">
          The flagship mode. Pear opens an interactive TUI, watches your files
          with fsnotify, and monitors your git diff. When you pause — while
          your agent thinks, after a save, between tasks — Pear reviews what
          changed and teaches you something relevant.
        </p>
        <p className="mt-3 text-(--fg)/70">
          Watch mode is designed to run alongside your editor and AI coding
          tools. It never interrupts; it waits for natural pauses in your
          workflow.
        </p>

        <h2 className="mt-12 font-(family-name:--font-jetbrains) text-xl font-semibold text-(--fg)">
          Ask
        </h2>
        <pre className="mt-2 overflow-x-auto rounded-lg bg-[var(--bg-subtle)] p-4 text-sm text-(--fg)">
          <code>pear ask &quot;what does this middleware do?&quot;</code>
        </pre>
        <p className="mt-4 text-(--fg)/70">
          Ask a question about your codebase. Pear reads your current context
          and answers with teaching intent — it explains concepts, not just
          facts. Use <code className="rounded bg-[var(--bg-subtle)] px-1.5 py-0.5 text-sm">@file</code> to
          include specific files as context.
        </p>

        <h2 className="mt-12 font-(family-name:--font-jetbrains) text-xl font-semibold text-(--fg)">
          Review
        </h2>
        <pre className="mt-2 overflow-x-auto rounded-lg bg-[var(--bg-subtle)] p-4 text-sm text-(--fg)">
          <code>pear review</code>
        </pre>
        <p className="mt-4 text-(--fg)/70">
          Review your uncommitted changes. Pear analyzes your current diff and
          gives feedback focused on learning — not just &ldquo;this is wrong&rdquo; but
          &ldquo;here&apos;s why this pattern matters and what to watch for.&rdquo;
        </p>

        <h2 className="mt-12 font-(family-name:--font-jetbrains) text-xl font-semibold text-(--fg)">
          Teach
        </h2>
        <pre className="mt-2 overflow-x-auto rounded-lg bg-[var(--bg-subtle)] p-4 text-sm text-(--fg)">
          <code>{`pear teach              # Pear picks a topic
pear teach "goroutines" # You pick a topic`}</code>
        </pre>
        <p className="mt-4 text-(--fg)/70">
          Deep-dive teaching grounded in your actual code. Without a topic,
          Pear picks one based on your recent changes. With a topic, it
          teaches that concept using examples from your codebase.
        </p>
      </div>
    </>
  );
}
