'use strict'
const path = require('path')
const name = "mirror"

// function resolve(dir) {
//   return path.join(__dirname, dir)
// }

module.exports = {
  configureWebpack: {
    // provide the app's title in webpack's name field, so that
    // it can be accessed in index.html to inject the correct title.
    resolve: {
      alias: {
        "@": path.join(__dirname, 'src')
      }
    }
  },
  "transpileDependencies": [
    "vuetify"
  ]
}