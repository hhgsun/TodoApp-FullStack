<template>
  <div class="todo-wrapper">

    <h2>To Do List App</h2>

    <task-input />

    <div class="task-list">
      <task-item
        v-for="(task, index) in $store.state.tasks"
        :key="index"
        :task="task"
      />
    </div>

  </div>
</template>

<script>
import TaskInput from "./components/TaskInput.vue";
import TaskItem from "./components/TaskItem.vue";

export default {
  name: "App",
  components: {
    TaskInput,
    TaskItem,
  },
  created(){
    this.$store.dispatch('getTasks');
  }
};
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  color: #2c3e50;
}
.todo-wrapper {
  max-width: 500px;
  padding: 20px 20px 50px 20px;
  margin: 20px auto;
  background-color: rgba(44, 62, 80, 0.1);
}
</style>
