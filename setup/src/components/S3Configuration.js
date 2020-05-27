import React from "react";
import saveConfig from "../saveConfig";

export default class S3Configuration extends React.Component {
    componentDidMount() {
        this.props.setNextCb(this.nextButtonPressed.bind(this));
    }

    async nextButtonPressed() {
        this.props.config.stage = this.props.stage;
        await saveConfig(this.props.config);
    }

    render() {
        return <div>
            <h1 className="title is-1">S3 Configuration</h1>
            <p>
                Ok, next we need to configure S3. This is a little more complicated, but we will guide you through this.
            </p>

            <div className="field">
                <br />
                <label><b>Region:</b></label>
                <p>The region which the S3 bucket is in.</p>
                <input className="input" type="text" placeholder="Region" value={this.props.config.s3Region} onChange={e => this.props.config.s3Region = e.target.value.trim()} />
            </div>

            <div className="field">
                <br />
                <label><b>Access Key ID:</b></label>
                <p>The access key ID for the S3 bucket.</p>
                <input className="input" type="text" placeholder="Access Key ID" value={this.props.config.s3AccessKeyId} onChange={e => this.props.config.s3AccessKeyId = e.target.value.trim()} />
            </div>

            <div className="field">
                <br />
                <label><b>Secret Access Key:</b></label>
                <p>The secret access key for the S3 bucket.</p>
                <input className="input" type="text" placeholder="Secret Access Key" value={this.props.config.s3SecretAccessKey} onChange={e => this.props.config.s3SecretAccessKey = e.target.value.trim()} />
            </div>

            <div className="field">
                <br />
                <label><b>Endpoint:</b></label>
                <p>The endpoint for the S3 bucket.</p>
                <input className="input" type="text" placeholder="Endpoint" value={this.props.config.s3Endpoint} onChange={e => this.props.config.s3Endpoint = e.target.value.trim()} />
            </div>

            <div className="field">
                <br />
                <label><b>Bucket:</b></label>
                <p>The bucket for the S3 bucket.</p>
                <input className="input" type="text" placeholder="Bucket" value={this.props.config.s3Bucket} onChange={e => this.props.config.s3Bucket = e.target.value.trim()} />
            </div>

            <div className="field">
                <br />
                <label><b>Bucket URL:</b></label>
                <p>The URL which leads to the S3 bucket.</p>
                <input className="input" type="text" placeholder="Bucket URL" value={this.props.config.s3BucketUrl} onChange={e => this.props.config.s3BucketUrl = e.target.value.trim()} />
            </div>
        </div>;
    }
}
