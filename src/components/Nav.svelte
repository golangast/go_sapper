<script>
  export let segment;
  import { fly } from "svelte-transitions";

  let visible;

  const onClick = () => {
    console.log(visible);
    if (visible == true) {
      visible = false;
    } else {
      visible = true;
    }
  };
</script>

<style>
  nav {
    border-bottom: 1px solid rgba(255, 62, 0, 0.1);
    font-weight: 300;
    padding: 0 1em;
    z-index: 1;
  }
  ul {
    margin: 0;
    padding: 0;
    z-index: 1;
  }
  /* clearfix */
  ul::after {
    content: "";
    display: block;
    clear: both;
  }
  li {
    display: block;
    float: left;
    z-index: 1;
  }
  [aria-current] {
    position: relative;
    display: inline-block;
  }
  [aria-current]::after {
    position: absolute;
    content: "";
    width: calc(100% - 1em);
    height: 2px;
    background-color: rgb(255, 62, 0);
    display: block;
    bottom: -1px;
  }
  a {
    text-decoration: none;
    padding: 1em 0.5em;
    display: block;
    z-index: -1;
  }
  .menus-s:hover div {
    color: purple;
    border: 1px solid purple;
  }

  .menuhalf {
    width: 30%;
    position: absolute;
    left: 10px;
    z-index: 9999999999999;
    box-shadow: 0px 0px 12px 0px rgba(0, 0, 0, 0.81);
    padding: 2%;
  }
  .menus-s {
    padding: 1em 0.5em;
    display: block;

    font-size: 1em;
    width: 5%;
  }
  .grid {
    flex-direction: row;
    width: 100%;
  }
  .grid-item {
    flex: 1;
    float: left;
    border: 1px solid #1a800078;
    text-decoration: none;
    margin: 1%;
    border-radius: 6px;
    font-size: 7pt;
    padding: 2%;
  }
</style>

<nav>
  <ul>
    <li>
      <div class="menus-s" on:click={onClick}>
        <div>&#9776;</div>
      </div>
      {#if visible}
        <div class="menuhalf" transition:fly={{ y: 200, duration: 2000 }}>
          <ul class="grid">
            <li class="grid-item">sdfa</li>
            <li class="grid-item">sdfa</li>
            <li class="grid-item">sdfa</li>
            <li class="grid-item">sdfa</li>
          </ul>
        </div>

        <ul />
      {/if}
    </li>
    <li>
      <a aria-current={segment === undefined ? 'page' : undefined} href=".">
        home
      </a>
    </li>
    <li>
      <a aria-current={segment === 'about' ? 'page' : undefined} href="about">
        about
      </a>
    </li>
    <li>
      <a aria-current={segment === 'form' ? 'page' : undefined} href="form">
        Form
      </a>
    </li>
    <li>
      <a aria-current={segment === 'show' ? 'page' : undefined} href="show">
        Show
      </a>
    </li>

    <!-- for the blog link, we're using rel=prefetch so that Sapper prefetches
		     the blog data when we hover over the link or tap it on a touchscreen -->
    <li>
      <a
        rel="prefetch"
        aria-current={segment === 'bloggy' ? 'page' : undefined}
        href="bloggy">
        Bloggy
      </a>
    </li>
  </ul>
</nav>
