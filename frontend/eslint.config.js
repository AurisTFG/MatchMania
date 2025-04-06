import eslint from "@eslint/js";
import importPlugin from "eslint-plugin-import";
import eslintPluginReact from "eslint-plugin-react";
import eslintPluginReactHooks from "eslint-plugin-react-hooks";
import tseslint from "typescript-eslint";

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

      "import/no-unresolved": "off",
      "import/order": [
        "error",
        {
          named: true,
          "newlines-between": "always",
          alphabetize: {
            order: "asc",
          },
          groups: [
            "builtin",
            ["external", "internal"],
            ["parent", "sibling", "index", "object"],
            "type",
          ],
          pathGroups: [
            {
              group: "builtin",
              pattern: "react",
              position: "before",
            },
            {
              group: "external",
              pattern: "@mui/icons-material",
              position: "after",
            },
          ],

          pathGroupsExcludedImportTypes: ["react"],
        },
      ],
    },
  }
);
