<template>
  <div class="tag">
    <!-- <span style="margin-left:0px;">tag:</span> -->
    <el-tag
      v-for="tag in dynamicTags"
      :key="tag"
      closable
      size="small"
      :disable-transitions="false"
      @close="handleClose(tag)"
    >{{ tag }}</el-tag>
    <el-input
      v-if="inputVisible"
      ref="saveTagInput"
      v-model="inputValue"
      class="input-new-tag"
      size="small"
      @keyup.enter.native="handleInputConfirm"
      @blur="handleInputConfirm"
    />
    <el-button v-else type="text" icon="el-icon-plus" size="small" @click="showInput" />
    <!-- <el-button v-else class="button-new-tag" size="small" @click="showInput">+ new</el-button> -->
  </div>
</template>

<style>
.el-tag + .el-tag {
  margin-left: 5px;
}
.button-new-tag {
  margin-left: 10px;
  height: 30px;
  line-height: 30px;
  padding-top: 0;
  padding-bottom: 0;
}
.input-new-tag {
  width: 90px;
  height: 30px;
  margin-left: 10px;
  vertical-align: bottom;
}
</style>

<script>
export default {
  data() {
    return {
      dynamicTags: ['tag'],
      inputVisible: false,
      inputValue: ''
    }
  },
  methods: {
    handleClose(tag) {
      this.dynamicTags.splice(this.dynamicTags.indexOf(tag), 1)
    },

    showInput() {
      this.inputVisible = true
      this.$nextTick(_ => {
        this.$refs.saveTagInput.$refs.input.focus()
      })
    },

    handleInputConfirm() {
      const inputValue = this.inputValue
      if (inputValue) {
        this.dynamicTags.push(inputValue)
      }
      this.inputVisible = false
      this.inputValue = ''
    }
  }
}
</script>
