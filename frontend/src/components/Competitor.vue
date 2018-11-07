<template>
  <v-flex>
    <img-inputer v-model="file" size="large" icon="img" theme="light"
      placeholder="image"></img-inputer>
  </v-flex>
</template>

<script>
export default {
  data: () => ({
    file: '',
  }),
};
</script>
