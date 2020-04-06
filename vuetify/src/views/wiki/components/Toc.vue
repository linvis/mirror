<template>
  <v-navigation-drawer
    id="documentation-toc"
    :floating="false"
    :right="true"
    clipped
  >
    <template>
      <ul class="pt-0 mb-6 documentation-toc">
        <li class="mb-2">
          <h3 class="body-1 text--primary">
            Contents
          </h3>
        </li>

        <template v-for="(item, i) in internalToc">
          <li
            v-if="item.visible"
            :key="i"
            :class="[
              `documentation-toc__link--${item.level}`,
              {
                'mb-2': i + 1 !== internalToc.length,
                'text--disabled': true
              }
            ]"
            class="documentation-toc__link"
          >
            <a
              :href="`#${item.id}`"
              class="d-block"
              @click.stop.prevent="goTo(`#${item.id}`)"
              v-html="item.text"
            />
          </li>
        </template>
      </ul>
    </template>
  </v-navigation-drawer>
</template>

<script>
import kebabCase from "lodash/kebabCase";
import { goTo } from "@/utils/helpers";
import { bus } from "@/utils/bus";

// import { get, sync } from "vuex-pathify";
export default {
  name: "DocumentationToc",
  data: () => ({
    internalToc: []
  }),
  created() {
    bus.$on("update-toc", this.updateToc);
    this.updateToc("");
  },
  beforeDestroy() {
    bus.$off("update-toc", this.updateToc);
  },
  methods: {
    goTo,
    updateToc(header) {
      this.internalToc = [];
      for (var i = 0; i < header.length; i++) {
        var text = "";
        var level = "";
        console.log("Tag name: " + header[i].textContent);
        if (header[i].tagName === "H1") {
          level = "level1";
          text = header[i].textContent;
        }
        if (header[i].tagName === "H2") {
          level = "level2";
          text = header[i].textContent;
        }
        if (header[i].tagName === "H3") {
          level = "level3";
          text = header[i].textContent;
        }
        if (level !== "") {
          this.internalToc.push({
            id: kebabCase(text),
            level,
            text,
            visible: true
          });
        }
      }
      this.$log.debug(this.internalToc);
    }
  }
};
</script>

<style lang="sass">
#documentation-toc
  .supporter-group__title
    padding-left: 8px
.documentation-toc
  list-style-type: none !important
  margin: 0
  padding: 32px 0 0
  text-align: left
  width: 100%
  li
    border-left: 2px solid transparent
    padding: 0 24px 0 8px
  li a
    color: inherit
    font-size: .875rem
    font-weight: 400
    text-decoration: none
    transition: color .1s ease-in
  .supporter-group
    justify-content: flex-start !important
  .documentation-toc__link--level1
    margin-left: 8px
  .documentation-toc__link--level2
    margin-left: 16px
  .documentation-toc__link--level3
    margin-left: 24px
</style>
