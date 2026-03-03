import Link from "next/link";
import { Github, Linkedin } from "lucide-react";

const linkColumns = [
  {
    heading: "Product",
    links: [
      { label: "About", href: "/about" },
      { label: "Pricing", href: "/pricing" },
      { label: "Compare", href: "/compare" },
    ],
  },
  {
    heading: "Resources",
    links: [
      { label: "Docs", href: "/docs" },
      { label: "FAQ", href: "/faq" },
      { label: "Blog", href: "/blog" },
    ],
  },
  {
    heading: "Legal",
    links: [
      { label: "Privacy Policy", href: "/privacy" },
      { label: "Terms of Use", href: "/terms" },
    ],
  },
];

const socials = [
  {
    label: "GitHub",
    href: "https://github.com/MitchTheStonky",
    icon: <Github className="h-3.5 w-3.5" />,
  },
  {
    label: "LinkedIn",
    href: "https://www.linkedin.com/in/mitchhazel/",
    icon: <Linkedin className="h-3.5 w-3.5" />,
  },
  {
    label: "X",
    href: "https://x.com/mitchthedev",
    icon: (
      <svg viewBox="0 0 24 24" fill="currentColor" className="h-3.5 w-3.5" aria-hidden="true">
        <path d="M18.244 2.25h3.308l-7.227 8.26 8.502 11.24H16.17l-5.214-6.817L4.99 21.75H1.68l7.73-8.835L1.254 2.25H8.08l4.713 6.231zm-1.161 17.52h1.833L7.084 4.126H5.117z" />
      </svg>
    ),
  },
];

export function Footer() {
  return (
    <footer className="relative z-10 border-t border-border bg-background/70 backdrop-blur-lg">
      <div className="mx-auto max-w-[1200px] px-6 py-10">
        {/* Main layout */}
        <div className="flex flex-wrap">
          {/* Brand column */}
          <div className="mb-6 flex-[1_1_240px] pr-12">
            <Link
              href="/"
              className="inline-flex items-center font-(family-name:--font-jetbrains) text-lg font-bold text-(--fg)"
            >
              <span className="mr-1.5 text-pear">●</span>
              pear
            </Link>

            <p className="mt-2.5 max-w-[240px] text-sm leading-relaxed text-(--fg-muted)">
              The learning system that makes sure you understand what AI writes.
            </p>

            <p className="mt-5 text-sm text-(--fg-muted)">
              &copy; 2026 pear. All rights reserved.
            </p>
          </div>

          {/* Link + social columns */}
          <nav aria-label="Footer navigation" className="flex flex-[1_1_500px] flex-wrap gap-12">
            {linkColumns.map((col) => (
              <div key={col.heading} className="min-w-[120px]">
                <h3 className="font-(family-name:--font-jetbrains) text-xs font-semibold uppercase tracking-wider text-(--fg-muted)">
                  {col.heading}
                </h3>
                <ul className="mt-3.5">
                  {col.links.map((link, i) => (
                    <li key={link.label} className={i > 0 ? "mt-2.5" : ""}>
                      <Link
                        href={link.href}
                        className="text-sm text-(--fg-muted) transition-colors hover:text-(--fg)"
                      >
                        {link.label}
                      </Link>
                    </li>
                  ))}
                </ul>
              </div>
            ))}

            {/* Socials column */}
            <div className="min-w-[120px]">
              <h3 className="font-(family-name:--font-jetbrains) text-xs font-semibold uppercase tracking-wider text-(--fg-muted)">
                Socials
              </h3>
              <ul className="mt-3.5">
                {socials.map((social, i) => (
                  <li key={social.label} className={i > 0 ? "mt-2.5" : ""}>
                    <a
                      href={social.href}
                      target="_blank"
                      rel="noopener noreferrer"
                      className="inline-flex items-center gap-2 text-sm text-(--fg-muted) transition-colors hover:text-(--fg)"
                    >
                      {social.icon}
                      {social.label}
                    </a>
                  </li>
                ))}
              </ul>
            </div>
          </nav>
        </div>
      </div>
    </footer>
  );
}
