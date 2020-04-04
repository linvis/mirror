const getters = {
  config: state => state.show.config,
  activeMenu: state => state.menu.activeMenu,
  catalog: state => state.editor.catalog,
  activeNote: state => state.editor.activeNote,
  token: state => state.user.token,
  avatar: state => state.user.avatar,
  name: state => state.user.name
};
export default getters;
