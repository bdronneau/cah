<template>
  <v-container>
    <v-row>
      <v-col>
        <RoomInformations :id=roomID />
      </v-col>
      <v-col>
        <v-simple-table>
          <template v-slot:default>
            <thead>
              <tr>
                <th class="text-left">Total</th>
                <th class="text-left">User</th>
              </tr>
            </thead>
            <tbody>
              <tr v-for="item in stats.classement" :key="item.user_name">
                <td>{{ item.total }}</td>
                <td>{{ item.user_name }}</td>
              </tr>
            </tbody>
          </template>
        </v-simple-table>
      </v-col>
    </v-row>

    <v-row>
      <v-col
        xs12
        sm12
        md12
      >
        <v-simple-table>
          <template v-slot:default>
            <thead>
              <tr>
                <th class="text-left">Black Card</th>
                <th class="text-left">White Card</th>
                <th class="text-left">User</th>
                <th class="text-left">Turn</th>
              </tr>
            </thead>
            <tbody>
              <tr v-for="item in stats.history" :key="item.turn">
                <td>{{ item.black_card_name }}</td>
                <td>{{ item.white_card_name }}</td>
                <td>{{ item.user_name }}</td>
                <td>{{ item.turn }}</td>
              </tr>
            </tbody>
          </template>
        </v-simple-table>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
  import RoomInformations from '@/components/RoomInformations.vue';

  import { getRoomStats } from '@/api/main';

  export default {
    components: {
      RoomInformations,
    },
    data() {
      return {
        stats: {
          history: [],
          classement: [],
        }
      }
    },
    computed: {
      userID () {
        return this.$store.getters.userID
      },
      roomID () {
        return this.$store.getters.getRoom.id
      }
    },
    created() {
      this.$store.commit('setRoom', { id: this.$route.params.id })
    },
    methods: {
    },
    mounted() {
      this.$store.dispatch("getRoomData")
      getRoomStats(this.roomID).then(data => {
        this.stats = data
      })
    },
  }
</script>
