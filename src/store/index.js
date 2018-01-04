import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

const store = new Vuex.Store({
	state: {
		login: false,
		user: {
			username: "",
			isAdmin: false,
			image: "",
			age: 0,
		},
		posts: []
	},
	mutations: {
		loginUser(state, data) {
			state.login = true
			state.user.username = data.username,
			state.user.isAdmin = data.isAdmin
		},
		logoutUser(state) {
			state.login = false
		},
		set(state, {type, data}) {
			state[type] = data
		},
		save(state, query) {
			this.state.posts.unshift(query)
		}
	},
	actions: {
		getPosts({commit}) {
			const url = 'https://jsonplaceholder.typicode.com/posts'
			Vue.http.get(url).then(response => {
				console.log('FROM VUEX');
				console.log(response.body);
				console.log('FROM VUEX');
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
		}
	},
})

export default store