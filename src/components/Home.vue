<template>
	<div>
		<h1>Your posts here</h1>
		<br>
	<div class="container">
		<div v-if="posts.length > 0">
			<md-card v-for="post in posts" class="content">
		        <md-card-header>
		        <md-avatar class="avatar">
		          <img src="https://pbs.twimg.com/profile_images/719228251168731137/61EfguCm.jpg" alt="Avatar">
		        </md-avatar>

				<div class="info">
			        <div class="md-title">{{post.creator}}</div>
			        <div class="md-subhead">{{post.isAdmin ? 'Admin': 'User'}}</div>
			    </div>
		      </md-card-header>

		      <md-card-media>
		        <img :src="post.image" alt="People">
		      </md-card-media>

		      <md-card-content>
		      	<h3>{{post.title}}</h3>
		      	<br>
		        {{post.body}}
		      </md-card-content>

		      <md-card-actions>
		        <md-button>Save</md-button>
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
				posts: [],
			}
		},
		computed: {
			downloadPosts: function() {
			}
		},
		methods: {
			refreshPosts: function() {
				this.$http.get('https://jsonplaceholder.typicode.com/posts').then(response => {
				console.log(response.body);
				console.log(typeof(response.body));
				console.log(response.body.length);
				response.body.map(el => {
					el['image'] = 'https://placeimg.com/640/300'
					el['creator'] = 'LikiPiki'
					el['isAdmin'] = true
					this.posts.push(el)
				});
			}, response => {
				console.log('Error loading json');
			})
			}
		},
		beforeCreate: function() {
			this.$http.get('https://jsonplaceholder.typicode.com/postskek').then(response => {
				console.log(response.body);
				console.log(typeof(response.body));
				console.log(response.body.length);
				response.body.map(el => {
					el['image'] = 'https://placeimg.com/640/300'
					el['creator'] = 'LikiPiki'
					el['isAdmin'] = true
					this.posts.push(el)
				});
			}, response => {
				console.log('Error loading json');
			})
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