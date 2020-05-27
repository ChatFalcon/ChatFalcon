import React from "react";
import saveConfig from "../saveConfig";

export default class InstallationKeyVerification extends React.Component {
    componentDidMount() {
        this.props.setNextCb(this.nextButtonPressed.bind(this));
    }

    async nextButtonPressed() {
        this.props.config.stage = this.props.stage;
        await saveConfig(this.props.config);
    }

    render() {
        return <div>
            <h1 className="title is-1">Welcome to ChatFalcon!</h1>
            <p>
                We'll get your forum up and running in no time!
                Firstly, what is the installation key for your forum? This will be printed in your console or likely shown during the setup if you used another deployment method.
            </p>
            <div className="field">
                <br />
                <label><b>Installation Key:</b></label>
                <input className="input" type="text" placeholder="Installation Key" value={this.props.config.installKey} onChange={e => this.props.config.installKey = e.target.value.trim()} />
            </div>
        </div>;
    }
}
