/** @type {import('tailwindcss').Config} */
module.exports = {
    content: ['./internal/templates/**/*.{html,js}', './public/**/*.html'],
    theme: {
        fontSize: {
            sm: '0.8rem',
            base: '1rem',
            lg: '1.25rem',
            xl: '1.563rem',
            '2xl': '1.953rem',
            '3xl': '2.441rem',
            '4xl': '3.052rem',
            '5xl': '4rem',
            '6xl': '5rem',
            '7xl': '6rem'
        },
        extend: {
            width: {
                content: '1280px',
                'content-small': '680px'
            },
            maxWidth: {
                content: '1280px',
                'content-small': '680px'
            }
        }
    },
    daisyui: {
        themes: ['night']
    },
    plugins: [require('@tailwindcss/typography'), require('daisyui')]
};
