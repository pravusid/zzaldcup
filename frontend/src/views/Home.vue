<script>
import MatchList from '../components/MatchList.vue';

export default {
  mixins: [
    MatchList,
  ],

  data: () => ({
    host: '/api/match',
  }),
};
</script>
