import Vue from "vue";
import Vuex from "vuex";
import getters from "./getters";
import show from "./modules/show";
import menu from "./modules/menu";
import editor from "./modules/editor";

Vue.use(Vuex);

const store = new Vuex.Store({
  modules: {
    show,
    menu,
    editor
  },
  getters
});

export default store;
