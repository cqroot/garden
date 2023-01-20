<script setup lang="ts">
import { ref } from "vue";
import NewTaskPanel from "./components/NewTaskPanel.vue";
import TaskList from "./components/TaskList.vue";
import EditModal from "./components/EditModal.vue";

const taskListRef = ref();
const editModalRef = ref();

const handleTaskClick = (id: number) => {
  editModalRef.value.open(id);
};

const handleTaskSubmit = (id: number) => {
  taskListRef.value.updateTask(id);
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
  <edit-modal ref="editModalRef" @on-task-submit="handleTaskSubmit" />
</template>
