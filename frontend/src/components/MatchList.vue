<template>
  <v-container>
    <v-layout align-center justify-center>
      <v-flex xs12>
        <v-card>
          <v-container fluid grid-list-md>
            <v-layout row wrap="">
              <v-flex xs6 lg3 v-for="(item, index) in items" :key="index">
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
                      <v-icon>share</v-icon>
                    </v-btn>
                    <v-btn icon :to="{ path: `/match/edit/${item.matchName}` }">
                      <v-icon>edit</v-icon>
                    </v-btn>
                    <v-btn icon>
                      <v-icon>clear</v-icon>
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
      <div class="my-4" slot="no-more">
        <span class="subheading">마지막입니다</span>
      </div>
      <div slot="no-results"></div>
    </infinite-loading>
  </v-container>
</template>

<script>
import InfiniteLoading from 'vue-infinite-loading';
import axios from '../libs/axios';

export default {
  components: {
    InfiniteLoading,
  },

  data: () => ({
    host: null,
    limit: 8,
    offset: 0,
    items: [],
  }),

  methods: {
    infiniteHandler($state) {
      axios.get(this.host, {
        params: {
          limit: this.limit,
          offset: this.offset,
        },
      }).then(({ data }) => {
        if (data.length !== 0) {
          this.offset += this.limit;
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
