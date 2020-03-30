import Vue from "vue";
import Vuex from "vuex";
import getters from "./getters";
import show from "./modules/show";
import menu from "./modules/menu";

Vue.use(Vuex);

const store = new Vuex.Store({
  modules: {
    show,
    menu
  },
  getters
});

export default store;
