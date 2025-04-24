import ts from "typescript-eslint";
import react from "eslint-plugin-react";
import reactHooks from "eslint-plugin-react-hooks";
import tanstackQuery from "@tanstack/eslint-plugin-query";
import prettier from "eslint-plugin-prettier/recommended";
import unusedImports from "eslint-plugin-unused-imports";
import * as pluginImportX from "eslint-plugin-import-x";
import tsParser from "@typescript-eslint/parser";

export default ts.config(
  {
    ignores: ["dist", "vite.config.ts", "eslint.config.js"],
  },
  ts.configs.strictTypeChecked,
  ...ts.configs.stylisticTypeChecked,
  react.configs.flat.recommended,
  react.configs.flat["jsx-runtime"],
  pluginImportX.flatConfigs.recommended,
  pluginImportX.flatConfigs.typescript,
  {
    files: ["**/*.{js,ts,tsx}"],
    settings: {
      react: {
        version: "detect",
      },
    },
    languageOptions: {
      parser: tsParser,
      parserOptions: {
        project: ["./tsconfig.app.json"],
        tsconfigRootDir: import.meta.dirname,
        ecmaVersion: "latest",
        sourceType: "module",
      },
    },
    plugins: {
      "unused-imports": unusedImports,
      "react-hooks": reactHooks,
      "@tanstack/query": tanstackQuery,
    },
    rules: {
      ...reactHooks.configs.recommended.rules,

      curly: "error", // force use of curly brackts on if statements
      "prefer-template": "error", // prefer template strings over string appends
      "unused-imports/no-unused-imports": "error", // remove unused imports
      "@typescript-eslint/consistent-type-definitions": ["error", "type"], // prefer type over interface
      "react/self-closing-comp": ["error", { component: true, html: true }], // force self closing tags for components and html elements
      "spaced-comment": ["error", "always", { markers: ["/"] }], // force space after comment markers
      "no-restricted-imports": [
        // restrict relative imports from parent directories
        "error",
        {
          patterns: [
            "**/../api/*",
            "**/../components/*",
            "**/../configs/*",
            "**/../constants/*",
            "**/../hocs/*",
            "**/../hooks/*",
            "**/../pages/*",
            "**/../providers/*",
            "**/../styles/*",
            "**/../types/*",
            "**/../utils/*",
            "**/../validators/*",
          ],
        },
      ],
      "import-x/extensions": ["error", "never"], // do not require file extensions on imports
      "import-x/order": [
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
    },
  },
  {
    // Prettier rules go last to override conflicting ESLint rules
    ...prettier,
    rules: {
      "prettier/prettier": [
        "error",
        {
          // https://prettier.io/docs/options
          endOfLine: "auto",
          singleQuote: true,
          trailingComma: "all",
          printWidth: 80,
          tabWidth: 2,
          useTabs: false,
          objectWrap: "preserve",
          singleAttributePerLine: true,
        },
      ],
    },
  }
);
