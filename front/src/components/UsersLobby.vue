<template>
  <v-simple-table>
    <template v-slot:default>
      <thead>
        <tr>
          <th class="text-left">UserID</th>
          <th class="text-left">Name</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="item in users" :key="item.id">
          <td>{{ item.id }}</td>
          <td>
            <v-badge v-if="item.id===userID"
              content="you"
              color="green"
            >
              {{ item.name }}
            </v-badge>
            <div v-if="item.id!==userID">{{ item.name }}</div>
          </td>
        </tr>
      </tbody>
    </template>
  </v-simple-table>
</template>

<script>
  export default {
    name: 'UsersLobby',
    computed: {
      users () {
        return this.$store.getters.users
      },
      userID () {
        return this.$store.getters.userID
      }
    }
  };
</script>
