import React from "react";
import saveConfig from "../saveConfig";

export default class ThankYou extends React.Component {
    componentDidMount() {
        this.props.setNextCb(this.nextButtonPressed.bind(this));
    }

    async nextButtonPressed() {
        this.props.config.stage = this.props.stage;
        await saveConfig(this.props.config);
    }

    render() {
        return <div>
            <h1 className="title is-1">Thank You!</h1>
            <p>
                The setup process is now complete and is ready for the automatic initialisation. Simply click finish to continue.
            </p>
        </div>;
    }
}
