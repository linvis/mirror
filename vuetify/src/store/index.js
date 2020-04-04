import Vue from "vue";
import Vuex from "vuex";
import getters from "./getters";
import show from "./modules/show";
import menu from "./modules/menu";
import editor from "./modules/editor";
import user from "./modules/user";

Vue.use(Vuex);

const store = new Vuex.Store({
  modules: {
    show,
    menu,
    editor,
    user
  },
  getters
});

export default store;
