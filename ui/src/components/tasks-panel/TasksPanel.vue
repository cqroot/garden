<script setup lang="ts">
import { ref } from "vue";
import TaskList from "./components/TaskList.vue";

const taskListRef = ref();
const taskDialogRef = ref();

function handleEditTask(id: number) {
  taskDialogRef.value.editTask(id);
}

function handleCreateTask() {
  taskDialogRef.value.createTask();
}

function updateTasks(dueStart: number, dueEnd: number) {
  taskListRef.value.updateTasks(dueStart, dueEnd);
}

function updateTask(id: number) {
  taskListRef.value.updateTask(id);
}

defineExpose({
  updateTasks,
  updateTask,
});
</script>

<template>
  <n-layout position="absolute">
    <n-layout
      :native-scrollbar="false"
      position="absolute"
      style="bottom: 64px; padding: 24px"
    >
      <task-list ref="taskListRef" @on-task-click="handleEditTask" />
    </n-layout>
    <n-layout-footer bordered position="absolute" style="height: 64px">
      <n-button
        strong
        secondary
        type="primary"
        class="w-full h-full rounded-none"
        style="font-weight: bold; font-size: 15px"
        @click="handleCreateTask"
      >
        Create a new task
      </n-button>
    </n-layout-footer>
  </n-layout>

  <task-dialog
    ref="taskDialogRef"
    @on-task-create="updateTasks"
    @on-task-update="updateTask"
  />
</template>
