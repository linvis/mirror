import Vue from "vue";
import App from "./App.vue";
import router from "./router";
import store from "./store";
import vuetify from "./plugins/vuetify";
import VueLogger from "vuejs-logger";
import Notifications from "vue-notification";

const isProduction = process.env.NODE_ENV === "production";

Vue.config.productionTip = false;

import "@/permission"; // permission control

// logger
const options = {
  isEnabled: true,
  logLevel: isProduction ? "error" : "debug",
  stringifyArguments: false,
  showLogLevel: true,
  showMethodName: true,
  separator: "|",
  showConsoleColors: true
};

// import { mockXHR } from "../mock";
// if (process.env.NODE_ENV === "development") {
//   mockXHR();
// }

Vue.use(Notifications);
Vue.use(VueLogger, options);

new Vue({
  router,
  store,
  vuetify,
  render: h => h(App)
}).$mount("#app");
