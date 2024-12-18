const { createApp, ref, onMounted, computed, h } = Vue


const app = createApp({
  setup() {
    const drawerData = ref([])

    const updateDrawer = async () => {
      try {
        const res = await fetch("/data/files.json", {
          method: "GET"
        })
        const data = await res.json()

        drawerData.value = [...data.data].filter((e) => ( !(e.fileName == "README.md" || e.fileName == "index.md" )))
        console.log(data);
      } catch (err) {
        console.error(err);
      }
    }

    onMounted(() => {
      updateDrawer()
    })


    const search = ref("");

    return {
      drawerData,
      search,
    }
  }
})


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


