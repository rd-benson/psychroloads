import { createRouter, createWebHashHistory } from "vue-router";
import HomeView from "@/views/Home.vue"
import WeatherView from "@/views/Weather.vue"

const router = createRouter({
  history: createWebHashHistory(),
  routes: [
    {
      path: "/home",
      name: "home",
      component: HomeView,
    },
    {
      path: "/weather",
      name: "weather",
      component: () => import('@/views/Weather.vue'),
    },
  ],
});

export default router;
