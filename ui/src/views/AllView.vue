<script setup lang="ts">
import { ref, onMounted } from "vue";
import TasksPanel from "@/components/tasks-panel/TasksPanel.vue";

const tasksPanelRef = ref();

onMounted(() => {
  tasksPanelRef.value.updateTasks(0, 0);
});
</script>

<template>
  <tasks-panel ref="tasksPanelRef" />
</template>
