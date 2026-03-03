import { Section } from "@/components/ui/section";
import { ScrollReveal } from "@/components/ui/scroll-reveal";
import { Check, X } from "lucide-react";

type CellValue = boolean | "partial" | string;

const rows: { label: string; tools: CellValue; free: CellValue; pro: CellValue }[] = [
  { label: "Auto-reviews changes", tools: false, free: true, pro: true },
  { label: "Explains why, not just what", tools: "partial", free: true, pro: true },
  { label: "Concept tagging", tools: false, free: true, pro: true },
  { label: "Remembers project context", tools: "partial", free: true, pro: true },
  { label: "Remembers what you understand", tools: false, free: false, pro: true },
  { label: "Adapts teaching to your level", tools: false, free: false, pro: true },
  { label: "Tracks knowledge gaps", tools: false, free: false, pro: true },
  { label: "Shows growth over time", tools: false, free: false, pro: true },
  { label: "Generates code for you", tools: true, free: false, pro: false },
];

function CellIcon({ value }: { value: CellValue }) {
  if (value === true) return <Check className="mx-auto h-4 w-4 text-pear" />;
  if (value === false) return <X className="mx-auto h-4 w-4 text-(--fg)/20" />;
  return <span className="block text-center text-xs text-(--fg)/50">Partial</span>;
}

export function TeachingDifference() {
  return (
    <Section id="comparison">
      <ScrollReveal>
        <h2 className="mb-4 text-center text-3xl font-bold text-(--fg) md:text-4xl font-(family-name:--font-jetbrains)">
          How Pear is different
        </h2>
        <p className="mx-auto mb-16 max-w-xl text-center text-lg text-(--fg-muted)">
          Pear doesn&rsquo;t replace your AI coding tool. It&rsquo;s the layer that makes sure you understand what it writes.
        </p>
      </ScrollReveal>

      {/* Desktop table */}
      <div className="mx-auto hidden max-w-3xl md:block">
        {/* Header */}
        <div className="mb-4 grid grid-cols-[1fr_8rem_8rem_8rem] items-end px-4 text-sm">
          <div />
          <div className="text-center text-xs font-semibold uppercase tracking-wider text-(--fg-muted)">
            Claude Code /<br />Cursor / OpenCode
          </div>
          <div className="text-center text-xs font-semibold uppercase tracking-wider text-(--fg-muted)">
            Pear Free
          </div>
          <div className="text-center text-xs font-semibold uppercase tracking-wider text-pear">
            Pear Pro
          </div>
        </div>

        {/* Rows */}
        <div className="space-y-2 text-sm">
          {rows.map((row, i) => (
            <ScrollReveal key={i} delay={i * 40}>
              <div className="grid grid-cols-[1fr_8rem_8rem_8rem] items-center rounded-lg border border-border bg-background px-4 py-3">
                <span className="text-(--fg)/80">{row.label}</span>
                <CellIcon value={row.tools} />
                <CellIcon value={row.free} />
                <CellIcon value={row.pro} />
              </div>
            </ScrollReveal>
          ))}
        </div>
      </div>

      {/* Mobile table */}
      <div className="mx-auto max-w-md md:hidden">
        <ScrollReveal>
          {/* Header */}
          <div className="grid grid-cols-[1fr_3.5rem_3.5rem_3.5rem] items-end px-3 pb-2 text-[10px] font-semibold uppercase tracking-wider">
            <div />
            <div className="text-center text-(--fg-muted)">AI</div>
            <div className="text-center text-(--fg-muted)">Free</div>
            <div className="text-center text-pear">Pro</div>
          </div>

          {/* Rows */}
          <div className="divide-y divide-border rounded-xl border border-border bg-background">
            {rows.map((row, i) => (
              <div key={i} className="grid grid-cols-[1fr_3.5rem_3.5rem_3.5rem] items-center px-3 py-3">
                <span className="text-xs text-(--fg)/80">{row.label}</span>
                <CellIcon value={row.tools} />
                <CellIcon value={row.free} />
                <CellIcon value={row.pro} />
              </div>
            ))}
          </div>
        </ScrollReveal>
      </div>

      <ScrollReveal delay={400}>
        <p className="mt-8 text-center text-sm text-(--fg-muted) italic">
          Pear doesn&rsquo;t replace your AI coding tool. It&rsquo;s the layer that makes sure you understand what it writes.
        </p>
      </ScrollReveal>
    </Section>
  );
}
