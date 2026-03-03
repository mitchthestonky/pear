import type { Metadata } from "next";

export const dynamic = "force-static";

export const metadata: Metadata = {
  title: "Providers | pear docs",
  description:
    "Set up Anthropic, OpenAI, or OpenRouter as your LLM provider for pear.",
  alternates: { canonical: "https://pearcode.dev/docs/providers" },
  openGraph: {
    title: "Providers | pear docs",
    description:
      "Set up Anthropic, OpenAI, or OpenRouter as your LLM provider for pear.",
    url: "https://pearcode.dev/docs/providers",
    siteName: "pear",
    type: "website",
  },
};

const jsonLd = {
  "@context": "https://schema.org",
  "@type": "WebPage",
  name: "Providers | pear docs",
  url: "https://pearcode.dev/docs/providers",
};

export default function ProvidersPage() {
  return (
    <>
      <script
        type="application/ld+json"
        dangerouslySetInnerHTML={{ __html: JSON.stringify(jsonLd) }}
      />

      <div className="prose-pear">
        <h1 className="font-(family-name:--font-jetbrains) text-3xl font-bold text-(--fg)">
          Providers
        </h1>
        <p className="mt-4 text-lg text-(--fg)/70">
          Pear is BYOK (Bring Your Own Key). You provide an API key from your
          preferred provider. Your code context is sent directly from your
          machine to the provider — pear never stores your code.
        </p>

        <div className="mt-4 overflow-x-auto">
          <table className="w-full text-sm">
            <thead>
              <tr className="border-b border-border text-left">
                <th className="pb-3 pr-6 font-(family-name:--font-jetbrains) font-semibold text-(--fg)">
                  Provider
                </th>
                <th className="pb-3 pr-6 font-(family-name:--font-jetbrains) font-semibold text-(--fg)">
                  Default Model
                </th>
                <th className="pb-3 font-(family-name:--font-jetbrains) font-semibold text-(--fg)">
                  Env Variable
                </th>
              </tr>
            </thead>
            <tbody className="text-(--fg)/70">
              <tr className="border-b border-border/50">
                <td className="py-3 pr-6 text-(--fg)">Anthropic</td>
                <td className="py-3 pr-6"><code className="text-xs">claude-sonnet-4-20250514</code></td>
                <td className="py-3"><code className="text-xs">ANTHROPIC_API_KEY</code></td>
              </tr>
              <tr className="border-b border-border/50">
                <td className="py-3 pr-6 text-(--fg)">OpenAI</td>
                <td className="py-3 pr-6"><code className="text-xs">gpt-4o</code></td>
                <td className="py-3"><code className="text-xs">OPENAI_API_KEY</code></td>
              </tr>
              <tr className="border-b border-border/50">
                <td className="py-3 pr-6 text-(--fg)">OpenRouter</td>
                <td className="py-3 pr-6"><code className="text-xs">anthropic/claude-sonnet-4-20250514</code></td>
                <td className="py-3"><code className="text-xs">OPENROUTER_API_KEY</code></td>
              </tr>
            </tbody>
          </table>
        </div>

        <h2 className="mt-12 font-(family-name:--font-jetbrains) text-xl font-semibold text-(--fg)">
          Anthropic
        </h2>
        <ol className="mt-4 list-decimal space-y-2 pl-6 text-(--fg)/70">
          <li>
            Go to{" "}
            <a
              href="https://console.anthropic.com/settings/keys"
              target="_blank"
              rel="noopener noreferrer"
              className="text-pear underline underline-offset-4 hover:text-pear-hover"
            >
              console.anthropic.com/settings/keys
            </a>
          </li>
          <li>Create a new API key</li>
          <li>
            Run <code className="rounded bg-[var(--bg-subtle)] px-1.5 py-0.5 text-sm">pear init</code> and
            select Anthropic, then paste your key
          </li>
        </ol>
        <p className="mt-3 text-sm text-(--fg)/50">
          Or set the <code className="text-xs">ANTHROPIC_API_KEY</code> environment
          variable and pear will detect it automatically.
        </p>

        <h2 className="mt-12 font-(family-name:--font-jetbrains) text-xl font-semibold text-(--fg)">
          OpenAI
        </h2>
        <ol className="mt-4 list-decimal space-y-2 pl-6 text-(--fg)/70">
          <li>
            Go to{" "}
            <a
              href="https://platform.openai.com/api-keys"
              target="_blank"
              rel="noopener noreferrer"
              className="text-pear underline underline-offset-4 hover:text-pear-hover"
            >
              platform.openai.com/api-keys
            </a>
          </li>
          <li>Create a new secret key</li>
          <li>
            Run <code className="rounded bg-[var(--bg-subtle)] px-1.5 py-0.5 text-sm">pear init</code> and
            select OpenAI, then paste your key
          </li>
        </ol>
        <p className="mt-3 text-sm text-(--fg)/50">
          Or set the <code className="text-xs">OPENAI_API_KEY</code> environment
          variable.
        </p>

        <h2 className="mt-12 font-(family-name:--font-jetbrains) text-xl font-semibold text-(--fg)">
          OpenRouter
        </h2>
        <ol className="mt-4 list-decimal space-y-2 pl-6 text-(--fg)/70">
          <li>
            Go to{" "}
            <a
              href="https://openrouter.ai/keys"
              target="_blank"
              rel="noopener noreferrer"
              className="text-pear underline underline-offset-4 hover:text-pear-hover"
            >
              openrouter.ai/keys
            </a>
          </li>
          <li>Create a new API key</li>
          <li>
            Run <code className="rounded bg-[var(--bg-subtle)] px-1.5 py-0.5 text-sm">pear init</code> and
            select OpenRouter, then paste your key
          </li>
        </ol>
        <p className="mt-3 text-sm text-(--fg)/50">
          OpenRouter gives you access to hundreds of models from multiple
          providers through a single API key. Change your model anytime in{" "}
          <code className="text-xs">~/.pear/config.toml</code>.
        </p>
      </div>
    </>
  );
}
