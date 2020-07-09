<template>
  <v-container>
    <v-row>
      <v-col>
        <RoomInformations :id=roomID />
      </v-col>
      <v-col>
        <CardBlack :id=roomID />
      </v-col>
    </v-row>

    <v-row>
      <v-col>
        <v-stepper
          v-model="stepState"
          vertical
        >
          <v-stepper-step :complete="stepState > 1" step="1">
            Select card to play
            <small>After confirm no edit</small>
          </v-stepper-step>

          <v-stepper-content step="1">
            <UserHand :cards=user.cards @played="playedCard" />
          </v-stepper-content>

          <v-stepper-step :complete="stepState > 2" step="2">
            Vote best card
            <small>Judge only</small>
          </v-stepper-step>

          <v-stepper-content step="2">
            <v-container v-if="!status.allPlayersAsPlayed">
              <v-row justify="center">
                <v-progress-circular
                  indeterminate
                  color="primary"
                ></v-progress-circular>
              </v-row>
              <v-row justify="center">
                Wait for others players to play ...
              </v-row>
            </v-container>

            <v-container v-if="status.allPlayersAsPlayed">
              <UserHand :cards=cardsPlayed @played="voteCard" />
            </v-container>
          </v-stepper-content>

          <v-stepper-step :complete="stepState > 3" step="3">Result</v-stepper-step>

          <v-stepper-content step="3">
            <v-container v-if="!status.voteIsDone">
              <v-row justify="center">
                <v-progress-circular
                  indeterminate
                  color="primary"
                ></v-progress-circular>
              </v-row>
              <v-row justify="center">
                Wait for vote ...
              </v-row>
            </v-container>
            <v-container v-if="status.voteIsDone">
              <v-row justify="center">
                <v-card
                  class="mx-auto"
                  color="#26c6da"
                  dark
                  min-width="200"
                >
                  <v-card-title>
                    <v-icon
                      large
                      left
                    >
                      mdi-forum
                    </v-icon>
                    <span class="title font-weight-light">Response</span>
                  </v-card-title>

                  <v-card-text class="headline font-weight-bold">
                    {{ room.current_response.card_name }}
                  </v-card-text>
                  <v-card-actions>
                    <v-list-item class="grow">

                      <v-list-item-content>
                        <v-list-item-title>by {{ room.current_response.user_name }}</v-list-item-title>
                      </v-list-item-content>
                    </v-list-item>
                  </v-card-actions>
                </v-card>
              </v-row>
              <v-btn v-if="judge && !haveWinner" color="primary" @click="nextTurn">Launch next run</v-btn>
              <v-btn v-if="judge && haveWinner" color="primary" @click="finishGame">Get summary</v-btn>
            </v-container>
            <!-- <v-btn color="primary" @click="nextStep('result')">Continue</v-btn>
            <v-btn text>Cancel</v-btn> -->
          </v-stepper-content>
        </v-stepper>

      </v-col>
    </v-row>
  </v-container>
</template>

<script>
  import RoomInformations from '@/components/RoomInformations.vue';
  import CardBlack from '@/components/CardBlack.vue';
  import UserHand from '@/components/UserHand.vue';

  import { getCardsByUserID, getJudgeHand, postUsersCards, postJudgeVote, postNextTurn, postEndRoom, retrieveRoomStatus, getRoomStats } from '@/api/main';

  export default {
    components: {
      RoomInformations,
      CardBlack,
      UserHand,
    },
    data() {
      return {
        stepState: 1,
        turn: 0,
        cardsPlayed: [],
        roomID: null,
        status: {
          voteIsDone: false,
          allPlayersAsPlayed: false,
        },
        user: {
          cards: []
        },
        room: {
          current_response: "",
          user_name: "",
        },
        intervalID: {
          room: null,
        },
        haveWinner: false,
      }
    },
    computed: {
      userID () {
        return this.$store.getters.userID
      },
      judge () {
        return this.$store.getters.isJudge
      }
    },
    created() {
        this.roomID = this.$route.params.id

        this.getUserCards()

        retrieveRoomStatus(this.roomID).then(data => {
          this.$log.debug(`Retrieve room status done`);
          this.$store.commit("setRoom", { turn: data.turn, name: data.name })
          this.room.turn = data.turn
          return this.computeRoomData(data)
        }).then(() => {
          return getRoomStats(this.roomID)
        }).then(data => {
          data.classement.forEach(player => {
            if (player.total >= 5) {
              this.haveWinner = true
              return
            }
          })
        })

        this.unsubscribe = this.$store.subscribe(mutation => {
          if (mutation.type === 'setJudge' && this.judge) {
            this.stepState = 2
          }

          if (mutation.type === 'setEndGame') {
            this.$router.push({ name: 'RoomEnd', params: { id: this.roomID }})
          }

          if (mutation.type === 'nextTurn') {
            this.room.turn = this.$store.getters.getRoom.turn
            this.cardsPlayed = []
            this.status.allPlayersAsPlayed = false

            if (this.judge) {
              this.stepState = 2
            } else {
              this.stepState = 1
            }

            this.getUserCards()
          }
        });
    },
    mounted() {
      clearInterval(this.intervalID.room)
      this.poolingData()
    },
    beforeDestroy () {
      clearInterval(this.intervalID.room)
    },
    methods: {
      getCardsByUserID,
      poolingData() {
        this.intervalID.room = setInterval(() => {
          this.getRoom().then(() => {
            if(this.$store.getters.getRoom.turn + 1 === this.room.turn) {
              this.$log.debug(`Next turn`);
              this.$store.commit("nextTurn")
            }
          })},
          10000
        )
      },
      playedCard(sendCard) {
        this.$log.info(sendCard.id, sendCard.name)
        postUsersCards(this.roomID, this.$store.getters.userID, sendCard.id).then(() => {
          this.nextStep("playedCard")
        })
      },
      voteCard(sendCard) {
        this.$log.info(`Voting ${sendCard.id} ${sendCard.name}`)
        postJudgeVote(this.roomID, this.$store.getters.userID, sendCard.id).then(() => {
          this.$log.debug("Done voting")
          this.status.voteIsDone = true
          return this.getRoom()
        }).then(() => {
          return getRoomStats(this.roomID)
        }).then(data => {
          data.classement.forEach(player => {
            if (player.total >= 5) {
              this.haveWinner = true
              return
            }
          })
        })
      },
      nextStep(fromStepName) {
        this.$log.info(this.stepState, fromStepName)
        if (fromStepName == "playedCard" && !this.judge) {
          this.stepState = 3
        } else {
          this.stepState++
        }
      },
      getRoom() {
        return retrieveRoomStatus(this.roomID).then(data => {
          return this.computeRoomData(data)
        })
      },
      finishGame() {
        postEndRoom(this.roomID).then(() => {
          this.$store.commit("setEndGame")
        })
      },
      computeRoomData(room) {
        this.$log.debug(`Start compute room data`);
        this.room = room
        this.$store.commit("setRoom", { raw: room })


        if(room.status === "finished") {
          this.$store.commit("setEndGame")
        }

        if(room.user_judge.id === this.userID && !this.judge) {
          this.$log.debug("You are judge")
          this.$store.commit("setJudge", true)
        }

        if(room.user_judge.id !== this.userID && this.judge) {
          this.$log.debug("You are not judge")
          this.$store.commit("setJudge", false)
        }


        if (room.current_turn_cards.count === room.users_count - 1 && this.judge && room.current_votes.vote === 0) {
          this.$log.debug("Ready to judge")
          this.status.allPlayersAsPlayed = true
          clearInterval(this.intervalID.room)
          this.getJudgeCards()
        }

        if (room.current_votes.vote === 1) {
          this.$log.debug("Vote is done go on 3")
          this.stepState = 3
          this.status.voteIsDone = true
        } else {
          this.$log.debug("Vote is not done")
          this.status.voteIsDone = false
        }

        this.$log.debug(`End compute room data`);
      },
      async getUserCards() {
        await getCardsByUserID(this.roomID, this.userID).then(data => {
          this.user.cards = data
          this.$log.debug(`User as ${this.user.cards.length} cards`)
          if (this.user.cards.length === 4) {
            this.nextStep("playedCard")
          }
        })
      },
      async getJudgeCards() {
        await getJudgeHand(this.roomID).then(resp => {
          this.cardsPlayed = resp.data
        })
      },
      nextTurn() {
        postNextTurn(this.roomID).then(() => {
          return this.getRoom()
        }).then(() => {
          this.poolingData()
          if(this.$store.getters.getRoom.turn + 1 === this.room.turn) {
            this.$log.debug(`Next turn`);
            this.$store.commit("nextTurn")
          }
        })
      }
    },
  }
</script>
