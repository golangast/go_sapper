<script>
  let me = "me.png";
</script>

<style>
  @import url("https://fonts.googleapis.com/css?family=Source+Code+Pro:400,700");
  .header {
    display: grid;
    grid-template-columns: repeat(auto-fit, minmax(1em, 1fr));
    width:100%;
  }
  .hitem {
    text-align: center;
    color: #e15258;
    font-family: monospace;
    display:inline-block;
    font-size:90%;
    padding:1%;
 
  }
  .h-mid img {
   border:1px solid black;
   border-radius: 147px;
   background:#e15258;
  }
  .half {
    width: 100%;
    display: grid;
    grid-template-columns: auto auto;
    margin-top: 2em;
  }
  .half .item {
    text-align: center;
    color: #199c94;
    font-family: "SourceSansPro", Helvetica, Arial, "Meiryo UI",
      "ＭＳ Ｐゴシック", "MS PGothic", sans-serif;
  }
  p {
    color: #f9845b;
    font-family: Monaco, Menlo, Consolas, "Courier New", monospace;
  }
  .grid {
    flex-direction: column;
    width: 100%;
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
  }
</style>

<div class="header">
<div class="hitem h-mid">
    <img  src={me} alt="" width="30%" />
  </div>
</div>

<div class="half">
  <div class="item grid" style="font-weight:bold;">
    Things Ive created.
    <br />
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

    <a class="grid-item" href="https://gitlab.com/zendrulat123" target="_blank">
      Gitlab
    </a>
  </div>
  <div class="item grid" style="font-weight:bold;">
    Tech that I use.
    <br />
    <a class="grid-item" href="https://ubuntu.com/" target="_blank">
      Ubuntu 20.04
    </a>

    <a class="grid-item" href="https://code.visualstudio.com/" target="_blank">
      VSCODE
    </a>

    <a class="grid-item" href="https://svelte.dev/" target="_blank">SVELTEJS</a>
    <a class="grid-item" href="https://sapper.svelte.dev/docs/" target="_blank">Sapper</a>

    <a class="grid-item" href="https://golang.org/" target="_blank">Go</a>
  </div>
</div>

<p style="width:100%; text-align:center;">
  I've been loving Go + Sveltejs incrediably a lot. Not much is simplier. I
  decided my next plan is to write blogs about what I am learning and thoughts
  about tech. There is a lot that I love to dig into and a lot that I try to
  learn.
</p>