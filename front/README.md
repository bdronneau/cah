# CAH - VueJS

## Features
- `/rooms` allow to create room
- UserID is stored in localstorage


## How to play

One player go to `/rooms` create a room and give ID to other users
Everybody join through `/`.
One player `start the party`, other players will see `enter in party` 10 seconds.
First player to get 10 selected cards end the game.

## Develop

You need to install node stuff with
```bash
npm install
```

I mainly use default [vue-cli](https://cli.vuejs.org/) presets combine with [`vuetify`](https://vuetifyjs.com/).

### Configuration

At start, the application request `/config` in order to retrieve configuration (see this behavior in [viws](https://github.com/ViBiOh/viws#environment-variables)). If the path does not exist, the default configuration is load in [`vuex`](https://vuex.vuejs.org/).

### Start local node HTTP server
```
npm run serve
```

### Compiles and minifies for production
```
npm run build
```

### Run (light) unit tests
```
npm run test:unit
```

### Lints and fixes files
```
npm run lint
```
