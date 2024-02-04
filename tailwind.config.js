/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./views/**/*.templ"],
  theme: {
    extend: {
      fontFamily: {
        workbench: ["Workbench", "sans-serif"],
      },
    },
  },
  plugins: [],
};
