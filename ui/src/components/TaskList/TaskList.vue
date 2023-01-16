<script setup lang="ts">
import { useTasks } from "./task";
const { tasks, updateTasks, updateTask } = useTasks();

const emit = defineEmits<{
  (e: "onTaskClick", id: number): void;
}>();

const onTaskClick = (id: number) => {
  emit("onTaskClick", id);
};

defineExpose({
  updateTasks,
  updateTask,
});
</script>

<template>
  <n-list hoverable clickable bordered>
    <n-list-item
      v-for="task in tasks"
      :key="task.id"
      @click="onTaskClick(task.id)"
    >
      <n-thing :title="task.title">
        {{ task.note }}
      </n-thing>
    </n-list-item>
  </n-list>
</template>
