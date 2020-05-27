import React from "react";
import saveConfig from "../saveConfig";

export default class MailgunConfiguration extends React.Component {
    componentDidMount() {
        this.props.setNextCb(this.nextButtonPressed.bind(this));
    }

    async nextButtonPressed() {
        this.props.config.stage = this.props.stage;
        await saveConfig(this.props.config);
    }

    render() {
        return <div>
            <h1 className="title is-1">Mailgun Configuration</h1>
            <p>
                Ok, lets setup Mailgun now. Simply enter the details below. <b>Note that your Mailgun instance needs to be in the EU region.</b>
            </p>

            <div className="field">
                <br />
                <label><b>Domain:</b></label>
                <input className="input" type="text" placeholder="Domain" value={this.props.config.mailgunDomain} onChange={e => this.props.config.mailgunDomain = e.target.value.trim()} />
            </div>

            <div className="field">
                <br />
                <label><b>Private Key:</b></label>
                <input className="input" type="text" placeholder="Private Key" value={this.props.config.mailgunPrivateKey} onChange={e => this.props.config.mailgunPrivateKey = e.target.value.trim()} />
            </div>

            <div className="field">
                <br />
                <label><b>Sender:</b></label>
                <input className="input" type="text" placeholder="Sender" value={this.props.config.mailgunFrom} onChange={e => this.props.config.mailgunFrom = e.target.value.trim()} />
            </div>
        </div>;
    }
}
