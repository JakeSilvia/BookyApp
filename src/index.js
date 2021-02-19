import Vue from 'vue';
import Buefy from 'buefy';
import App from './App.vue';
import router from './router';
import ajax from './requests/requests'
import messaging from './messaging/messaging'

Vue.config.productionTip = false;
Vue.use(Buefy, {
  defaultIconPack: 'mdi',
});
let APP = {}
APP.Ajax = ajax
APP.Messaging = messaging
window.APP = APP

Vue.prototype.$ajax = ajax
Vue.prototype.$messageCallbacks = []
Vue.prototype.$listenForEvents = function (callback) {
  Vue.prototype.$messageCallbacks.push(callback)
}

messaging.setupWebSocketConn(function (event){
  for (let i = 0; i < Vue.prototype.$messageCallbacks.length ; i++) {
    Vue.prototype.$messageCallbacks[i](event)
  }
})

Vue.prototype.$notification = function (type, message) {
  this.$buefy.toast.open({
    message: message,
    type: type
  })
}


new Vue({
  el: '#app',
  router,
  render: (h) => h(App),
}).$mount('#app');
