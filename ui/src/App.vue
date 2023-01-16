<script setup lang="ts">
import MainContent from "@/components/MainContent/MainContent.vue";
</script>

<template>
  <n-notification-provider>
    <n-layout>
      <n-layout-header bordered> Garden </n-layout-header>
      <n-layout-content
        :native-scrollbar="false"
        content-style="padding: 0 24px;"
      >
        <main-content />
      </n-layout-content>
    </n-layout>
  </n-notification-provider>
</template>

<style scoped>
.n-layout-header {
  height: 64px;
  padding: 24px;
  font-weight: bold;
  font-size: 20px;
}

.n-layout-content {
  max-width: 1200px;
  margin: 0 auto;
  padding: 24px;
  height: calc(100vh - 64px);
}
</style>
