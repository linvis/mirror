export function goTo(id) {
  this.$vuetify.goTo(id).then(() => {
    if (!id) return (document.location.hash = "");

    if (history.replaceState) {
      history.replaceState(null, null, id);
    } else {
      document.location.hash = id;
    }
  });
}
