<template>
	<div>
		<h1>Registration new user</h1>
		<form class="login-form">
			  <md-field>
                <label for="username">Username</label>
                <md-input name="username" v-model="form.username"/>
            </md-field>
            <md-field>
                <label for="username">Password</label>
                <md-input type="password" name="password" v-model="form.password"/>
            </md-field>
            <md-button class="md-primary" @click="clicked">Register</md-button>
		</form>
	</div>
</template>

<script>
	export default {
		name: 'Register',
		data() {
			return {
				form: {
					username: '', 
					password: '',
				},
			}
		},
		methods: {
			clicked: function (event) {
				// register new user here
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