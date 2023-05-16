<template>
    <div>
        <label class="bulma-label bulma-is-small">{{ props.message.name }}</label>
        <template  v-if="props.message.Form.element =='input' && props.message.Form.type == 'text' " >
            <input class="bulma-input bulma-is-small" @input="emit('update:message', {...props.message, Data: {
                data: ($event.target as HTMLInputElement).value,
                type: 'string'
            } } )" :value="props.message.Data!.data" />
        </template>
        <template v-else-if="props.message.Form.element =='input' && props.message.Form.type == 'number' ">       
            <input type="number" class="bulma-input bulma-is-small" @input="emit('update:message', {...props.message, Data: {
                data: parseInt(($event.target as HTMLInputElement).value),
                type: 'number'
            } } )" :value="props.message.Data!.data" />
        </template>
        <template v-else-if="props.message.Form.element =='input' && props.message.Form.type == 'textarea' && props.message.Data!.type == 'object'" >
            <textarea class="bulma-textarea" :value="objectToJson(props.message.Data!.data as object)" @input="updateObject(($event.target as HTMLInputElement).value)"> </textarea>
        </template>
        <template v-else>
            <label >unknown element</label>
            <pre>
                {{ props.message }}
            </pre>
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

const objectToJson = (object : object) => {
    try {
        return JSON.stringify(object)
    } catch (e) {
        console.warn(`cant't stringify object = `, object, 'message =',props.message, e)
        return `{"error": "cant't stringify object, see logs"}`
    }
}
const updateObject = (string : string) => {
    try {
        const obj = JSON.parse(string)
         emit('update:message', {...props.message, Data: {
                data:  obj,
                type: 'object'
            } } )
    } catch (e) {
        console.warn(`cant't parse string = `, string, 'message =',props.message, e)
    }
    
}

</script>