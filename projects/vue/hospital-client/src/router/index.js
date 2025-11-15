import { createRouter, createWebHistory } from 'vue-router'
import Layout from '../layout/Layout.vue'
import Home from '../views/Home.vue'
import Registration from '../views/Registration.vue'

const routes = [
    {
        path: '/',
        component: Layout,
        children: [
            {
                path: '',
                name: 'Home',
                component: Home
            },
            {
                path: '/registration',
                name: 'Registration',
                component: Registration
            }
        ]
    }
]

const router = createRouter({
    history: createWebHistory(),
    routes
})

export default router
