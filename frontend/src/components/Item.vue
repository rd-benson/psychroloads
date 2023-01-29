<script lang="ts">
import router from '@/router/router'

export default {
    props: {
        name: {
            type: String,
            default: "/"
        },
    },
    data() {
        return {
            hover: false
        }
    }
}
</script>

<template>
    <router-link :to="name" @mouseover="hover = true" @mouseleave="hover = false">
        {{ (hover && $route.name != name) ? name + " >" : name}}
    </router-link>
</template>