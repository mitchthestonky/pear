import { NextResponse } from "next/server";
import { Redis } from "@upstash/redis";
import { Ratelimit } from "@upstash/ratelimit";
import { Resend } from "resend";
import WaitlistWelcome from "@/emails/waitlist-welcome";
import WaitlistAdminNotify from "@/emails/waitlist-admin-notify";

let _resend: Resend | null = null;
function getResend() {
  if (!_resend && process.env.RESEND_API_KEY) {
    _resend = new Resend(process.env.RESEND_API_KEY);
  }
  return _resend;
}

const redis = new Redis({
  url: process.env.KV_REST_API_URL!,
  token: process.env.KV_REST_API_TOKEN!,
});

const ratelimit = new Ratelimit({
  redis,
  limiter: Ratelimit.slidingWindow(2, "24 h"),
  prefix: "waitlist:rl",
});

const EMAIL_REGEX = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
const FROM_EMAIL = process.env.RESEND_FROM_EMAIL ?? "noreply@pearcode.dev";
const ADMIN_EMAIL = process.env.RESEND_ADMIN_EMAIL;

export async function POST(request: Request) {
  try {
    const ip =
      request.headers.get("x-forwarded-for")?.split(",")[0]?.trim() ??
      "unknown";

    const { success } = await ratelimit.limit(ip);
    if (!success) {
      return NextResponse.json(
        { error: "Too many requests. Please try again later." },
        { status: 429 }
      );
    }

    const { email, website } = await request.json();

    // Honeypot: if the hidden "website" field has a value, it's a bot
    if (website) {
      return NextResponse.json({ ok: true });
    }

    if (!email || !EMAIL_REGEX.test(email)) {
      return NextResponse.json(
        { error: "Invalid email address." },
        { status: 400 }
      );
    }

    const normalized = email.toLowerCase().trim();

    // SADD returns 1 if the member was added (new), 0 if already existed
    const added = await redis.sadd("waitlist:emails", normalized);
    const isNew = added === 1;

    // Send emails (only for new signups, don't block response on failure)
    const resend = getResend();
    if (isNew && resend) {
      const timestamp = new Date().toISOString();
      const totalCount = await redis.scard("waitlist:emails");

      Promise.allSettled([
        resend.emails.send({
          from: `Pear <${FROM_EMAIL}>`,
          to: normalized,
          subject: "You're on the Pear waitlist",
          react: WaitlistWelcome({ email: normalized }),
        }),

        ...(ADMIN_EMAIL
          ? [
              resend.emails.send({
                from: `Pear Waitlist <${FROM_EMAIL}>`,
                to: ADMIN_EMAIL,
                subject: `New waitlist signup: ${normalized}`,
                react: WaitlistAdminNotify({
                  email: normalized,
                  totalCount,
                  ip,
                  timestamp,
                }),
              }),
            ]
          : []),
      ]).catch(() => {
        // Silently fail — signup is already saved
      });
    }

    return NextResponse.json({ ok: true });
  } catch {
    return NextResponse.json(
      { error: "Something went wrong. Please try again." },
      { status: 500 }
    );
  }
}
