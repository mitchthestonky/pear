import {
  Body,
  Container,
  Head,
  Heading,
  Hr,
  Html,
  Preview,
  Section,
  Text,
} from "@react-email/components";

interface WaitlistWelcomeProps {
  email?: string;
}

export default function WaitlistWelcome({ email }: WaitlistWelcomeProps) {
  return (
    <Html>
      <Head>
        <style>{`@import url('https://fonts.googleapis.com/css2?family=JetBrains+Mono:wght@400;700&family=Inter:wght@400;500;600&display=swap');`}</style>
      </Head>
      <Preview>You&apos;re on the pear waitlist. The learning engine for AI-enabled builders.</Preview>
      <Body style={body}>
        <Container style={container}>
          {/* Logo */}
          <Section style={logoSection}>
            <Text style={logo}>
              <span style={logoDot}>●</span> pear
            </Text>
          </Section>

          <Heading style={heading}>You&apos;re on the list.</Heading>

          <Text style={paragraph}>
            Thanks for signing up, {email ?? "there"}. You&apos;re now on the
            pear waitlist.
          </Text>

          <Text style={paragraph}>
            pear is the only learning engine for AI-enabled builders. It sits
            in your terminal, watches your coding sessions, and teaches you
            the engineering concepts behind the code you ship. Not after the
            fact. Right when it matters.
          </Text>

          <Text style={paragraph}>
            While your AI tools make you faster, pear makes sure you actually
            understand what you are building. It remembers what you know,
            adapts how it teaches, and tracks your growth over time.
          </Text>

          <Text style={paragraph}>
            We&apos;re launching in March 2026. You&apos;ll be among the
            first to try it.
          </Text>

          <Hr style={hr} />

          <Text style={subheading}>Know someone who should be on this list?</Text>

          <Text style={paragraph}>
            If you know a developer who ships with AI but wants to understand
            their code better, send them to pearcode.dev. The best engineers
            invest in understanding, not just output.
          </Text>

          <Hr style={hr} />

          <Text style={footer}>
            pear. The only learning engine for AI-enabled builders.
          </Text>
        </Container>
      </Body>
    </Html>
  );
}

const body: React.CSSProperties = {
  backgroundColor: "#0a0a0a",
  fontFamily: "'Inter', system-ui, sans-serif",
  margin: 0,
  padding: 0,
};

const container: React.CSSProperties = {
  backgroundColor: "#0a0a0a",
  margin: "0 auto",
  padding: "48px 24px",
  maxWidth: "520px",
};

const logoSection: React.CSSProperties = {
  marginBottom: "32px",
};

const logo: React.CSSProperties = {
  fontFamily: "'JetBrains Mono', monospace",
  fontSize: "20px",
  fontWeight: 700,
  color: "#fafafa",
  margin: 0,
};

const logoDot: React.CSSProperties = {
  color: "#22c55e",
};

const heading: React.CSSProperties = {
  fontFamily: "'JetBrains Mono', monospace",
  fontSize: "28px",
  fontWeight: 700,
  color: "#fafafa",
  lineHeight: "1.3",
  margin: "0 0 24px 0",
};

const subheading: React.CSSProperties = {
  fontFamily: "'JetBrains Mono', monospace",
  fontSize: "16px",
  fontWeight: 700,
  color: "#fafafa",
  lineHeight: "1.4",
  margin: "0 0 12px 0",
};

const paragraph: React.CSSProperties = {
  fontSize: "16px",
  lineHeight: "1.6",
  color: "#9ca3af",
  margin: "0 0 16px 0",
};

const hr: React.CSSProperties = {
  borderColor: "#ffffff18",
  margin: "32px 0",
};

const footer: React.CSSProperties = {
  fontSize: "13px",
  color: "#6b7280",
  margin: 0,
};
