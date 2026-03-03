"use client";

import { useEffect, useState, useCallback } from "react";

interface Line {
  type: "command" | "output" | "voice" | "context" | "response" | "insight";
  text: string;
  delay: number;
}

const lines: Line[] = [
  { type: "command", text: "$ pear watch", delay: 0 },
  {
    type: "output",
    text: "🍐 pear watching. Learning while you build.",
    delay: 800,
  },
  {
    type: "context",
    text: "📎 Diff detected: src/middleware/auth.ts (+12 -3)",
    delay: 2500,
  },
  {
    type: "output",
    text: "🔍 Concepts identified: RBAC · JWT claims · array vs string comparison",
    delay: 3800,
  },
  {
    type: "output",
    text: "🧠 Checking your learning state...",
    delay: 5000,
  },
  {
    type: "insight",
    text: `💡 pear noticed something:

You're checking req.user.role === "admin", but role
is an array in JWT claims, so this always returns false.

This is standard RBAC: users can have multiple roles,
so claims store them as arrays. Use
req.user.roles.includes("admin") instead.

You've seen array vs string bugs twice before, but
this is your first time hitting it in auth middleware
- where a silent false means unauthorized access
slips through.`,
    delay: 6200,
  },
  {
    type: "output",
    text: `📊 Memory updated:
   RBAC patterns ······ familiar → solid
   JWT claim structure · new → learning
   Auth middleware ····· seen → reinforced`,
    delay: 11500,
  },
];

export function TerminalMockup({ isVisible }: { isVisible: boolean }) {
  const [visibleLines, setVisibleLines] = useState<number>(0);
  const [typedChars, setTypedChars] = useState<Record<number, number>>({});
  const [key, setKey] = useState(0);

  const reset = useCallback(() => {
    setVisibleLines(0);
    setTypedChars({});
    setKey((k) => k + 1);
  }, []);

  useEffect(() => {
    if (!isVisible) return;
    reset();
  }, [isVisible, reset]);

  useEffect(() => {
    if (!isVisible) return;

    const timeouts: NodeJS.Timeout[] = [];

    lines.forEach((line, i) => {
      const t = setTimeout(() => {
        setVisibleLines((v) => Math.max(v, i + 1));

        if (
          line.type === "command" ||
          line.type === "response" ||
          line.type === "insight"
        ) {
          let charIdx = 0;
          const speed =
            line.type === "command" ? 80 : line.type === "insight" ? 10 : 12;
          const interval = setInterval(() => {
            charIdx++;
            setTypedChars((prev) => ({ ...prev, [i]: charIdx }));
            if (charIdx >= line.text.length) clearInterval(interval);
          }, speed);
          timeouts.push(interval as unknown as NodeJS.Timeout);
        }
      }, line.delay);
      timeouts.push(t);
    });

    return () => timeouts.forEach(clearTimeout);
  }, [isVisible, key]);

  function renderLine(line: Line, index: number) {
    if (index >= visibleLines) return null;

    const isTyping =
      line.type === "command" ||
      line.type === "response" ||
      line.type === "insight";
    const chars = typedChars[index] ?? (isTyping ? 0 : line.text.length);
    const displayText = isTyping ? line.text.slice(0, chars) : line.text;

    const colorClass =
      line.type === "command"
        ? "text-green-400 dark:text-green-600"
        : line.type === "insight"
          ? "text-green-300/90 dark:text-green-600/90"
          : line.type === "output" || line.type === "context"
            ? "text-gray-400 dark:text-gray-500"
            : "text-gray-100 dark:text-gray-800";

    return (
      <div
        key={`${key}-${index}`}
        className={`${colorClass} ${!isTyping ? "animate-in fade-in duration-300" : ""} whitespace-pre-wrap`}
      >
        {displayText}
        {isTyping && chars < line.text.length && (
          <span className="cursor-blink text-green-400 dark:text-green-600">▊</span>
        )}
      </div>
    );
  }

  return (
    <div className="mx-auto max-w-3xl overflow-hidden rounded-xl border border-white/10 dark:border-black/10 shadow-2xl">
      {/* Title bar */}
      <div className="flex h-8 items-center gap-2 bg-[#1a1a1a] dark:bg-[#e8e8e8] px-4">
        <div className="h-3 w-3 rounded-full bg-[#ff5f57]" />
        <div className="h-3 w-3 rounded-full bg-[#febc2e]" />
        <div className="h-3 w-3 rounded-full bg-[#28c840]" />
      </div>
      {/* Content */}
      <div className="terminal-scrollbar relative h-[360px] overflow-y-auto bg-[#0a0a0a] dark:bg-[#fafafa] p-4 font-(family-name:--font-jetbrains) text-xs leading-relaxed sm:h-[480px] sm:p-6 sm:text-sm">
        <div className="space-y-3">
          {lines.map((line, i) => renderLine(line, i))}
        </div>
        <button
          onClick={reset}
          className="absolute right-4 top-4 text-xs text-gray-500 transition-colors hover:text-gray-300 dark:text-gray-400 dark:hover:text-gray-600"
        >
          ↻ Replay
        </button>
      </div>
    </div>
  );
}
