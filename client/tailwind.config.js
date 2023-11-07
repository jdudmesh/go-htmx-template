/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./src/**/*.js", "../views/*.html", "../views/*.templ"],
  darkMode: "class",
  theme: {
    extend: {}
  },
  plugins: [require("daisyui")]
};
