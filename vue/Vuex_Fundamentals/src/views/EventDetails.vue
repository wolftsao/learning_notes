<template>
  <div v-if="event.currentEvent">
    <h1>{{ event.currentEvent.title }}</h1>
    <p>{{ event.currentEvent.time }} on {{ event.currentEvent.date }} @ {{ event.currentEvent.location }}</p>
    <p>{{ event.currentEvent.description }}</p>
  </div>
</template>

<script>
import { mapActions, mapState } from 'vuex'
export default {
  props: ['id'],
  created() {
    this.fetchEvent(this.id).catch(error => {
      this.$router.push({
        name: 'ErrorDisplay',
        params: { error: error },
      })
    })
  },
  computed: {
    ...mapState(['event'])
  },
  methods: {
    ...mapActions('event', ['fetchEvent'])
  }
}
</script>