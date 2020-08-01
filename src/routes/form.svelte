<script>
  let name = "name";
  let email = "email";
  let pass = "pass";
  let result = null;

  async function doPost() {
    const res = await fetch("http://localhost:8081/post", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        name,
        email,
        pass,
      }),
    })
      .then((res) => res.json())
      .then((data) => {
        console.log(data);
      })
      .catch((error) => {
        console.error(error);
      });

    const json = await res.json();
    result = JSON.stringify(json);
  }
</script>

<input bind:value={name} />
<input bind:value={email} />
<input bind:value={pass} />
<button type="button" on:click={doPost}>Post it.</button>
<p>Result:</p>
<pre>{result}</pre>
