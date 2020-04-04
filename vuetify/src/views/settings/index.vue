<template>
  <v-card min-width="30%" class="mx-4">
    <v-row>
      <v-col cols="6">
        <v-card-title>Settings</v-card-title>
      </v-col>
      <v-col col="6">
        <v-card-actions align-end class="pr-4">
          <v-spacer></v-spacer>
          <v-btn color="green accent-4" @click="submitSettingProgram">
            Update
          </v-btn>
        </v-card-actions>
      </v-col>
    </v-row>
    <v-container>
      <form>
        <v-text-field
          v-model="leetcodeName"
          label="Leetcode User Name"
          outlined
          clearable
        ></v-text-field>
        <v-text-field
          v-model="leetcodeCookies"
          label="Leetcode Cookies"
          outlined
          clearable
        ></v-text-field>
        <v-text-field
          v-model="githubName"
          label="Github User Name"
          outlined
          clearable
        ></v-text-field>
      </form>
    </v-container>
  </v-card>
</template>

<script>
import { submitSettingProgram, querySettingProgram } from "@/api/setting";

export default {
  data() {
    return {
      leetcodeName: "",
      leetcodeCookies: "",
      githubName: ""
    };
  },
  mounted() {
    this.featchData();
  },
  methods: {
    featchData() {
      querySettingProgram().then(response => {
        this.leetcodeName = response.data.leetcode_user_name;
        this.leetcodeCookies = response.data.cookie;
        this.githubName = response.data.github_user_name;
      });
    },
    handleSubmit: function(event) {
      var setting = {
        leetcode_user_name: this.leetcodeName,
        cookie: this.leetcodeCookies,
        github_user_name: this.githubName
      };
      submitSettingProgram(setting).then(response => {
        // this.$notify({
        //   title: "成功",
        //   message: "",
        //   type: "success",
        //   duration: 800
        // });
      });
    }
  }
};
</script>
