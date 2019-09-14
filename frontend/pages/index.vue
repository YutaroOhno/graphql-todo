<template>
  <ul>
    <li v-for="item in todos" :key="item.id">{{ item.title }}</li>
  </ul>
</template>

<script>
import getTodosGql from "~/apollo/gql/getTodos.gql";

export default {
  data() {
    return {
      todos: []
    };
  },

  apollo: {
    todos: {
      query: getTodosGql
    }
  }
};
</script>
