const homeComponent = {
  setup() {

    const currentHash = ref(window.location.hash);
    const hashType = ref("");
    const data = ref(null);

    const fetchData = async (id) => {
      try {
        data.value = null
        const res = await fetch(`/data/call/${id}.json`, {
          method: "GET",
        })
        const resData = await res.json()
        data.value = resData

      } catch (err) {
        console.error(err);
      }
    }

    // Event listener function for hash changes
    const onHashChange = () => {
      
      let id = window.location.hash;

      if (!!!id) {
        fetchData(btoa("README.md"))
        hashType.value = "#home"
        return
      }

      let idArr = String(id).split(":")

      hashType.value = idArr[0]
      currentHash.value = idArr[1]

      console.log(idArr[1]);
      
      if (idArr[0] == "#call") {
        fetchData(idArr[1])
      }
    };

    Vue.onMounted(() => {
      console.log("home mounted");
      onHashChange()
      window.addEventListener('hashchange', onHashChange);
    });

    Vue.onUnmounted(() => {
      console.log("home un mounted");

      window.removeEventListener('hashchange', onHashChange);
    });

    const selected = ref(0);

    const headers = ["one", "two", "three"];

    return {
      currentHash: currentHash,
      hashType: hashType,
      data: data,
      selected,
      headers,
    }
  },
  template: `<div style="height:100%;">
    <div v-if="!!data && hashType == '#call'" class="home-layout">
      <div class="home-layout__request">
        <h1>[[data.name]]</h1>
        <request-component :request="data.request" />
        <docs-component :docs="data.docs" />
      </div>
      <div class="home-layout__response">
        <response-component :response="data.response" />
      </div>
    </div>
    <div v-if="hashType == '#home'" class="index-page">
      <h1 class="index-page__title">[[data.request.title]]</h1>
      <div class="index-page__contents">
        <div class="index-page__contents__content" v-for="(item, i) in data.request.content" :key="i">
          <h2>[[item.name]]</h2>
          <p>[[item.content]]</p>
        </div>
      </div>
      <docs-component :docs="data.docs" />
    </div>
  </div>`
}

    // <div v-if="!!data">
    //   <div v-if="!!data.request">
    //      <request-component :request={data.request} />
    //   </div>
    // </div>
