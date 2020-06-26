<script>
  import { fade } from "svelte/transition";
  import { elasticOut } from "svelte/easing";
  let me = "me.png";
  let visible;
  let visible2;
  function check() {
    console.log(visible);
    if (visible == true) {
      visible = false;
    } else {
      visible = true;
    }
  }
  function check2() {
    console.log(visible2);
    if (visible2 == true) {
      visible2 = false;
    } else {
      visible2 = true;
    }
  }
  function spin(node, { duration }) {
    return {
      duration,
      css: (t) => {
        const eased = elasticOut(t);

        return `
					transform: scale(${eased}) rotate(${eased * 1080}deg);
					color: hsl(
						${~~(t * 360)},
						${Math.min(100, 1000 - 1000 * t)}%,
						${Math.min(50, 500 - 500 * t)}%
					);`;
      },
    };
  }
</script>

<style>
  @import url("https://fonts.googleapis.com/css?family=Source+Code+Pro:400,700");
  .header {
    display: grid;
    grid-template-columns: repeat(auto-fit, minmax(1em, 1fr));
    width: 100%;
  }
  .hitem {
    text-align: center;
    color: #e15258;
    font-family: monospace;
    display: inline-block;
    font-size: 90%;
    padding: 1%;
  }
  .h-mid img {
    border: 1px solid black;
    border-radius: 147px;
    background: #e15258;
  }

  p {
    color: #f9845b;
    font-family: Monaco, Menlo, Consolas, "Courier New", monospace;
  }
  .grid {
    flex-direction: column;
    width: 100%;
  }
    .gridtitle{
    flex-direction: row;
    width: 50%;
    float:left;
    text-align:center;

  }
  .grid-item {
    flex: 1;
    float: left;
    width: 23%;
    border: 1px solid #1a800078;
    text-decoration: none;
    margin: 1%;
    border-radius: 6px;
    font-size: 7pt;
    padding: 2%;
    text-align:center;
  }
  section > div {
    margin-bottom: 40px;
  }
  .button {
    width: 50%;
    margin: 0 auto;
    border: 1px solid purple;
  }
</style>

<div class="header">
  <div class="hitem h-mid">
    <img src={me} alt="" width="30%" />
  </div>
</div>

<div class="item gridtitle" style="font-weight:bold;" on:click={check}>
  <div class="button">Things Ive created.</div>
  <br />
  {#if visible}
    <div in:spin={{ duration: 8000 }} out:fade>
      <a
        class="grid-item"
        href="https://docs.google.com/document/d/1Zb9GCWPKeEJ4Dyn2TkT-O3wJ8AFc-IMxZzTugNCjr-8/edit?usp=sharing"
        target="_blank">
        Resources
      </a>

      <a
        class="grid-item"
        href="https://medium.com/jexia/master-the-world-of-golang-issue-9-2bdfec563c2e"
        target="_blank">
        Tutorials
      </a>

      <a
        class="grid-item"
        href="https://gitlab.com/zendrulat123"
        target="_blank">
        Gitlab
      </a>
    </div>
  {/if}
</div>

<div class="item gridtitle" style="font-weight:bold;" on:click={check2}>
  <div class="button">Tech that I use.</div>
  <br />
  {#if visible2}
    <div class="grid" in:spin={{ duration: 8000 }} out:fade>
      <a class="grid-item" href="https://ubuntu.com/" target="_blank">
        Ubuntu 20.04
      </a>

      <a
        class="grid-item"
        href="https://code.visualstudio.com/"
        target="_blank">
        VSCODE
      </a>

      <a class="grid-item" href="https://svelte.dev/" target="_blank">
        SVELTEJS
      </a>
      <a
        class="grid-item"
        href="https://sapper.svelte.dev/docs/"
        target="_blank">
        Sapper
      </a>

      <a class="grid-item" href="https://golang.org/" target="_blank">Go</a>
    </div>
  {/if}
</div>

<p style="width:100%; text-align:center;">
  I've been loving Go + Sveltejs incrediably a lot. Not much is simplier. I
  decided my next plan is to write blogs about what I am learning and thoughts
  about tech. There is a lot that I love to dig into and a lot that I try to
  learn.
</p>
