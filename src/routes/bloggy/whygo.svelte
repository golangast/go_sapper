<style>
.body{border:1px solid rgb(255,62,0);background: #676778;color: #ffffffbf; width:100%;padding:2%; font-family: 'Work Sans', sans-serif;}
pre{
    tab-size: 4;
    font-size: 90%;
    display: block;
    overflow-x: auto;
    margin: 1.5em 0;
    padding: 1em 1.5em;
    border: 1px solid #c4c5c0;
    border-radius: 4px;
    background:#f0f0f0;
    color: #735e5ebf;
}
code{font-size: 100%;font-family: 'Work Sans', sans-serif;}
</style>

<h3>I Love Go because of it's simplicity.</h3> <br>
<div class="body">
<p>What does this mean? I am referring to its api and it's ability to uncouple code.
Why is this important? Many developers are bound to the constraints
to their language and tech stack.  The point is so that when your 
project gets bigger you are not constrained to what you have already
created.  
<br>
An example of what I am referring to.</p>
<pre>
<code>
{`type Guitarist struct{Guitarist } //object
func(m Musician) Run(){} // method
type Musician interface { Run()} //interface that uses the method
`}
</code>
</pre>


The three bits of code that are stated are pretty simple.
but they are the building blocks to uncoupling code.
<ul>
<li>
1. struct
</li>
<li>
2. method
</li>
<li>
3. interface
</li>
</ul>


These are the components needed to uncouple your code.  

Lets say you are building a Service on Musican.
<pre>
<code>
{`func Service(Musician){
    fmt.Println(Musician)
}
`}
</code>
</pre>

We used the interface Musician to allow for any struct with run() to be passed.
So if we needed to pass another struct of data in we can.

This allows for not having to think about your data or concrete data type.
Something many languages you are tied to.

I will go further in the next blog about packages and utilizing this method. 


</div>