const DocsComponent = {
  props: ["docs"],
  setup(props) {
    return {
      docs: unescapeHtml(props.docs)
    }
  },
  template: `<div>
    <div v-html="[[docs]]"></div> 
  </div>`
}