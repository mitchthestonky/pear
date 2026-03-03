import type { Metadata } from "next";

export const dynamic = "force-static";

export const metadata: Metadata = {
  title: "Configuration | pear docs",
  description:
    "Configure pear — config file structure, settings, and per-repo overrides.",
  alternates: { canonical: "https://pearcode.dev/docs/configuration" },
  openGraph: {
    title: "Configuration | pear docs",
    description:
      "Configure pear — config file structure, settings, and per-repo overrides.",
    url: "https://pearcode.dev/docs/configuration",
    siteName: "pear",
    type: "website",
  },
};

const jsonLd = {
  "@context": "https://schema.org",
  "@type": "WebPage",
  name: "Configuration | pear docs",
  url: "https://pearcode.dev/docs/configuration",
};

export default function ConfigurationPage() {
  return (
    <>
      <script
        type="application/ld+json"
        dangerouslySetInnerHTML={{ __html: JSON.stringify(jsonLd) }}
      />

      <div className="prose-pear">
        <h1 className="font-(family-name:--font-jetbrains) text-3xl font-bold text-(--fg)">
          Configuration
        </h1>
        <p className="mt-4 text-lg text-(--fg)/70">
          All pear configuration lives in <code className="rounded bg-[var(--bg-subtle)] px-1.5 py-0.5 text-sm">~/.pear/</code>.
          Nothing is stored in your project repository except optional git hooks.
        </p>

        <h2 className="mt-12 font-(family-name:--font-jetbrains) text-xl font-semibold text-(--fg)">
          Directory structure
        </h2>
        <pre className="mt-4 overflow-x-auto rounded-lg bg-[var(--bg-subtle)] p-4 text-sm text-(--fg)">
          <code>{`~/.pear/
├── config.toml          # Main configuration
├── learning.json        # Learning progress and concept tracking
├── codebases/<slug>.toml # Per-repo overrides
└── logs/<timestamp>.log  # Session logs`}</code>
        </pre>

        <h2 className="mt-12 font-(family-name:--font-jetbrains) text-xl font-semibold text-(--fg)">
          config.toml
        </h2>
        <p className="mt-4 text-(--fg)/70">
          Created by <code className="rounded bg-[var(--bg-subtle)] px-1.5 py-0.5 text-sm">pear init</code>.
          You can also edit it directly.
        </p>
        <pre className="mt-4 overflow-x-auto rounded-lg bg-[var(--bg-subtle)] p-4 text-sm text-(--fg)">
          <code>{`name = "Mitch"
languages = "go, typescript"
level = "intermediate"

[provider]
active = "anthropic"

[provider.anthropic]
api_key = "sk-ant-..."
model = "claude-sonnet-4-20250514"

[provider.openai]
api_key = ""
model = "gpt-4o"

[provider.openrouter]
api_key = ""
model = "anthropic/claude-sonnet-4-20250514"

[watch]
# Pause detection threshold in seconds
pause_seconds = 30`}</code>
        </pre>

        <h2 className="mt-12 font-(family-name:--font-jetbrains) text-xl font-semibold text-(--fg)">
          Settings reference
        </h2>
        <div className="mt-4 overflow-x-auto">
          <table className="w-full text-sm">
            <thead>
              <tr className="border-b border-border text-left">
                <th className="pb-3 pr-6 font-(family-name:--font-jetbrains) font-semibold text-(--fg)">
                  Key
                </th>
                <th className="pb-3 pr-6 font-(family-name:--font-jetbrains) font-semibold text-(--fg)">
                  Type
                </th>
                <th className="pb-3 font-(family-name:--font-jetbrains) font-semibold text-(--fg)">
                  Description
                </th>
              </tr>
            </thead>
            <tbody className="text-(--fg)/70">
              <tr className="border-b border-border/50">
                <td className="py-3 pr-6"><code className="text-(--fg) text-xs">name</code></td>
                <td className="py-3 pr-6">string</td>
                <td className="py-3">Your display name</td>
              </tr>
              <tr className="border-b border-border/50">
                <td className="py-3 pr-6"><code className="text-(--fg) text-xs">languages</code></td>
                <td className="py-3 pr-6">string</td>
                <td className="py-3">Comma-separated languages you use</td>
              </tr>
              <tr className="border-b border-border/50">
                <td className="py-3 pr-6"><code className="text-(--fg) text-xs">level</code></td>
                <td className="py-3 pr-6">string</td>
                <td className="py-3">Experience level: beginner, intermediate, advanced</td>
              </tr>
              <tr className="border-b border-border/50">
                <td className="py-3 pr-6"><code className="text-(--fg) text-xs">provider.active</code></td>
                <td className="py-3 pr-6">string</td>
                <td className="py-3">Active provider: anthropic, openai, or openrouter</td>
              </tr>
              <tr className="border-b border-border/50">
                <td className="py-3 pr-6"><code className="text-(--fg) text-xs">provider.&lt;name&gt;.api_key</code></td>
                <td className="py-3 pr-6">string</td>
                <td className="py-3">API key for the provider</td>
              </tr>
              <tr className="border-b border-border/50">
                <td className="py-3 pr-6"><code className="text-(--fg) text-xs">provider.&lt;name&gt;.model</code></td>
                <td className="py-3 pr-6">string</td>
                <td className="py-3">Model identifier</td>
              </tr>
              <tr className="border-b border-border/50">
                <td className="py-3 pr-6"><code className="text-(--fg) text-xs">watch.pause_seconds</code></td>
                <td className="py-3 pr-6">int</td>
                <td className="py-3">Seconds of inactivity before triggering a review</td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>
    </>
  );
}
