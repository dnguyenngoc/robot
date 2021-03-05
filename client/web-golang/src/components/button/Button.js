import React from 'react';
import './Button.scss'


export default class Button extends React.Component {

    constructor(props) {
        super(props);
        this.state = {
            name: this.props.name
        };
    }
    
    render(){
        return (<button className="btn btn-primary">{this.state.name}</button>);
    }
}