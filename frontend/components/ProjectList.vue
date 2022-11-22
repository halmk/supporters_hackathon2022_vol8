<template>
  <v-card tile>
    <v-list-item
      v-for="(project, i) in projects"
      :key="i"
      :to="project.name"
      append
      router
      exact
    >
      <v-list-item-content>
        <v-list-item-title>{{ project.name }}</v-list-item-title>
      </v-list-item-content>
    </v-list-item>
  </v-card>
</template>
<script>
export default {
  name: 'ProjectList',
  props: {
    // projects: Array,
  },
  data: () => ({
    projects: [
      {
        name: 'project1',
      },
      {
        name: 'project2',
      },
    ],
  }),
  computed: {},
  methods: {},
}
</script>
