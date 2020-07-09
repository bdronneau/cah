<template>
  <v-row
    align-center
    justify-center
  >
    <v-col
      xs12
      sm8
      md4
    >
      <v-card class="elevation-12">
        <v-toolbar
          color="primary"
          dark
          flat
        >
          <v-toolbar-title>Join room</v-toolbar-title>
          <v-spacer></v-spacer>
        </v-toolbar>
        <v-card-text>
          <v-form
            v-model="valid"
            ref="form"
          >
            <v-text-field
              v-model="name"
              label="Name"
              :rules="nameRules"
              :counter="10"
              prepend-icon="mdi-message-text"
              type="text"
              required
            ></v-text-field>

            <v-text-field
              id="password"
              label="Password"
              name="password"
              prepend-icon="mdi-lock"
              type="password"
              disabled
            ></v-text-field>
          </v-form>
        </v-card-text>
        <v-card-actions>
          <v-spacer></v-spacer>
          <v-btn
            color="primary"
            @click="validate"
            :disabled="!valid"
          >
            Join
          </v-btn>
        </v-card-actions>
      </v-card>
    </v-col>
  </v-row>
</template>

<script>
    import { roomExist } from "@/helpers/room"

  export default {
    data: () => ({
      valid: false,
      name: '',
      nameRules: [
        v => !!v || 'Name is required',
        v => v.length === 10 || 'Name must equal to 10 characters',
      ],
    }),
     methods: {
      async validate () {
        if (this.$refs.form.validate() && await roomExist(this.name)) {

          this.$router.push({ name: 'LobbyRoom', params: { id: this.name }})
        }
      },
    },
  }
</script>
