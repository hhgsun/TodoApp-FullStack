<template>
  <div class="task-input">
    <form @submit.prevent="handlerInput">
      <input
        v-model="text"
        placeholder="......"
        :disabled="isWait"
        ref="text"
      />
      <small>press enter</small>
    </form>
  </div>
</template>

<script>
export default {
  name: "TaskInput",
  data: () => {
    return {
      isWait: false,
      text: "",
    };
  },
  methods: {
    handlerInput: function () {
      if (this.text.length > 0) {
        this.isWait = true;
        const taskData = {
          text: this.text,
          done: 0,
        };
        this.$store.dispatch("addTask", taskData).then(() => {
          this.text = "";
          this.isWait = false;
          this.$nextTick(() => {
            this.focusInput();
          });
        });
      }
    },
    focusInput: function () {
      this.$refs.text.focus();
    },
  },
};
</script>

<style>
.task-input form {
  margin-bottom: 10px;
  text-align: right;
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