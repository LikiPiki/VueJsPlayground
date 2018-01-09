import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

const store = new Vuex.Store({
	state: {
		login: true,
		user: {
			username: "",
			isAdmin: false,
		},
		posts: []
	},
	mutations: {
		loginUser(state, data) {
			console.log('logggging user');
			this.state.login = true
			this.state.user = data
			console.log(this.state);
		},
		logoutUser(state) {
			state.login = false
		},
		set(state, {type, data}) {
			state[type] = data
		},
		save(state, query) {
			this.state.posts.unshift(query)
		},
	},
	actions: {
		getPosts({commit}) {
			const url = 'https://jsonplaceholder.typicode.com/posts'
			// change response to new
			Vue.http.get(url).then(response => {
				response.body.map(el => {
					el['image'] = 'https://placeimg.com/640/300'
					el['creator'] = 'LikiPiki'
					el['isAdmin'] = true
			});
				commit('set', {type: 'posts', data: response.body})
			}, response => {
				console.log('Error loading json VUEX');
			})
		},
		savePost({commit}, query) {
			commit('save', query);
		},
	},
})

export default store