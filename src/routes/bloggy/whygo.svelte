<script>
import Index from "../index.svelte";

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
</style>

<h3 style="text-decoration:underline;">I Love Go because of it's simplicity.</h3> <br>
<div class="body">
I Love Go because of it's simplicity.
<p>What does this mean? I am referring to it's api and it's ability to uncouple code.
Why is this important? Many developers are bound to the constraints
to their language and tech stack.  The point is so that when your 
project gets bigger you are not constrained to what you have already
created.  
<br>
An example of what I am referring to.</p>
<pre>
<code>
{`type Guitarist struct{ Name string } //object

func (g Guitarist) Run() { fmt.Println(g) } // method

type Musician interface{ Run() } //interface that uses the method
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


Lets say you are building a Service on Musican.
<pre>
<code>
{
`type Musician interface{ Run() } //interface that uses the method
`}
</code>
</pre>
Meaning....<br>
We used the interface Musician to allow for any struct with run() to be passed in.
A lot of poeple miss this point when seeing the word interface because of it's
different meaning in other languages.  

This allows for not having to think about your data or concrete data type.
Something many languages you are tied to.

<br>
Another thing a lot of people miss from other languages is that you can do embedding 
on this interface type.

<pre>
<code>
{`type Musicman interface{ Musician }
`}
</code>
</pre>

By doing so you are attempting to turn an interface "type" into a value.
A value that initally does not have a value but carries methods.

<br><br>
<p style="color: #95e152;text-shadow: 0 0 3px #000000, 0 0 5px #f0f0f0;">Why would you want to do such a thing?</p>
<hr>
<br>

Really it is more of the later open decision to extend a type.
Meaning if you have a lot of methods tied up to this type and so
it is bound by the adding of methods, then you need a way to utilize 
this type but still not couple it to a hierachy. 
<pre>
<code>
{`
1. type Guitarist struct{ Name string } //object

2. func (g Guitarist) Run() { fmt.Println(g) } // method

3. func Running(m Musicman) { fmt.Println("Musician works") } // function

4. type Musician interface{ Run() } //interface that has the method

5. type Musicman interface{ Musician } //interface that is embedded

`}
</code>
</pre>
5. is just a function that takes in the embedded interface type.
The reason I added this was to demonstrate the use of embedded networkInterface types.
<br>
<br>
3. is utilizing the embedded interface as a parameter. 
<br>
<br>
Here is a working <a style="color:#f9845b" href="https://play.golang.org/p/Fs_PKx1mN6V" target="_blank">example</a> of the code. 

<br>
<hr>
An abstract class will only get you so far in uncoupling code.  Mainly because it is an object
and it becomes coupled when you want to start or change instances of it.  It also becomes 
heavy when you attempt to remove or add properties/methods and create objects of it.
This is why I personally love go.  It really shines in a future uncoupling sense. 
I hope people who read this can get a good idea of the importance of uncoupling code. 

</div>