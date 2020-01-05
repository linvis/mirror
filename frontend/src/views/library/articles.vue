<template>
  <div class="app-container">
    <div class="article-card">
      <h2 class="article-head">
        <a href="https://www.baidu.com">Test1</a>
      </h2>
      <p class="article-date">2020年1月1日</p>
      <a href="https://www.baidu.com">
        <span class="tag-code">go</span>
      </a>
      <a class="more" href="<%- url_for(post.path) %>" />
    </div>
    <div class="article-card">
      <h2 class="article-head">
        <a href="https://www.baidu.com">Test1</a>
      </h2>
      <p class="article-date">2020年1月1日</p>
      <a href="https://www.baidu.com">
        <span class="tag-code">go</span>
      </a>
      <a class="more" href="<%- url_for(post.path) %>" />
    </div>
  </div>
</template>

<script>

export default {
  methods: {
  }
}
</script>

<style lang="scss" scoped>
@import "~@/styles/article.scss";
</style>
