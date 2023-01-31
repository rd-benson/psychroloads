<script lang="ts">

export default {
    props: ['prompt', 'text'],
    data() {
        return {
            active: false,
            height: 100,
        }
    },
    methods: {
        toggleCollapse() {
            this.active = !this.active
            this.height = (this.height == 100) ? 200 : 100
        }
    },
}

</script>

<template>
    <button @click="toggleCollapse" type="button" class="collapsible">{{ prompt + " " + active + " " + height }}</button>
    <div v-show="active" class="content" :styles="{minHeight: height}">
        <p>{{ text }}</p>
    </div>
</template>

<style>
.collapsible {
  background-color: #777;
  color: white;
  cursor: pointer;
  padding: 18px;
  width: 100%;
  border: none;
  text-align: left;
  outline: none;
  font-size: 15px;
}

.active, .collapsible:hover {
  background-color: #555;
}

.content {
  padding: 0 18px;
  overflow: hidden;
  min-height: 0;
  transition: max-height 0.2s ease-out;
  background-color: #f1f1f1;
}
</style>