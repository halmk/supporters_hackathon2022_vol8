<template>
  <v-row justify="center" align="center">
    <v-col cols="12">
      <div class="ma-1 pa-3 grey lighten-4">
        <span class="text-h4"> プロジェクトを作成する </span>
        <div class="grey lighten-4">
          <ProjectCreateForm />
        </div>
      </div>
    </v-col>
  </v-row>
</template>

<script>
export default {
  name: 'IndexPage',
}
</script>
