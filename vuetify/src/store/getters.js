const getters = {
  config: state => state.show.config,
  activeMenu: state => state.menu.activeMenu,
  catalog: state => state.editor.catalog,
  activeNote: state => state.editor.activeNote
};
export default getters;
