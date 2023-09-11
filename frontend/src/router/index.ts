import { createRouter, createWebHashHistory,RouterOptions,Router, RouteRecordRaw } from 'vue-router'
import HelloWorld from '../components/HelloWorld.vue'
import Login from '../components/Login.vue'

const routes: RouteRecordRaw[] = [
  { path: '/', component: HelloWorld },
    { path: '/login', component: Login },
]
const options: RouterOptions = {
 history: createWebHashHistory(),
 routes,
}
const router: Router = createRouter(options)
export default router