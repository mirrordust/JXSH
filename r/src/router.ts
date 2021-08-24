import { createRouter, createWebHashHistory } from 'vue-router'

import TheHome from './components/TheHome.vue'
import TheTagged from './components/TheTagged.vue'
import TheAboutMe from './components/TheAboutMe.vue'


const routes = [
    { path: "/", component: TheHome },
    { path: "/tagged", component: TheTagged },
    { path: "/about", component: TheAboutMe }
]

export default createRouter({
    history: createWebHashHistory(),
    routes
})
