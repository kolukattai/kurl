const { createApp, ref, onMounted, computed, h, defineEmits } = Vue

const isDev = document.querySelector("#dev-mode")

const loadDevMode = async () => {
  try {
    if (!!isDev) {
      let res = await fetch("/data/env.json")
      res = await res.json()
      let envArr = [];
      if (res.env) {
        for (const key in res.env) {
          envArr.push({key, value: res.env[key]})
        }
        sessionStorage.setItem("env", JSON.stringify(envArr))
      }
    }
  } catch (err) {
    console.error(err);
  }
}


const app = createApp({
  setup() {
    const drawerData = ref([])
    const env = ref({});

    const updateDrawer = async () => {
      try {
        loadDevMode()
        const res = await fetch("/data/files.json", {
          method: "GET"
        })
        const data = await res.json()

        drawerData.value = [...data.data].filter((e) => ( !(e.fileName == "README.md" || e.fileName == "index.md" || !/\.md/gi.test(e.fileName) )))
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

app.component('code-view', CodeView)
app.component('env-component', EnvComponent)
app.component('drawer-component', drawerComponent);
app.component('home-component', homeComponent)
app.component('request-component', requestComponent)
app.component('response-component', responseComponent)
app.component('each-response', eachResponse)
app.component('docs-component', DocsComponent)
app.component('tool-tip', ToolTip)
app.component('request-template-component', requestTemplateComponent)
app.component('expandable-component', expandableComponent)
app.component('api-call', ApiCall)
app.component('MethodUrl', MethodUrl)
app.component('ReqHeader', ReqHeader)
app.component('ReqBody', ReqBody)


app.config.compilerOptions.delimiters = ['[[', ']]'];
app.mount('#app')


