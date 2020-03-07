<!--
 * @Author: your name
 * @Date: 2020-01-31 15:33:44
 * @LastEditTime: 2020-03-07 19:55:05
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /frontend/src/views/settings/index.vue
 -->
<template>
  <div class="mixin-components-container">
    <el-card class="box-card" shadow="hover">
      <div slot="header" class="clearfix">
        <span>编程配置</span>
      </div>
      <el-form ref="form" label-position="left" :model="form" label-width="160px">
        <el-form-item label="Leetcode User Name">
          <el-input v-model="leetcodeName" clearable style="width: 40%;" />
        </el-form-item>
        <el-form-item label="Leetcode Cookies">
          <el-input v-model="leetcodeCookies" clearable style="width: 40%;" />
        </el-form-item>
        <el-form-item label="Github User Name">
          <el-input v-model="githubName" clearable style="width: 40%;" />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleSubmit">立即创建</el-button>
          <el-button>取消</el-button>
        </el-form-item>
      </el-form>
    </el-card>
  </div>
</template>

<script>

import { submitSettingProgram, querySettingProgram } from '@/api/setting'

export default {
  data() {
    return {
      leetcodeName: '',
      leetcodeCookies: '',
      githubName: '',
      form: {

      }
    }
  },
  mounted() {
    this.featchData()
  },
  methods: {
    featchData() {
      querySettingProgram().then(response => {
        this.leetcodeName = response.data.leetcode_user_name
        this.leetcodeCookies = response.data.cookie
        this.githubName = response.data.github_user_name
      })
    },
    handleSubmit: function(event) {
      var setting = {
        'leetcode_user_name': this.leetcodeName,
        'cookie': this.leetcodeCookies,
        'github_user_name': this.githubName
      }
      submitSettingProgram(setting).then(response => {
        this.$notify({
          title: '成功',
          message: '',
          type: 'success',
          duration: 800
        })
      })
    }
  }

}
</script>

<style scoped>
.mixin-components-container {
  background-color: #f0f2f5;
  padding: 30px;
  min-height: calc(100vh - 84px);
}
.component-item {
  min-height: 100px;
}
</style>
