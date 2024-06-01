/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./**/*.templ"],
  theme: {
    extend: {},
  },
  plugins: [],
  safelist: [
    'left-0', 'left-full', 'right-0', 'right-full', 'before:left-0', 'before:left-full', 'before:right-0', 'before:right-full'
  ]
}

