const expandableComponent = {
  props: ["label", "expanded"],
  setup(props) {
    const expanded = ref(props.expanded);



    return {
      expanded,
      label: props.label
    }
  },
  template: `<div class="expandable">
    <h3 class="title" @click="expanded = !expanded">
      <span>[[label]]</span>
      <img :class="{
        'arrow': true,
        'active': expanded
      }" src="/static/images/left-arrow.svg" height="20px" />
    </h3>
    <div class="content" v-if="expanded">
      <slot />
    </div>
  </div>`
}