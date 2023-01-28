<script setup lang="ts">
import { computed, ref } from "vue";
import { CalendarSharp as CalendarIcon } from "@vicons/ionicons5";
import { formattedDate } from "@/api/date";

const props = defineProps({
  timestamp: { type: Number, default: 0 },
});
const emit = defineEmits<{
  (e: "update:timestamp", timestamp: number): void;
}>();
const timestamp = ref(props.timestamp);

const dueButtonText = computed(() => {
  if (timestamp.value === 0 || timestamp.value === null) {
    return "Due Date";
  } else {
    return formattedDate(timestamp.value);
  }
});

function handleTimestampUpdate(ts: number) {
  timestamp.value = ts;
  emit("update:timestamp", timestamp.value);
}
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
    <n-date-picker
      panel
      type="date"
      clearable
      :value="timestamp"
      :on-update:value="handleTimestampUpdate"
    />
  </n-popover>
</template>
