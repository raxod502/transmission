// https://github.com/sveltejs/template/blob/431bd4d58e59b46ebfa1f4fc2c1ab55853fc1521/rollup.config.js

import commonjs from "@rollup/plugin-commonjs";
import livereload from "rollup-plugin-livereload";
import nodePolyfills from "rollup-plugin-node-polyfills";
import postcss from "rollup-plugin-postcss";
import resolve from "@rollup/plugin-node-resolve";
import svelte from "rollup-plugin-svelte";
import { terser } from "rollup-plugin-terser";
import sveltePreprocess from "svelte-preprocess";

const production = !process.env.ROLLUP_WATCH;

export default {
  input: "frontend/src/svelte/index.js",
  output: {
    sourcemap: !production,
    format: "iife",
    name: "index",
    file: "frontend/out/index.js",
  },
  plugins: [
    svelte({
      dev: !production,
      css: (css) => {
        css.write("index.css", !production);
      },
      preprocess: sveltePreprocess(),
    }),
    postcss(),
    nodePolyfills(),
    resolve({
      browser: true,
      dedupe: ["svelte"],
    }),
    commonjs(),
    !production && livereload("frontend/out"),
    production && terser(),
  ],
  watch: {
    clearScreen: false,
  },
};
