import Link from "next/link";

export default function NotFound() {
  return (
    <div className="flex min-h-screen flex-col items-center justify-center px-6 text-center">
      <p className="font-(family-name:--font-jetbrains) text-6xl font-bold text-pear">
        404
      </p>
      <p className="mt-4 text-lg text-muted-foreground">
        This page doesn&apos;t exist.
      </p>
      <Link
        href="/"
        className="mt-6 rounded-full bg-pear px-6 py-2 text-sm font-medium text-white transition-colors hover:bg-pear-hover"
      >
        Back to home
      </Link>
    </div>
  );
}
