<template>
  <div class="search">
    <el-autocomplete
      v-model="keyword"
      :fetch-suggestions="querySearchAsync"
      placeholder="search"
      clearable="true"
      style="width:100%;"
      @select="handleSelect"
    />
  </div>
</template>

<script>
export default {
  data() {
    return {
      keyword: ''
    }
  },
  methods: {
    querySearch(queryString, cb) {
      // var restaurants = this.restaurants
      // var results = queryString ? restaurants.filter(this.createFilter(queryString)) : restaurants
      // // 调用 callback 返回建议列表的数据
      // cb(results)
    },
    handleSelect(item) {
      console.log(item)
    }
  }
}
</script>
<style>
.search {
  padding: 0%;
  margin: 0%;
  width: auto;
}
</style>
