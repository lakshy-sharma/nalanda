<template>
    <div class="grid grid-cols-2 text-right">
        <label :for="name" class="mx-1 my-1">{{ label }}</label>
        <input
            :type="type"
            :name="name"
            :placeholder="placeholder"
            :required="required"
            :min="min"
            :max="max"
            :value="modelValue"
            :autocomplete="name + '-new'"
            @input="$emit('update:modelValue', $event.target.value)"
            class="mx-1 my-1 border-2 text-gray-700 max-w-min"/>
    </div>
</template>

<script>
export default {
    name: "TextInput",
    props: {
        name: String,
        type: String,
        label: String,
        placeholder: String,
        required: String,
        min: String,
        max: String,
        modelValue: String
    },
}
</script>