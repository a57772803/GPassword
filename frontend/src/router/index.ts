import { createRouter, createWebHashHistory,RouterOptions,Router, RouteRecordRaw } from 'vue-router'
import HelloWorld from '../components/HelloWorld.vue'
import Login from '../components/Login.vue'
import DashBoard from '../components/DashBoard.vue'

const routes: RouteRecordRaw[] = [
  { path: '/', component: HelloWorld },
  { path: '/login', component: Login },
        { path: '/DashBoard', component: DashBoard },
]
const options: RouterOptions = {
 history: createWebHashHistory(),
 routes,
}
const router: Router = createRouter(options)
export default router