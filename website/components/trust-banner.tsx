"use client";

import Image from "next/image";

const providers = [
  { name: "Anthropic", src: "/logos/providers/anthropic.svg" },
  { name: "OpenAI", src: "/logos/providers/openai.svg" },
  { name: "OpenRouter", src: "/logos/providers/openrouter.svg" },
  { name: "Google Gemini", src: "/logos/providers/gemini.svg" },
  { name: "Mistral", src: "/logos/providers/mistral.svg" },
];

const tools = [
  { name: "Cursor", src: "/logos/tools/cursor.svg" },
  { name: "Claude Code", src: "/logos/tools/claude-code.svg" },
  { name: "VS Code", src: "/logos/tools/vscode.svg" },
  { name: "Copilot", src: "/logos/tools/copilot.svg" },
  { name: "Windsurf", src: "/logos/tools/windsurf.svg" },
  { name: "Zed", src: "/logos/tools/zed.svg" },
];

function LogoIcon({ name, src }: { name: string; src: string }) {
  return (
    <div className="group relative" title={name}>
      <Image
        src={src}
        alt={name}
        width={40}
        height={40}
        className="h-9 w-9 sm:h-10 sm:w-10 opacity-30 transition-opacity duration-200 group-hover:opacity-60 dark:invert-0 invert"
      />
      <span className="pointer-events-none absolute -bottom-6 left-1/2 -translate-x-1/2 whitespace-nowrap rounded bg-foreground/10 px-1.5 py-0.5 text-[10px] text-(--fg-muted) opacity-0 backdrop-blur-sm transition-opacity duration-200 group-hover:opacity-100">
        {name}
      </span>
    </div>
  );
}

export function TrustBanner() {
  return (
    <div className="mt-8 sm:mt-10">
      <p className="mb-4 text-center text-[10px] font-medium uppercase tracking-[0.2em] text-(--fg-muted)/35 sm:text-[11px]">
        Pear works with
      </p>

      <div className="flex items-end justify-center gap-8 sm:gap-12">
        {/* Providers */}
        <div className="flex flex-col items-center gap-2">
          <div className="flex items-center gap-5 sm:gap-7">
            {providers.map((p) => (
              <LogoIcon key={p.name} {...p} />
            ))}
          </div>
          <span className="text-[9px] font-medium uppercase tracking-[0.2em] text-(--fg-muted)/35 sm:text-[10px]">
            Providers
          </span>
        </div>

        {/* Divider */}
        <div className="hidden sm:block h-8 w-px bg-(--fg-muted)/10" />

        {/* Tooling */}
        <div className="flex flex-col items-center gap-2">
          <div className="flex items-center gap-5 sm:gap-7">
            {tools.map((t) => (
              <LogoIcon key={t.name} {...t} />
            ))}
          </div>
          <span className="text-[9px] font-medium uppercase tracking-[0.2em] text-(--fg-muted)/35 sm:text-[10px]">
            Tooling
          </span>
        </div>
      </div>
    </div>
  );
}
