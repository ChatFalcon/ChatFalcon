import React from "react";
import ReactDOM from "react-dom";
import InstallationKeyVerification from "./components/InstallationKeyVerification";
import RedisConfiguration from "./components/RedisConfiguration";
import S3Configuration from "./components/S3Configuration";
import MailgunConfiguration from "./components/MailgunConfiguration";
import ThankYou from "./components/ThankYou";
import UserConfiguration from "./components/UserConfiguration";

// Defines the various stages.
const stages = [
    InstallationKeyVerification,
    RedisConfiguration,
    S3Configuration,
    MailgunConfiguration,
    UserConfiguration,
    ThankYou,
];

// The base application class.
class App extends React.Component {
    constructor(props) {
        super(props);
        this.state = {err: null, stage: 0, nextCb: async() => {}, config: {}};
    }

    async manageStage() {
        this.setState({loading: true, err: null});
        try {
            try {
                await this.state.nextCb();
            } catch (err) {
                this.setState({err: err.message ? err.message : String(err)});
                return;
            }

            const stage = this.state.stage;
            if (stage + 1 === stages.length) {
                window.location.replace("/");
            } else {
                this.state.stage++;
            }
        } finally {
            this.setState({loading: false});
        }
    }

    render() {
        const Stage = stages[this.state.stage];
        return <section className="hero">
            <div className="hero-body">
                <div className="container">
                    {this.state.err ? <div className="notification is-danger">{this.state.err}</div> : undefined}
                    <Stage setNextCb={cb => this.setState({nextCb: cb})} config={this.state.config} stage={this.state.stage} />
                    <hr />
                    <p>
                        <a className={`button is-primary ${this.state.loading ? 'is-loading' : ''}`} onClick={this.manageStage.bind(this)}>{
                            stages.length - 1 === this.state.stage ? "Finish" : "Next"
                        }</a>
                    </p>
                </div>
            </div>
        </section>;
    }
}

// Render the application.
ReactDOM.render(<App />, document.getElementById("app"));
