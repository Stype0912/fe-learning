import Vue from 'vue'
import Router from 'vue-router'
import Counter from '@/components/HelloWorld'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      name: 'HelloWorld',
      component: Counter,
    }
  ]
})
