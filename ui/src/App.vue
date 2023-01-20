<script setup lang="ts">
import { RouterView } from "vue-router";
import MainLayout from "@/layouts/MainLayout.vue";
</script>

<template>
  <n-notification-provider>
    <n-message-provider>
      <main-layout>
        <router-view />
      </main-layout>
    </n-message-provider>
  </n-notification-provider>
</template>
