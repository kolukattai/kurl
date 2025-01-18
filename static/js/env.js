const EnvComponent = {
  props: [],
  setup(props) {
    
    const envVariables = ref([]);

    const isDev = document.querySelector("#dev-mode")
    
    const saveValue = () => {
      sessionStorage.setItem("env", JSON.stringify(envVariables.value))

      let obj = {}

      for (const el of envVariables.value) {
        obj[el.key] = el.value
      }

      if (!!isDev) {
        fetch("/data/env.json", {
          method: "POST",
          headers: {
            "Content-Type": "application/json"
          },
          body: JSON.stringify(obj)
        }).then(res => {
          console.log(res);
        }).catch(err => {
          console.error(err);
        })
      }

      return initData()
    }
    
    Vue.onMounted(() => {
      initData()      
    })

    const initData = () => {
      let value = sessionStorage.getItem("env")
      if (!!!value) {
        let pl = [{key: "", value: ""}];
        envVariables.value = pl
        sessionStorage.setItem("env", JSON.stringify(pl))
        return
      }
      try {
        envVariables.value = JSON.parse(value)
      } catch (err) {
        sessionStorage.clear()
        window.location.reload()
      }
    }


    const addVariable = () => {
      envVariables.value = [...envVariables.value, {key: "", value: ""}]
    }

    return {
      saveValue,
      envVariables,
      addVariable,
    }
  },
  template: `
        <form class="env-view" @submit.prevent="saveValue()">
          <div class="env-view-header">
            <h1>Env Variables</h1>
            <button class="save-btn" type="submit">SAVE</button>
          </div>
          <div class="env-input-layout" v-for="(item, i) in envVariables" :key="i">
            <input v-model="item.key" placeholder="Key" type="text" class="flex-1">
            <input v-model="item.value" placeholder="Value" type="text" class="flex-1">
          </div>
          <div >
            <button type="button" @click="addVariable()" class="add-env" style="display: flex; align-items: center;">
              <img src="/static/images/add.svg" height="20px" /> 
              <span style="padding-left: 10px">ADD VARIABLE</span>
            </button>
          </div>
        </form>
`
}
