const state = {
  config: {
    nav: true,
    menu: true,
    tree: false,
    editor: true,
    reminder: false,
    dialogSleep: false
  }
};

const mutations = {
  CHANGE_CONFIG: (state, { key, val }) => {
    state.config[key] = val;
  }
};
const actions = {
  changeConfig({ commit }, data) {
    commit("CHANGE_CONFIG", data);
  }
};

export default {
  namespaced: true,
  state,
  mutations,
  actions
};
