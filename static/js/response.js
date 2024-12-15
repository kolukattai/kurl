const eachResponse = {

  props: ["data", "index"],

  setup(props) {

    const headers = Vue.computed(() => {
      const headerList = []

      for (const key in props.data.headers) {
        if (Object.prototype.hasOwnProperty.call(props.data.headers, key)) {
          const value = props.data.headers[key];
          headerList.push({key: key, value: value})
        }
      }

      return headerList
    })

    const body = Vue.computed(() => {
      if (JSON.stringify(props.data.headers).toLowerCase().includes("application/json")) {
        return `<pre>
  <code class="language-json">
    ${JSON.stringify(props.data.body, null, 4)}
  </code>
</pre>`
      }
      return ""
    })

    return {
      data: props.data,
      headers: headers,
      body: body,
    }
  },

  template: `
    <div>

      <h3>Headers</h3>

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




      <div v-if="!!body">
        <h3>Body</h3>
        <div v-html="body"></div>
      </div>
    </div>
  `

}

const responseComponent = {
  props: ["response"],

  setup(props) {
    const selected = ref(0);

    const headers = ["one", "two", "three"];

    return {
      response: props.response,
      selected,
    }
  },

  template: `<div>

    <h2>Saved Responses</h2>

    <div class="tabs">
      <div :class="{
        'tab': true,
        'tab--active': selected == i
      }" 
        v-for="(item, i) in response" :key="i"
        @click="selected = i" 
      >
        Status [[item.status]]
      </div>
    </div>


    <each-response 
      v-for="(item, i) in response" :key="i"
      :data="item"
      :index="i"
      v-show="i == selected"
    />
  </div>`
}