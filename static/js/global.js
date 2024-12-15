const { createApp, ref, onMounted, computed, h } = Vue


const app = createApp({
  setup() {
    const drawerData = ref([])

    const updateDrawer = async () => {
      try {
        const res = await fetch("/data/files", {
          method: "GET"
        })
        const data = await res.json()

        drawerData.value = data.data
        console.log(data);
      } catch (err) {
        console.error(err);
      }
    }

    onMounted(() => {
      updateDrawer()
    })


    return {
      drawerData
    }
  }
})

const drawerComponent = {
  props: ["item"],
  setup(props) {

    const fileName = computed(() => {
      let n = String(props.item.fileName).replace(/\-/g, " ").replace(".md", "")
      return `${n[0].toUpperCase()}${n.substring(1,n.length)}`;
    });

    const navigate = computed(() => {
      let id = btoa(props.item.filePath)
      return `#call:${id}`
    })

    return {
      fileName: fileName,
      values: props.item,
      navigate: navigate,
    }
  },
  template: `
    <li>
      <span v-if="values.isFolder">[[fileName]]</span>
      <a v-else :href="navigate">[[fileName]]</a>
        <ul v-if="!!item.files.length">
          <drawer-component :item="item" v-for="(item, i) in item.files" :key="i" />
        </ul>
    </li>
  `
}

function unescapeHtml(str) {
  const element = document.createElement('div');
  element.innerHTML = str;  // Decode HTML entities to their characters
  return element.innerHTML; // Retains HTML tags intact 
}

app.component('drawer-component', drawerComponent);
app.component('home-component', homeComponent)
app.component('request-component', requestComponent)
app.component('response-component', responseComponent)
app.component('each-response', eachResponse)
app.component('docs-component', DocsComponent)

app.config.compilerOptions.delimiters = ['[[', ']]'];
app.mount('#app')


