var Box = React.createClass({
	render: function() {
		return (
			<div className="Box">
				<input placeholder="Test"/>
			</div>
		);
	}
});

React.render(
	<Box />,
	$("#content").get(0)
);
