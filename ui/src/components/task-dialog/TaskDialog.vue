<script setup lang="ts">
import DatePickButton from "./DatePickButton.vue";
import { useTaskDialog } from "./task-dialog";

const emit = defineEmits<{
  (e: "onTaskCreate"): void;
  (e: "onTaskUpdate", id: number): void;
}>();
const { isModalVisible, timestamp, task, submitTask, createTask, editTask } =
  useTaskDialog(emit);

defineExpose({
  createTask,
  editTask,
});

const rules = {
  title: {
    required: true,
    message: "Please input task title",
    trigger: "blur",
  },
  note: {
    required: false,
    message: "Please input task note",
    trigger: ["input"],
  },
};
</script>

<template>
  <n-modal v-model:show="isModalVisible">
    <n-card
      style="width: 600px"
      :bordered="false"
      size="huge"
      @keydown.enter.prevent="submitTask"
    >
      <n-form ref="formRef" :label-width="80" :model="task" :rules="rules">
        <n-form-item label="Title" path="user.name">
          <n-input v-model:value="task.title" placeholder="Input title" />
        </n-form-item>
        <n-form-item label="Note" path="user.age">
          <n-input v-model:value="task.note" placeholder="Input note" />
        </n-form-item>
        <n-grid :x-gap="12" :y-gap="8" :cols="4">
          <n-grid-item>
            <date-pick-button :timestamp="timestamp" />
          </n-grid-item>
          <n-grid-item> </n-grid-item>
          <n-grid-item> </n-grid-item>
          <n-grid-item>
            <div style="display: flex; justify-content: flex-end">
              <n-button type="primary" @click="submitTask"> Confirm </n-button>
            </div>
          </n-grid-item>
        </n-grid>
      </n-form>
    </n-card>
  </n-modal>
</template>
