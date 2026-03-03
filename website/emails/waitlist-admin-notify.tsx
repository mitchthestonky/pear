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

interface WaitlistAdminNotifyProps {
  email?: string;
  totalCount?: number;
  ip?: string;
  timestamp?: string;
}

export default function WaitlistAdminNotify({
  email,
  totalCount,
  ip,
  timestamp,
}: WaitlistAdminNotifyProps) {
  return (
    <Html>
      <Head>
        <style>{`@import url('https://fonts.googleapis.com/css2?family=JetBrains+Mono:wght@400;700&family=Inter:wght@400;500;600&display=swap');`}</style>
      </Head>
      <Preview>{`New waitlist signup: ${email ?? "unknown"}`}</Preview>
      <Body style={body}>
        <Container style={container}>
          {/* Logo */}
          <Section style={logoSection}>
            <Text style={logo}>
              <span style={logoDot}>●</span> pear
            </Text>
          </Section>

          <Heading style={heading}>New waitlist signup</Heading>

          <Section style={detailsBox}>
            <Text style={detailRow}>
              <span style={detailLabel}>Email</span>
              <br />
              <span style={detailValue}>{email ?? "unknown"}</span>
            </Text>
            <Text style={detailRow}>
              <span style={detailLabel}>Total signups</span>
              <br />
              <span style={detailValue}>{totalCount ?? "-"}</span>
            </Text>
            <Text style={detailRow}>
              <span style={detailLabel}>IP</span>
              <br />
              <span style={detailValue}>{ip ?? "unknown"}</span>
            </Text>
            <Text style={{ ...detailRow, marginBottom: 0 }}>
              <span style={detailLabel}>Time</span>
              <br />
              <span style={detailValue}>{timestamp ?? new Date().toISOString()}</span>
            </Text>
          </Section>

          <Hr style={hr} />

          <Text style={footer}>
            Pear waitlist notifications
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
  fontSize: "24px",
  fontWeight: 700,
  color: "#fafafa",
  lineHeight: "1.3",
  margin: "0 0 24px 0",
};

const detailsBox: React.CSSProperties = {
  backgroundColor: "#111111",
  borderRadius: "12px",
  border: "1px solid #ffffff18",
  padding: "24px",
};

const detailRow: React.CSSProperties = {
  margin: "0 0 16px 0",
  fontSize: "14px",
  lineHeight: "1.5",
};

const detailLabel: React.CSSProperties = {
  color: "#6b7280",
  fontSize: "12px",
  textTransform: "uppercase" as const,
  letterSpacing: "0.05em",
  fontWeight: 600,
};

const detailValue: React.CSSProperties = {
  color: "#fafafa",
  fontFamily: "'JetBrains Mono', monospace",
  fontSize: "15px",
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
