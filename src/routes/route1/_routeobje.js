const routeob = [{Title:"great title", Slug:"thisroute", Html:"<p>fdsaf</p>"},{Title:"great title", Slug:"thisroute", Html:"<p>fdsaf</p>"}]


routeob.forEach(singleroute => {
	singleroute.html = singleroute.html.replace(/^\t{3}/gm, '');
});

export default routeob;