import { createRouter, createWebHistory } from 'vue-router'
import Layout from '../layout/Layout.vue'
import Home from '../views/Home.vue'
import Registration from '../views/Registration.vue'
import MedicalRecord from '../views/MedicalRecord.vue'
import Medicine from '../views/Medicine.vue'
import Statistics from '../views/Statistics.vue'

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
            },
            {
                path: '/medical-record',
                name: 'MedicalRecord',
                component: MedicalRecord
            },
            {
                path: '/medicine',
                name: 'Medicine',
                component: Medicine
            },
            {
                path: '/statistics',
                name: 'Statistics',
                component: Statistics
            }
        ]
    }
]

const router = createRouter({
    history: createWebHistory(),
    routes
})

export default router
