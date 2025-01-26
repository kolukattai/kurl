const MethodUrl = {
  props: ["request"],
  setup(props, ctx) {

    const onChange = (key, value) => {
      ctx.emit('changeVal', {...props.request, [key]: value})
    }

    const request = computed(() => props.request)

    const makeCall = () => {
      ctx.emit("submitEve")
    }

    return {
      request,
      onChange,
      makeCall,
    }
  },
  template: `<form class="api-call-method-call" @submit.prevent="makeCall()">
    <select class="" :value="request.method" @change="(e) => onChange('method', e.target.value)">
      <option v-for="(item, i) in ['GET', 'PUT', 'POST', 'DELETE']" :value="item">[[item]]</option>
    </select>
    <input :value="request.url" @change="(e) => onChange('url', e.target.value)" placeholder="url goes here" />
    <button type="submit" class="btn">send</button>
  </form>`
}


const ReqHeader = {
  props: ["request"],
  setup(props, ctx) {

    const headers = computed(() => ObjToArr(props.request.headers))

    const changeEve = (key, val, index) => {
      let reqHeader = [...headers.value]
      reqHeader[index] = {key: key, value: val}
      let obj = {}
      for (const ele of reqHeader) {
        obj[ele.key] = ele.value
      }
      let req = {...props.request, headers: obj}
      console.log("::", req);
      
      ctx.emit("changeVal", req)
    }

    const removeHeader = (index) => {
      let reqHeader = [...headers.value]
      reqHeader.splice(index,1)
      let obj = {}
      for (const ele of reqHeader) {
        obj[ele.key] = ele.value
      }
      let req = {...props.request, headers: obj}
      ctx.emit("changeVal", req)
    }
    
    const addHeader = () => {
      ctx.emit("changeVal", {...props.request, headers: {...props.request.headers, "": ""}})
    }

    return {
      headers,
      changeEve,
      addHeader,
      removeHeader,
    }
  },
  template: `<div>
    <table class="table">
      <thead>
        <tr>
          <th>Key</th>
          <th>Value</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="(item, i) in headers" :key="i">
          <td>
            <input :value="item.key" placeholder="key" @change="changeEve($event.target.value, item.value, i)" />
          </td>
          <td style="display: flex;">
            <input :value="item.value" placeholder="value" @change="changeEve(item.key, $event.target.value, i)" />
            <img 
            @click="removeHeader(i)"
            v-if="request.length != 1" class="close" src="/static/images/close.svg" alt="close" height="24px" />
          </td>
        </tr>
      </tbody>
    </table>
    <button class="btn" style="display:flex;align-items: center;" @click="addHeader()">
      <span style="display: inline-block;padding-right: 8px;">ADD</span>
      <img src="/static/images/add.svg" alt="close" height="20px" />
    </button>
  </div>`
}

const ReqBody = {
  props: ["request"],
  setup(props, ctx) {
    const body = computed(() => props.request.body)
    return {
      body,
      request: props.request
    }
  },
  template: `<div>
    <textarea :disabled="['GET','DELETE'].includes(request.method)" :value="JSON.stringify(request.body, '', ' ')"></textarea>
  </div>`
}

const ApiCall = {
  props: ["request", "response"],
  setup(props, ctx) {

    const changeHandle = (e) => {
      console.log("ApiCall", e);
      
      ctx.emit("changeVal", e)
    }

    const request = computed(() => props.request)

    const submitHandle = () => {
      ctx.emit("submitEve", props.request)
    }

    const activeTab = ref(0);

    const responseHeader = computed(() => !!props.response ? ObjToArr(props.response.headers) : [])
    const responseBody = computed(() => !!props.response ? props.response.body_str : "")
    const responseStatus = computed(() => !!props.response ? props.response.status : "")
    const isResponse = computed(() => !!props.response)

    const responseTab = ref(0);

    return {
      request,
      changeHandle,
      submitHandle,
      activeTab,
      responseBody,
      responseHeader,
      responseTab,
      isResponse,
      responseStatus,
    }
  },
  template: `<div>
    <MethodUrl 
      :request="request" 
      @changeVal="changeHandle($event)"
      @submitEve="submitHandle()"
    />
    <div class="tabs" style="margin-top: 15px;">
      <button 
        v-for="(item, i) in ['Headers', 'Body']"
        :class="{
          'type-1': true,
          'tab': true,
          'active': i == activeTab
        }"
        @click="activeTab = i"
      >[[item]]</button>
    </div>
    <ReqHeader
      v-if="activeTab == 0"
      :request="request"
      @changeVal="changeHandle($event)"
    />
    <ReqBody
      v-if="activeTab == 1"
      :request="request"
      @changeVal="changeHandle($event)"
    />

    <div v-if="isResponse" style="padding-top: 15px;">
      <div style="display: flex;align-items: center;justify-content: space-between;">
        <h1>Response</h1>
        <button class="btn">SAVE</button>
      </div>

      <div class="tabs">
        <div :class="{
          'tab': true,
          'type-1': true,
          'active': i == responseTab,
        }" v-for="(item, i) in ['Body', 'Headers']" :key="i"
          @click="responseTab=i"
        >
          [[item]]
        </div>
        <div style="flex: 1;"></div>
        <div style="display: flex;align-items: center;">
          <span>[[responseStatus]]</span>
        </div>
      </div>

      <table class="response-header" v-if="responseTab == 1">
        <thead>
          <tr>
            <th>Key</th>
            <th>Value</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="(item, i) in responseHeader">
            <td>[[item.key]]</td>
            <td>[[item.value]]</td>
          </tr>
        </tbody>
      </table>

      <code-view :text="response.body" v-if="responseTab == 0">


    </div>

    </div>`
}