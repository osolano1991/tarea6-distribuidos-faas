function getAll(entity) {
	fetch('https://tarea6-distribuidos-faas-oscarsolanomora.netlify.app/api/' + entity)
	  .then((response) => response.json())
		.then((data) => {
			fetch('/template/list/' + entity + '.html')
				.then((response) => response.text())
				.then((template) => {
					var rendered = Mustache.render(template, data);
					document.getElementById('content').innerHTML = rendered;
				});
		})
}

function getById(query, entity) {
	var params = new URLSearchParams(query);
	fetch('https://tarea6-distribuidos-faas-oscarsolanomora.netlify.app/api/' + entity + '/?id=' + params.get('id'))
	  .then((response) => response.json())
		.then((data) => {
			fetch('/template/detail/' + entity + '.html')
				.then((response) => response.text())
				.then((template) => {
					var rendered = Mustache.render(template, data);
					document.getElementById('content').innerHTML = rendered;
				});
		})
}

function home() {
	fetch('/template/home.html')
		.then((response) => response.text())
		.then((template) => {
			var rendered = Mustache.render(template, {});
			document.getElementById('content').innerHTML = rendered;
		});
}

function init() {
	router = new Navigo(null, false, '#!');
	router.on({
		'/doctores': function() {
			getAll('doctores');
		},
		'/citas': function() {
			getAll('citas');
		},
		'/pacientes': function() {
			getAll('pacientes');
		},
		'/doctorById': function(_, query) {
			getById(query, 'doctores');
		},
		'/citaById': function(_, query) {
			getById(query, 'citas');
		},
		'/pacienteById': function(_, query) {
			getById(query, 'pacientes');
		}
	});
	router.on(() => home());
	router.resolve();
}
