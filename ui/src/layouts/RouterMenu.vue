<script setup lang="ts">
import { computed, h } from "vue";
import type { Component } from "vue";
import { RouterLink, useRoute } from "vue-router";
import { NIcon } from "naive-ui";
import type { MenuOption } from "naive-ui";
import router from "@/router";

function renderIcon(icon: Component) {
  return () => h(NIcon, null, { default: () => h(icon) });
}

const route = useRoute();
const routes = router.getRoutes();
const menuOptions = [] as MenuOption[];

for (const r of routes) {
  if (r.path === "/") {
    continue;
  }

  menuOptions.push({
    label: () =>
      h(
        RouterLink,
        {
          to: {
            name: r.path.substring(1),
          },
        },
        { default: () => r.meta.title }
      ),
    key: r.path,
    icon: renderIcon(r.meta.icon as Component),
  });

  if (r.meta.withDivider === true) {
    menuOptions.push({
      type: "divider",
    });
  }
}

const currentPath = computed(() => route.path);
</script>

<template>
  <n-menu :options="menuOptions" :value="currentPath" />
</template>
