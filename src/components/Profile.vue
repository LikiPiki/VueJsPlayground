<template>
	<div>
		<h1>Profile {{getUser.username}}</h1>
		<md-avatar class="md-large">
	<img v-if="getUser.userImage" :src="getUser.userImage" alt="People">
		<img v-else src="https://pbs.twimg.com/profile_images/719228251168731137/61EfguCm.jpg" alt="People">
	  </md-avatar>
		<br>
		<br>Edit profile icon
	<form class="login-form" enctype="multipart/form-data">
	  <md-field>
		<label>Choose jpg/png file only!</label>
		<md-file
		  @change="imageChanged"
		  v-model="imageForm.imageName"
		/>
	  </md-field>
	  <md-button class="md-primary" @click="changeProfileImage">
		Change/set image
	  </md-button>
	</form>
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
			return {				form: {

					title: '',
					content: '',
					imagelink: '',
				},
		imageForm: {
		    imageData: null,
		    dimageName: '',
		}
			}
		},
		computed: {
			getUser() {
				return this.$store.state.user
			},
		},
		methods: {
			changeProfileImage() {
				if (this.imageForm.imageName && (this.imageForm.imageName.endsWith(".jpg") || (this.imageForm.imageName.endsWith(".png")))) {
					console.log("send", this.imageForm)
					let data = {
						...this.$store.state.user,
						...this.imageForm
					}
					console.log(data)
					this.$http.post('/load_profile_image', data).then(response => {
					}, response => {
						console.log("Recive", response)
					})
				}
			},
			imageChanged(e) {
				let fileReader = new FileReader()
				fileReader.readAsDataURL(e.target.files[0])
				fileReader.onload = (e) => {
					this.imageForm.imageData = e.target.result
				}
			},
			savePost: function() {

				if (this.$store.state.user) {
				  console.log("Yeap")
					let data = {
						...this.$store.state.user,
						...this.form,
					}
					console.log('data is ', data)
					this.$http.post('/add_post', data).then(response => {
						console.log("Success", response.body);
						this.$store.dispatch('savePost', response.data)

					}, response => {
						console.log("Error", response.body);
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
