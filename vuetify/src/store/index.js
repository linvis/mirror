import Vue from "vue";
import Vuex from "vuex";
import getters from "./getters";
import show from "./modules/show";

Vue.use(Vuex);

const store = new Vuex.Store({
  modules: {
    show
  },
  getters
});

export default store;
