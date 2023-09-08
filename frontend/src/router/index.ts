import { createRouter, createWebHashHistory,RouterOptions,Router, RouteRecordRaw } from 'vue-router'
import HelloWorld from '../components/HelloWorld.vue'
import  Encrypt from '../components/Encrypt.vue'

const Home = { template: '<div>Home</div>' }
const About = { template: '<div>About</div>' }

const routes: RouteRecordRaw[] = [
  { path: '/', component: HelloWorld },
    { path: '/encrypt', component: Encrypt },
  { path: '/about', component: About },
]
const options: RouterOptions = {
 history: createWebHashHistory(),
 routes,
}
const router: Router = createRouter(options)
export default router