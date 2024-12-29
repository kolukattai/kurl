const requestComponent = {
  props: ["request"],

  setup(props) {

    const selected = ref("curl");
    const templates = ref([]);

    Vue.onMounted(() => {
      updateTemplates()
    });

    const updateTemplates = () => {
      let arr = []
      arr.push({
        key: "curl",
        value: curlTemplate()
      })
      arr.push({
        key: "JavaScript",
        value: fetchTemplate()
      })
      templates.value = arr.map((e) => {
        let el = e
        el.value = updateVariables(e.value)
        return el
      });
    }

    const updateVariables = (val) => {
      let env = sessionStorage.getItem("env")
      if (!!!env) {
        return val
      }
      try {
        let arr = []
        arr = JSON.parse(env)
        arr.forEach((e) => {
          let fromKey = `{{${e.key}}}`
          val = val.replaceAll(fromKey, e.value)
        })
        return val
      } catch (err) {
        console.error(err);
        return val
      }
    }

    const curlTemplate = () => {
      let result = `curl -X ${props.request.method} ${props.request.url}`

      if (!!props.request.headers) {
        for (const key in props.request.headers) {
          result += `\\\n\t-H "${key}: ${JSON.stringify(props.request.headers[key]).replace(/\"/g, '')}"`
        }
      }

      if (!!props.request.body) {
        let bodyType = typeof props.request.body
        if (bodyType == "object") {
          result += `\\\n\t-d '${JSON.stringify(props.request.body)}'`
        } else {
          result += `\\\n\t-d '${props.request.body}'`
        }
      }

      return result
    }

    const fetchTemplate = () => {

      let requestHeader = () => {
        if (!!!props.request.headers) {
          return ""
        }
        let headObj = ""
        for (const key in props.request.headers) {
          headObj += `\t"${key}": "${props.request.headers[key]}"\n`
        }
        let headerVal = `,\n\theaders: {\n\t${headObj}\t}`
        return headerVal
      }
      let requestBody = () => {
        if (!!!props.request.body) {
          return ""
        }
        let obj = JSON.stringify(props.request.body, "", "\t\t")
        return `,\n\tbody: \`${obj}\``.replace("}", "\t}") + "\n"
      }

      let result = `fetch("${props.request.url}", {
        method: "${String(props.request.method).toUpperCase()}"${requestHeader()}${requestBody()}})\n\t.then(res => {\n\t\tconsole.log(res)\n\t})\n\t.catch(err => {\n\t\tconsole.error(err)\n\t})`
      return result
    }


    return {
      selected,
      templates,
    }
  },
  template: `
  <div style="margin-bottom: 70px">
    <div class="tabs">
      <div :class="{
        'tab': true,
        'tab--active': selected == item.key,
        'type-1': true,
      }" v-for="(item, i) in templates" :key="i" @click="selected = item.key">
        <span >[[item.key]]</span>
      </div>
    </div>
    <code-view 
      v-for="(item, i) in templates" 
      :key="i" v-show="item.key == selected"
      :text="item.value"
    />
  </div>
  `
}