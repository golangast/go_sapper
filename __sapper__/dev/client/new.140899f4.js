import { S as SvelteComponent, a as init, s as safe_not_equal, t as text, f as claim_text, k as insert, y as noop, g as detach } from './client.e22c960e.js';

/* src/routes/new.svelte generated by Svelte v3.23.2 */

function create_fragment(ctx) {
	let t0;
	let t1;

	return {
		c() {
			t0 = text("learning ");
			t1 = text(name);
		},
		l(nodes) {
			t0 = claim_text(nodes, "learning ");
			t1 = claim_text(nodes, name);
		},
		m(target, anchor) {
			insert(target, t0, anchor);
			insert(target, t1, anchor);
		},
		p: noop,
		i: noop,
		o: noop,
		d(detaching) {
			if (detaching) detach(t0);
			if (detaching) detach(t1);
		}
	};
}

let name = "world";

class New extends SvelteComponent {
	constructor(options) {
		super();
		init(this, options, null, create_fragment, safe_not_equal, {});
	}
}

export default New;
//# sourceMappingURL=data:application/json;charset=utf-8;base64,eyJ2ZXJzaW9uIjozLCJmaWxlIjoibmV3LjE0MDg5OWY0LmpzIiwic291cmNlcyI6WyIuLi8uLi8uLi9zcmMvcm91dGVzL25ldy5zdmVsdGUiXSwic291cmNlc0NvbnRlbnQiOlsiPHNjcmlwdCBsYW5nPVwidHNcIj5cbiAgICBsZXQgbmFtZSA9ICd3b3JsZCc7XG4gICAgXG4gICAgXG48L3NjcmlwdD5cblxuXG5sZWFybmluZyB7bmFtZX0iXSwibmFtZXMiOltdLCJtYXBwaW5ncyI6Ijs7Ozs7Ozs7Ozs7YUFPVSxJQUFJOzs7OzBCQUFKLElBQUk7Ozs7Ozs7Ozs7Ozs7Ozs7SUFOTixJQUFJLEdBQUcsT0FBTzs7Ozs7Ozs7Ozs7In0=
