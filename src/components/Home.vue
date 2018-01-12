<template>
	<div>
		<h1>Your posts here</h1>
		<br>
	<div class="container">
		<div v-if="allPosts.length > 0">
			<md-card v-for="(post, index) in allPosts" :key="index" class="content">
		        <md-card-header>
		        <md-avatar class="avatar">
		          <img src="https://pbs.twimg.com/profile_images/719228251168731137/61EfguCm.jpg" alt="Avatar">
		        </md-avatar>

				<div class="info">
			        <div class="md-title">{{post.user.username}}</div>
			        <div class="md-subhead">{{post.user.isAdmin ? 'Admin': 'User'}}</div>
			    </div>
		      </md-card-header>

		      <md-card-media>
		        <img :src="post.imageLink" alt="People">
		      </md-card-media>

		      <md-card-content>
		      	<h3>{{post.title}}</h3>
		      	<br>
		        {{post.content}}
		      </md-card-content>

		      <md-card-actions>
		        <md-button>Read</md-button>
		      </md-card-actions>
		    </md-card>
			</div>
 			<div v-else>
				<h1>Failed load posts</h1>
 				<md-button class="md-fab" @click="refreshPosts">
        		<md-icon>cached</md-icon>
      			</md-button>
			</div>
 		</div>
	</div>
</template>


<script>
	export default {
		name: 'Home',
		data() {
			return {
			}
		},
		computed: {
			allPosts() {
				return this.$store.state.posts
			}
		},
		methods: {
			refreshPosts: function() {
				this.$store.dispatch('getPosts')
			}
		},
	}
</script>

<style>
	.container {
		display: block;
		width: 600px;
		margin: 0 auto;
	}	
	.content {
		margin-bottom: 20px;
	}
</style>