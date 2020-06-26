<script>
import { has_prop } from "svelte/internal";

</script>
<style>
.body{border:1px solid rgb(255,62,0);background: #676778;color: #ffffffbf; width:100%;padding:2%; font-family: 'Work Sans', sans-serif;}
pre{
    tab-size: 4;
    font-size: 90%;
    display: block;
    overflow-x: auto;
    margin: 1.5em 0;
    padding: 1em 1.5em;
       border: 1px solid rgb(255,62,0);
    border-radius: 4px;
    background:#f0f0f0;
    color: #735e5ebf;
}
code{font-size: 100%;font-family: 'Work Sans', sans-serif;}
span{
    color:rgb(255,62,0);
    background:#f0f0f0;
    border-radius: 3px;
    padding: .5%;
    margin:.5%;

}
a{
        color:rgb(255,62,0);
}
</style>

<h3 style="text-decoration:underline;">What do I do with Go?.</h3> <br>
<div class="body">
What do I do with Go?.
<p>
I usually build servers.  I enjoy the nice API that it has.
A good example is the <a href="https://golang.org/pkg/net/http/">http package</a>
The <span>ListenAndServe </span>method is really where the server begins.
</p>
<pre>
<code>
{`http.HandleFunc("/hello", helloHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
`}
</code>
</pre>
The <span>HandleFunc</span> is how routing is setup.
I think a lot of people dont realize how easy this is to read.
They just presume this is how go does it and dont stop to 
realize how simple this is.
</div>