<template>
	<div>
		Menu is here
		<br>
		 <md-button class="md-accent" to="/login" v-if="!ifLogin">Login</md-button>
		 <md-button to="/register" v-if="!ifLogin">Register</md-button>
		 <md-button to="/">Home</md-button>
		 <md-button to="/profile" class="md-primary" v-if="ifLogin">Profile</md-button>
		 <md-button class="md-primary" v-if="ifLogin" @click="logout">Logout</md-button>
	</div>
</template>

<script>
	export default {
		name: 'Menu',
		data() {
			return {
				name: 'vue',
			}
		},
		computed: {
			ifLogin() {
				return this.$store.state.login
			}
		},
		methods: {
			logout: function () {
				console.log("logout");
				this.$store.commit('logoutUser')
			}
		}
	}
</script>