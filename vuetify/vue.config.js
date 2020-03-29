"use strict";
const path = require("path");
const SpritePlugin = require("svg-sprite-loader/plugin");

// function resolve(dir) {
//   return path.join(__dirname, dir)
// }

module.exports = {
  outputDir: "docs",
  configureWebpack: {
    // provide the app's title in webpack's name field, so that
    // it can be accessed in index.html to inject the correct title.
    resolve: {
      alias: {
        "@": path.join(__dirname, "src")
      },
      extensions: ["", ".js", ".json", ".vue", ".scss", ".css"]
    }
  },
  transpileDependencies: ["vuetify"],
  chainWebpack: config => {
    const svgRule = config.module.rule("svg");

    // clear all existing loaders.
    // if you don't do this, the loader below will be appended to
    // existing loaders of the rule.
    svgRule.uses.clear();

    // add replacement loader(s)
    svgRule
      .use("svg-sprite-loader")
      .loader("svg-sprite-loader")
      .options({
        extract: true,
        publicPath: "/"
      })
      .end();

    config.plugin("spritePlugin").use(SpritePlugin);
  }
};
