import Vue from 'vue';
import VueRouter from 'vue-router';
import Home from '../views/Home.vue';

Vue.use(VueRouter);

const routes = [
  {
    path: '/',
    name: 'Home',
    component: Home,
  },
  {
    path: '/about',
    name: 'About',
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () => import(/* webpackChunkName: "about" */ '../views/About.vue'),
  },
  {
    path: '/rooms',
    name: 'Rooms',
    component: () => import(/* webpackChunkName: "rooms" */ '../views/ListRooms.vue'),
  },
  {
    path: '/rooms/new',
    name: 'NewRoom',
    component: () => import(/* webpackChunkName: "newroom" */ '../views/CreateRoom.vue'),
  },
  {
    path: '/room/:id',
    name: 'Room',
    component: () => import(/* webpackChunkName: "room" */ '../views/Room.vue'),
  },
  {
    path: '/room/:id/lobby',
    name: 'LobbyRoom',
    component: () => import(/* webpackChunkName: "roomlobby" */ '../views/RoomLobby.vue'),
  },
  {
    path: '/room/:id/end',
    name: 'RoomEnd',
    component: () => import(/* webpackChunkName: "roomend" */ '../views/RoomEnd.vue'),
  },
];

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes,
});

export default router;
