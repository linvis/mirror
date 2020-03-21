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
  </div>
</template>

<script>
export default {
  data() {
    return {
      inputVisible: false,
      inputValue: ''
    }
  },
  computed: {
    activeFile() {
      return this.$store.state.editor.activeFile
    },
    dynamicTags: {
      get: function() {
        return this.$store.state.editor.activeFile.metadata.tag
      },
      set: function(newVal) {
        this.$store.state.editor.activeFile.metadata.tag = newVal
      }
    }
  },
  watch: {
    activeFile: function(newVal, oldVal) {
      this.dynamicTags = this.$store.state.editor.activeFile.metadata.tag
    }
  },
  methods: {
    // delete
    handleClose(tag) {
      this.dynamicTags.splice(this.dynamicTags.indexOf(tag), 1)

      this.activeFile.metadata.tag = this.dynamicTags
      this.$store.dispatch('editor/submitCatalog')
    },

    showInput() {
      this.inputVisible = true
      this.$nextTick(_ => {
        this.$refs.saveTagInput.$refs.input.focus()
      })
    },

    // add
    handleInputConfirm() {
      console.log(this.activeFile)
      const inputValue = this.inputValue
      if (inputValue) {
        this.dynamicTags.push(inputValue)
      }
      this.inputVisible = false
      this.inputValue = ''

      this.activeFile.metadata.tag = this.dynamicTags
      this.$store.dispatch('editor/submitCatalog')
    }
  }
}
</script>

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
