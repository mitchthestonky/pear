"use client";

import { Fragment, useState, useEffect, useRef } from "react";
import { cn } from "@/lib/utils";

export type AnimationStep =
  | { action: "type"; text: string; speed?: number }
  | { action: "pause"; duration: number }
  | { action: "delete"; count: number; speed?: number };

interface TypewriterTextProps {
  steps: AnimationStep[];
  loopFrom?: number;
  defaultSpeed?: number;
  defaultDeleteSpeed?: number;
  onCompleteStep?: number;
  onComplete?: () => void;
  className?: string;
}

export interface Highlight {
  word: string;
  className: string;
}

export function TypewriterText({
  steps,
  loopFrom,
  defaultSpeed = 50,
  defaultDeleteSpeed = 30,
  onCompleteStep,
  onComplete,
  className,
  highlights,
}: TypewriterTextProps & { highlights?: Highlight[] }) {
  const [displayText, setDisplayText] = useState("");
  const onCompleteRef = useRef(onComplete);
  onCompleteRef.current = onComplete;

  useEffect(() => {
    let cancelled = false;
    let stepIdx = 0;
    let charIdx = 0;
    let completeFired = false;
    let text = "";
    let timeoutId: ReturnType<typeof setTimeout> | null = null;

    const schedule = (fn: () => void, delay: number) => {
      if (!cancelled) {
        timeoutId = setTimeout(() => {
          if (!cancelled) fn();
        }, delay);
      }
    };

    const fireCompleteIfNeeded = (currentStep: number) => {
      if (
        onCompleteStep !== undefined &&
        currentStep === onCompleteStep &&
        !completeFired
      ) {
        completeFired = true;
        onCompleteRef.current?.();
      }
    };

    const runStep = () => {
      if (cancelled) return;

      if (stepIdx >= steps.length) {
        if (loopFrom !== undefined) {
          stepIdx = loopFrom;
          charIdx = 0;
          runStep();
        }
        return;
      }

      const step = steps[stepIdx];

      if (step.action === "type") {
        const speed = step.speed ?? defaultSpeed;

        if (charIdx < step.text.length) {
          text += step.text[charIdx];
          setDisplayText(text);
          charIdx++;
          schedule(runStep, speed);
        } else {
          fireCompleteIfNeeded(stepIdx);
          stepIdx++;
          charIdx = 0;
          runStep();
        }
      } else if (step.action === "pause") {
        fireCompleteIfNeeded(stepIdx);
        schedule(() => {
          stepIdx++;
          charIdx = 0;
          runStep();
        }, step.duration);
      } else if (step.action === "delete") {
        const speed = step.speed ?? defaultDeleteSpeed;

        if (charIdx < step.count) {
          text = text.slice(0, -1);
          setDisplayText(text);
          charIdx++;
          schedule(runStep, speed);
        } else {
          stepIdx++;
          charIdx = 0;
          runStep();
        }
      }
    };

    runStep();

    return () => {
      cancelled = true;
      if (timeoutId) clearTimeout(timeoutId);
    };
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);

  const highlightLine = (line: string) => {
    if (!highlights?.length) return line;
    const parts: React.ReactNode[] = [];
    let remaining = line;
    let key = 0;
    while (remaining.length > 0) {
      let earliestIdx = remaining.length;
      let matched: Highlight | null = null;
      for (const h of highlights) {
        const idx = remaining.indexOf(h.word);
        if (idx !== -1 && idx < earliestIdx) {
          earliestIdx = idx;
          matched = h;
        }
      }
      if (!matched) {
        parts.push(remaining);
        break;
      }
      if (earliestIdx > 0) parts.push(remaining.slice(0, earliestIdx));
      parts.push(
        <span key={key++} className={matched.className}>
          {matched.word}
        </span>
      );
      remaining = remaining.slice(earliestIdx + matched.word.length);
    }
    return <>{parts}</>;
  };

  return (
    <span className={cn(className)}>
      {displayText.split("\n").map((line, i) => (
        <Fragment key={i}>
          {i > 0 && <br />}
          {highlightLine(line)}
        </Fragment>
      ))}
      <span
        className="cursor-blink inline-block h-[1.1em] w-[0.55em] translate-y-[0.1em] bg-foreground"
        aria-hidden="true"
      />
    </span>
  );
}
