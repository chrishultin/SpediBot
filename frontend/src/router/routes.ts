import type { RouteRecordRaw } from 'vue-router';

const routes: RouteRecordRaw[] = [
  {
    path: '/',
    component: () => import('layouts/MainLayout.vue'),
    meta: {
      requiresAuth: true,
    },
    children: [
      { path: '', component: () => import('pages/IndexPage.vue') },
      { path: '/servers/:serverID', component: () => import('pages/ManageServerPage.vue') },
    ],
  },
  {
    path: '/login',
    component: () => import('layouts/UnauthenticatedLayout.vue'),
    children: [{ path: '', component: () => import('pages/LoginPage.vue') }],
  },

  // Always leave this as last one,
  // but you can also remove it
  {
    path: '/:catchAll(.*)*',
    component: () => import('pages/ErrorNotFound.vue'),
  },
];

export default routes;
