import ts from "typescript-eslint";
import react from "eslint-plugin-react";
import imports from "eslint-plugin-import";
import reactHooks from "eslint-plugin-react-hooks";
import tanstackQuery from "@tanstack/eslint-plugin-query";
import prettier from "eslint-plugin-prettier/recommended";
import unusedImports from "eslint-plugin-unused-imports";

export default ts.config(
  {
    ignores: ["**/dist", "**/vite.config.ts", "**/eslint.config.js"],
  },
  ...ts.configs.recommended,
  ...ts.configs.stylisticTypeChecked,
  ...ts.configs.strictTypeChecked,
  react.configs.flat.recommended,
  react.configs.flat["jsx-runtime"],
  imports.flatConfigs.recommended,
  imports.flatConfigs.typescript,
  {
    settings: {
      react: {
        version: "detect",
      },
    },
    languageOptions: {
      parserOptions: {
        project: ["./tsconfig.app.json"],
        tsconfigRootDir: import.meta.dirname,
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
      "import/extensions": ["error", "never"], // do not require file extensions on imports
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
