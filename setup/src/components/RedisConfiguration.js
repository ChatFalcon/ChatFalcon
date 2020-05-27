import React from "react";
import saveConfig from "../saveConfig";

export default class RedisConfiguration extends React.Component {
    componentDidMount() {
        this.props.setNextCb(this.nextButtonPressed.bind(this));
    }

    async nextButtonPressed() {
        this.props.config.stage = this.props.stage;
        await saveConfig(this.props.config);
    }

    render() {
        return <div>
            <h1 className="title is-1">Redis Configuration</h1>
            <p>
                Ok awesome! Next question, what is your Redis hostname/password if you are using it? If not, leave this blank.
            </p>
            <div className="field">
                <br />
                <label><b>Hostname:</b></label>
                <input className="input" type="text" placeholder="Hostname" value={this.props.config.redisHostname} onChange={e => this.props.config.redisHostname = e.target.value.trim()} />
            </div>
            <div className="field">
                <br />
                <label><b>Password:</b></label>
                <input className="input" type="text" placeholder="Password" value={this.props.config.redisPassword} onChange={e => this.props.config.redisPassword = e.target.value.trim()} />
            </div>
        </div>;
    }
}
