import EventService from '@/services/EventService.js'

export default {
    namespaced: true,
    state: {
        events: [],
        currentEvent: {}
    },
    mutations: {
        ADD_EVENT(state, event) {
            state.events.push(event)
        },
        SET_EVENTS(state, events) {
            state.events = events
        },
        SET_EVENT(state, event) {
            state.currentEvent = event
        }
    },
    actions: {
        createEvent({ commit }, event) {
            EventService.postEvent(event)
                .then(() => {
                    commit('ADD_EVENT', event)
                })
                .catch(error => {
                    throw (error)
                })
        },
        fetchEvents({ commit }) {
            return EventService.getEvents()
                .then(response => {
                    commit('SET_EVENTS', response.data)
                })
                .catch(error => {
                    throw (error)
                })
        },
        fetchEvent({ commit, state }, id) {
            const existingEvent = state.events.find(event => event.id === id)
            if (existingEvent) {
                commit('SET_EVENT', existingEvent)
            } else {
                return EventService.getEvent(id)
                    .then(response => {
                        commit('SET_EVENT', response.data)
                    })
                    .catch(error => {
                        throw (error)
                    })
            }
        }
    }
}