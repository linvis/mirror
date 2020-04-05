"use strict";
const path = require("path");

// function resolve(dir) {
//   return path.join(__dirname, dir)
// }
const port = 9527;

module.exports = {
  publicPath: "/",
  outputDir: "dist",
  assetsDir: "static",
  devServer: {
    port: port,
    overlay: {
      warnings: false,
      errors: true
    },
    proxy: {
      // change xxx-api/login => mock/login
      // detail: https://cli.vuejs.org/config/#devserver-proxy
      "/dev-api": {
        target: `http://127.0.0.1:9528/`,
        changeOrigin: true,
        pathRewrite: {
          "^/dev-api": ""
        }
      }
    }
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
