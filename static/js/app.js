var Box = React.createClass({displayName: "Box",
	render: function() {
		return (
			React.createElement("div", {className: "Box"}, 
				React.createElement("input", {placeholder: "Test"}), 
				React.createElement("a", {className: "waves-effect waves-light btn"}, "Click")
			)
		);
	}
});

React.render(
	React.createElement(Box, null),
	$("#content").get(0)
);
