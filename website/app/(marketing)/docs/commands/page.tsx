import type { Metadata } from "next";

export const dynamic = "force-static";

export const metadata: Metadata = {
  title: "Commands | pear docs",
  description:
    "Full command reference for the pear CLI and TUI slash commands.",
  alternates: { canonical: "https://pearcode.dev/docs/commands" },
  openGraph: {
    title: "Commands | pear docs",
    description:
      "Full command reference for the pear CLI and TUI slash commands.",
    url: "https://pearcode.dev/docs/commands",
    siteName: "pear",
    type: "website",
  },
};

const jsonLd = {
  "@context": "https://schema.org",
  "@type": "WebPage",
  name: "Commands | pear docs",
  url: "https://pearcode.dev/docs/commands",
};

const cliCommands = [
  { cmd: "pear init", desc: "Initialize configuration — name, languages, level, API key" },
  { cmd: "pear watch", desc: "Watch files and teach proactively in an interactive TUI" },
  { cmd: 'pear ask "question"', desc: "Ask a question with full codebase context" },
  { cmd: "pear review", desc: "Review uncommitted changes with teaching feedback" },
  { cmd: "pear teach [topic]", desc: "Deep-dive teaching on a topic (optional)" },
  { cmd: "pear doctor", desc: "Check config, API keys, and provider connectivity" },
  { cmd: "pear hooks install", desc: "Install post-commit git hook" },
  { cmd: "pear hooks uninstall", desc: "Remove post-commit git hook" },
  { cmd: "pear progress", desc: "Show learning progress across sessions" },
];

const tuiCommands = [
  { cmd: "/help", desc: "Show all available commands" },
  { cmd: "/watch", desc: "Start the file watcher" },
  { cmd: "/review", desc: "Review current changes" },
  { cmd: "/settings", desc: "Configure provider and model" },
  { cmd: "/status", desc: "Show session info" },
  { cmd: "/copy", desc: "Copy last response to clipboard" },
  { cmd: "/export", desc: "Export conversation to file" },
  { cmd: "/clear", desc: "Reset conversation history" },
  { cmd: "/quit", desc: "Exit pear" },
];

function CommandTable({ commands }: { commands: { cmd: string; desc: string }[] }) {
  return (
    <div className="mt-4 overflow-x-auto">
      <table className="w-full text-sm">
        <thead>
          <tr className="border-b border-border text-left">
            <th className="pb-3 pr-8 font-(family-name:--font-jetbrains) font-semibold text-(--fg)">
              Command
            </th>
            <th className="pb-3 font-(family-name:--font-jetbrains) font-semibold text-(--fg)">
              Description
            </th>
          </tr>
        </thead>
        <tbody>
          {commands.map((c) => (
            <tr key={c.cmd} className="border-b border-border/50">
              <td className="py-3 pr-8">
                <code className="rounded bg-[var(--bg-subtle)] px-1.5 py-0.5 text-sm text-(--fg)">
                  {c.cmd}
                </code>
              </td>
              <td className="py-3 text-(--fg)/70">{c.desc}</td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
}

export default function CommandsPage() {
  return (
    <>
      <script
        type="application/ld+json"
        dangerouslySetInnerHTML={{ __html: JSON.stringify(jsonLd) }}
      />

      <div className="prose-pear">
        <h1 className="font-(family-name:--font-jetbrains) text-3xl font-bold text-(--fg)">
          Commands
        </h1>
        <p className="mt-4 text-lg text-(--fg)/70">
          Full reference for the pear CLI and interactive TUI slash commands.
        </p>

        <h2 className="mt-12 font-(family-name:--font-jetbrains) text-xl font-semibold text-(--fg)">
          CLI commands
        </h2>
        <CommandTable commands={cliCommands} />

        <h2 className="mt-12 font-(family-name:--font-jetbrains) text-xl font-semibold text-(--fg)">
          TUI slash commands
        </h2>
        <p className="mt-2 text-(--fg)/70">
          These commands are available inside the interactive TUI during{" "}
          <code className="rounded bg-[var(--bg-subtle)] px-1.5 py-0.5 text-sm">pear watch</code>.
        </p>
        <CommandTable commands={tuiCommands} />
      </div>
    </>
  );
}
