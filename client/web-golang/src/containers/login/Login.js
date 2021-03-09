import React from 'react';
import './Login.scss'
import Loading from '../../components/loading/Loading'
import { Redirect } from "react-router-dom";
import Notification from '../../components/popup/Notification'
import{ userServiceLoginAccessToken } from '../../services/UserService'


export default class Login extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            isLoading: false,
            username: 'duynguyenngoc@hotmail.com',
            password: '1q2w3e4r',
            submitted: false,
            status_code: 0,
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
        if (!this.state.username) this.setState({authenticate: false})
        if (!this.state.password) this.setState({authenticate: false})
        userServiceLoginAccessToken(this.state.username, this.state.password) 
        // console.log()
    }
    

    componentDidMount() {
   
    }

    goToHome(){
        return (<Redirect to= "/"/>)
    }

    render(){
        const { username, password, submitted, status_code } = this.state;
        if(this.state.isLoading) {
            return (<Loading/>)
        }else {
                return (
                    <div>
                        <div>
                            {status_code===200 ? this.goToHome() : ""}
                            {status_code!==200 && status_code !== 0 ? <Notification/> : "" }
                        </div>
                        <div className="login">
                        <div className="title">
                            <h2>Login</h2>
                        </div>
                        <form name="form" onChange={this.handleChange}>
                            <div className="">
                                <input type="text" name="username" placeholder="Username" value={username} required/>
                            </div>
                            <div className="">
                                <input type="password" className="" name="password" placeholder="Password" value={password} required/>
                            </div>
                            <div>
                                <input type="button" value="Login" onClick={this.handleSubmit}></input>
                            </div>
                        </form>
                        <div className="booter"></div>
                    </div>
                     </div>
                    
                   
                );
        }
                
    }
}