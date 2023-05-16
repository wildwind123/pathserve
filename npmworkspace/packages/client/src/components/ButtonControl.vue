<template>
    <div>
        <template v-if="props.message.Form && props.message.Form.element == 'button'">
            <button class="bulma-button bulma-is-small" @click="emit('clicked')">{{ props.message.name }}</button>
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