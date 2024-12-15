const requestComponent = {
  props: ["request"],

  setup(props) {

    Vue.onMounted(() => {
      console.log("request mounted", props.request);
    });


    const className = Vue.computed(() => {
      return `request-bar__method request-bar__method--${String(props.request.method).toLowerCase()}` 
    })

    const requestParams = Vue.computed(() => {
      let req = []
      
      if (!!props.request.headers) {
        req.push("headers")
      }
      
      if (!!props.request.body) {
        req.push("body")
      }

      if (!!req.length) {
        selected.value = req[0]
      }

      return req
    })

    const selected = ref("");


    const headers = Vue.computed(() => {
      const headerList = []

      for (const key in props.request.headers) {
        if (Object.prototype.hasOwnProperty.call(props.request.headers, key)) {
          const value = props.request.headers[key];
          headerList.push({key: key, value: value})
        }
      }

      return headerList
    })


    const body = Vue.computed(() => {
      if (JSON.stringify(props.request.headers).toLowerCase().includes("application/json")) {
        return `<pre>
  <code class="language-json">
    ${JSON.stringify(props.request.body, null, 4)}
  </code>
</pre>`
      }
      return ""
    })


    Vue.onMounted(() => {
      if (!!requestParams.length) {
        selected.value = requestParams[0]
      }
    })


    return {
      request: props.request,
      className: className,
      requestParams,
      selected,
      headers,
      body,
    }
  },
  template: `
  <div>

    <h2>Request</h2>

    <div class="request-bar">
      <div :class="className">[[request.method]]</div>
      <div class="request-bar__url">
        [[request.url]]
      </div>
    </div>

    <div class="tabs">
      <div v-for="(item, i) in requestParams" :key="i" 
      :class="{
        'tab': true,
        'tab--active': selected == item
      }" 
      @click="selected = item">
        [[item]]
      </div>
    </div>

    <div v-if="selected == 'headers'">
      <table class="response-header">
        <thead>
          <tr>
            <th>Key</th>
            <th>Value</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="(item, i) in headers" :key="i">
            <td>[[item.key]]</td>
            <td>[[item.value]]</td>
          </tr>
        </tbody>
      </table>
    </div>


    <div v-if="selected == 'body'">

      <div v-html="body"></div>

    </div>

  </div>
  `
}