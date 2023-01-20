<script setup lang="ts">
import { ref, onMounted } from "vue";
import { tomorrowTimestamp } from "@/api/date";
import TasksPanel from "@/components/tasks-panel/TasksPanel.vue";

const tasksPanelRef = ref();

onMounted(() => {
  const tomorrowTs = tomorrowTimestamp();
  tasksPanelRef.value.updateTasks(tomorrowTs, 0);
});
</script>

<template>
  <tasks-panel ref="tasksPanelRef" />
</template>
