<script setup lang="ts">
import { ref, onMounted } from "vue";
import { todayTimestamp } from "@/api/date";
import TasksPanel from "@/components/tasks-panel/TasksPanel.vue";

const tasksPanelRef = ref();

onMounted(() => {
  const todayTs = todayTimestamp();
  tasksPanelRef.value.updateTasks(todayTs, todayTs);
});
</script>

<template>
  <tasks-panel ref="tasksPanelRef" />
</template>
