<script setup lang="ts">
import { computed, ref } from "vue";
import { CalendarSharp as CalendarIcon } from "@vicons/ionicons5";
import { formattedDate } from "@/api/date";

const timestamp = ref(0);
const dueButtonText = computed(() => {
  if (timestamp.value === 0 || timestamp.value === null) {
    return "Due Date";
  } else {
    return formattedDate(timestamp.value);
  }
});

defineExpose({
  timestamp,
});
</script>

<template>
  <n-popover placement="bottom" trigger="click">
    <template #trigger>
      <n-button class="mr-2">
        <template #icon>
          <n-icon>
            <calendar-icon />
          </n-icon>
        </template>
        {{ dueButtonText }}
      </n-button>
    </template>
    <n-date-picker panel type="date" clearable v-model:value="timestamp" />
  </n-popover>
</template>
