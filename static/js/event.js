
// Global Event Bus using reactive object
const EventBus = Vue.reactive({
  listeners: {},
  $on(event, callback) {
    if (!this.listeners[event]) {
      this.listeners[event] = [];
    }
    this.listeners[event].push(callback);
  },
  $emit(event, data) {
    if (this.listeners[event]) {
      this.listeners[event].forEach(callback => callback(data));
    }
  },
  $off(event, callback) {
    if (this.listeners[event]) {
      const index = this.listeners[event].indexOf(callback);
      if (index !== -1) {
        this.listeners[event].splice(index, 1);
      }
    }
  }
});

const ObjToArr = (obj) => {
  let arr = [];
  for (const key in obj) {
    arr.push({key: key, value: obj[key]})
  }
  return arr
}
