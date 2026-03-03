import { ImageResponse } from "next/og";

export const runtime = "edge";
export const alt = "pear | The Learning Engine for AI-Enabled Engineers";
export const size = { width: 1200, height: 630 };
export const contentType = "image/png";

export default function OGImage() {
  return new ImageResponse(
    (
      <div
        style={{
          width: "100%",
          height: "100%",
          display: "flex",
          flexDirection: "column",
          alignItems: "center",
          justifyContent: "center",
          backgroundColor: "#0a0a0a",
          fontFamily: "system-ui, sans-serif",
        }}
      >
        <div
          style={{
            display: "flex",
            alignItems: "center",
            gap: "8px",
            marginBottom: "40px",
          }}
        >
          <div
            style={{
              width: "28px",
              height: "28px",
              borderRadius: "50%",
              backgroundColor: "#22c55e",
            }}
          />
          <span
            style={{
              fontSize: "36px",
              fontWeight: 700,
              color: "#fafafa",
            }}
          >
            pear
          </span>
        </div>
        <div
          style={{
            fontSize: "56px",
            fontWeight: 700,
            color: "#fafafa",
            textAlign: "center",
            lineHeight: 1.2,
            maxWidth: "800px",
          }}
        >
          AI makes shipping faster.
        </div>
        <div
          style={{
            fontSize: "56px",
            fontWeight: 700,
            color: "#fafafa",
            textAlign: "center",
            lineHeight: 1.2,
            maxWidth: "800px",
            display: "flex",
            gap: "16px",
          }}
        >
          <span style={{ color: "#22c55e" }}>pear</span>
          <span>makes you smarter.</span>
        </div>
        <div
          style={{
            fontSize: "22px",
            color: "#9ca3af",
            marginTop: "32px",
            textAlign: "center",
            maxWidth: "600px",
          }}
        >
          The learning engine for AI-enabled engineers.
        </div>
      </div>
    ),
    { ...size }
  );
}
