<template>
	<div>
		<h1>Profile {{getUser.username}}</h1>
		<md-avatar class="md-large">
        <img src="https://pbs.twimg.com/profile_images/719228251168731137/61EfguCm.jpg" alt="People">
      </md-avatar>
        <br>
      <div>
      	Change profile icon 
      </div>
      <md-button id="edit-icon" class="md-icon-button md-raised md-accent"> <md-icon>edit</md-icon>
      	</md-button>
	</div>
</template>

<script>
	export default {
		name: 'Profile',
		data() {
			return {

			}
		},
		computed: {
			getUser() {
				console.log(this.$store.state.user);
				return this.$store.state.user
			}
		}
	}
</script>

<style>
</style>