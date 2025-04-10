import eslint from "@eslint/js";
import importPlugin from "eslint-plugin-import";
import eslintPluginReact from "eslint-plugin-react";
import eslintPluginReactHooks from "eslint-plugin-react-hooks";
import tseslint from "typescript-eslint";
import pluginQuery from "@tanstack/eslint-plugin-query";

export default tseslint.config(
  {
    ignores: ["**/dist", "**/vite.config.ts", "eslint.config.js"],
  },
  {
    languageOptions: {
      parserOptions: {
        project: ["./tsconfig.app.json"],
        tsconfigRootDir: import.meta.dirname,
      },
    },

    settings: {
      react: {
        version: "detect",
      },
    },
  },
  eslint.configs.recommended,
  ...tseslint.configs.recommended,
  ...tseslint.configs.stylisticTypeChecked,
  ...tseslint.configs.strictTypeChecked,
  importPlugin.flatConfigs.recommended,
  eslintPluginReact.configs.flat.recommended,
  {
    plugins: {
      "react-hooks": eslintPluginReactHooks,
      "@tanstack/query": pluginQuery,
    },
    rules: eslintPluginReactHooks.configs.recommended.rules,
  },
  {
    rules: {
      curly: "error", // force use of curly brackts on if statements
      "spaced-comment": [
        "error",
        "always",
        {
          markers: ["/"],
        },
      ], // force space after comments
      "prefer-template": "error", // prefer template strings over string appends
      "react/react-in-jsx-scope": "off", // react is always in scope with vite
      "import/named": "off", // fails to resolve imports from react-router-dom
      "import/no-unresolved": "off", // fails to resolve imports from local files
      "import/order": [
        "error",
        {
          named: true,
          "newlines-between": "never",
          alphabetize: { order: "asc" },
          groups: [
            "builtin",
            "external",
            "internal",
            "parent",
            "sibling",
            "index",
            "object",
            "type",
          ],
        },
      ],

      // TODO: remove these rules when the code is fixed
      "@typescript-eslint/no-unsafe-member-access": "off",
      "@typescript-eslint/no-unsafe-argument": "off",
    },
  }
);
