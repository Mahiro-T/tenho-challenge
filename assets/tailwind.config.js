/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "../*.{go,html}",
    "../assets/*.{js,ts}"
  ],
  theme: {
    extend: {},
  },
  plugins: [],
  mode: 'jit',
  darkMode: 'media'
}
