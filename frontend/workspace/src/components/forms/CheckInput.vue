<template>
    <div class="mb-3">
        <input :checked="checked" :required="required" type="check-input"/>
        <label :for="name">{{ label }}</label>
    </div>
</template>

<script>
export default {
    name: "CheckInput",
    props: ["label", "required", "name", "value", "checked"]
}
</script>