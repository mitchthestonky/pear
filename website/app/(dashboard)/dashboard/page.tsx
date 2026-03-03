import type { Metadata } from "next";

export const metadata: Metadata = {
  robots: { index: false, follow: false },
};

export default function Dashboard() {
  return (
    <div className="flex min-h-screen items-center justify-center">
      <p className="text-muted-foreground">Dashboard coming soon.</p>
    </div>
  );
}
