<script setup lang="ts">
import { ref } from "vue";
import type { Task } from "@/api/types";
import { reqGetTask, reqUpdateTask } from "@/api/requests";

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
const emit = defineEmits<{
  (e: "onTaskSubmit", id: number): void;
}>();

function open(id: number) {
  reqGetTask(id).then((response) => {
    currTask.value = response.data;
  });
  showEditModal.value = true;
}

function submitTask() {
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
        <n-row :gutter="[0, 24]">
          <n-col :span="24">
            <div style="display: flex; justify-content: flex-end">
              <n-button type="primary" @click="submitTask"> Confirm </n-button>
            </div>
          </n-col>
        </n-row>
      </n-form>
    </n-card>
  </n-modal>
</template>
