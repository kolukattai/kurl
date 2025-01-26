const DocsComponent = {
  props: ["docs", "editDoc"],
  setup(props) {
    const docData = ref("");

    Vue.onMounted(() => {
      let converted = htmlToMarkdown(props.docs)

      console.log(props.docs);
      console.log("\n\n");
      console.log(converted, );
      
      docData.value = converted;
    })

    return {
      docs: unescapeHtml(props.docs),
      docData,
    }
  },
  template: `<div>
    <div v-html="[[docs]]"></div> 
  </div>`
}