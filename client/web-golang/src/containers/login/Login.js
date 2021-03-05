import React from 'react';
import './Login.scss'
import Loading from '../../components/loading/Loading'
import Button from '../../components/button/Button';
import { Redirect } from "react-router-dom";



export default class Login extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            isLoading: false,
            username: '',
            password: '',
            submitted: false,
            authenticate: false,
        };
        this.handleChange = this.handleChange.bind(this);
        this.handleSubmit = this.handleSubmit.bind(this);
    }

    handleChange(e) {
        const { name, value } = e.target;
        this.setState({ [name]: value });
    }

    handleSubmit(e) {
        e.preventDefault();
        this.setState({ submitted: true });

    }

    componentDidMount() {
   
    }

    render(){

        const { username, password, submitted, authenticate } = this.state;
        if(this.state.isLoading) {
            return (<Loading/>)
        }else {
            if (authenticate) {
                return <Redirect to={'/xxx'} />
            }else {
                return (
                    <div className="col-md-offset-3 col-md-6">
                        <h2>Login</h2>
                        <form name="form" onSubmit={this.handleSubmit}>
                            <div className={'form-group' + (submitted && !username ? ' has-error' : '')}>
                                <label htmlFor="username">Username</label>
                                <input type="text" className="form-control" name="username" value={username} onChange={this.handleChange} />
                                {submitted && !username &&
                                    <div className="help-block">Username is required</div>
                                }
                            </div>
                            <div className={'form-group' + (submitted && !password ? ' has-error' : '')}>
                                <label htmlFor="password">Password</label>
                                <input type="password" className="form-control" name="password" value={password} onChange={this.handleChange} />
                                {submitted && !password &&
                                    <div className="help-block">Password is required</div>
                                }
                            </div>
                            <div className="form-group">
                                <Button  key="name" name="Login"/>
                            </div>
                        </form>
                    </div>
                );
            }
        }
       
    }
}