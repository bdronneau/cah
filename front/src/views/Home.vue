<template>
  <div class="home">
    <JoinRoom />
  </div>
</template>

<script>
import JoinRoom from '@/components/JoinRoom.vue';

export default {
  name: 'Home',
  components: {
    JoinRoom,
  },
};
</script>
