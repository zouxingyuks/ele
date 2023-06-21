const gulp = require('gulp');
const postcss = require('gulp-postcss');
const tailwindcss = require('tailwindcss');

function css() {
    return gulp.src('./src/css/*.css')
        .pipe(postcss([
            // require other plugins here if needed
            tailwindcss('./tailwind.config.js'),
            require('autoprefixer'),
        ]))
        .pipe(gulp.dest('./assets/css/'));
}

function watch() {
    gulp.watch('./src/css/*.css', css);
}

exports.css = css;
exports.watch = watch;
