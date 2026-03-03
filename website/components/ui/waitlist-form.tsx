"use client";

import { useState, type FormEvent } from "react";
import { cn } from "@/lib/utils";
import { toast } from "sonner";

type FormState = "idle" | "submitting" | "success";

const EMAIL_REGEX = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;

export function WaitlistForm({ className }: { className?: string }) {
  const [email, setEmail] = useState("");
  const [state, setState] = useState<FormState>("idle");

  async function handleSubmit(e: FormEvent<HTMLFormElement>) {
    e.preventDefault();

    const trimmed = email.trim();
    if (!trimmed) {
      toast.error("Please enter your email address.");
      return;
    }
    if (!EMAIL_REGEX.test(trimmed)) {
      toast.error("Please enter a valid email address.");
      return;
    }

    setState("submitting");

    try {
      const honeypot = (
        document.getElementById("website-hp") as HTMLInputElement
      )?.value;
      const res = await fetch("/api/waitlist", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ email: trimmed, website: honeypot || "" }),
      });
      if (!res.ok) {
        const data = await res.json().catch(() => ({}));
        throw new Error(data.error || "Something went wrong. Please try again.");
      }
      setState("success");
      setEmail("");
      toast.success("You're on the list. We'll be in touch.");
    } catch (err) {
      setState("idle");
      toast.error(
        err instanceof Error ? err.message : "Something went wrong. Please try again."
      );
    }
  }

  return (
    <form
      onSubmit={handleSubmit}
      noValidate
      className={cn(
        "flex w-full flex-col gap-3 sm:flex-row sm:gap-2",
        className
      )}
    >
      {/* Honeypot field — hidden from humans, traps bots */}
      <input
        id="website-hp"
        name="website"
        type="text"
        tabIndex={-1}
        autoComplete="off"
        aria-hidden="true"
        className="absolute left-[-9999px] h-0 w-0 opacity-0"
      />
      <input
        type="email"
        value={email}
        onChange={(e) => setEmail(e.target.value)}
        placeholder="you@example.com"
        aria-label="Email address"
        required
        disabled={state === "success"}
        className={cn(
          "flex-1 rounded-lg border border-border bg-(--bg-subtle) px-4 py-3",
          "font-(family-name:--font-inter) text-(--fg)",
          "placeholder:text-(--fg-muted)",
          "outline-none transition-colors",
          "focus:border-pear",
          "disabled:opacity-60"
        )}
      />
      <button
        type="submit"
        disabled={state === "submitting" || state === "success"}
        className={cn(
          "whitespace-nowrap rounded-lg bg-pear px-6 py-3",
          "font-(family-name:--font-inter) font-semibold text-white",
          "cursor-pointer transition-colors",
          "hover:bg-pear-hover",
          "disabled:cursor-not-allowed disabled:opacity-70"
        )}
      >
        {state === "submitting"
          ? "..."
          : state === "success"
            ? "✓ You're in"
            : "Join the waitlist →"}
      </button>
    </form>
  );
}
