import { createRouter, createWebHistory } from 'vue-router'
import Board from '../views/Board.vue'
import NotFound from '../views/NotFound.vue'

const routes = [
  {
    path: '/board/:id',
    name: 'Home',
    component: Board,
    props: true
  },
  {
    path: "/:catchAll(.*)",
    name: 'NotFound',
    component: NotFound,
  },
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

export default router
