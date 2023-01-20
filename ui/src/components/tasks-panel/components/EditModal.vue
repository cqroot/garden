<script setup lang="ts">
import { ref } from "vue";
import type { Task } from "@/api/types";
import { reqGetTask, reqUpdateTask } from "@/api/requests";
import DatePickButton from "./DatePickButton.vue";

const showEditModal = ref(false);
const currTask = ref({} as Task);
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
const datePickButtonRef = ref();

const emit = defineEmits<{
  (e: "onTaskSubmit", id: number): void;
}>();

function open(id: number) {
  reqGetTask(id).then((response) => {
    currTask.value = response.data;
    datePickButtonRef.value.setTimestamp(currTask.value.due);
  });
  showEditModal.value = true;
}

function submitTask() {
  currTask.value.due = datePickButtonRef.value.timestamp;
  reqUpdateTask(currTask.value.id, currTask.value).then(() => {
    emit("onTaskSubmit", currTask.value.id);
  });
  showEditModal.value = false;
}

defineExpose({
  open,
});
</script>

<template>
  <n-modal v-model:show="showEditModal">
    <n-card
      style="width: 600px"
      :bordered="false"
      size="huge"
      @keydown.enter.prevent="submitTask"
    >
      <n-form ref="formRef" :label-width="80" :model="currTask" :rules="rules">
        <n-form-item label="Title" path="user.name">
          <n-input v-model:value="currTask.title" placeholder="Input title" />
        </n-form-item>
        <n-form-item label="Note" path="user.age">
          <n-input v-model:value="currTask.note" placeholder="Input note" />
        </n-form-item>
        <n-grid :x-gap="12" :y-gap="8" :cols="4">
          <n-grid-item>
            <date-pick-button ref="datePickButtonRef" />
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
