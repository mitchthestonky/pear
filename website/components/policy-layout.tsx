import { Section } from "@/components/ui/section";

interface PolicyLayoutProps {
  title: string;
  lastUpdated: string;
  children: React.ReactNode;
}

export function PolicyLayout({ title, lastUpdated, children }: PolicyLayoutProps) {
  return (
    <main className="relative z-10 pt-16">
      <Section>
        <div className="mx-auto max-w-2xl">
          <h1 className="text-center font-(family-name:--font-jetbrains) text-3xl font-bold text-(--fg) md:text-4xl">
            {title}
          </h1>
          <p className="mt-4 text-center text-sm text-(--fg-muted)">
            Last updated: {lastUpdated}
          </p>
          <div className="mt-16 space-y-10 text-lg leading-relaxed text-(--fg)/70">
            {children}
          </div>
        </div>
      </Section>
    </main>
  );
}
