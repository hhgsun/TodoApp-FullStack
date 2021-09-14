import { createStore } from 'vuex'

const API_URL = "http://localhost:8090"
const ENDPOINT_TASK = "/tasks"

export default createStore({
  state() {
    return {
      tasks: []
    }
  },
  mutations: {
    setAll(state, payload) {
      state.tasks = payload;
    },
    add(state, payload) {
      const data = [
        payload,
        ...state.tasks,
      ]
      state.tasks = data;
    },
    remove(state, payload) {
      const index = state.tasks.findIndex(t => t.ID == payload.ID);
      if (index !== -1) {
        state.tasks.splice(index, 1);
      }
    },
    done(state, payload) {
      const index = state.tasks.findIndex(t => t.ID == payload.ID);
      if (index !== -1) {
        state.tasks[index] = {
          ...state.tasks[index],
          Done: true
        };
      }
    }
  },
  actions: {
    async getTasks({ commit }) {
      const res = await fetch(API_URL + ENDPOINT_TASK, {
        method: "GET"
      });
      const data = await res.json();
      commit('setAll', data.reverse());
      return data;
    },
    async addTask({ commit }, payload) {
      const res = await fetch(API_URL + ENDPOINT_TASK, {
        method: "POST",
        body: JSON.stringify(payload)
      });
      const data = await res.json();
      commit('add', data);
      return data;
    },
    async removeTask({ commit }, payload) {
      await fetch(API_URL + ENDPOINT_TASK + "/" + payload.ID, {
        method: "DELETE",
      });
      commit('remove', payload);
      return payload;
    },
    async doneTask({ commit }, payload) {
      const updateData = { ...payload, Done: true }
      console.log(payload, updateData)
      const res = await fetch(API_URL + ENDPOINT_TASK + "/" + payload.ID, {
        method: "PUT",
        body: JSON.stringify(updateData)
      });
      const data = await res.json();
      commit('done', data);
      return data;
    }
  }
})
