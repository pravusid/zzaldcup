<template>
  <v-container>
    <v-layout align-center justify-center>
      <v-flex xs12 v-if="loading" class="text-xs-center">
        <v-progress-circular :size="50" color="primary" indeterminate></v-progress-circular>
      </v-flex>
      <v-flex xs12 v-else>
        <v-card>
          <v-container fluid grid-list-md>
            <v-layout row wrap="">
              <v-flex xs6 lg3 v-for="item in items" :key="item.id">
                <v-card>
                  <v-img height="200px">
                    <v-container fill-height fluid pa-2>
                      <v-layout fill-height>
                        <v-flex xs12 align-end flexbox>
                          <span class="headline" v-text="item.matchName"></span>
                        </v-flex>
                      </v-layout>
                    </v-container>
                  </v-img>
                  <v-card-actions>
                    <v-spacer></v-spacer>
                    <v-btn icon>
                      <v-icon>favorite</v-icon>
                    </v-btn>
                    <v-btn icon>
                      <v-icon>bookmark</v-icon>
                    </v-btn>
                    <v-btn icon>
                      <v-icon>share</v-icon>
                    </v-btn>
                  </v-card-actions>
                </v-card>
              </v-flex>
            </v-layout>
          </v-container>
        </v-card>
      </v-flex>
    </v-layout>
    <infinite-loading @infinite="infiniteHandler" spinner="spiral">
      <div slot="no-more">데이터가 없습니다</div>
    </infinite-loading>
  </v-container>
</template>

<script>
import axios from 'axios';
import InfiniteLoading from 'vue-infinite-loading';

export default {
  components: {
    InfiniteLoading,
  },

  created() {
    axios.get('/api/match', {
      params: {
        limit: 4,
        offset: 0,
      },
    }).then((res) => {
      this.items = res.data;
      this.loading = false;
    });
  },

  data: () => ({
    loading: true,
    limit: 4,
    offset: 4,
    items: [],
  }),

  methods: {
    infiniteHandler($state) {
      axios.get('/api/match', {
        params: {
          limit: this.limit,
          offset: this.offset,
        },
      }).then(({ data }) => {
        if (data.length !== 0) {
          this.offset += 4;
          this.items.push(...data);
          $state.loaded();
        } else {
          $state.complete();
        }
      });
    },
  },
};
</script>

<style scoped>
</style>
