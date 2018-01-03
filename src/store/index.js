import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

const store = new Vuex.Store({
	state: {
		login: true,
		user: {
			username: "",
			login: "",
			image: "",
			age: 0,
		}
	},
	mutations: {
		loginUser(state) {
			state.login = true
		},
		logoutUser(state) {
			state.login = false
		}
	},
	actions: {
	}
})

export default store