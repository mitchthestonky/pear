"use client";

import { useState } from "react";
import Link from "next/link";
import { usePathname } from "next/navigation";

const links = [
  { label: "Getting Started", href: "/docs" },
  { label: "Workflows", href: "/docs/usage" },
  { label: "Commands", href: "/docs/commands" },
  { label: "Configuration", href: "/docs/configuration" },
  { label: "Providers", href: "/docs/providers" },
];

export function DocsSidebar() {
  const pathname = usePathname();
  const [open, setOpen] = useState(false);

  return (
    <>
      {/* Mobile toggle */}
      <button
        onClick={() => setOpen(!open)}
        className="flex w-full items-center justify-between rounded-lg border border-border bg-[var(--bg-subtle)] px-4 py-2.5 text-sm font-medium text-(--fg) md:hidden"
      >
        Docs menu
        <span className={`transition-transform duration-200 ${open ? "rotate-180" : ""}`}>
          ▾
        </span>
      </button>

      {/* Sidebar nav */}
      <nav
        className={`${open ? "block" : "hidden"} md:block md:w-52 shrink-0`}
      >
        <ul className="space-y-1 py-2 md:sticky md:top-24">
          {links.map((link) => {
            const active = pathname === link.href;
            return (
              <li key={link.href}>
                <Link
                  href={link.href}
                  onClick={() => setOpen(false)}
                  className={`block rounded-md px-3 py-2 text-sm transition-colors ${
                    active
                      ? "bg-pear/10 font-medium text-pear"
                      : "text-(--fg-muted) hover:text-(--fg) hover:bg-[var(--bg-subtle)]"
                  }`}
                >
                  {link.label}
                </Link>
              </li>
            );
          })}
        </ul>
      </nav>
    </>
  );
}
