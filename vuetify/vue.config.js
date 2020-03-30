"use strict";
const path = require("path");

// function resolve(dir) {
//   return path.join(__dirname, dir)
// }
const port = 9528;

module.exports = {
  publicPath: "/",
  outputDir: "dist",
  assetsDir: "static",
  devServer: {
    port: port
  },
  configureWebpack: {
    // provide the app's title in webpack's name field, so that
    // it can be accessed in index.html to inject the correct title.
    resolve: {
      alias: {
        "@": path.join(__dirname, "src")
      }
      // extensions: [".js", ".json", ".vue", ".scss", ".css"]
    }
  },
  transpileDependencies: ["vuetify"]
};
