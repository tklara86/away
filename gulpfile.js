const { src, dest,watch, series } = require('gulp');
const sass = require('gulp-sass')(require('sass'));
const postcss = require('gulp-postcss');
const cssnano = require('cssnano');
const terser = require('gulp-terser');
const autoprefixer = require('autoprefixer');
const browserSync = require('browser-sync').create();



// Sass task
function scssTask() {
    return src('ui/static/styles/scss/main.scss', { sourcemaps: true })
        .pipe(sass())
        .pipe(postcss([cssnano, autoprefixer]))
        .pipe(sass().on('error', sass.logError))
        .pipe(dest('ui/static/dist', { sourcemaps: '.' }));
}


// JS Task
function jsTask() {
   return src('ui/static/js/app.js', { sourcemaps: true })
       .pipe(terser())
       .pipe(dest('ui/static/dist', { sourcemaps: '.' }));
}

// BrowserSync task

// Watch task
function watchTask() {
 //  watch('*.html', browserSyncReload);
    watch(['ui/static/styles/**/*.scss', 'ui/static/js/**/*.js'], series(scssTask, jsTask))
}

// Default Gulp task
exports.default = series(
    scssTask,
    jsTask,
  //  browserSyncServe,
    watchTask
)

