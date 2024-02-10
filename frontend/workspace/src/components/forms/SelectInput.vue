<template>
    <div class="mb-3">
        <label :for="name">{{ label }}</label>
        <select :id="name" :name="name" :required="required" >
            <option v-for="option in items" :value="option.value" :key="option.value">
                {{ option.text }}
            </option>
        </select>
    </div>
</template>

<script>
export default {
    name: "SelectInput",
    props: ["items", "name", "required", "label"]
}
</script>