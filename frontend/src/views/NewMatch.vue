<template>
  <v-container>
    <v-layout row wrap align-center justify-center>
      <v-flex xs12 sm12 md8>
        <h1>New Match</h1>
        <v-form>
          <v-text-field prepend-icon="label" name="name" label="대결명"
                        type="text"></v-text-field>
          <v-text-field prepend-icon="more_vert" name="quota"
                        label="최대 개수" type="text"></v-text-field>
        </v-form>
      </v-flex>
    </v-layout>
    <v-layout row wrap align-center justify-center>
      <v-flex xs12 sm12 md8>
        <v-btn color="primary">전송</v-btn>
      </v-flex>
    </v-layout>
  </v-container>
</template>

<script>
export default {
};
</script>

<style scoped>

</style>
