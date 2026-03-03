import { cn } from "@/lib/utils";

interface SectionProps {
  id?: string;
  children: React.ReactNode;
  alternate?: boolean;
  className?: string;
}

export function Section({ id, children, alternate, className }: SectionProps) {
  return (
    <section
      id={id}
      className={cn(
        alternate ? "bg-[var(--bg-subtle)]/60" : "",
        className
      )}
    >
      <div className="mx-auto max-w-[1200px] px-6 py-16 md:py-24 lg:py-32">
        {children}
      </div>
    </section>
  );
}
