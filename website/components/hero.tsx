"use client";

import { useState, useMemo } from "react";
import { motion, AnimatePresence } from "framer-motion";
import { TypewriterText } from "@/components/ui/typewriter-text";
import type { AnimationStep } from "@/components/ui/typewriter-text";
import { WaitlistForm } from "@/components/ui/waitlist-form";
import { TrustBanner } from "@/components/trust-banner";

export function Hero() {
  const [typingDone, setTypingDone] = useState(false);

  const steps = useMemo<AnimationStep[]>(
    () => [
      // 0: Type first line
      { action: "type", text: "AI makes you fast." },
      // 1: Pause
      { action: "pause", duration: 1000 },
      // 2: Type second line prefix
      { action: "type", text: "\npear makes you\n" },
      // 3: Type first cycling word
      { action: "type", text: "sharp." },
      // 4: Pause (onComplete fires after step 3 to fade in subtitle)
      { action: "pause", duration: 2000 },
      // 5: Loop restarts here — delete "sharp." (6 chars)
      { action: "delete", count: 6 },
      // 6: Type "irreplaceable."
      { action: "type", text: "irreplaceable." },
      // 7: Pause
      { action: "pause", duration: 2000 },
      // 8: Delete "irreplaceable." (14 chars)
      { action: "delete", count: 14 },
      // 9: Type "smart."
      { action: "type", text: "smart." },
      // 10: Pause
      { action: "pause", duration: 2000 },
      // 11: Delete "smart." (6 chars)
      { action: "delete", count: 6 },
      // 12: Type "fluent."
      { action: "type", text: "fluent." },
      // 13: Pause
      { action: "pause", duration: 2000 },
      // 14: Delete "fluent." (7 chars)
      { action: "delete", count: 7 },
      // 15: Type "sharp."
      { action: "type", text: "sharp." },
      // 16: Pause
      { action: "pause", duration: 2000 },
    ],
    []
  );

  return (
    <section
      id="hero"
      className="flex min-h-[calc(100dvh-4rem)] items-center justify-center px-4 pt-12 sm:px-6 sm:pt-16"
    >
      <div className="w-full max-w-3xl text-center">
        <h1 className="font-(family-name:--font-jetbrains) text-[2.25rem] font-bold leading-tight text-foreground sm:text-5xl" suppressHydrationWarning>
          <TypewriterText
            steps={steps}
            loopFrom={5}
            defaultSpeed={80}
            defaultDeleteSpeed={80}
            onCompleteStep={3}
            onComplete={() => {
              setTypingDone(true);
              window.dispatchEvent(new Event("hero-ready"));
            }}
            highlights={[{
              word: "pear",
              className: "highlight-pear",
            }]}
          />
        </h1>

        <AnimatePresence>
          {typingDone && (
            <motion.div
              initial={{ opacity: 0, y: 8 }}
              animate={{ opacity: 1, y: 0 }}
              transition={{ duration: 1, ease: [0.16, 1, 0.3, 1] }}
            >
              <p className="mt-4 font-(family-name:--font-inter) text-base text-(--fg-muted) sm:mt-6 sm:text-xl">
                The pair programmer that teaches, not just builds.
              </p>

              <WaitlistForm className="mx-auto mt-6 max-w-md sm:mt-8" />

              <TrustBanner />
            </motion.div>
          )}
        </AnimatePresence>
      </div>
    </section>
  );
}
