<template>
    <div>
        <template v-if="props.message.Form.element =='input' && props.message.Form.type == 'text' ">
            <label class="bulma-label bulma-is-small">{{ props.message.Form.name }}</label>
            <input class="bulma-input bulma-is-small" @input="emit('update:message', {...props.message, Data: {
                data: ($event.target as HTMLInputElement).value,
                type: 'string'
            } } )" :value="props.message.Data.data" />
        </template>
    </div>
</template>
<script lang="ts" setup>
import { Message } from "@pathserve/messenger";
const props = defineProps<{
    message : Message
}>()
const emit = defineEmits<{
    (e: 'update:message', message : Message) : void
}>()
</script>