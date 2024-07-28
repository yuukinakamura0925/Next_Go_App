module.exports = {
  extends: [
    "stylelint-config-standard",
    "stylelint-config-recommended-scss",
    "stylelint-config-tailwindcss",
    "stylelint-config-prettier"
  ],
  rules: {
    "at-rule-no-unknown": [
      true,
      {
        "ignoreAtRules": ["tailwind", "apply", "variants", "responsive", "screen", "layer"]
      }
    ],
    // その他のルールを追加することができます
  }
};