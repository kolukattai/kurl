
const drawerComponent = {
  props: ["item", "search"],
  setup(props) {

    const fileName = computed(() => {
      let n = String(props.item.fileName).replace(/\-/g, " ").replace(".md", "")
      return `${n[0].toUpperCase()}${n.substring(1,n.length)}`;
    });

    const navigate = computed(() => {
      let id = btoa(props.item.filePath)
      return `#call:${id}`
    })

    const showList = computed(() => {
      let va = JSON.stringify(props.item).toLowerCase()
      .indexOf(String(props.search).toLowerCase()) != -1
      let va1 = JSON.stringify(props.item).toLowerCase().replaceAll("-", " ")
      .indexOf(String(props.search).toLowerCase()) != -1
      return va || va1
    })



    const showFolder = computed(() => {
      let va = JSON.stringify(props.item.files).toLowerCase()
      .indexOf(String(props.search).toLowerCase()) != -1
      let va1 = JSON.stringify(props.item.files).toLowerCase().replaceAll("-", " ")
      .indexOf(String(props.search).toLowerCase()) != -1
      return va || va1
    })

    const currentHash = ref("");
    

    // Event listener function for hash changes
    const onHashChange = () => {
      currentHash.value = window.location.hash
    };



    Vue.onMounted(() => {
      currentHash.value = window.location.hash
      window.addEventListener('hashchange', onHashChange);
    })
    
    Vue.onUnmounted(() => {
      window.addEventListener('hashchange', onHashChange);
    })

    return {
      fileName: fileName,
      values: props.item,
      navigate: navigate,
      showList, showFolder,
      currentHash,
    }
  },
  template: `
    <li class="drawer__item" >
      <span class="drawer__item__folder" v-if="values.isFolder && showFolder">[[fileName]]</span>
      <a :class="{
      'drawer__item__link': true,
      'drawer__item__link--active': currentHash == navigate,
      }" v-else :href="navigate" v-show="showList">[[fileName]]</a>
      <ul v-if="!!item.files.length">
        <drawer-component :search="search" :item="itm" v-for="(itm, i) in item.files" :key="i" />
      </ul>
    </li>
  `
}