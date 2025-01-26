const CodeView = {
  props: ["text"],
  setup(props) {
    const message = ref("copy");

    const copyVal = () => {
      if (message == "Copied!") {
        return
      }
      message.value = "Copied!"
      window.navigator.clipboard.writeText(props.text)
      .catch(err => console.error(err))
      setTimeout(() => {
        message.value = "copy"
      }, 3000);
    }

    Vue.onMounted(() => {

    })

    return {message, copyVal, text: props.text}
  },
  template: `<div class="code-view">
    <button class="code-view__copy" @click="copyVal()">
      <tool-tip :class="message != 'copy' ? 'active' : ''">[[message]]</tool-tip>
      <img v-if="message == 'copy'" src="/static/images/copy.svg" height="24px" width="24px" alt="copy" />
      <img v-else src="/static/images/check.svg" height="24px" width="24px" alt="copy" />
    </button>
    <pre><code>[[text]]</code></pre>
  </div>`
}

const ToolTip = {
  props: ["class"],
  setup(props) {

    const className = Vue.computed(() => `tooltip ${props.class}`)

    return {
      className,
    }
  },
  template: `<span :class="className"><slot/></span>`
}
