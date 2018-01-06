<template>
	<div>
		<h1>Hello i am login like {{name}}</h1>
		<form class="login-form md-layout-row md-gutter">
			  <md-field>
                <label for="username">Username</label>
                <md-input name="username" v-model="form.username"/>
            </md-field>
            <md-field>
                <label for="username">Password</label>
                <md-input type="password" name="password" v-model="form.password"/>
            </md-field>
            <md-button class="md-primary" @click="clicked">Login</md-button>
		</form>
	</div>
</template>

<script>
	export default {
		name: 'Login',
		data() {
			return {
				name: 'Likipiki',
				form: {
					username: '', 
					password: '',
				},
			}
		},
		methods: {
			clicked: function (event) {
					let data = {
						username: this.form.username,
						password: this.form.password,
					}

					this.$http.post('/login', data).then(response => {
						console.log(response);
						this.$store.commit('loginUser', response)
						console.log('Logining');
					}, response => {
						console.log('Login error')
					})
			}
		}
	}
</script>

<style>
	.login-form {
		display: block;
		width: 400px;
		margin: 0 auto;
	}
</style>