<template>
  <v-app id="inspire">
    <v-content>
      <v-container fluid fill-height>
        <v-layout align-center justify-center>
          <v-flex xs12 sm8 md4>
            <v-card class="elevation-12">
              <v-toolbar color="primary" dark flat>
                <v-toolbar-title>Login</v-toolbar-title>
              </v-toolbar>
              <v-card-text>
                <v-form>
                  <v-text-field
                    label="Login"
                    name="login"
                    prepend-icon="mdi-account"
                    type="text"
                    v-model="loginForm.username"
                  ></v-text-field>

                  <v-text-field
                    id="password"
                    label="Password"
                    name="password"
                    prepend-icon="mdi-lock"
                    type="password"
                    v-model="loginForm.password"
                  ></v-text-field>
                </v-form>
              </v-card-text>
              <v-card-actions>
                <v-spacer></v-spacer>
                <v-btn color="primary" @click="handleLogin">Login</v-btn>
              </v-card-actions>
            </v-card>
          </v-flex>
        </v-layout>

        <v-overlay :value="overlay">
          <v-progress-circular indeterminate size="64"></v-progress-circular>
        </v-overlay>

        <notifications group="loginFail" position="top center">
          <template slot="body" slot-scope="props">
            <v-alert v-bind:type="props.item.type">
              {{ props.item.text }}
            </v-alert>
          </template>
        </notifications>
      </v-container>
    </v-content>
  </v-app>
</template>

<script>
import md5 from "blueimp-md5";

export default {
  data() {
    return {
      overlay: false,
      loginForm: {
        username: "",
        password: ""
      },
      redirect: undefined
    };
  },
  watch: {
    overlay(val) {
      val &&
        setTimeout(() => {
          this.overlay = false;
        }, 3000);
    },
    $route: {
      handler: function(route) {
        this.redirect = route.query && route.query.redirect;
      },
      immediate: true
    }
  },
  methods: {
    handleLogin() {
      this.overlay = true;
      this.$log.debug(this.loginForm);
      this.$log.debug("fuck");

      var password = md5(this.loginForm.password);
      this.$store
        .dispatch("user/login", { username: this.loginForm.username, password })
        .then(() => {
          this.$router.push({ path: this.redirect || "/" });
          this.overlay = false;
        })
        .catch(() => {
          this.overlay = false;
          this.$notify({
            group: "loginFail",
            type: "error",
            text: "Login Fail",
            duration: 1000
          });
        });
    }
  }
};
</script>

<style lang="scss">
.n-light {
  margin: 10px;
  margin-bottom: 0;
  border-radius: 3px;
  font-size: 13px;
  padding: 10px 20px;
  color: #495061;
  background: #eaf4fe;
  border: 1px solid #d4e8fd;
}
</style>
