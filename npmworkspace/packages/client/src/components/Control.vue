<template>
    <div>
        <template v-for="message in props.messages">
            <InputControl @update:message="emit('update:message', $event)" :message="message"></InputControl>
        </template>
    </div>
</template>
<script lang="ts" setup>
import { Message } from "@pathserve/messenger";
import InputControl from "@/components/InputControl.vue";
const props = defineProps<{
    messages: Message[]
}>()
const emit = defineEmits<{
    (e: 'update:message', message : Message) : void
}>()
</script>