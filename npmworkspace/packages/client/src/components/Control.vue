<template>
  <div class="control">
    <template v-for="message in props.messages">
      <template v-if="message.Form">
        <template v-if="message.Form.element == 'input'">
          <InputControl
            @update:message="emit('update:message', $event)"
            :message="message"
          ></InputControl>
        </template>
        <template v-if="message.Form.element == 'button'">
          <ButtonControl
            @clicked="emit('update:message', message)"
            :message="message"
          ></ButtonControl>
        </template>
      </template>
    </template>
  </div>
</template>
<script lang="ts" setup>
import { Message } from "@pathserve/messenger";
import InputControl from "@/components/InputControl.vue";
import ButtonControl from "@/components/ButtonControl.vue";
const props = defineProps<{
  messages: Message[];
}>();
const emit = defineEmits<{
  (e: "update:message", message: Message): void;
}>();
</script>
<style scoped>
.control div:not(.control div:last-child) {
  margin-bottom: 10px;
}
</style>
