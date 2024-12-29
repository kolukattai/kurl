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
      if (typeof props.data.body == "object") {
        return JSON.stringify(props.data.body, null, 4)
      }
      return ""
    })

    const showHeaders = ref(false);

    return {
      data: props.data,
      headers: headers,
      body: body, showHeaders,
    }
  },

  template: `
    <div>

      <request-component :request="data.request" />

      <hr />

      <h1>Response</h1>

      <h3>Status: [[data.status]]</h3>


      <div v-if="!!body">
        <h3>Body</h3>
        <code-view :text="body"></code-view>
      </div>


      <h3 @click="showHeaders = !showHeaders" style="cursor: pointer;
        background: #2d2d2d;
        padding: 10px 15px;">
        <span> [[showHeaders ? 'Hide' : 'Show']] Headers</span>
      </h3>

      <table class="response-header" v-if="showHeaders">
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
  `

}

const responseComponent = {
  props: ["response"],

  setup(props) {
    const emit = defineEmits()
    const selected = ref(0);
    const showDeletable = ref(false);

    const headers = ["one", "two", "three"];

    Vue.onMounted(() => {
      let isDev = document.querySelector("#dev-mode")
      if (!!!isDev) {
        return
      }
      showDeletable.value = true
    })

    const saveTemplate = (index) =>{
      return `Saved ${index+1}`
    }
    
    const deleteSavedResponse = (item, i) => {
      const yes = confirm(`Do your rely want to delete this response "${!!item.request.name ? item.request.name : 'Response ' + (i + 1)}"`)
      if (!yes) {
        return
      }

      try {
        fetch(`/saved/${item.request.refID}/${i}`, {
          method: "DELETE"
        }).then(res => res.json())
        .then(res => {
          console.log(res, Vue, emit);
          // emit("newSaved", res)
          window.location.reload()
        }).catch(err => {
          console.error(err);
        })
      } catch (err) {
        console.error(err);
      }
    }

    return {
      response: props.response,
      selected, saveTemplate,
      deleteSavedResponse,
      showDeletable,
    }
  },

  template: `<div>
    <div class="tabs" v-show="response.length > 1" style="padding-bottom: 10px;">
      <div :class="{
        'tab': true,
        'tab--active': selected == i
      }" 
        v-for="(item, i) in response" :key="i"
      >
        <span @click="selected = i" >[[!!item.request.name ? item.request.name : saveTemplate(i)]]</span>
        <button v-if="showDeletable" class="tab-close-btn" @click="deleteSavedResponse(item, i)" >
          <img src="/static/images/close.svg" height="24px" />
        </button>
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

const ResponseAction = {
  props: ["id"],
  setup() {
    return {}
  },
  template: `<div>
    
  </div>`
}

