<template>
  <v-row justify="center" align="center">
    <v-col cols="12">
      <div class="ma-1 pa-3 grey lighten-4">
        <span class="text-h4"> プロジェクトを作成する </span>
        <div class="mt-1">
          <ProjectCreateForm />
        </div>
      </div>
    </v-col>
    <v-col cols="12">
      <div class="ma-1 pa-3 grey lighten-4">
        <span class="text-h4"> プロジェクトリスト </span>
        <div class="mt-2">
          <ProjectList />
        </div>
      </div>
    </v-col>
  </v-row>
</template>

<script>
export default {
  name: 'ProjectIndexPage',
}
</script>
