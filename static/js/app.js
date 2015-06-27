var Box = React.createClass({displayName: "Box",
	render: function() {
		return (
			React.createElement("div", {className: "Box"}, 
				React.createElement("input", {placeholder: "Test"})
			)
		);
	}
});

React.render(
	React.createElement(Box, null),
	$("#content").get(0)
);
