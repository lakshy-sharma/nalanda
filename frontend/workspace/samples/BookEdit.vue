<template>
    <div class="container">
        <h1 class="mt-3">
            Book Edit
        </h1>
    </div>
</template>

<script>
import Security from './security';
export default {
    beforeMount() {
        Security.requireToken();
    }
}
</script>