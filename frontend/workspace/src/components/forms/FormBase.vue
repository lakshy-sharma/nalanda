<template>
    <form
        v-on:submit.prevent="submit"
        :ref="name"
        :event="event"
        autocomplete="off"
        :method="method"
        :action="action"
        class="flex-grid">
        <slot></slot>
    </form>
</template>

<script>
export default {
    name: "FormBase",
    props: ["method", "action", "name", "event"],
    methods: {
        submit() {
            console.log("Submit clicked.")
            this.$emit(this.$props["event"]);
        }
    }
}
</script>