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
  imports.flatConfigs.recommended,
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
      "unused-imports/no-unused-imports": "error",
      "@typescript-eslint/consistent-type-definitions": ["error", "type"],
      "react/self-closing-comp": ["error", { component: true, html: true }],
      "spaced-comment": ["error", "always", { markers: ["/"] }],
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

      "react/react-in-jsx-scope": "off", // react is always in scope with vite
      "import/named": "off", // fails to resolve imports from react-router-dom
      "import/no-unresolved": "off", // fails to resolve imports from local files

      // TODO: turn these on when the code is fixed
      "@typescript-eslint/no-unsafe-member-access": "off",
      "@typescript-eslint/no-unsafe-argument": "off",
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
