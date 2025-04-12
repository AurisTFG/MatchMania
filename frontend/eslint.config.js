import ts from "typescript-eslint";
import react from "eslint-plugin-react";
import imports from "eslint-plugin-import";
import reactHooks from "eslint-plugin-react-hooks";
import tanstackQuery from "@tanstack/eslint-plugin-query";
import prettier from "eslint-plugin-prettier/recommended";

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
      "react-hooks": reactHooks,
      "@tanstack/query": tanstackQuery,
    },
    rules: {
      ...reactHooks.configs.recommended.rules,
      curly: "error", // force use of curly brackts on if statements
      "spaced-comment": [
        // force space after comments
        "error",
        "always",
        {
          markers: ["/"],
        },
      ],
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
  },
  {
    // Prettier rules go last to override conflicting ESLint rules
    ...prettier,
    rules: {
      "prettier/prettier": [
        "error",
        {
          // https://prettier.io/docs/options
          endOfLine: "lf",
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
