import { Navbar } from "@/components/navbar";
import { Footer } from "@/components/footer";
import { FloatingDots } from "@/components/ui/floating-dots";

export default function MarketingLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <div className="flex min-h-screen flex-col">
      <FloatingDots />
      <Navbar />
      <div className="flex-1">{children}</div>
      <Footer />
    </div>
  );
}
