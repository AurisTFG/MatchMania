/// <reference types="vitest/config" />
import { defineConfig } from "vite";
import react from "@vitejs/plugin-react-swc";
import tsconfigPaths from 'vite-tsconfig-paths';

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [react(), tsconfigPaths()],
  envPrefix: "MATCHMANIA",
  test: {
    globals: true,
    environment: "jsdom",
    setupFiles: "./src/configs/vitestSetup.ts",
  },
});
