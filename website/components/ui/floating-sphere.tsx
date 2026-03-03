"use client";

import { useEffect, useRef } from "react";
import {
  Scene,
  PerspectiveCamera,
  WebGLRenderer,
  SphereGeometry,
  ShaderMaterial,
  Mesh,
  Color,
} from "three";

/* ── Vertex shader ── */
const VERT = /* glsl */ `
varying vec3 vNormal;
varying vec3 vViewPos;
varying vec3 vObjPos;

void main() {
  vObjPos   = position;
  vNormal   = normalize(normalMatrix * normal);
  vec4 mv   = modelViewMatrix * vec4(position, 1.0);
  vViewPos  = -mv.xyz;
  gl_Position = projectionMatrix * mv;
}
`;

/* ── Fragment shader — fuzzy tennis-ball felt ── */
const FRAG = /* glsl */ `
uniform float uTime;
uniform vec3  uColor;
uniform float uOpacity;

varying vec3 vNormal;
varying vec3 vViewPos;
varying vec3 vObjPos;

/* ---- noise helpers ---- */
float hash(vec3 p) {
  p = fract(p * vec3(443.897, 441.423, 437.195));
  p += dot(p, p.yzx + 19.19);
  return fract((p.x + p.y) * p.z);
}

float vnoise(vec3 p) {
  vec3 i = floor(p);
  vec3 f = fract(p);
  f = f * f * (3.0 - 2.0 * f);

  return mix(
    mix(mix(hash(i),                hash(i + vec3(1,0,0)), f.x),
        mix(hash(i + vec3(0,1,0)),  hash(i + vec3(1,1,0)), f.x), f.y),
    mix(mix(hash(i + vec3(0,0,1)),  hash(i + vec3(1,0,1)), f.x),
        mix(hash(i + vec3(0,1,1)),  hash(i + vec3(1,1,1)), f.x), f.y),
    f.z
  );
}

void main() {
  vec3 N = normalize(vNormal);
  vec3 V = normalize(vViewPos);
  float NdotV = max(dot(N, V), 0.0);

  /* ── Lighting ── */
  vec3  L    = normalize(vec3(-0.3, 0.8, 0.5));
  float diff = max(dot(N, L), 0.0);
  float light = 0.55 + diff * 0.45;

  /* ── Felt texture — LOW frequencies so it's visible at screen size ── */
  float n1 = vnoise(vObjPos * 6.0);
  float n2 = vnoise(vObjPos * 12.0);
  float n3 = vnoise(vObjPos * 24.0);
  float n4 = vnoise(vObjPos * 48.0);
  float fuzz = n1 * 0.35 + n2 * 0.30 + n3 * 0.20 + n4 * 0.15;

  /* ── Pear base with felt variation ── */
  vec3 bright = uColor * 1.18 + vec3(0.04, 0.06, 0.01);
  vec3 dark   = uColor * 0.90;
  vec3 feltColor = mix(dark, bright, fuzz);

  vec3 col = feltColor * light;

  /* ── Fiber sparkle (subtle grain) ── */
  float grain = hash(vObjPos * 60.0);
  col += smoothstep(0.65, 1.0, grain) * diff * 0.07;

  /* ── Soft edge darkening ── */
  col *= 0.82 + 0.18 * NdotV;

  /* ── Fuzzy silhouette — noise perturbs alpha at edges ── */
  float edgeNoise = vnoise(vObjPos * 10.0) * 0.14;
  float fuzzyEdge = smoothstep(0.0, 0.3, NdotV + edgeNoise - 0.03);
  float alpha = fuzzyEdge * uOpacity;

  /* ── Soft fuzz rim ── */
  float rim = 1.0 - NdotV;
  col += smoothstep(0.5, 0.9, rim) * 0.06 * bright;

  gl_FragColor = vec4(col, alpha);
}
`;

export function FloatingSphere() {
  const ref = useRef<HTMLDivElement>(null);

  useEffect(() => {
    const el = ref.current;
    if (!el) return;

    // Skip WebGL on mobile — saves bandwidth, battery, and avoids GPU issues
    if (window.innerWidth < 768) return;

    /* ── Renderer ── */
    const renderer = new WebGLRenderer({
      alpha: true,
      antialias: true,
      powerPreference: "high-performance",
    });
    renderer.setPixelRatio(Math.min(window.devicePixelRatio, 2));
    renderer.setSize(window.innerWidth, window.innerHeight);
    el.appendChild(renderer.domElement);

    /* ── Scene / Camera ── */
    const scene = new Scene();
    const camera = new PerspectiveCamera(
      50,
      window.innerWidth / window.innerHeight,
      0.1,
      100
    );
    camera.position.z = 8;

    /* ── Sphere ── */
    const sphereGeo = new SphereGeometry(0.35, 24, 24);
    const sphereMat = new ShaderMaterial({
      transparent: true,
      depthWrite: false,
      uniforms: {
        uTime: { value: 0 },
        uColor: { value: new Color(0x4ade80) },
        uOpacity: { value: 0 },
      },
      vertexShader: VERT,
      fragmentShader: FRAG,
    });
    const sphere = new Mesh(sphereGeo, sphereMat);
    scene.add(sphere);

    /* ── State ── */
    let scrollY = window.scrollY;
    let smoothScroll = scrollY;
    let smoothX = 4.5;
    let smoothY = 0;
    let animId = 0;
    let running = true;
    let heroReady = false;
    let introStart = -1;

    /* ── Events ── */
    function onHeroReady() {
      heroReady = true;
    }
    function onScroll() {
      scrollY = window.scrollY;
    }
    function onResize() {
      camera.aspect = window.innerWidth / window.innerHeight;
      camera.updateProjectionMatrix();
      renderer.setSize(window.innerWidth, window.innerHeight);
    }
    function onVisibility() {
      if (document.hidden) {
        running = false;
        cancelAnimationFrame(animId);
      } else {
        running = true;
        animId = requestAnimationFrame(loop);
      }
    }

    /* ── Render loop ── */
    function loop(time: number) {
      if (!running) return;
      const t = time * 0.001;

      // Don't render until hero typing finishes
      if (!heroReady) {
        animId = requestAnimationFrame(loop);
        return;
      }

      // Track intro start time
      if (introStart < 0) introStart = t;
      const introElapsed = t - introStart;
      const introRaw = Math.min(introElapsed / 1.8, 1);
      const introEase = 1 - Math.pow(1 - introRaw, 3); // ease-out cubic

      // Slow lerp — sphere drifts to catch up after fast scrolls
      smoothScroll += (scrollY - smoothScroll) * 0.025;
      const maxScroll =
        document.documentElement.scrollHeight - window.innerHeight;
      const frac = maxScroll > 0 ? smoothScroll / maxScroll : 0;

      // Smooth sine wave through outer gutters — alternates sides naturally
      const xTarget = Math.cos(frac * Math.PI * 5) * 4.5;
      const yTarget = Math.cos(frac * Math.PI * 3) * 0.8;

      // Lerp towards target so sphere lags behind scroll
      smoothX += (xTarget - smoothX) * 0.03;
      smoothY += (yTarget - smoothY) * 0.03;

      // Constant gentle float
      const naturalX = smoothX + Math.sin(t * 0.5) * 0.08;
      const naturalY = smoothY + Math.sin(t * 0.4) * 0.08;

      // Intro: fly in from offscreen right to natural position
      const offscreenX = 10;
      const x = offscreenX + (naturalX - offscreenX) * introEase;
      const y = naturalY;

      // Position + slow rotation (visible via felt texture)
      sphere.position.set(x, y, 0);
      sphere.rotation.y = t * 0.2;
      sphere.rotation.x = t * 0.1;

      // Opacity fades in with intro
      sphereMat.uniforms.uOpacity.value = introEase * 0.93;
      sphereMat.uniforms.uTime.value = t;

      renderer.render(scene, camera);
      animId = requestAnimationFrame(loop);
    }

    /* ── Bind ── */
    window.addEventListener("hero-ready", onHeroReady);
    window.addEventListener("scroll", onScroll, { passive: true });
    window.addEventListener("resize", onResize);
    document.addEventListener("visibilitychange", onVisibility);
    animId = requestAnimationFrame(loop);

    /* ── Cleanup ── */
    return () => {
      running = false;
      cancelAnimationFrame(animId);
      window.removeEventListener("hero-ready", onHeroReady);
      window.removeEventListener("scroll", onScroll);
      window.removeEventListener("resize", onResize);
      document.removeEventListener("visibilitychange", onVisibility);
      renderer.dispose();
      sphereGeo.dispose();
      sphereMat.dispose();
      if (el.contains(renderer.domElement)) el.removeChild(renderer.domElement);
    };
  }, []);

  return (
    <div
      ref={ref}
      className="pointer-events-none fixed inset-0 z-[5]"
      aria-hidden="true"
    />
  );
}
