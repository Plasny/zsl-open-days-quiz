/** @type {import('tailwindcss').Config} */
module.exports = {
    content: ["./**/*.html", "./**/*.templ", "./**/*.go",],
    theme: {
        extend: {
            keyframes: {
                cursorblink: {
                    '0%': {
                        opacity: 0
                    },
                },
            },
            animations: {
                'cursorblink': 'cursorblink 1.5s steps(2) infinite'
            }
        },
    },
    plugins: [],
}

