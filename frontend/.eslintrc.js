module.exports = {
  parser: "@typescript-eslint/parser",
  extends: [
    "next",
    "next/core-web-vitals",
    "plugin:@typescript-eslint/recommended"
  ],
  plugins: ["@typescript-eslint"],
  rules: {
    "@typescript-eslint/no-explicit-any": "warn" // ここを 'warn' に変更
  }
};
