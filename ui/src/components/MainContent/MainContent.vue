<script setup lang="ts">
import { ref, onMounted } from "vue";
import { useNotification } from "./notification";
import TaskList from "@/components/TaskList/TaskList.vue";
import NewTaskPanel from "@/components/NewTaskPanel.vue";
import EditModal from "@/components/EditModal.vue";

useNotification();

const containerRef = ref();

const taskListRef = ref();
const editModalRef = ref();

const onTaskClick = (id: number) => {
  editModalRef.value.open(id);
};

const onTaskSubmit = (id: number) => {
  taskListRef.value.updateTask(id);
};

const onTaskPut = () => {
  taskListRef.value.updateTasks();
};

onMounted(() => {
  taskListRef.value.updateTasks();
});
</script>

<template>
  <div ref="containerRef">
    <new-task-panel @onTaskPut="onTaskPut" class="mb-6" />

    <!-- <n-divider /> -->

    <task-list ref="taskListRef" @onTaskClick="onTaskClick" />
    <edit-modal ref="editModalRef" @onTaskSubmit="onTaskSubmit" />
  </div>
</template>
