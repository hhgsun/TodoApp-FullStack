<template>
  <div class="task-input">
    <form @submit.prevent="handlerInput">
      <input v-model="text" />
      <small>press enter</small>
    </form>
  </div>
</template>

<script>
export default {
  name: "TaskInput",
  data: () => {
    return {
      text: "",
    };
  },
  methods: {
    handlerInput() {
      if (this.text.length > 0) {
        const taskData = {
          text: this.text,
          done: 0,
        };
        this.$store.dispatch("addTask", taskData);
        this.text = "";
      }
    },
  },
};
</script>

<style>
.task-input form {
  margin-bottom: 10px;
}
.task-input input {
  box-sizing: border-box;
  padding: 10px;
  font-size: 1.2em;
  width: 100%;
}
.task-input small {
  font-style: italic;
  opacity: 0.5;
}
</style>