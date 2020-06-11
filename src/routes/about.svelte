<script>
import Index from "./index.svelte";
</script>



<svelte:head>
	<title>About</title>
</svelte:head>

<style>
.body {
	background: #384047;
	padding:2%;
	border-radius:12px;
    display: grid;
    grid-template-rows: repeat(auto-fit, minmax(50px, 1fr));
      box-shadow: 7px 2px 14px -1px #888888;
}
.body .item{
	font-weight: bold;
    color: #c0c1c5;
}
</style>


<h1>About Me</h1>

<div class="body">
  <div class="item">Hi, I just wanted a way to 
  communicate the things I love and What I enjoy
  doing/learning.  I think a lot of what people call
  intelligence really boils down to curiousity.</div>
</div>

<br>

<div class="body">
  <div class="item">Poeple's minds are like muscles and 
  people have to be careful when they don't exercise them.
  If you embrace the growth mindset and understand that learning
   is a process that requires effort, you're better prepared to
    take on new challenges.
  </div>
</div>
<br>
<div class="body">
  <div class="item">This is why i always enjoy sharing and teaching.
  It is what continues to be the bases for computer enjoyment.
  </div>
</div>