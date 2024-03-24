/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    'views//**/*.templ',
  ],
  darkMode: 'class',
  theme:
  {
    mytheme: {
      "primary": "#009ae4",
      "secondary": "#00b99f",
      "accent": "#006f00",
      "neutral": "#0e051d",
      "base-100": "#262626",
      "info": "#00a9ff",
      "success": "#8bf763",
      "warning": "#d97f00",
      "error": "#fb075c",
    },
  },
  plugins: [
    require("@tailwindcss/typography"),
    require("daisyui")
  ],
  corePlugins: {
    preflight: true,
  }
}