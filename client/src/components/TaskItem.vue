<template>
  <div class="task-item" :class="task.Done ? 'done' : ''">
    <span>{{ task.Text }}</span>
    <div class="task-actions">
      <button @click="doneItem(task)">&#10004;</button>
      <button @click="removeItem(task)">&#10005;</button>
    </div>
  </div>
</template>

<script>
export default {
  name: "TaskItem",
  props: {
    task: Object,
  },
  methods: {
    doneItem: function (task) {
      this.$store.dispatch("doneTask", task);
    },
    removeItem: function (task) {
      this.$store.dispatch("removeTask", task);
    },
  },
};
</script>

<style>
.task-item {
  position: relative;
  border: 1px solid;
  margin: 10px 0;
  padding: 15px;
  font-size: 1.2em;
  line-height: 1;
}
.task-item.done {
  opacity: 0.5;
}
.task-item.done span {
  font-style: oblique;
  text-decoration: line-through;
}
.task-item.done .task-actions button:first-child {
  display: none;
}
.task-item .task-actions {
  position: absolute;
  top: 12px;
  right: 10px;
  display: none;
}
.task-item:hover > .task-actions {
  display: block;
}
.task-item .task-actions button {
  cursor: pointer;
  margin-left: 5px;
}
</style>