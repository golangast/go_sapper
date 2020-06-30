<script>
	import Apis, {actualapi, getResult} from "./_getapi.svelte";
	let actualtt = getResult("http://localhost:8080/api");
	let ff = [];
	var promise = Promise.resolve(actualtt).then(function(val) { 
			console.log(val); 
		ff = val;
	});  
</script>

{#each ff as {name, email, pass}, i }
 <ul>
 <li>{i}: {name} {email} {pass}</li>
 </ul>
 {/each}