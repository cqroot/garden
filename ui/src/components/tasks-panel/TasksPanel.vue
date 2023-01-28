<script setup lang="ts">
import { ref } from "vue";
import NewTaskPanel from "./components/NewTaskPanel.vue";
import TaskList from "./components/TaskList.vue";

const taskListRef = ref();
const taskDialogRef = ref();

const handleTaskClick = (id: number) => {
  taskDialogRef.value.editTask(id);
};

const updateTasks = (dueStart: number, dueEnd: number) => {
  taskListRef.value.updateTasks(dueStart, dueEnd);
};

const updateTask = (id: number) => {
  taskListRef.value.updateTask(id);
};

defineExpose({
  updateTasks,
  updateTask,
});
</script>

<template>
  <new-task-panel @on-task-put="updateTasks" />
  <task-list ref="taskListRef" @on-task-click="handleTaskClick" />
  <task-dialog
    ref="taskDialogRef"
    @on-task-create="updateTasks"
    @on-task-update="updateTask"
  />
</template>
