import type { Metadata } from "next";
import { JetBrains_Mono, Inter } from "next/font/google";
import { ThemeProvider } from "@/lib/theme-provider";
import { Toaster } from "@/components/ui/sonner";
import "./globals.css";
import { Analytics } from "@vercel/analytics/next";
import { SpeedInsights } from "@vercel/speed-insights/next";

const jetbrainsMono = JetBrains_Mono({
  subsets: ["latin"],
  variable: "--font-jetbrains",
  display: "swap",
});

const inter = Inter({
  subsets: ["latin"],
  variable: "--font-inter",
  display: "swap",
});

export const metadata: Metadata = {
  metadataBase: new URL("https://pearcode.dev"),
  title: "pear | Pair Programmer for AI-Enabled Engineers",
  description:
    "pear is a pair programmer that watches your code and tells you what matters. Free and open source. Pro from $20/mo.",
  alternates: {
    canonical: "https://pearcode.dev",
  },
  openGraph: {
    title: "pear | Pair Programmer for AI-Enabled Engineers",
    description:
      "AI writes your code. pear makes sure you understand it. A pair programmer that watches your code, surfaces insights, and teaches you what matters.",
    url: "https://pearcode.dev",
    siteName: "pear",
    type: "website",
  },
  twitter: {
    card: "summary_large_image",
    title: "pear | Pair Programmer for AI-Enabled Engineers",
    description:
      "AI writes your code. pear makes sure you understand it. A pair programmer that watches your code, surfaces insights, and teaches you what matters.",
  },
};

export default function RootLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <html lang="en" suppressHydrationWarning>
      <head>
        <link
          rel="preconnect"
          href="https://fonts.gstatic.com"
          crossOrigin="anonymous"
        />
        <link rel="author" href="https://pearcode.dev/llms.txt" type="text/plain" />
        <script
          dangerouslySetInnerHTML={{
            __html: `(function(){try{var t=localStorage.getItem("theme");if(t==="dark"||(!t&&window.matchMedia("(prefers-color-scheme:dark)").matches)){document.documentElement.classList.add("dark")}}catch(e){}})()`,
          }}
        />
      </head>
      <body
        className={`${jetbrainsMono.variable} ${inter.variable} antialiased`}
      >
        <ThemeProvider>
          {children}
          <Toaster position="bottom-center" />
        </ThemeProvider>
        <Analytics />
        <SpeedInsights />
      </body>
    </html>
  );
}
