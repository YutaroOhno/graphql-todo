<template>
  <div>
    <ul>
      <li v-for="item in todos" :key="item.id">{{ item.title }}{{item.body}}</li>
    </ul>
    <div>
      <input type="text" name="title" v-model="input.title" />
      <input type="text" name="body" v-model="input.body" />
      <button type="button" @click="createTodo">add todo</button>
    </div>
  </div>
</template>

<script>
import getTodosGql from "~/apollo/gql/getTodos.gql";
import createTodoGql from "~/apollo/gql/createTodo.gql";

export default {
  data() {
    return {
      todos: [],
      input: {
        title: "",
        body: ""
      }
    };
  },

  apollo: {
    todos: {
      query: getTodosGql
    }
  },

  methods: {
    async createTodo() {
      const { title, body } = this.input;
      const { data, error } = await this.$apollo.mutate({
        mutation: createTodoGql,
        variables: {
          title,
          body
        }
      });

      if (error) {
        console.log(error);
      } else {
        // Todo
        // refetch
      }
    }
  }
};
</script>
