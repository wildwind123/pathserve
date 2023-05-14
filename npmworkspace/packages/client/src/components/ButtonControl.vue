<template>
    <div>
        <template v-if="props.message.Form.element == 'button'">
            <button @click="emit('clicked')">{{ props.message.Form.name }}</button>
        </template>
    </div>
</template>
<script lang="ts" setup>
import { Message } from '@pathserve/messenger';
const props = defineProps<{
    message : Message
}>()
const emit = defineEmits<{
    (e: 'clicked') : void
}>()
</script>