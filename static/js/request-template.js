const requestTemplateComponent = {
  props: ["request"],
  setup(props) {
    const selected = ref("");
    const templates = ref([]);
    
    Vue.onMounted(() => {
      const arr = [];
      let req = {...props.request}
      for (const key in req) {
        let fVal = String(req[key])
        .replaceAll(/\s\-/g,"\\\n\t-")
        .replaceAll(/\\n/g, "\n")
        .replace("{\"", "{\n\t\"")
        .replace("},", "},\n")

        let val = {key: key, value: fVal}
        console.log("val",val);
        arr.push(val)
      }
      templates.value = arr;
      if (!!arr.length) {
        selected.value = arr[0].key
      }
      console.log("templates 1", arr);
      console.log("templates 2", req);
    })


    return {
      templates, selected,
    }
  },
  template: `<div style="padding: 15px 0px">
    <div class="tabs">
      <div :class="{
        'tab': true,
        'tab--active': selected == item.key,
        'type-1': true,
      }" v-for="(item, i) in templates" :key="i" @click="selected = item.key">
        [[item.key]]
      </div>
    </div>
    <code-view 
      v-for="(item, i) in templates" 
      :key="i" v-show="item.key == selected"
      :text="item.value"
    />
  </div>`
}