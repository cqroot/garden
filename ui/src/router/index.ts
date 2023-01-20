import { createRouter, createWebHistory } from "vue-router";
import {
  Albums as AllIcon,
  Sunny as TodayIcon,
  Calendar as UpcomingIcon,
} from "@vicons/ionicons5";

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: "/",
      redirect: "/all",
    },
    {
      path: "/all",
      name: "all",
      component: () => import("@/views/AllView.vue"),
      meta: {
        title: "All",
        icon: AllIcon,
      },
    },
    {
      path: "/today",
      name: "today",
      component: () => import("@/views/AllView.vue"),
      meta: {
        title: "Today",
        icon: TodayIcon,
        withDivider: true,
      },
    },
    {
      path: "/upcoming",
      name: "upcoming",
      component: () => import("@/views/AllView.vue"),
      meta: {
        title: "Upcoming",
        icon: UpcomingIcon,
        withDivider: true,
      },
    },
  ],
});

export default router;
