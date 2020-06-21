<script context="module">
let arrays = [];
export async function getResult(s) {
        let response = await fetch(s);
        let text = await response.json();
				 if (response.ok) {
					 console.log(text);
				 		return text
				 } else {
					 throw new Error(text);
				 }
	}
 arrays = getResult();
 export const actualapi = arrays;	
</script>