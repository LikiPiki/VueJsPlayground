<template>
	<div>
		<h1>Profile {{getUser.username}}</h1>
		<md-avatar class="md-large">
		<img src="https://pbs.twimg.com/profile_images/719228251168731137/61EfguCm.jpg" alt="People">
	  </md-avatar>
		<br>
		<br>Edit profile icon
	  <md-button id="edit-icon" class="md-icon-button md-raised md-accent"> <md-icon>edit</md-icon>
		</md-button>
	  <br>
	  <form class="login-form md-layout-row md-gutter">
		<md-card class="md-flex-50 md-flex-small-100">
			<md-card-header>
			  <div class="md-title">Create new post</div>
			</md-card-header>
			<md-card-content>
				<md-field>
					<label for="title">Title</label>
					<md-input name="title" v-model="form.title" required/>
				</md-field>
				<md-field>
					<label for="content">Content</label>
					<md-input type='textarea' name="content" v-model="form.content" required/>
				</md-field>
				<md-field>
					<label for="imagelink">Image link</label>
					<md-input name="imagelink" v-model="form.imagelink" required/>
				</md-field>
				<md-button class="md-primary" @click="savePost">Save Post</md-button>
			</md-card-content>
		</md-card>
	  </form>
	</div>
</template>

<script>

	export default {
		name: 'Profile',
		data() {
			return {
				form: {
					title: '',
					content: '',
					imagelink: '',
				}
			}
		},
		computed: {
			getUser() {
				return this.$store.state.user
			},
		},
		methods: {
			savePost: function() {
				this.$store.dispatch('savePost', {
					creator: this.$store.state.user.username,
					image: this.form.imagelink,
					body: this.form.content,
					title: this.form.title
				})
				if (this.$store.state.user) {
					let data = {
						...this.$store.state.user,
						...this.form,
					}
					console.log('data is ', data)
					this.$http.post('/add_post', data).then(response => {
						console.log("SUCCESS");
						console.log(response.body);
					}, response => {
						console.log('Error');
						console.log(response.body);
					})
				} else {
					console.log("NOT");
				}
			}
		},
	}
</script>

<style>
	.login-form {
		display: block;
		width: 400px;
		margin: 0 auto;
		padding: 20px;
	}
</style>