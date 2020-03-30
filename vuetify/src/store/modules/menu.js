const state = {
  activeMenu: null
};

const mutations = {
  CHANGE_ACTIVEMENU: (state, val) => {
    state.activeMenu = val;
  }
};
const actions = {
  changeActiveMenu({ commit }, data) {
    commit("CHANGE_ACTIVEMENU", data);
  }
};

export default {
  namespaced: true,
  state,
  mutations,
  actions
};
