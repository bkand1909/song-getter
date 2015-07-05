var Box = React.createClass({
	render: function() {
		return (
			<div className="Box">
				<input placeholder="Test"/>
				<a className="waves-effect waves-light btn">Click</a>
			</div>
		);
	}
});

React.render(
	<Box />,
	$("#content").get(0)
);
