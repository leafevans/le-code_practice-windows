import { createRouter, createWebHistory } from 'vue-router'
import Layout from '../layout/Layout.vue'

const Home = () => import('../views/Home.vue')
const Registration = () => import('../views/Registration.vue')
const RegistrationList = () => import('../views/RegistrationList.vue')
const MedicalRecord = () => import('../views/MedicalRecord.vue')
const Medicine = () => import('../views/Medicine.vue')
const Doctor = () => import('../views/Doctor.vue')
const DoctorWorkbench = () => import('../views/DoctorWorkbench.vue')
const Charge = () => import('../views/Charge.vue')
const Department = () => import('../views/Department.vue')
const Statistics = () => import('../views/Statistics.vue')
const Login = () => import('../views/Login.vue')
const NotFound = () => import('../views/NotFound.vue')

const routes = [
    {
        path: '/login',
        name: 'Login',
        component: Login
    },
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
                path: '/registration-list',
                name: 'RegistrationList',
                component: RegistrationList
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
                path: '/doctor',
                name: 'Doctor',
                component: Doctor
            },
            {
                path: '/doctor-workbench',
                name: 'DoctorWorkbench',
                component: DoctorWorkbench
            },
            {
                path: '/charge',
                name: 'Charge',
                component: Charge
            },
            {
                path: '/department',
                name: 'Department',
                component: Department
            },
            {
                path: '/statistics',
                name: 'Statistics',
                component: Statistics
            }
        ]
    },
    {
        path: '/:pathMatch(.*)*',
        name: 'NotFound',
        component: NotFound
    }
]

const router = createRouter({
    history: createWebHistory(),
    routes
})

router.beforeEach((to, from, next) => {
    const token = localStorage.getItem('token')
    if (to.path !== '/login' && !token) {
        next('/login')
    } else {
        next()
    }
})

export default router
